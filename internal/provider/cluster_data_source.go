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
	_                     datasource.DataSource = &ClusterDataSource{}
	loginNodesObjectTypes                       = map[string]attr.Type{
		"status":         types.StringType,
		"address":        types.StringType,
		"scheme":         types.StringType,
		"healthyNodes":   types.StringType,
		"unhealthyNodes": types.StringType,
	}
	clusterDescriptionObjectTypes = map[string]attr.Type{
		"clusterName":               types.StringType,
		"region":                    types.StringType,
		"version":                   types.StringType,
		"cloudFormationStackStatus": types.StringType,
		"clusterStatus":             types.StringType,
		"scheduler":                 types.ObjectType{AttrTypes: schedulerObjectTypes},
		"cloudformationStackArn":    types.StringType,
		"creationTime":              types.StringType,
		"lastUpdatedTime":           types.StringType,
		"clusterConfiguration":      types.StringType,
		"computeFleetStatus":        types.StringType,
		"tags": types.ListType{
			ElemType: types.MapType{ElemType: types.StringType},
		},
		"headNode":   types.MapType{ElemType: types.StringType},
		"failures":   types.ListType{ElemType: types.MapType{ElemType: types.StringType}},
		"loginNodes": types.ObjectType{AttrTypes: loginNodesObjectTypes},
	}
)

func NewClusterDataSource() datasource.DataSource {
	return &ClusterDataSource{}
}

// ClusterDataSource defines the data source implementation.
type ClusterDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ClusterDataSourceModel describes the data source data model.
type ClusterDataSourceModel struct {
	ClusterName types.String `tfsdk:"cluster_name"`
	Cluster     types.Object `tfsdk:"cluster"`
	LogStreams  types.List   `tfsdk:"log_streams"`
	Filters     types.List   `tfsdk:"filters"`
	Region      types.String `tfsdk:"region"`
	StackEvents types.List   `tfsdk:"stack_events"`
}

