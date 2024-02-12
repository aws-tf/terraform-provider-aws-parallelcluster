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
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

const (
	ClusterReadyTimeout = 180 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_                      resource.Resource                = &ClusterResource{}
	_                      resource.ResourceWithImportState = &ClusterResource{}
	failedToFindClusterErr                                  = "404 Not Found"
)

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ClusterResourceModel describes the resource data model.
type ClusterResourceModel struct {
	ClusterName               types.String `tfsdk:"cluster_name"`
	ClusterConfiguration      types.String `tfsdk:"cluster_configuration"`
	Region                    types.String `tfsdk:"region"`
	RollbackOnFailure         types.Bool   `tfsdk:"rollback_on_failure"`
	SuppressValidators        types.List   `tfsdk:"suppress_validators"`
	ValidationFailureLevel    types.String `tfsdk:"validation_failure_level"`
	Id                        types.String `tfsdk:"id"`
	CloudformationStackArn    types.String `tfsdk:"cloudformation_stack_arn"`
	CloudformationStackStatus types.String `tfsdk:"cloudformation_stack_status"`
	ClusterStatus             types.String `tfsdk:"cluster_status"`
	Version                   types.String `tfsdk:"version"`
}

func (r *ClusterResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *ClusterResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Create a managed cluster in an AWS Region.",

		Attributes: map[string]schema.Attribute{
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "The name of the cluster to create. The name must start with an alphabetical character. The name can have up to 60 characters. If Slurm accounting is enabled, the name can have up to 40 characters.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"cluster_configuration": schema.StringAttribute{
				MarkdownDescription: "The cluster configuration as a YAML document.",
				Required:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that the cluster is in.",
				Optional:            true,
				Computed:            true,
			},
			"rollback_on_failure": schema.BoolAttribute{
				MarkdownDescription: "If set to true, cluster stack rollback occurs if the cluster fails to create. The default is true.",
				Optional:            true,
			},
			"suppress_validators": schema.ListAttribute{
				MarkdownDescription: "Identify one or more configuration validators to suppress.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"validation_failure_level": schema.StringAttribute{
				MarkdownDescription: "The minimum validation level that causes cluster create to fail. The default is ERROR.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Cluster identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cloudformation_stack_arn": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "The Amazon Resource Name (ARN) of the main CloudFormation stack.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cloudformation_stack_status": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "Status of the cloudformation stack.",
			},
			"cluster_status": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "Status of the cluster.",
			},
			"version": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "The AWS ParallelCluster version that's used to create the cluster.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ClusterResource) waitClusterReady(
	ctx context.Context,
	id string,
) (*openapi.DescribeClusterResponseContent, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(openapi.CLUSTERSTATUS_CREATE_IN_PROGRESS),
			string(openapi.CLUSTERSTATUS_DELETE_IN_PROGRESS),
			string(openapi.CLUSTERSTATUS_UPDATE_IN_PROGRESS),
			"UPDATE_COMPLETE_CLEANUP_IN_PROGRESS",
		},
		Target: []string{
			string(openapi.CLUSTERSTATUS_CREATE_COMPLETE),
			string(openapi.CLUSTERSTATUS_CREATE_FAILED),
			string(openapi.CLUSTERSTATUS_UPDATE_COMPLETE),
			string(openapi.CLUSTERSTATUS_UPDATE_FAILED),
			string(openapi.CLUSTERSTATUS_DELETE_COMPLETE),
			string(openapi.CLUSTERSTATUS_DELETE_FAILED),
		},
		Refresh: r.clusterStatus(ctx, id),
		Timeout: ClusterReadyTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if output, ok := outputRaw.(openapi.DescribeClusterResponseContent); ok {
		return &output, err
	}
	return nil, err
}

func (r *ClusterResource) clusterStatus(
	ctx context.Context,
	id string,
) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cluster, err := r.getCluster(ctx, id)
		if err != nil {
			return nil, "", err
		}

		return cluster, string(cluster.ClusterStatus), err
	}
}

func (r *ClusterResource) getCluster(
	ctx context.Context, clusterName string,
) (openapi.DescribeClusterResponseContent, error) {
	cluster, _, err := r.client.ClusterOperationsAPI.DescribeCluster(
		ctx,
		clusterName,
	).Execute()
	if err != nil {
		return openapi.DescribeClusterResponseContent{}, err
	}

	return *cluster, err
}

