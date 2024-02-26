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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

const (
	imageReadyTimeout = 180 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_                    resource.Resource                = &ImageResource{}
	_                    resource.ResourceWithImportState = &ImageResource{}
	failedToFindImageErr                                  = "404 Not Found"
)

func (r *ImageResource) waitImageReady(
	ctx context.Context,
	id string,
) (*openapi.DescribeImageResponseContent, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(openapi.IMAGEBUILDSTATUS_BUILD_IN_PROGRESS),
			string(openapi.IMAGEBUILDSTATUS_DELETE_IN_PROGRESS),
		},
		Target: []string{
			string(openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE),
			string(openapi.IMAGEBUILDSTATUS_BUILD_FAILED),
			string(openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE),
			string(openapi.IMAGEBUILDSTATUS_DELETE_FAILED),
		},
		Refresh: r.imageStatus(ctx, id),
		Timeout: imageReadyTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if output, ok := outputRaw.(openapi.DescribeImageResponseContent); ok {
		return &output, nil
	}
	return nil, err
}

func (r *ImageResource) imageStatus(
	ctx context.Context,
	id string,
) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		image, err := r.getImage(ctx, id)
		if err != nil {
			return nil, "", err
		}
		return image, string(image.ImageBuildStatus), err
	}
}

func (r *ImageResource) getImage(
	ctx context.Context,
	imageId string,
) (openapi.DescribeImageResponseContent, error) {
	image, _, err := r.client.ImageOperationsAPI.DescribeImage(ctx, imageId).Execute()
	if err != nil {
		return openapi.DescribeImageResponseContent{}, err
	}

	return *image, err
}

func NewImageResource() resource.Resource {
	return &ImageResource{}
}

// ImageResource defines the resource implementation.
type ImageResource struct {
	client     *openapi.APIClient
	awsv4      openapi.AWSv4
	role       string
	expiration time.Time
}

// ImageResourceModel describes the resource data model.
type ImageResourceModel struct {
	ImageId                types.String `tfsdk:"image_id"`
	Id                     types.String `tfsdk:"id"`
	ImageConfiguration     types.String `tfsdk:"image_configuration"`
	SuppressValidators     types.List   `tfsdk:"suppress_validators"`
	ValidationFailureLevel types.String `tfsdk:"validation_failure_level"`
	RollbackOnFailure      types.Bool   `tfsdk:"rollback_on_failure"`
	Region                 types.String `tfsdk:"region"`

	// Image Builder Response Content
	Version                   types.String `tfsdk:"version"`
	CloudformationStackArn    types.String `tfsdk:"cloudformation_stack_arn"`
	CloudformationStackStatus types.String `tfsdk:"cloudformation_stack_status"`
	ImageBuildStatus          types.String `tfsdk:"image_build_status"`

	// Ec2 Info Response Content
	AmiId types.String `tfsdk:"ami_id"`
}

func (r *ImageResource) getClient() *openapi.APIClient {
	return r.client
}

func (r *ImageResource) getAWSv4() openapi.AWSv4 {
	return r.awsv4
}

func (r *ImageResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_image"
}

func (r *ImageResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Create a custom AWS ParallelCluster image in an AWS Region.",

		Attributes: map[string]schema.Attribute{
			"image_id": schema.StringAttribute{
				MarkdownDescription: "Id of the Image that will be built.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the Image that will be built.",
				Computed:            true,
				Optional:            false,
				Required:            false,
			},

			"image_configuration": schema.StringAttribute{
				MarkdownDescription: "The image configuration as a YAML document.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that the image is in.",
				Optional:            true,
				Computed:            true,
			},
			"rollback_on_failure": schema.BoolAttribute{
				MarkdownDescription: "If set to true, image stack rollback occurs if the image fails to create. The default is true.",
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"suppress_validators": schema.ListAttribute{
				MarkdownDescription: "Identify one or more configuration validators to suppress.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"validation_failure_level": schema.StringAttribute{
				MarkdownDescription: "The minimum validation level that causes image create to fail. The default is ERROR.",
				Optional:            true,
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
			"image_build_status": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "Status of the image.",
			},
			"version": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "The AWS ParallelImage version that's used to create the image.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ami_id": schema.StringAttribute{
				Computed:            true,
				Optional:            false,
				Required:            false,
				MarkdownDescription: "EC2 AMI Id",
			},
		},
	}
}

func (r *ImageResource) Configure(
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
	r.role = config.role
	r.expiration = config.expiration
}