func (d *ClusterDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (d *ClusterDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Get detailed information about an existing cluster.",

		Attributes: map[string]schema.Attribute{
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "The name of the cluster.",
				Optional:            false,
				Required:            true,
			},
			"cluster": schema.ObjectAttribute{
				MarkdownDescription: "the cluster.",
				Optional:            false,
				Required:            false,
				Computed:            true,
				AttributeTypes:      clusterDescriptionObjectTypes,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that the cluster is in.",
				Optional:            true,
			},
			"log_streams": schema.ListAttribute{
				MarkdownDescription: "List of logstreams.",
				Optional:            false,
				Required:            false,
				Computed:            true,
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
			"filters": schema.ListAttribute{
				MarkdownDescription: "Filter the log streams.",
				Optional:            true,
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
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

func (d *ClusterDataSource) Configure(
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

func populateScheduler(
	scheduler *openapi.Scheduler,
	data *ClusterDataSourceModel,
	resp *datasource.ReadResponse,
) {
}

func populateClusterDataSource(
	cluster *openapi.DescribeClusterResponseContent,
	data *ClusterDataSourceModel,
	resp *datasource.ReadResponse,
) {
	ctx := context.Background()
	clusterMap, err := cluster.ToMap()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure retrieving cluster description.",
			fmt.Sprintf("%v", err),
		)
	}
	delete(clusterMap, "creationTime")
	delete(clusterMap, "lastUpdatedTime")
	delete(clusterMap, "scheduler")
	delete(clusterMap, "headNode")
	delete(clusterMap, "clusterConfiguration")
	delete(clusterMap, "tags")
	delete(clusterMap, "failures")
	delete(clusterMap, "loginNodes")

	tfClusterMap, diags := types.MapValueFrom(ctx, types.StringType, clusterMap)
	tfClusterMapElements := tfClusterMap.Elements()

	if creationTime, ok := cluster.GetCreationTimeOk(); ok {
		tfClusterMapElements["creationTime"] = types.StringValue(creationTime.String())
	} else {
		tfClusterMapElements["creationTime"] = types.StringNull()
	}
	if lastUpdatedTime, ok := cluster.GetLastUpdatedTimeOk(); ok {
		tfClusterMapElements["lastUpdatedTime"] = types.StringValue(lastUpdatedTime.String())
	} else {
		tfClusterMapElements["lastUpdatedTime"] = types.StringNull()
	}

	// Populate HeadNode
	if headNode, ok := cluster.GetHeadNodeOk(); ok {
		headNodeMap, err := headNode.ToMap()
		if err != nil {
		}

		headNodeMap["launchTime"] = headNode.GetLaunchTime().String()
		headNodeMap["state"] = string(headNode.GetState())

		tfHeadNodeMap, diags := types.MapValueFrom(ctx, types.StringType, headNodeMap)
		resp.Diagnostics.Append(diags...)

		tfClusterMapElements["headNode"] = tfHeadNodeMap
	} else {
		tfClusterMapElements["headNode"] = types.MapNull(types.StringType)
	}

	// Populate Tags
	if cluster.Tags != nil {
		tagList := make([]types.Map, 0)
		for _, tag := range cluster.Tags {
			tagMap, err := tag.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving cluster list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			tfTagMap, diags := types.MapValueFrom(ctx, types.StringType, tagMap)
			resp.Diagnostics.Append(diags...)
			tagList = append(tagList, tfTagMap)
		}
		tfClusterMapElements["tags"], diags = types.ListValueFrom(
			ctx,
			types.MapType{ElemType: types.StringType},
			tagList,
		)
		resp.Diagnostics.Append(diags...)
	} else {
		tfClusterMapElements["tags"] = types.ListNull(types.MapType{ElemType: types.StringType})
	}

	// Populate Failures
	if failures, ok := cluster.GetFailuresOk(); ok {
		failureList := make([]types.Map, 0)
		for _, failure := range failures {
			failureMap, err := failure.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving cluster list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			tfTagMap, diags := types.MapValueFrom(ctx, types.StringType, failureMap)
			resp.Diagnostics.Append(diags...)
			failureList = append(failureList, tfTagMap)
		}
		tfClusterMapElements["failures"], diags = types.ListValueFrom(
			ctx,
			types.MapType{ElemType: types.StringType},
			failureList,
		)
		resp.Diagnostics.Append(diags...)

	} else {
		tfClusterMapElements["failures"] = types.ListNull(types.MapType{ElemType: types.StringType})
	}

	// Populate ClusterConfig
	if clusterConfig, ok := cluster.GetClusterConfigurationOk(); ok {
		var configBytes []byte
		s3Resp, err := http.Get(clusterConfig.GetUrl())
		if err != nil {
			resp.Diagnostics.AddWarning(
				"Failed to download cluster configuration from s3.",
				err.Error(),
			)
		} else {
			defer s3Resp.Body.Close()
			configBytes, err = io.ReadAll(s3Resp.Body)
			if err != nil {
				resp.Diagnostics.AddWarning(
					"Failed to download cluster configuration from s3.",
					err.Error(),
				)
			}
		}

		tfClusterMapElements["clusterConfiguration"] = types.StringValue(string(configBytes))
	} else {
		tfClusterMapElements["clusterConfiguration"] = types.StringNull()
	}

	// Populate Scheduler
	var tfMetadataObject types.Object
	var tfSchedulerObject types.Object
	if scheduler, ok := cluster.GetSchedulerOk(); ok {
		schedulerMap, err := scheduler.ToMap()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error occured while retrieving cluster list",
				fmt.Sprintf("Error: %v", err),
			)
		}
		delete(schedulerMap, "metadata")

		tfSchedulerMap, diags := types.MapValueFrom(ctx, types.StringType, schedulerMap)
		resp.Diagnostics.Append(diags...)
		tfSchedulerMapElements := tfSchedulerMap.Elements()

		if metadata, ok := cluster.Scheduler.GetMetadataOk(); ok {
			metadataMap, err := metadata.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving cluster list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			tfMetadataMap, diags := types.MapValueFrom(ctx, types.StringType, metadataMap)
			resp.Diagnostics.Append(diags...)
			tfMetadataObject, diags = types.ObjectValue(
				metadataObjectTypes,
				tfMetadataMap.Elements(),
			)
			resp.Diagnostics.Append(diags...)

		} else {
			tfMetadataObject = types.ObjectNull(metadataObjectTypes)
		}
		tfSchedulerMapElements["metadata"] = tfMetadataObject
		tfSchedulerObject, diags = types.ObjectValue(
			schedulerObjectTypes,
			tfSchedulerMapElements,
		)
		resp.Diagnostics.Append(diags...)

	} else {
		tfSchedulerObject = types.ObjectNull(schedulerObjectTypes)
	}

	tfClusterMapElements["scheduler"] = tfSchedulerObject

	// Populate LoginNodes
	if loginNodes, ok := cluster.GetLoginNodesOk(); ok {
		loginNodesMap, err := loginNodes.ToMap()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error occured while retrieving cluster list",
				fmt.Sprintf("Error: %v", err),
			)
		}
		loginNodesObject, diags := types.ObjectValueFrom(
			ctx,
			loginNodesObjectTypes,
			loginNodesMap,
		)
		resp.Diagnostics.Append(diags...)
		tfClusterMapElements["loginNodes"] = loginNodesObject
	} else {
		tfClusterMapElements["loginNodes"] = types.ObjectNull(loginNodesObjectTypes)
	}
	clusterObject, diags := types.ObjectValue(
		clusterDescriptionObjectTypes,
		tfClusterMapElements,
	)
	resp.Diagnostics.Append(diags...)
	data.Cluster = clusterObject
}

