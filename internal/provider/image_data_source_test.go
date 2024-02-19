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

func TestUnitNewImageDataSource(t *testing.T) {
	t.Parallel()

	obj := NewImageDataSource()
	if _, ok := obj.(*ImageDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ImageDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitImageDataSourceMetadata(t *testing.T) {
	d := NewImageDataSource()
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

func TestUnitImageDataSourceSchema(t *testing.T) {
	d := ImageDataSource{}
	dataSourceModel := ImageDataSourceModel{}
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

func TestUnitImageDataSourceConfigure(t *testing.T) {
	d := ImageDataSource{}

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitImageDataSourceRead(t *testing.T) {
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

	d := ImageDataSource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	d.Schema(ctx, mReq, &mResp)

	req.Config = tfsdk.Config{
		Schema: mResp.Schema,
	}

	d.Read(context.TODO(), req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	req.Config = tfsdk.Config{
		Raw: tftypes.NewValue(
			tftypes.Object{},
			map[string]tftypes.Value{
				"image": {},
				"image_id": tftypes.NewValue(
					tftypes.String,
					"some_id",
				),
				"log_streams":  {},
				"stack_events": {},
				"region":       {},
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

	someArn := "some_arn"
	someReason := "some_reason"
	someName := "some_name"
	someArchitecture := "some_architecture"
	someDescription := "some_description"
	someProp := "some_prop"
	someToken := "some_token"
	someUrl := "some_url"
	someTime := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	someKey := "some_key"
	someValue := "some_value"

	image := openapi.DescribeImageResponseContent{
		ImageId:                         "some_id",
		Region:                          "some_region",
		Version:                         "some_version",
		ImageBuildStatus:                openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
		ImageBuildLogsArn:               &someArn,
		CloudformationStackStatus:       openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE.Ptr(),
		CloudformationStackStatusReason: &someReason,
		CloudformationStackArn:          &someArn,
		CreationTime:                    &someTime,
		CloudformationStackCreationTime: &someTime,
		CloudformationStackTags: []openapi.Tag{
			{
				Key:   &someKey,
				Value: &someValue,
			},
		},
		ImageConfiguration:            openapi.ImageConfigurationStructure{},
		ImagebuilderImageStatus:       openapi.IMAGEBUILDERIMAGESTATUS_AVAILABLE.Ptr(),
		ImagebuilderImageStatusReason: &someReason,
		Ec2AmiInfo: &openapi.Ec2AmiInfo{
			AmiId: "some_id",
			Tags: []openapi.Tag{{
				Key:   &someKey,
				Value: &someValue,
			}},
			AmiName:      &someName,
			Architecture: &someArchitecture,
			State:        openapi.EC2AMISTATE_AVAILABLE.Ptr(),
			Description:  &someDescription,
		},
	}

	logStreams := openapi.ListImageLogStreamsResponseContent{
		LogStreams: []openapi.LogStream{{
			LogStreamName:       "some_log_name",
			CreationTime:        someTime,
			FirstEventTimestamp: someTime,
			LastEventTimestamp:  someTime,
			LastIngestionTime:   someTime,
			LogStreamArn:        "some_log_arn",
		}},
	}

	stackEvents := openapi.GetImageStackEventsResponseContent{
		Events: []openapi.StackEvent{{
			StackId:              "some_id",
			EventId:              "some_id",
			StackName:            "some_name",
			LogicalResourceId:    "some_resource_id",
			PhysicalResourceId:   "some_phy_resource_id",
			ResourceType:         "some_type",
			Timestamp:            someTime,
			ResourceStatus:       openapi.CLOUDFORMATIONRESOURCESTATUS_CREATE_COMPLETE,
			ResourceStatusReason: &someReason,
			ResourceProperties:   &someProp,
			ClientRequestToken:   &someToken,
		}},
	}

	pathsToTest := []string{
		"images/custom/some_id",
		"images/custom/some_id/logstreams",
		"images/custom/some_id/stackevents",
	}

	someConfig := "some_config"
	s3Server, err := mockJsonWithTextServer([]string{someUrl}, someConfig)
	if err != nil {
		t.Fatal(err)
	}
	url := s3Server.URL + "/" + someUrl
	image.ImageConfiguration.Url = &url

	server, err = mockJsonServer(
		pathsToTest,
		image,
		logStreams,
		stackEvents,
	)
	if err != nil {
		t.Fatal(err)
	}

	cfg = openapi.NewConfiguration()
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

	expectedOutput := ImageDataSourceModel{
		ImageId: types.StringValue(image.ImageId),
		Image: types.ObjectValueMust(imageDescriptionObjectTypes, map[string]attr.Value{
			"imageId":           types.StringValue(image.ImageId),
			"region":            types.StringValue(image.Region),
			"version":           types.StringValue(image.Version),
			"imageBuildStatus":  types.StringValue(string(image.ImageBuildStatus)),
			"imageBuildLogsArn": types.StringValue(*image.ImageBuildLogsArn),
			"cloudformationStackStatus": types.StringValue(
				string(*image.CloudformationStackStatus),
			),
			"cloudformationStackStatusReason": types.StringValue(
				*image.CloudformationStackStatusReason,
			),
			"cloudformationStackArn": types.StringValue(*image.CloudformationStackArn),
			"creationTime":           types.StringValue(image.GetCreationTime().String()),
			"cloudformationStackCreationTime": types.StringValue(
				image.GetCloudformationStackCreationTime().String(),
			),
			"cloudformationStackTags": types.ListValueMust(
				types.MapType{ElemType: types.StringType},
				[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{
					"Key":   types.StringValue(*image.CloudformationStackTags[0].Key),
					"Value": types.StringValue(*image.CloudformationStackTags[0].Value),
				})},
			),
			"imageConfiguration":      types.StringValue(someConfig),
			"imagebuilderImageStatus": types.StringValue(string(*image.ImagebuilderImageStatus)),
			"imagebuilderImageStatusReason": types.StringValue(
				*image.ImagebuilderImageStatusReason,
			),
			"ec2AmiInfo": types.ObjectValueMust(
				ec2AmiInfoObjectTypes,
				map[string]attr.Value{
					"amiId": types.StringValue(image.Ec2AmiInfo.AmiId),
					"tags": types.ListValueMust(
						types.MapType{ElemType: types.StringType},
						[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{
							"Key":   types.StringValue(*image.Ec2AmiInfo.Tags[0].Key),
							"Value": types.StringValue(*image.Ec2AmiInfo.Tags[0].Value),
						})},
					),
					"amiName":      types.StringValue(*image.Ec2AmiInfo.AmiName),
					"architecture": types.StringValue(*image.Ec2AmiInfo.Architecture),
					"state":        types.StringValue(string(*image.Ec2AmiInfo.State)),
					"description":  types.StringValue(*image.Ec2AmiInfo.Description),
				},
			),
		}),
		Region: types.StringNull(),
		LogStreams: types.ListValueMust(
			types.MapType{ElemType: types.StringType},
			[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{
				"logStreamName": types.StringValue(logStreams.LogStreams[0].GetLogStreamName()),
				"creationTime": types.StringValue(
					logStreams.LogStreams[0].GetCreationTime().String(),
				),
				"firstEventTimestamp": types.StringValue(
					logStreams.LogStreams[0].GetFirstEventTimestamp().String(),
				),
				"lastEventTimestamp": types.StringValue(
					logStreams.LogStreams[0].GetLastEventTimestamp().String(),
				),
				"lastIngestionTime": types.StringValue(
					logStreams.LogStreams[0].GetLastIngestionTime().String(),
				),
				"uploadSequenceToken": types.StringNull(),
				"logStreamArn": types.StringValue(
					logStreams.LogStreams[0].GetLogStreamArn(),
				),
			})},
		),
		StackEvents: types.ListValueMust(
			types.MapType{ElemType: types.StringType},
			[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{
				"stackId":   types.StringValue(stackEvents.Events[0].GetStackId()),
				"eventId":   types.StringValue(stackEvents.Events[0].GetEventId()),
				"stackName": types.StringValue(stackEvents.Events[0].GetStackName()),
				"logicalResourceId": types.StringValue(
					stackEvents.Events[0].GetLogicalResourceId(),
				),
				"physicalResourceId": types.StringValue(
					stackEvents.Events[0].GetPhysicalResourceId(),
				),
				"resourceType": types.StringValue(stackEvents.Events[0].GetResourceType()),
				"timestamp": types.StringValue(
					stackEvents.Events[0].GetTimestamp().String(),
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
			})},
		),
	}

	var output ImageDataSourceModel
	resp.State.Get(ctx, &output)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics.Errors())
	}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match output from read operation. \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
}
