// Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "LICENSE.txt" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
// implied. See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &ComputeFleetStatusResource{}
	_ resource.ResourceWithImportState = &ComputeFleetStatusResource{}
)

func NewComputeFleetStatusResource() resource.Resource {
	return &ComputeFleetStatusResource{}
}

// ComputeFleetStatusResource defines the resource implementation.
type ComputeFleetStatusResource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ComputeFleetStatusResourceModel describes the resource data model.
type ComputeFleetStatusResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	ClusterName          types.String `tfsdk:"cluster_name"`
	StatusRequest        types.String `tfsdk:"status_request"`
	Status               types.String `tfsdk:"status"`
	Region               types.String `tfsdk:"region"`
	LastStatusUpdateTime types.String `tfsdk:"last_status_update_time"`
}

func (r *ComputeFleetStatusResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_compute_fleet_status"
}

func requestedComputeFleetStatusValidator(
	ctx context.Context,
	req validator.StringRequest,
	resp *validator.StringResponse,
) {
	_, err := openapi.NewRequestedComputeFleetStatusFromValue(req.ConfigValue.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Validation Error",
			fmt.Sprintf("Valid values: %v", openapi.AllowedRequestedComputeFleetStatusEnumValues),
		)
	}
}

func (r *ComputeFleetStatusResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Update the status of the cluster compute fleet.",

		Attributes: map[string]schema.Attribute{
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "The name of the cluster.",
				Optional:            false,
				Required:            true,
			},
			"status_request": schema.StringAttribute{
				MarkdownDescription: "The compute fleet status.",
				Optional:            false,
				Required:            true,
				Validators: []validator.String{
					&AttributeValidator{
						description: fmt.Sprintf(
							"Valid values: %v",
							openapi.AllowedRequestedComputeFleetStatusEnumValues,
						),
						markdownDescription: fmt.Sprintf(
							"Valid values: `%v`",
							openapi.AllowedRequestedComputeFleetStatusEnumValues,
						),
						validatorFunction: requestedComputeFleetStatusValidator,
					},
				},
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The compute fleet status.",
				Optional:            false,
				Required:            false,
				Computed:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that the cluster is in.",
				Optional:            true,
				Required:            false,
			},
			"last_status_update_time": schema.StringAttribute{
				MarkdownDescription: "The timestamp that represents the last status update time.",
				Computed:            true,
				Optional:            false,
				Required:            false,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ComputeFleetStatus identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ComputeFleetStatusResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(configData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf(
				"Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)

		return
	}

	r.client = client.client
	r.awsv4 = client.awsv4
}

func (r *ComputeFleetStatusResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data ComputeFleetStatusResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = data.ClusterName
	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	computeFleetReqContent := openapi.UpdateComputeFleetRequestContent{}
	status, err := openapi.NewRequestedComputeFleetStatusFromValue(data.StatusRequest.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure while updating compute fleet status",
			fmt.Sprintf("%v", err),
		)
		return
	}

	computeFleetReqContent.SetStatus(*status)
	computeFleetReq := r.client.ClusterComputeFleetAPI.UpdateComputeFleet(
		reqCtx,
		data.ClusterName.ValueString(),
	)
	computeFleetReq = computeFleetReq.UpdateComputeFleetRequestContent(computeFleetReqContent)

	respContent, rawHttp, err := computeFleetReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("%v", err), fmt.Sprintf("%v", rawHttp.Body))
		return
	}

	if respContent != nil {
		data.Status = types.StringValue(string(respContent.GetStatus()))
		data.LastStatusUpdateTime = types.StringValue(
			respContent.GetLastStatusUpdatedTime().String(),
		)
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeFleetStatusResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data ComputeFleetStatusResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)
	name := data.ClusterName.ValueString()

	computeFleet, rawHttp, err := r.client.ClusterComputeFleetAPI.DescribeComputeFleet(reqCtx, name).
		Execute()
	if err != nil {
		resp.Diagnostics.AddWarning(fmt.Sprintf("%v", err.Error()), fmt.Sprintf("%v", rawHttp.Body))
	}

	data.Status = types.StringValue(string(computeFleet.GetStatus()))
	data.LastStatusUpdateTime = types.StringValue(
		computeFleet.GetLastStatusUpdatedTime().String(),
	)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeFleetStatusResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var data ComputeFleetStatusResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	computeFleetReqContent := openapi.UpdateComputeFleetRequestContent{}
	status, err := openapi.NewRequestedComputeFleetStatusFromValue(data.StatusRequest.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure while updating compute fleet status",
			fmt.Sprintf("%v", err),
		)
		return
	}

	computeFleetReqContent.SetStatus(*status)
	computeFleetReq := r.client.ClusterComputeFleetAPI.UpdateComputeFleet(
		reqCtx,
		data.ClusterName.ValueString(),
	)
	computeFleetReq = computeFleetReq.UpdateComputeFleetRequestContent(computeFleetReqContent)

	respContent, rawHttp, err := computeFleetReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("%v", err), fmt.Sprintf("%v", rawHttp.Body))
		return
	}

	if respContent != nil {
		data.Status = types.StringValue(string(respContent.GetStatus()))
		data.LastStatusUpdateTime = types.StringValue(
			respContent.GetLastStatusUpdatedTime().String(),
		)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeFleetStatusResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data ComputeFleetStatusResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ComputeFleetStatusResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
