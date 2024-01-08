// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_                     datasource.DataSource = &ImageDataSource{}
	ec2AmiInfoObjectTypes                       = map[string]attr.Type{
		"amiId":        types.StringType,
		"tags":         types.ListType{ElemType: types.MapType{ElemType: types.StringType}},
		"amiName":      types.StringType,
		"architecture": types.StringType,
		"state":        types.StringType,
		"description":  types.StringType,
	}
)

var imageDescriptionObjectTypes = map[string]attr.Type{
	"imageId":                         types.StringType,
	"region":                          types.StringType,
	"version":                         types.StringType,
	"imageBuildStatus":                types.StringType,
	"imageBuildLogsArn":               types.StringType,
	"cloudformationStackStatus":       types.StringType,
	"cloudformationStackStatusReason": types.StringType,
	"cloudformationStackArn":          types.StringType,
	"creationTime":                    types.StringType,
	"cloudformationStackCreationTime": types.StringType,
	"cloudformationStackTags": types.ListType{
		ElemType: types.MapType{ElemType: types.StringType},
	},
	"imageConfiguration":            types.StringType,
	"imagebuilderImageStatus":       types.StringType,
	"imagebuilderImageStatusReason": types.StringType,
	"ec2AmiInfo":                    types.ObjectType{AttrTypes: ec2AmiInfoObjectTypes},
}

func NewImageDataSource() datasource.DataSource {
	return &ImageDataSource{}
}

// ImageDataSource defines the data source implementation.
type ImageDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ImageDataSourceModel describes the data source data model.
type ImageDataSourceModel struct {
	ImageId     types.String `tfsdk:"image_id"`
	Image       types.Object `tfsdk:"image"`
	Region      types.String `tfsdk:"region"`
	LogStreams  types.List   `tfsdk:"log_streams"`
	StackEvents types.List   `tfsdk:"stack_events"`
}

func (d *ImageDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_image"
}

func (d *ImageDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Get detailed information about an existing image.",

		Attributes: map[string]schema.Attribute{
			"image": schema.ObjectAttribute{
				MarkdownDescription: "Detailed information about an existing image.",
				Optional:            false,
				Required:            false,
				Computed:            true,
				AttributeTypes:      imageDescriptionObjectTypes,
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: "Image identifier",
				Required:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region of the images.",
				Optional:            true,
			},
			"log_streams": schema.ListAttribute{
				MarkdownDescription: "List of log streams that's associated with an image.",
				Required:            false,
				Optional:            false,
				Computed:            true,
				ElementType:         types.MapType{ElemType: types.StringType},
			},
			"stack_events": schema.ListAttribute{
				MarkdownDescription: "Events that are associated with the stack for an image build.",
				Required:            false,
				Optional:            false,
				Computed:            true,
				ElementType:         types.MapType{ElemType: types.StringType},
			},
		},
	}
}

