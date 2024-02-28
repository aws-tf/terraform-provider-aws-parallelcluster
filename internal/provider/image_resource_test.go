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
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewImageResource(t *testing.T) {
	t.Parallel()

	obj := NewImageResource()
	if _, ok := obj.(*ImageResource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ImageResource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitImageResourceMetadata(t *testing.T) {
	r := NewImageResource()
	p := PclusterProvider{}
	resp := resource.MetadataResponse{}
	req := resource.MetadataRequest{}
	providerResp := provider.MetadataResponse{}
	providerReq := provider.MetadataRequest{}

	p.Metadata(context.TODO(), providerReq, &providerResp)
	req.ProviderTypeName = providerResp.TypeName

	r.Metadata(context.TODO(), req, &resp)

	if !strings.HasPrefix(resp.TypeName, providerResp.TypeName) {
		t.Fatalf(
			"Error provider typename expected as the prefix for resource or resource name. \nO: %#v\nE: %#v",
			resp.TypeName,
			providerResp.TypeName,
		)
	}
}

func TestUnitImageResourceSchema(t *testing.T) {
	r := ImageResource{}
	dataSourceModel := ImageResourceModel{}
	resp := resource.SchemaResponse{}
	req := resource.SchemaRequest{}

	r.Schema(context.TODO(), req, &resp)

	rResource := reflect.TypeOf(dataSourceModel)
	numFields := rResource.NumField()
	numAttributes := len(resp.Schema.Attributes)

	for i := 0; i < numFields; i++ {
		tag := rResource.Field(i).Tag
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

func TestUnitImageResourceConfigure(t *testing.T) {
	r := ImageResource{}

	err := standardResourceConfigureTests(&r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitImageResource(t *testing.T) {
	t.Parallel()

	// Mock Server Setup
	somePath := "some_path"
	someConfig := "some_config"
	s3Server, err := mockJsonServer(
		mockCfg{path: somePath, outText: someConfig, useJsonable: false},
	)
	if err != nil {
		t.Fatal(err)
	}
	someUrl := s3Server.URL + "/v3/" + somePath
	someArn := "some_arn"
	someReason := "some_reason"
	someTime := time.Now()
	someId := "some_id"
	imageResponse := openapi.DescribeImageResponseContent{
		ImageId:                         someId,
		Region:                          "some_other_region",
		Version:                         "some_version",
		ImageBuildStatus:                openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
		ImageBuildLogsArn:               &someArn,
		CloudformationStackStatus:       openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE.Ptr(),
		CloudformationStackStatusReason: &someReason,
		CloudformationStackArn:          &someArn,
		CreationTime:                    &someTime,
		CloudformationStackCreationTime: &someTime,
		CloudformationStackTags:         []openapi.Tag{},
		ImageConfiguration: openapi.ImageConfigurationStructure{
			Url: &someUrl,
		},
		ImagebuilderImageStatus:       openapi.IMAGEBUILDERIMAGESTATUS_AVAILABLE.Ptr(),
		ImagebuilderImageStatusReason: &someReason,
		Ec2AmiInfo: &openapi.Ec2AmiInfo{
			AmiId: "some_ami_id",
		},
	}

	buildImageResponse := openapi.BuildImageResponseContent{
		Image: openapi.ImageInfoSummary{
			ImageId:                   someId,
			ImageBuildStatus:          openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
			Region:                    imageResponse.Region,
			Version:                   imageResponse.Version,
			CloudformationStackArn:    imageResponse.CloudformationStackArn,
			CloudformationStackStatus: imageResponse.CloudformationStackStatus,
			Ec2AmiInfo: &openapi.Ec2AmiInfoSummary{
				AmiId: imageResponse.Ec2AmiInfo.AmiId,
			},
		},
	}

	deleteResponse := openapi.DeleteImageResponseContent{
		Image: openapi.ImageInfoSummary{
			ImageBuildStatus: openapi.IMAGEBUILDSTATUS_DELETE_FAILED,
		},
	}

	server, err := mockJsonServer(
		[]mockCfg{
			{path: "images/custom/" + someId, out: &imageResponse},
			{path: "images/custom", out: &buildImageResponse, method: http.MethodPost},
			{path: "images/custom/" + someId, out: &deleteResponse, method: http.MethodDelete},
		}...)
	if err != nil {
		t.Fatal(err)
	}

	resp := resource.ReadResponse{}
	req := resource.ReadRequest{}
	mResp := resource.SchemaResponse{}
	mReq := resource.SchemaRequest{}
	ctx := context.TODO()

	serverNotFound := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}),
	)

	cfgNotFound := openapi.NewConfiguration()
	cfgNotFound.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: serverNotFound.URL,
		},
	}

	r := ImageResource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfgNotFound),
	}

	r.Schema(ctx, mReq, &mResp)
	tfPlan := tfsdk.Plan{
		Schema: mResp.Schema,
	}
	tfState := tfsdk.State{
		Schema: mResp.Schema,
	}

	req.State = tfState
	r.Read(ctx, req, &resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	rawTfState := tftypes.NewValue(
		tftypes.Object{},
		map[string]tftypes.Value{
			"image_id": tftypes.NewValue(
				tftypes.String,
				"some_id",
			),
			"id": tftypes.NewValue(
				tftypes.String,
				"some_id",
			),
			"image_configuration": tftypes.NewValue(
				tftypes.String,
				"some_config",
			),
			"suppress_validators": tftypes.NewValue(
				tftypes.List{ElementType: tftypes.String},
				[]tftypes.Value{
					tftypes.NewValue(tftypes.String, string(openapi.VALIDATIONLEVEL_WARNING)),
				},
			),
			"validation_failure_level": tftypes.NewValue(
				tftypes.String,
				string(openapi.VALIDATIONLEVEL_INFO),
			),
			"rollback_on_failure": tftypes.NewValue(
				tftypes.Bool,
				true,
			),
			"region": tftypes.NewValue(
				tftypes.String,
				"some_region",
			),
			"version": tftypes.NewValue(
				tftypes.String,
				"some_version",
			),
			"cloudformation_stack_arn": tftypes.NewValue(
				tftypes.String,
				"some_stack_arn",
			),
			"cloudformation_stack_status": tftypes.NewValue(
				tftypes.String,
				string(openapi.CLOUDFORMATIONRESOURCESTATUS_CREATE_COMPLETE),
			),
			"image_build_status": tftypes.NewValue(
				tftypes.String,
				string(openapi.IMAGEBUILDERIMAGESTATUS_AVAILABLE),
			),
			"ami_id": tftypes.NewValue(
				tftypes.String,
				"some_id",
			),
		},
	)
	tfPlanRaw := tfsdk.Plan{
		Raw:    rawTfState,
		Schema: mResp.Schema,
	}
	tfStateRaw := tfsdk.State{
		Raw:    rawTfState,
		Schema: mResp.Schema,
	}

	req.State = tfStateRaw

	resp = resource.ReadResponse{}
	resp.State = tfState
	r.Read(ctx, req, &resp)

	output := ImageResourceModel{}
	resp.State.Get(ctx, &output)

	if output.ImageId.String() != `""` {
		t.Fatal("Expecting image id to be empty.")
	}
	if resp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", resp.Diagnostics.Errors())
	}

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	r.client = openapi.NewAPIClient(cfg)

	expectedOutput := ImageResourceModel{
		ImageId:          types.StringValue(imageResponse.ImageId),
		Id:               types.StringValue(imageResponse.ImageId),
		Region:           types.StringValue(imageResponse.Region),
		Version:          types.StringValue(imageResponse.Version),
		ImageBuildStatus: types.StringValue(string(imageResponse.ImageBuildStatus)),
		CloudformationStackStatus: types.StringValue(
			string(*imageResponse.CloudformationStackStatus),
		),
		CloudformationStackArn: types.StringValue(*imageResponse.CloudformationStackArn),
		ImageConfiguration:     types.StringValue(someConfig),
		AmiId:                  types.StringValue(imageResponse.Ec2AmiInfo.AmiId),
		RollbackOnFailure:      types.BoolValue(true),
		ValidationFailureLevel: types.StringValue(string(openapi.VALIDATIONLEVEL_INFO)),
		SuppressValidators: types.ListValueMust(
			types.StringType,
			[]attr.Value{types.StringValue(string(openapi.VALIDATIONLEVEL_WARNING))},
		),
	}

	resp = resource.ReadResponse{
		State: tfState,
	}
	r.Read(ctx, req, &resp)
	output = ImageResourceModel{}
	resp.State.Get(ctx, &output)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", resp.Diagnostics.Errors())
	}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	// Delete Testing
	dResp := resource.DeleteResponse{}
	dReq := resource.DeleteRequest{}

	dReq.State = tfState
	r.client = openapi.NewAPIClient(cfgNotFound)
	r.Delete(ctx, dReq, &dResp)
	if !dResp.Diagnostics.HasError() {
		t.Fatal("Expecting delete operation to return error.")
	}

	dReq.State = tfStateRaw
	dResp = resource.DeleteResponse{}

	r.client = openapi.NewAPIClient(cfg)
	r.Delete(ctx, dReq, &dResp)
	if !dResp.Diagnostics.HasError() {
		t.Fatal("Expecting delete operation to return error.")
	}

	deleteResponse.Image.ImageBuildStatus = openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE
	dReq.State = tfStateRaw
	dResp = resource.DeleteResponse{}
	r.client = openapi.NewAPIClient(cfg)
	r.Delete(ctx, dReq, &dResp)
	if dResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during delete operation: %v", dResp.Diagnostics.Errors())
	}

	// Create Testing
	cResp := resource.CreateResponse{}
	cReq := resource.CreateRequest{}

	cReq.Plan = tfPlan
	r.Create(ctx, cReq, &cResp)

	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}

	cResp = resource.CreateResponse{
		State: tfState,
	}
	cReq.Plan = tfPlanRaw
	r.client = openapi.NewAPIClient(cfgNotFound)
	r.Create(ctx, cReq, &cResp)
	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}

	cResp = resource.CreateResponse{
		State: tfState,
	}
	cReq.Plan = tfPlanRaw
	r.client = openapi.NewAPIClient(cfg)
	r.Create(ctx, cReq, &cResp)
	if cResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during delete operation: %v", dResp.Diagnostics.Errors())
	}

	output = ImageResourceModel{}
	cResp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	cResp = resource.CreateResponse{
		State: tfState,
	}
	imageResponse.ImageBuildStatus = openapi.IMAGEBUILDSTATUS_BUILD_FAILED
	r.Create(ctx, cReq, &cResp)
	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}
	imageResponse.ImageBuildStatus = openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE

	// Update Testing
	uResp := resource.UpdateResponse{}
	uReq := resource.UpdateRequest{}

	uReq.Plan = tfPlan
	uReq.State = tfState
	r.Update(ctx, uReq, &uResp)

	if !uResp.Diagnostics.HasError() {
		t.Fatal("Expecting update operation to return error.")
	}

	uResp = resource.UpdateResponse{
		State: tfState,
	}
	uReq = resource.UpdateRequest{
		Plan: tfPlanRaw,
	}
	r.Update(ctx, uReq, &uResp)
	if uResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during update operation: %v", dResp.Diagnostics.Errors())
	}

	output = ImageResourceModel{}
	uResp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	// Import Testing
	iResp := resource.ImportStateResponse{}
	iReq := resource.ImportStateRequest{ID: someId}

	iResp.State = tfState

	r.client = openapi.NewAPIClient(cfgNotFound)
	r.ImportState(ctx, iReq, &iResp)

	if !iResp.Diagnostics.HasError() {
		t.Fatal("Expecting import state operation to return error.")
	}

	iResp = resource.ImportStateResponse{
		State: tfState,
	}

	r.client = openapi.NewAPIClient(cfg)
	r.ImportState(ctx, iReq, &iResp)
	output = ImageResourceModel{}
	iResp.State.Get(ctx, &output)

	if iResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v.", iResp.Diagnostics.Errors())
	}

	expectedOutput.RollbackOnFailure = types.BoolNull()
	expectedOutput.ValidationFailureLevel = types.StringNull()
	expectedOutput.SuppressValidators = types.ListNull(types.StringType)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
	server.Close()
	serverNotFound.Close()
}
