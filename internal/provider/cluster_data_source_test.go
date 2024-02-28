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
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewClusterDataSource(t *testing.T) {
	t.Parallel()

	obj := NewClusterDataSource()
	if _, ok := obj.(*ClusterDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ClusterDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitClusterDataSourceMetadata(t *testing.T) {
	t.Parallel()
	d := NewClusterDataSource()
	p := PclusterProvider{}
	resp := datasource.MetadataResponse{}
	req := datasource.MetadataRequest{}
	providerResp := provider.MetadataResponse{}
	providerReq := provider.MetadataRequest{}

	p.Metadata(context.TODO(), providerReq, &providerResp)
	req.ProviderTypeName = providerResp.TypeName

	d.Metadata(context.TODO(), req, &resp)

	if !strings.HasPrefix(resp.TypeName, providerResp.TypeName) {
		t.Fatalf(
			"Error provider typename expected as the prefix for resource or datasource name. \nO: %#v\nE: %#v",
			resp.TypeName,
			providerResp.TypeName,
		)
	}
}

func TestUnitClusterDataSourceSchema(t *testing.T) {
	t.Parallel()
	d := ClusterDataSource{}
	dataSourceModel := ClusterDataSourceModel{}
	resp := datasource.SchemaResponse{}
	req := datasource.SchemaRequest{}

	d.Schema(context.TODO(), req, &resp)

	rDataSource := reflect.TypeOf(dataSourceModel)
	numFields := rDataSource.NumField()
	numAttributes := len(resp.Schema.Attributes)

	for i := 0; i < numFields; i++ {
		tag := rDataSource.Field(i).Tag
		if _, ok := resp.Schema.Attributes[tag.Get("tfsdk")]; !ok {
			t.Fatalf(
				"Error expected attribute missing in schema. O: %#v\nE: %#v",
				resp.Schema.Attributes,
				tag,
			)
		}
	}

	if numAttributes != numFields {
		t.Fatalf(
			"Error extra attributes defined in schema. O: %#v E: %#v",
			numFields,
			numAttributes,
		)
	}
}

func TestUnitClusterDataSourceConfigure(t *testing.T) {
	t.Parallel()
	d := ClusterDataSource{}

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitClusterDataSourceRead(t *testing.T) {
	t.Parallel()

	resp := datasource.ReadResponse{}
	req := datasource.ReadRequest{}
	mResp := datasource.SchemaResponse{}
	mReq := datasource.SchemaRequest{}
	ctx := context.TODO()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	d := ClusterDataSource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	d.Schema(ctx, mReq, &mResp)

	req.Config = tfsdk.Config{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	req.Config = tfsdk.Config{
		Raw: tftypes.NewValue(
			tftypes.Object{},
			map[string]tftypes.Value{
				"cluster_name": tftypes.NewValue(
					tftypes.String,
					"some_name",
				),
				"cluster":     {},
				"log_streams": {},
				"filters":     {},
				"region": tftypes.NewValue(
					tftypes.String,
					"some_region",
				),
				"stack_events": {},
			},
		),
		Schema: mResp.Schema,
	}

	resp = datasource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	mName := "some_name"
	mVersion := "some_version"
	// var mHealth int32 = 1
	// var mUnhealthy int32 = 0
	cluster := openapi.DescribeClusterResponseContent{
		ClusterName: mName,
		ClusterConfiguration: openapi.ClusterConfigurationStructure{
			Url: &server.URL,
		},
		Scheduler: &openapi.Scheduler{
			Type: "some_type",
			Metadata: &openapi.Metadata{
				Name:    &mName,
				Version: &mVersion,
			},
		},
		HeadNode: &openapi.EC2Instance{
			InstanceId:       "some_id",
			InstanceType:     "some_type",
			LaunchTime:       time.Now(),
			PrivateIpAddress: "some_address",
			State:            openapi.INSTANCESTATE_PENDING,
		},
		LoginNodes: &openapi.LoginNodesPool{
			Status:  openapi.LOGINNODESSTATE_ACTIVE,
			Address: &mName,
			Scheme:  &mName,
		},
		Tags:                      []openapi.Tag{{}},
		Failures:                  []openapi.Failure{{}},
		CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
		ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_ENABLED,
	}

	stackEvents := openapi.GetClusterStackEventsResponseContent{
		Events: []openapi.StackEvent{{
			StackId:            "some_stack_id",
			EventId:            "some_event_id",
			StackName:          "some_stack_name",
			LogicalResourceId:  "some_logical_resource_id",
			PhysicalResourceId: "some_physical_resource_id",
			ResourceType:       "some_resource_type",
			Timestamp:          time.Now(),
			ResourceStatus:     openapi.CLOUDFORMATIONRESOURCESTATUS_CREATE_COMPLETE,
		}},
	}

	logStreams := openapi.ListClusterLogStreamsResponseContent{
		LogStreams: []openapi.LogStream{{
			LogStreamName:       "some_name",
			CreationTime:        time.Now(),
			FirstEventTimestamp: time.Now(),
			LastEventTimestamp:  time.Now(),
			LastIngestionTime:   time.Now(),
			UploadSequenceToken: "some_token",
			LogStreamArn:        "some_arn",
		}},
	}
	expectedOutputs := []ClusterDataSourceModel{
		{
			ClusterName: types.StringValue(cluster.ClusterName),
			Region:      types.StringValue("some_region"),
			Filters:     types.ListNull(types.MapType{ElemType: types.StringType}),
			Cluster: types.ObjectValueMust(clusterDescriptionObjectTypes, map[string]attr.Value{
				"clusterName": types.StringValue(cluster.ClusterName),
				"region":      types.StringValue(cluster.Region),
				"version":     types.StringValue(cluster.Version),
				"cloudFormationStackStatus": types.StringValue(
					string(cluster.CloudFormationStackStatus),
				),
				"clusterStatus": types.StringValue(string(cluster.ClusterStatus)),
				"scheduler": types.ObjectValueMust(schedulerObjectTypes, map[string]attr.Value{
					"metadata": types.ObjectValueMust(metadataObjectTypes, map[string]attr.Value{
						"name":    types.StringValue(*cluster.Scheduler.Metadata.Name),
						"version": types.StringValue(*cluster.Scheduler.Metadata.Version),
					}),
					"type": types.StringValue(cluster.Scheduler.Type),
				}),
				"cloudformationStackArn": types.StringValue(cluster.CloudformationStackArn),
				"creationTime":           types.StringValue(cluster.CreationTime.Round(0).String()),
				"lastUpdatedTime": types.StringValue(
					cluster.LastUpdatedTime.Round(0).String(),
				),
				"clusterConfiguration": types.StringValue(""),
				"computeFleetStatus":   types.StringValue(string(cluster.ComputeFleetStatus)),
				"tags": types.ListValueMust(
					types.MapType{ElemType: types.StringType},
					[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{})},
				),
				"headNode": types.MapValueMust(types.StringType, map[string]attr.Value{
					"instanceId":   types.StringValue(cluster.HeadNode.InstanceId),
					"instanceType": types.StringValue(cluster.HeadNode.InstanceType),
					"launchTime": types.StringValue(
						cluster.HeadNode.LaunchTime.Round(0).String(),
					),
					"privateIpAddress": types.StringValue(cluster.HeadNode.PrivateIpAddress),
					"state":            types.StringValue(string(cluster.HeadNode.State)),
				}),
				"failures": types.ListValueMust(
					types.MapType{ElemType: types.StringType},
					[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{})},
				),
				"loginNodes": types.ObjectValueMust(loginNodesObjectTypes, map[string]attr.Value{
					"status":  types.StringValue(string(cluster.LoginNodes.Status)),
					"address": types.StringValue(*cluster.LoginNodes.Address),
					"scheme":  types.StringValue(*cluster.LoginNodes.Scheme),
					"healthyNodes": types.NumberValue(
						big.NewFloat(float64(cluster.LoginNodes.GetHealthyNodes())),
					),
					"unhealthyNodes": types.NumberValue(
						big.NewFloat(float64(cluster.LoginNodes.GetUnhealthyNodes())),
					),
				}),
			}),
			StackEvents: types.ListValueMust(
				types.MapType{ElemType: types.StringType},
				[]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{
						"stackId":   types.StringValue(stackEvents.Events[0].GetStackId()),
						"eventId":   types.StringValue(stackEvents.Events[0].GetEventId()),
						"stackName": types.StringValue(stackEvents.Events[0].GetStackName()),
						"logicalResourceId": types.StringValue(
							stackEvents.Events[0].GetLogicalResourceId(),
						),
						"physicalResourceId": types.StringValue(
							stackEvents.Events[0].GetPhysicalResourceId(),
						),
						"resourceType": types.StringValue(
							stackEvents.Events[0].GetResourceType(),
						),
						"timestamp": types.StringValue(
							stackEvents.Events[0].GetTimestamp().Round(0).String(),
						),
						"resourceStatus": types.StringValue(
							string(stackEvents.Events[0].GetResourceStatus()),
						),
						"resourceStatusReason": types.StringValue(
							stackEvents.Events[0].GetResourceStatusReason(),
						),
						"resourceProperties": types.StringValue(
							stackEvents.Events[0].GetResourceProperties(),
						),
						"clientRequestToken": types.StringValue(
							stackEvents.Events[0].GetClientRequestToken(),
						),
					}),
				},
			),
			LogStreams: types.ListValueMust(
				types.MapType{ElemType: types.StringType},
				[]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{
						"logStreamName": types.StringValue(logStreams.LogStreams[0].LogStreamName),
						"creationTime": types.StringValue(
							logStreams.LogStreams[0].GetCreationTime().Round(0).String(),
						),
						"firstEventTimestamp": types.StringValue(
							logStreams.LogStreams[0].GetFirstEventTimestamp().Round(0).String(),
						),
						"lastEventTimestamp": types.StringValue(
							logStreams.LogStreams[0].GetLastEventTimestamp().Round(0).String(),
						),
						"lastIngestionTime": types.StringValue(
							logStreams.LogStreams[0].GetLastIngestionTime().Round(0).String(),
						),
						"logStreamArn": types.StringValue(logStreams.LogStreams[0].LogStreamArn),
					}),
				},
			),
		},
	}

	mocks := []mockCfg{
		{
			path: "clusters/some_name",
			out:  cluster,
		},
		{
			path: "clusters/some_name/stackevents",
			out:  stackEvents,
		},
		{
			path: "clusters/some_name/logstreams",
			out:  logStreams,
		},
	}
	server, err := mockJsonServer(mocks...)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}
	d.client = openapi.NewAPIClient(cfg)

	resp = datasource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics)
	}

	output := ClusterDataSourceModel{}
	diags := resp.State.Get(ctx, &output)
	if diags.HasError() {
		t.Fatalf("Read operation returned unexpected error while retrieving state data: %v", diags)
	}

	if !reflect.DeepEqual(expectedOutputs[0], output) {
		t.Fatalf(
			"Expected output did not match actual output: \nE: %v\nO: %v\n",
			expectedOutputs[0],
			output,
		)
	}
}