func (d *ImageDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(configData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf(
				"Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)

		return
	}

	d.client = client.client
	d.awsv4 = client.awsv4
}

func (d *ImageDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ImageDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(ctx, openapi.ContextAWSv4, d.awsv4)

	imageReq := d.client.ImageOperationsAPI.DescribeImage(reqCtx, data.ImageId.ValueString())

	if !data.Region.IsNull() {
		imageReq = imageReq.Region(data.Region.ValueString())
	}

	image, rawHttp, err := imageReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("%v", err), fmt.Sprintf("%v", rawHttp.Body))
	}

	// Retrieve image configuration
	s3Resp, err := http.Get(*image.GetImageConfiguration().Url)
	if err != nil {
		resp.Diagnostics.AddError("Failed to download image configuration from s3.", err.Error())
	}
	defer s3Resp.Body.Close()
	configBytes, err := io.ReadAll(s3Resp.Body)
	if err != nil {
		resp.Diagnostics.AddError("Failed to download image configuration from s3.", err.Error())
	}

	imageConfiguration := types.StringValue(string(configBytes))

	// Populate cloudformation stack tags
	var cloudformationStackTags types.List
	if tags, ok := image.GetCloudformationStackTagsOk(); ok {
		tagList := make([]attr.Value, 0)
		for _, tag := range tags {
			tagMap, diags := types.MapValue(
				types.StringType,
				map[string]attr.Value{
					"Key":   types.StringValue(tag.GetKey()),
					"Value": types.StringValue(tag.GetValue()),
				},
			)
			resp.Diagnostics.Append(diags...)
			tagList = append(tagList, tagMap)
		}
		cfStackTags, diags := types.ListValue(
			types.MapType{ElemType: types.StringType},
			tagList,
		)
		resp.Diagnostics.Append(diags...)
		cloudformationStackTags = cfStackTags
	} else {
		cloudformationStackTags = types.ListNull(types.MapType{ElemType: types.StringType})
	}

	// Populate ec2AmiInfo
	var ec2AmiInfo types.Object
	if info, ok := image.GetEc2AmiInfoOk(); ok {
		var tagsList types.List
		if tags, ok := info.GetTagsOk(); ok {
			tagList := make([]attr.Value, 0)
			for _, tag := range tags {
				tagMap, diags := types.MapValue(
					types.StringType,
					map[string]attr.Value{
						"Key":   types.StringValue(tag.GetKey()),
						"Value": types.StringValue(tag.GetValue()),
					},
				)
				resp.Diagnostics.Append(diags...)
				tagList = append(tagList, tagMap)
			}
			tl, diags := types.ListValue(
				types.MapType{ElemType: types.StringType},
				tagList,
			)
			resp.Diagnostics.Append(diags...)
			tagsList = tl
		} else {
			tagsList = types.ListNull(types.MapType{ElemType: types.StringType})
		}

		ec2AmiInfoMap := map[string]attr.Value{
			"amiId":        types.StringValue(info.GetAmiId()),
			"tags":         tagsList,
			"amiName":      types.StringValue(info.GetAmiName()),
			"architecture": types.StringValue(info.GetArchitecture()),
			"state":        types.StringValue(string(info.GetState())),
			"description":  types.StringValue(info.GetDescription()),
		}
		t, diags := types.ObjectValue(ec2AmiInfoObjectTypes, ec2AmiInfoMap)
		resp.Diagnostics.Append(diags...)
		ec2AmiInfo = t
	} else {
		ec2AmiInfo = types.ObjectNull(ec2AmiInfoObjectTypes)
	}

	// Populate Image Description
	imageDescriptionMap := map[string]attr.Value{
		"imageId":           types.StringValue(image.GetImageId()),
		"region":            types.StringValue(image.GetRegion()),
		"version":           types.StringValue(image.GetVersion()),
		"imageBuildStatus":  types.StringValue(string(image.GetImageBuildStatus())),
		"imageBuildLogsArn": types.StringValue(image.GetImageBuildLogsArn()),
		"cloudformationStackStatus": types.StringValue(
			string(image.GetCloudformationStackStatus()),
		),
		"cloudformationStackStatusReason": types.StringValue(
			image.GetCloudformationStackStatusReason(),
		),
		"cloudformationStackArn": types.StringValue(image.GetCloudformationStackArn()),
		"creationTime":           types.StringValue(image.GetCreationTime().String()),
		"cloudformationStackCreationTime": types.StringValue(
			image.GetCloudformationStackCreationTime().String(),
		),
		"cloudformationStackTags": cloudformationStackTags,
		"imageConfiguration":      imageConfiguration,
		"imagebuilderImageStatus": types.StringValue(
			string(image.GetImagebuilderImageStatus()),
		),
		"imagebuilderImageStatusReason": types.StringValue(
			image.GetImagebuilderImageStatusReason(),
		),
		"ec2AmiInfo": ec2AmiInfo,
	}
	imageDescription, diags := types.ObjectValue(imageDescriptionObjectTypes, imageDescriptionMap)
	resp.Diagnostics.Append(diags...)
	data.Image = imageDescription

	// Populate Image Log Streams
	imageLogStreamReq := d.client.ImageLogsAPI.ListImageLogStreams(
		reqCtx,
		data.ImageId.ValueString(),
	)

	if !data.Region.IsNull() {
		imageLogStreamReq = imageLogStreamReq.Region(data.Region.ValueString())
	}

	logStreams, rawHttp, err := imageLogStreamReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("%v", err), fmt.Sprintf("%v", rawHttp.Body))
	}

	if logStreams != nil {
		logStreamList := make([]attr.Value, 0)
		for _, logStream := range logStreams.GetLogStreams() {
			logStreamMap := map[string]attr.Value{
				"logStreamName": types.StringValue(logStream.GetLogStreamName()),
				"creationTime":  types.StringValue(logStream.GetCreationTime().String()),
				"firstEventTimestamp": types.StringValue(
					logStream.GetFirstEventTimestamp().String(),
				),
				"lastEventTimestamp": types.StringValue(
					logStream.GetLastEventTimestamp().String(),
				),
				"lastIngestionTime":   types.StringValue(logStream.GetLastIngestionTime().String()),
				"uploadSequenceToken": types.StringNull(),
				"logStreamArn":        types.StringValue(logStream.GetLogStreamArn()),
			}
			tfLogStreamMap, diags := types.MapValue(types.StringType, logStreamMap)
			resp.Diagnostics.Append(diags...)
			logStreamList = append(logStreamList, tfLogStreamMap)
		}
		tfLogStreamList, diags := types.ListValue(
			types.MapType{ElemType: types.StringType},
			logStreamList,
		)
		resp.Diagnostics.Append(diags...)
		data.LogStreams = tfLogStreamList
	} else {
		data.LogStreams = types.ListNull(types.MapType{ElemType: types.StringType})
	}

	// Populate Image Log Events
	imageStackEventsReq := d.client.ImageLogsAPI.GetImageStackEvents(
		reqCtx,
		data.ImageId.ValueString(),
	)

	if !data.Region.IsNull() {
		imageStackEventsReq = imageStackEventsReq.Region(data.Region.ValueString())
	}

	stackEvents, _, _ := imageStackEventsReq.Execute()

	if stackEvents != nil {
		stackEventsList := make([]attr.Value, 0)
		for _, stackEvent := range stackEvents.GetEvents() {
			stackEventMap := map[string]attr.Value{
				"stackId":              types.StringValue(stackEvent.GetStackId()),
				"eventId":              types.StringValue(stackEvent.GetEventId()),
				"stackName":            types.StringValue(stackEvent.GetStackName()),
				"logicalResourceId":    types.StringValue(stackEvent.GetLogicalResourceId()),
				"physicalResourceId":   types.StringValue(stackEvent.GetPhysicalResourceId()),
				"resourceType":         types.StringValue(stackEvent.GetResourceType()),
				"timestamp":            types.StringValue(stackEvent.GetTimestamp().String()),
				"resourceStatus":       types.StringValue(string(stackEvent.GetResourceStatus())),
				"resourceStatusReason": types.StringValue(stackEvent.GetResourceStatusReason()),
				"resourceProperties":   types.StringValue(stackEvent.GetResourceProperties()),
				"clientRequestToken":   types.StringValue(stackEvent.GetClientRequestToken()),
			}
			tfStackEventMap, diags := types.MapValue(types.StringType, stackEventMap)
			resp.Diagnostics.Append(diags...)
			stackEventsList = append(stackEventsList, tfStackEventMap)
		}
		tfStackEventsList, diags := types.ListValue(
			types.MapType{ElemType: types.StringType},
			stackEventsList,
		)
		resp.Diagnostics.Append(diags...)
		data.StackEvents = tfStackEventsList
	} else {
		data.StackEvents = types.ListNull(types.MapType{ElemType: types.StringType})
	}

	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