func (r *ClusterResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(configData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf(
				"Expected *openapi.Client, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)

		return
	}

	r.client = config.client
	r.awsv4 = config.awsv4
}

func populateClusterResourceDesc(
	desc *openapi.DescribeClusterResponseContent,
	data *ClusterResourceModel,
) {
	if desc != nil {
		if desc.ClusterName == "" {
			data.ClusterName = types.StringNull()
			data.Id = types.StringNull()
		} else {
			data.ClusterName = types.StringValue(desc.ClusterName)
			data.Id = types.StringValue(desc.ClusterName)
		}
		if desc.Region == "" {
			data.Region = types.StringNull()
		} else {
			data.Region = types.StringValue(desc.Region)
		}
		if desc.Version == "" {
			data.Version = types.StringNull()
		} else {
			data.Version = types.StringValue(desc.Version)
		}
		if desc.CloudformationStackArn == "" {
			data.CloudformationStackArn = types.StringNull()
		} else {
			data.CloudformationStackArn = types.StringValue(desc.CloudformationStackArn)
		}
		if desc.ClusterStatus == "" {
			data.ClusterStatus = types.StringNull()
		} else {
			data.ClusterStatus = types.StringValue(string(desc.ClusterStatus))
		}
		if desc.CloudFormationStackStatus == "" {
			data.CloudformationStackStatus = types.StringNull()
		} else {
			data.CloudformationStackStatus = types.StringValue(string(desc.CloudFormationStackStatus))
		}
		if url, ok := desc.ClusterConfiguration.GetUrlOk(); ok {
			if *url == "" {
				data.ClusterConfiguration = types.StringNull()
			} else {
				s3Resp, err := http.Get(*url)
				if err == nil {
					defer s3Resp.Body.Close()
					configBytes, err := io.ReadAll(s3Resp.Body)
					if err == nil {
						data.ClusterConfiguration = types.StringValue(string(configBytes))
					}

				}
			}
		}
	}
}

func populateClusterResourceInfo(cluster *openapi.ClusterInfoSummary, data *ClusterResourceModel) {
	if cluster != nil {
		if cluster.ClusterName == "" {
			data.ClusterName = types.StringNull()
			data.Id = types.StringNull()
		} else {
			data.ClusterName = types.StringValue(cluster.ClusterName)
			data.Id = types.StringValue(cluster.ClusterName)
		}
		if cluster.Region == "" {
			data.Region = types.StringNull()
		} else {
			data.Region = types.StringValue(cluster.Region)
		}
		if cluster.Version == "" {
			data.Version = types.StringNull()
		} else {
			data.Version = types.StringValue(cluster.Version)
		}
		if cluster.CloudformationStackArn == "" {
			data.CloudformationStackArn = types.StringNull()
		} else {
			data.CloudformationStackArn = types.StringValue(cluster.CloudformationStackArn)
		}
		if cluster.ClusterStatus == "" {
			data.ClusterStatus = types.StringNull()
		} else {
			data.ClusterStatus = types.StringValue(string(cluster.ClusterStatus))
		}
		if cluster.CloudformationStackStatus == "" {
			data.CloudformationStackStatus = types.StringNull()
		} else {
			data.CloudformationStackStatus = types.StringValue(string(cluster.CloudformationStackStatus))
		}
	}
}

func (r *ClusterResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	createClusterRequestContent := *openapi.NewCreateClusterRequestContent(data.ClusterName.ValueString(), data.ClusterConfiguration.ValueString()) // CreateClusterRequestContent |

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	createClusterReq := r.client.ClusterOperationsAPI.CreateCluster(reqCtx)

	if !data.Region.IsNull() {
		createClusterReq = createClusterReq.Region(data.Region.ValueString())
	}

	if !data.ValidationFailureLevel.IsNull() {
		createClusterReq = createClusterReq.ValidationFailureLevel(
			openapi.ValidationLevel(data.ValidationFailureLevel.ValueString()),
		)
	}

	suppressValidators := make([]string, 0)
	for _, s := range data.SuppressValidators.Elements() {
		suppressValidators = append(suppressValidators, s.String())
	}

	tflog.Info(ctx, "Creating Parallel Cluster")
	httpResp, rawHttp, err := createClusterReq.
		CreateClusterRequestContent(createClusterRequestContent).
		SuppressValidators(suppressValidators).
		RollbackOnFailure(data.RollbackOnFailure.ValueBool()).
		Execute()
	if err != nil {
		if rawHttp != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("%v", err),
				fmt.Sprintf("%v", rawHttp.Body),
			)
		} else {
			resp.Diagnostics.AddError(
				"Error while creating cluster.",
				fmt.Sprintf("%v", err),
			)
		}
		return
	}
	// response from `CreateCluster`: CreateClusterResponseContent
	tflog.Trace(ctx, fmt.Sprintf("Response from `ClusterOperationsAPI.CreateCluster`: %v\n", resp))

	if cluster, ok := httpResp.GetClusterOk(); ok {
		populateClusterResourceInfo(cluster, &data)

		clusterDesc, err := r.waitClusterReady(
			reqCtx,
			data.ClusterName.ValueString(),
		)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error while waiting for cluster to finish updating.",
				fmt.Sprintf("%v", err),
			)
		} else {
			populateClusterResourceDesc(clusterDesc, &data)
			if data.ClusterStatus.ValueString() == string(openapi.CLUSTERSTATUS_CREATE_FAILED) {
				resp.Diagnostics.AddError(
					"Error while updating cluster.",
					fmt.Sprintf("Cluster Status: %v", data.ClusterStatus.ValueString()),
				)
			}
		}
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	clusterDesc, err := r.getCluster(reqCtx, data.Id.ValueString())
	if err != nil {
		if err.Error() == failedToFindClusterErr {
			resp.Diagnostics.AddWarning("Failed to find cluster.", err.Error())
		} else {
			resp.Diagnostics.AddError("Failure while retrieving cluster.", err.Error())
		}
	}

	populateClusterResourceDesc(&clusterDesc, &data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var planData ClusterResourceModel
	var stateData ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	clusterUpdateRequestContent := openapi.NewUpdateClusterRequestContent(
		planData.ClusterConfiguration.ValueString(),
	)

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	clusterUpdateRequest := r.client.ClusterOperationsAPI.UpdateCluster(
		reqCtx,
		planData.ClusterName.ValueString(),
	)

	clusterUpdateRequest = clusterUpdateRequest.UpdateClusterRequestContent(
		*clusterUpdateRequestContent,
	)

	if !planData.Region.IsNull() {
		clusterUpdateRequest = clusterUpdateRequest.Region(planData.Region.ValueString())
	}

	if !planData.ValidationFailureLevel.IsNull() {
		clusterUpdateRequest = clusterUpdateRequest.ValidationFailureLevel(
			openapi.ValidationLevel(planData.ValidationFailureLevel.ValueString()),
		)
	}

	// If the cluster config and region are unchanged nothing is left to do.
	if planData.ClusterConfiguration == stateData.ClusterConfiguration {
		if planData.Region.Equal(stateData.Region) {
			resp.Diagnostics.Append(req.State.Set(ctx, &planData)...)
			readResp := &resource.ReadResponse{
				State:       resp.State,
				Private:     resp.Private,
				Diagnostics: resp.Diagnostics,
			}
			r.Read(
				ctx,
				resource.ReadRequest{
					State:        req.State,
					Private:      req.Private,
					ProviderMeta: req.ProviderMeta,
				},
				readResp,
			)
			resp.Diagnostics.Append(readResp.Diagnostics...)
			resp.State = readResp.State
			return
		}
	}

	content, fullResp, err := clusterUpdateRequest.Execute()
	if err != nil || content == nil {
		if fullResp != nil {
			resp.Diagnostics.AddError("Failure while updating cluster.",
				fmt.Sprintf("Error: %v\nMessage: %v\n", err, fullResp.Body),
			)
		} else {
			resp.Diagnostics.AddError("Failure while updating cluster.",
				fmt.Sprintf("Error: %v\n", err),
			)
		}
		return
	}

	if cluster, ok := content.GetClusterOk(); ok {
		populateClusterResourceInfo(cluster, &planData)

		clusterDesc, err := r.waitClusterReady(
			reqCtx,
			planData.ClusterName.ValueString(),
		)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error while waiting for cluster to finish updating.",
				fmt.Sprintf("%v", err),
			)
		} else {
			populateClusterResourceDesc(clusterDesc, &planData)
			if planData.ClusterStatus.ValueString() == string(openapi.CLUSTERSTATUS_UPDATE_FAILED) {
				resp.Diagnostics.AddError(
					"Error while updating cluster.",
					fmt.Sprintf("Cluster Status: %v", planData.ClusterStatus.ValueString()),
				)
			}
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *ClusterResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)
	deleteClusterRequest := r.client.ClusterOperationsAPI.DeleteCluster(
		reqCtx,
		data.ClusterName.ValueString(),
	)

	if !data.Region.IsNull() {
		deleteClusterRequest = deleteClusterRequest.Region(data.Region.ValueString())
	}

	_, _, err := deleteClusterRequest.Execute()
	if err != nil && err.Error() != failedToFindClusterErr {
		resp.Diagnostics.AddError("Failure occurred while deleting cluster", err.Error())
	}
	clusterDesc, err := r.waitClusterReady(reqCtx, data.ClusterName.ValueString())
	if err != nil && err.Error() != failedToFindClusterErr {
		resp.Diagnostics.AddError(
			"Cluster delete failed to complete.",
			err.Error(),
		)
	}
	if clusterDesc != nil {
		populateClusterResourceDesc(clusterDesc, &data)
		if data.ClusterStatus.ValueString() == string(openapi.CLUSTERSTATUS_DELETE_FAILED) {
			resp.Diagnostics.AddError(
				"Error while updating cluster.",
				fmt.Sprintf("Cluster Status: %v", data.ClusterStatus.ValueString()),
			)
		}

		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}
}

func (r *ClusterResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	var data ClusterResourceModel

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	clusterDesc, err := r.getCluster(reqCtx, req.ID)

	if err != nil {
		resp.Diagnostics.AddError("Failed to find cluster.", err.Error())
		return
	}

	populateClusterResourceDesc(&clusterDesc, &data)

	data.SuppressValidators = types.ListNull(types.StringType)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