func (r *ImageResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data ImageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = data.ImageId

	createImageRequestContent := *openapi.NewBuildImageRequestContent(data.ImageConfiguration.ValueString(), data.ImageId.ValueString())

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	createImageReq := r.client.ImageOperationsAPI.BuildImage(reqCtx)

	if !data.Region.IsNull() {
		createImageReq = createImageReq.Region(data.Region.ValueString())
	}

	if !data.ValidationFailureLevel.IsNull() {
		createImageReq = createImageReq.ValidationFailureLevel(
			openapi.ValidationLevel(data.ValidationFailureLevel.ValueString()),
		)
	}

	suppressValidators := make([]string, 0)
	for _, s := range data.SuppressValidators.Elements() {
		suppressValidators = append(suppressValidators, s.String())
	}

	tflog.Info(ctx, "Building Image")
	httpResp, fullResp, err := createImageReq.
		BuildImageRequestContent(createImageRequestContent).
		SuppressValidators(suppressValidators).
		RollbackOnFailure(data.RollbackOnFailure.ValueBool()).
		Execute()
	if err != nil {
		msg := ""
		if fullResp != nil {
			b, ioErr := io.ReadAll(fullResp.Body)
			if ioErr != nil {
				resp.Diagnostics.AddError(
					"Error when calling `ImageOperationsAPI.CreateImage",
					fmt.Sprintf(
						"Error: %v\nError encounted while handling the above: %v\n",
						err,
						ioErr,
					),
				)
				return
			}
			msg = string(b)
		}
		resp.Diagnostics.AddError(
			"Error when calling `ImageOperationsAPI.CreateImage",
			fmt.Sprintf("Error: %v\nMessage: %v\n", err, msg),
		)
		return
	}
	// response from `CreateImage`: CreateImageResponseContent
	tflog.Trace(ctx, fmt.Sprintf("Response from `ImageOperationsAPI.CreateImage`: %v\n", resp))

	data.ImageId = types.StringValue(httpResp.Image.GetImageId())
	data.CloudformationStackArn = types.StringValue(httpResp.Image.GetCloudformationStackArn())
	data.CloudformationStackStatus = types.StringValue(
		string(httpResp.Image.GetCloudformationStackStatus()),
	)
	data.ImageBuildStatus = types.StringValue(string(httpResp.Image.GetImageBuildStatus()))
	data.Region = types.StringValue(httpResp.Image.GetRegion())
	data.Version = types.StringValue(httpResp.Image.GetVersion())
	data.AmiId = types.StringValue(httpResp.Image.Ec2AmiInfo.GetAmiId())

	imageSummary, err := r.waitImageReady(reqCtx, httpResp.Image.ImageId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Image create failed to complete.",
			fmt.Sprintf("Error: %v", err),
		)
	}

	if imageSummary != nil {
		data.ImageBuildStatus = types.StringValue(string(imageSummary.GetImageBuildStatus()))
		if imageSummary.ImageBuildStatus == openapi.IMAGEBUILDSTATUS_BUILD_FAILED {
			resp.Diagnostics.AddError(
				"Image create failed to complete.",
				fmt.Sprintf("Reason: %v", imageSummary.GetImagebuilderImageStatusReason()),
			)
		}

		if imageSummary.Ec2AmiInfo != nil {
			data.AmiId = types.StringValue(imageSummary.Ec2AmiInfo.GetAmiId())
		}
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ImageResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data ImageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	data.Id = data.ImageId
	imageSummary, err := r.getImage(reqCtx, data.ImageId.ValueString())
	if err != nil && err.Error() != failedToFindImageErr {
		resp.Diagnostics.AddError("Error while retrieving image.", err.Error())
	}

	if id, ok := imageSummary.GetImageIdOk(); ok {
		data.ImageId = types.StringValue(*id)
	}
	if region, ok := imageSummary.GetRegionOk(); ok {
		data.Region = types.StringValue(*region)
	}
	if version, ok := imageSummary.GetVersionOk(); ok {
		data.Version = types.StringValue(*version)
	}
	if stackArn, ok := imageSummary.GetCloudformationStackArnOk(); ok {
		data.CloudformationStackArn = types.StringValue(*stackArn)
	}
	if stackStatus, ok := imageSummary.GetCloudformationStackStatusOk(); ok {
		data.CloudformationStackStatus = types.StringValue(string(*stackStatus))
	}
	if buildStatus, ok := imageSummary.GetImageBuildStatusOk(); ok {
		data.ImageBuildStatus = types.StringValue(string(*buildStatus))
	}
	if ec2AmiInfo, ok := imageSummary.GetEc2AmiInfoOk(); ok {
		data.AmiId = types.StringValue(ec2AmiInfo.GetAmiId())
	}
	if config, ok := imageSummary.GetImageConfigurationOk(); ok {
		if url, ok := config.GetUrlOk(); ok {
			s3Resp, err := http.Get(*url)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}
			defer s3Resp.Body.Close()
			configBytes, err := io.ReadAll(s3Resp.Body)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}

			data.ImageConfiguration = types.StringValue(string(configBytes))
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ImageResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var data ImageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	data.Id = data.ImageId
	imageSummary, err := r.getImage(reqCtx, data.ImageId.ValueString())
	if err != nil && err.Error() != failedToFindImageErr {
		resp.Diagnostics.AddError("Failed to find image.", err.Error())
	}

	if id, ok := imageSummary.GetImageIdOk(); ok {
		data.ImageId = types.StringValue(*id)
	}
	if region, ok := imageSummary.GetRegionOk(); ok {
		data.Region = types.StringValue(*region)
	}

	if version, ok := imageSummary.GetVersionOk(); ok {
		data.Version = types.StringValue(*version)
	}

	if stackArn, ok := imageSummary.GetCloudformationStackArnOk(); ok {
		data.CloudformationStackArn = types.StringValue(*stackArn)
	} else {
		data.CloudformationStackArn = types.StringNull()
	}

	if stackStatus, ok := imageSummary.GetCloudformationStackStatusOk(); ok {
		data.CloudformationStackStatus = types.StringValue(string(*stackStatus))
	} else {
		data.CloudformationStackStatus = types.StringNull()
	}

	if buildStatus, ok := imageSummary.GetImageBuildStatusOk(); ok {
		data.ImageBuildStatus = types.StringValue(string(*buildStatus))
	}

	if ec2AmiInfo, ok := imageSummary.GetEc2AmiInfoOk(); ok {
		data.AmiId = types.StringValue(ec2AmiInfo.GetAmiId())
	}

	if config, ok := imageSummary.GetImageConfigurationOk(); ok {
		if url, ok := config.GetUrlOk(); ok {
			s3Resp, err := http.Get(*url)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}
			defer s3Resp.Body.Close()
			configBytes, err := io.ReadAll(s3Resp.Body)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}

			data.ImageConfiguration = types.StringValue(string(configBytes))
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ImageResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data ImageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)
	deleteImageRequest := r.client.ImageOperationsAPI.DeleteImage(
		reqCtx,
		data.ImageId.ValueString(),
	)

	if !data.Region.IsNull() {
		deleteImageRequest = deleteImageRequest.Region(data.Region.ValueString())
	}

	deleteResponse, _, err := deleteImageRequest.Execute()
	if err != nil && err.Error() != failedToFindImageErr {
		resp.Diagnostics.AddError("Failure occurred while deleting image", err.Error())
	}
	if image, ok := deleteResponse.GetImageOk(); ok {
		buildStatus := image.ImageBuildStatus

		if buildStatus == openapi.IMAGEBUILDSTATUS_DELETE_FAILED {
			resp.Diagnostics.AddError(
				"Image delete failed to complete.",
				fmt.Sprintf("Error %v", err),
			)
		} else if buildStatus == openapi.IMAGEBUILDSTATUS_DELETE_IN_PROGRESS {
			out, err := r.waitImageReady(reqCtx, data.ImageId.ValueString())
			if err != nil && err.Error() != failedToFindImageErr {
				resp.Diagnostics.AddError("Image delete failed to complete.", err.Error())
			}
			if out != nil && out.ImageBuildStatus == openapi.IMAGEBUILDSTATUS_DELETE_FAILED {
				resp.Diagnostics.AddError("Build Status: "+string(out.ImageBuildStatus), "")
			}
		}
	}
}