func (d *ClusterDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ClusterDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, d.awsv4)

	cluster, rawHttp, err := d.client.ClusterOperationsAPI.DescribeCluster(
		reqCtx,
		data.ClusterName.ValueString(),
	).Execute()
	if err != nil {
		if rawHttp != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("%v", err.Error()),
				fmt.Sprintf("%v", rawHttp.Body),
			)
		} else {
			resp.Diagnostics.AddError(
				"Error while describing cluster.",
				fmt.Sprintf("%v", err.Error()),
			)
		}
		return
	}

	// Populate Cluster
	if cluster != nil {
		populateClusterDataSource(cluster, &data, resp)
	}

	logStreams, rawHttp, err := d.client.ClusterLogsAPI.ListClusterLogStreams(
		reqCtx,
		data.ClusterName.ValueString(),
	).Execute()
	if err != nil {
		if rawHttp != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("%v", err.Error()),
				fmt.Sprintf("%v", rawHttp.Body),
			)
		} else {
			resp.Diagnostics.AddError(
				"Error while listing cluster log streams.",
				fmt.Sprintf("%v", err.Error()),
			)
		}
	}

	data.LogStreams = types.ListNull(types.StringType)
	// Populate ClusterLogStreams
	if logStreams != nil {
		var logStreamsList []types.Map
		for _, logStream := range logStreams.GetLogStreams() {
			logStreamMap, err := logStream.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure retrieving cluster log streams",
					fmt.Sprintf("%v", err),
				)
			}
			delete(logStreamMap, "uploadSequenceToken")

			logStreamMap["creationTime"] = logStream.GetCreationTime().String()
			logStreamMap["firstEventTimestamp"] = logStream.GetFirstEventTimestamp().String()
			logStreamMap["lastEventTimestamp"] = logStream.GetLastEventTimestamp().String()
			logStreamMap["lastIngestionTime"] = logStream.GetLastIngestionTime().String()

			tfLogStreamMap, diags := types.MapValueFrom(ctx, types.StringType, logStreamMap)
			resp.Diagnostics.Append(diags...)

			logStreamsList = append(logStreamsList, tfLogStreamMap)
		}

		tfLogStreamsList, diags := types.ListValueFrom(
			ctx,
			types.MapType{ElemType: types.StringType},
			logStreamsList,
		)
		resp.Diagnostics.Append(diags...)
		data.LogStreams = tfLogStreamsList
	} else {
		data.LogStreams = types.ListNull(types.StringType)
	}

	// Populate Image Log Events
	imageStackEventsReq := d.client.ClusterLogsAPI.GetClusterStackEvents(
		reqCtx,
		data.ClusterName.ValueString(),
	)

	if !data.Region.IsNull() {
		imageStackEventsReq = imageStackEventsReq.Region(data.Region.ValueString())
	}

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
