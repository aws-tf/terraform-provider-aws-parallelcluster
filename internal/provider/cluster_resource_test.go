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

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewClusterResource(t *testing.T) {
	t.Parallel()

	obj := NewClusterResource()
	if _, ok := obj.(*ClusterResource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ClusterResource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitClusterResourceMetadata(t *testing.T) {
	r := NewClusterResource()
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

func TestUnitClusterResourceSchema(t *testing.T) {
	r := ClusterResource{}
	dataSourceModel := ClusterResourceModel{}
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

func TestUnitClusterResourceConfigure(t *testing.T) {
	r := ClusterResource{}

	err := standardResourceConfigureTests(&r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitClusterResource(t *testing.T) {
	t.Parallel()
	resp := resource.ReadResponse{}
	req := resource.ReadRequest{}
	mResp := resource.SchemaResponse{}
	mReq := resource.SchemaRequest{}
	ctx := context.TODO()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	cfgNotFound := openapi.NewConfiguration()
	cfgNotFound.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	r := ClusterResource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfgNotFound),
	}

	r.Schema(ctx, mReq, &mResp)

	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	req.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.Read(ctx, req, &resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	rawTfValue := tftypes.NewValue(
		tftypes.Object{},
		map[string]tftypes.Value{
			"cluster_name": tftypes.NewValue(
				tftypes.String,
				"some_name",
			),
			"region": tftypes.NewValue(
				tftypes.String,
				"some_region",
			),
			"cluster_configuration": tftypes.NewValue(tftypes.String, "some_config"),
			"rollback_on_failure":   {},
			"suppress_validators": tftypes.NewValue(
				tftypes.List{ElementType: tftypes.String},
				[]tftypes.Value{},
			),
			"validation_failure_level": {},
			"id":                       tftypes.NewValue(tftypes.String, "some_name"),
			"cloudformation_stack_arn": tftypes.NewValue(tftypes.String, "some_arn"),
			"cloudformation_stack_status": tftypes.NewValue(
				tftypes.String,
				string(openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_IN_PROGRESS),
			),
			"cluster_status": tftypes.NewValue(
				tftypes.String,
				string(openapi.CLUSTERSTATUS_CREATE_IN_PROGRESS),
			),
			"version": tftypes.NewValue(tftypes.String, "some_verison"),
		},
	)
	req.State = tfsdk.State{
		Raw:    rawTfValue,
		Schema: mResp.Schema,
	}

	resp = resource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.Read(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", resp.Diagnostics.Errors())
	}

	somePath := "some_url"
	s3Server, err := mockJsonServer(
		mockCfg{path: somePath, outText: "some_config", useJsonable: false},
	)
	if err != nil {
		t.Fatal(err)
	}

	configUrl := s3Server.URL + "/v3/" + somePath
	input := openapi.DescribeClusterResponseContent{
		ClusterName:   "some_name",
		ClusterStatus: openapi.CLUSTERSTATUS_CREATE_COMPLETE,
		Region:        "some_other_region",
		ClusterConfiguration: openapi.ClusterConfigurationStructure{
			Url: &configUrl,
		},
		CloudformationStackArn:    "some_other_arn",
		CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		Version:                   "some_other_version",
		ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_ENABLED,
	}

	infoSummary := openapi.ClusterInfoSummary{
		ClusterName:               input.ClusterName,
		Region:                    input.Region,
		Version:                   input.Version,
		CloudformationStackArn:    input.CloudformationStackArn,
		CloudformationStackStatus: input.CloudFormationStackStatus,
		ClusterStatus:             input.ClusterStatus,
		Scheduler:                 input.Scheduler,
	}
	createInput := openapi.CreateClusterResponseContent{
		Cluster: infoSummary,
	}

	updateInput := openapi.UpdateClusterResponseContent{
		Cluster: infoSummary,
	}

	server, err = mockJsonServer(
		[]mockCfg{
			{path: "clusters/some_name", out: input},
			{path: "clusters", out: createInput, method: http.MethodPost},
			{path: "clusters/some_name", out: updateInput, method: http.MethodPut},
		}...)
	if err != nil {
		t.Fatal(err)
	}

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}
	r.client = openapi.NewAPIClient(cfg)

	resp = resource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	expectedOutput := ClusterResourceModel{
		ClusterName:               types.StringValue(input.ClusterName),
		Id:                        types.StringValue(input.ClusterName),
		ClusterConfiguration:      types.StringValue("some_config"),
		CloudformationStackArn:    types.StringValue(input.CloudformationStackArn),
		ClusterStatus:             types.StringValue(string(input.ClusterStatus)),
		CloudformationStackStatus: types.StringValue(string(input.CloudFormationStackStatus)),
		Region:                    types.StringValue(input.Region),
		RollbackOnFailure:         types.BoolNull(),
		SuppressValidators:        types.ListValueMust(types.StringType, []attr.Value{}),
		Version:                   types.StringValue(input.Version),
	}

	output := ClusterResourceModel{}
	r.Read(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", resp.Diagnostics.Errors())
	}

	resp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	// Delete Unit Testing
	for _, c := range []struct {
		resp      resource.DeleteResponse
		req       resource.DeleteRequest
		client    *openapi.APIClient
		shouldErr bool
	}{
		{
			resp: resource.DeleteResponse{},
			req: resource.DeleteRequest{
				State: tfsdk.State{
					Schema: mResp.Schema,
				},
			},
			client:    openapi.NewAPIClient(cfgNotFound),
			shouldErr: true,
		},
		{
			resp: resource.DeleteResponse{
				State: tfsdk.State{},
			},
			req: resource.DeleteRequest{
				State: tfsdk.State{
					Raw:    rawTfValue,
					Schema: mResp.Schema,
				},
			},
			client:    openapi.NewAPIClient(cfgNotFound),
			shouldErr: false,
		},
	} {
		r.client = c.client
		r.Delete(ctx, c.req, &c.resp)

		if c.shouldErr && !c.resp.Diagnostics.HasError() {
			t.Fatal("Expecting delete operation to return error.")
		}

		if !c.shouldErr {
			if c.resp.Diagnostics.HasError() {
				t.Fatalf(
					"Unexpected error during delete operation: %v",
					c.resp.Diagnostics.Errors(),
				)
			}
			if !c.resp.State.Raw.IsNull() {
				output = ClusterResourceModel{}
				c.resp.State.Get(ctx, &output)
				t.Fatalf("Expected nil state after delete operation: \nO: %v", output)
			}
		}
	}

	// Create Unit Testing
	cResp := resource.CreateResponse{}
	cReq := resource.CreateRequest{}

	cReq.Plan = tfsdk.Plan{
		Schema: mResp.Schema,
	}

	r.Create(ctx, cReq, &cResp)

	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}

	cReq.Plan.Raw = rawTfValue
	cResp = resource.CreateResponse{}
	r.Create(ctx, cReq, &cResp)

	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}

	cResp = resource.CreateResponse{}
	cResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.client = openapi.NewAPIClient(cfg)

	r.Create(ctx, cReq, &cResp)
	if cResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during create operation: %v", cResp.Diagnostics.Errors())
	}

	output = ClusterResourceModel{}
	cResp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	// Update Unit Testing
	r.client = openapi.NewAPIClient(cfgNotFound)
	uResp := resource.UpdateResponse{}
	uReq := resource.UpdateRequest{}

	uReq.Plan = tfsdk.Plan{
		Schema: mResp.Schema,
	}

	uReq.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.Update(ctx, uReq, &uResp)

	if !uResp.Diagnostics.HasError() {
		t.Fatal("Expecting update operation to return error.")
	}

	uReq.Plan.Raw = rawTfValue
	uReq.State.Raw = rawTfValue

	uResp = resource.UpdateResponse{}
	uResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	r.Update(ctx, uReq, &uResp)

	if uResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during update operation: %v", uResp.Diagnostics.Errors())
	}

	uReq.Plan.Raw = rawTfValue
	uReq.State.Raw = cResp.State.Raw

	uResp = resource.UpdateResponse{}
	uResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	r.client = openapi.NewAPIClient(cfg)
	r.Update(ctx, uReq, &uResp)

	if uResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during update operation: %v", uResp.Diagnostics.Errors())
	}

	output = ClusterResourceModel{}
	uResp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}

	// Import Unit Testing
	iResp := resource.ImportStateResponse{}
	iReq := resource.ImportStateRequest{ID: "some_name"}

	iResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.client = openapi.NewAPIClient(cfgNotFound)
	r.ImportState(ctx, iReq, &iResp)

	if !iResp.Diagnostics.HasError() {
		t.Fatal("Expecting import state operation to return error.")
	}

	r.client = openapi.NewAPIClient(cfg)
	iResp = resource.ImportStateResponse{
		State: tfsdk.State{Schema: mResp.Schema},
	}

	r.ImportState(ctx, iReq, &iResp)
	output = ClusterResourceModel{}
	iResp.State.Get(ctx, &output)

	if iResp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", uResp.Diagnostics.Errors())
	}

	expectedOutput.SuppressValidators = types.ListNull(types.StringType)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
}