func (r *ImageResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	var data ImageResourceModel
	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, r.awsv4)

	imageSummary, err := r.getImage(reqCtx, req.ID)
	if err != nil && err.Error() == failedToFindImageErr {
		resp.Diagnostics.AddError("Failed to find image.", err.Error())
	}

	if id, ok := imageSummary.GetImageIdOk(); ok {
		data.ImageId = types.StringValue(*id)
		data.Id = data.ImageId
	}
	if region, ok := imageSummary.GetRegionOk(); ok {
		data.Region = types.StringValue(*region)
	}

	if version, ok := imageSummary.GetVersionOk(); ok {
		data.Version = types.StringValue(*version)
	}

	if stackArn, ok := imageSummary.GetCloudformationStackArnOk(); ok {
		data.CloudformationStackArn = types.StringValue(*stackArn)
	} else {
		data.CloudformationStackArn = types.StringNull()
	}

	if stackStatus, ok := imageSummary.GetCloudformationStackStatusOk(); ok {
		data.CloudformationStackStatus = types.StringValue(string(*stackStatus))
	} else {
		data.CloudformationStackStatus = types.StringNull()
	}

	if buildStatus, ok := imageSummary.GetImageBuildStatusOk(); ok {
		data.ImageBuildStatus = types.StringValue(string(*buildStatus))
	}

	if ec2AmiInfo, ok := imageSummary.GetEc2AmiInfoOk(); ok {
		data.AmiId = types.StringValue(ec2AmiInfo.GetAmiId())
	}

	if config, ok := imageSummary.GetImageConfigurationOk(); ok {
		if url, ok := config.GetUrlOk(); ok {
			s3Resp, err := http.Get(*url)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}
			defer s3Resp.Body.Close()
			configBytes, err := io.ReadAll(s3Resp.Body)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failed to download image configuration from s3.",
					err.Error(),
				)
			}

			data.ImageConfiguration = types.StringValue(string(configBytes))
		}
	}

	data.SuppressValidators = types.ListNull(types.StringType)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	resource.ImportStatePassthroughID(ctx, path.Root("image_id"), req, resp)
}
