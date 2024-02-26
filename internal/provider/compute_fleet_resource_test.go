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
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewComputeFleetStatusResource(t *testing.T) {
	t.Parallel()

	obj := NewComputeFleetStatusResource()
	if _, ok := obj.(*ComputeFleetStatusResource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ComputeFleetStatusResource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitComputeFleetStatusResourceMetadata(t *testing.T) {
	r := NewComputeFleetStatusResource()
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

func TestUnitComputeFleetStatusResourceSchema(t *testing.T) {
	r := ComputeFleetStatusResource{}
	dataSourceModel := ComputeFleetStatusResourceModel{}
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

func TestUnitComputeFleetStatusResourceConfigure(t *testing.T) {
	r := ComputeFleetStatusResource{}

	err := standardResourceConfigureTests(&r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitComputeFleetStatusResourceRead(t *testing.T) {
	t.Parallel()
	resp := resource.ReadResponse{}
	req := resource.ReadRequest{}
	mResp := resource.SchemaResponse{}
	mReq := resource.SchemaRequest{}
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

	r := ComputeFleetStatusResource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	r.Schema(ctx, mReq, &mResp)

	req.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.Read(ctx, req, &resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	req.State = tfsdk.State{
		Raw: tftypes.NewValue(
			tftypes.Object{},
			map[string]tftypes.Value{
				"id": tftypes.NewValue(
					tftypes.String,
					"some_id",
				),
				"cluster_name": tftypes.NewValue(
					tftypes.String,
					"some_name",
				),
				"status_request": {},
				"status":         {},
				"region": tftypes.NewValue(
					tftypes.String,
					"some_region",
				),
				"last_status_update_time": tftypes.NewValue(
					tftypes.String,
					"some_time",
				),
			},
		),
		Schema: mResp.Schema,
	}

	someTime := time.Now()
	emptyFleet := openapi.DescribeComputeFleetResponseContent{}
	computeFleet := openapi.DescribeComputeFleetResponseContent{
		Status:                openapi.COMPUTEFLEETSTATUS_ENABLED,
		LastStatusUpdatedTime: &someTime,
	}

	tests := map[string]openapi.DescribeComputeFleetResponseContent{
		"invalid/path":                    emptyFleet,
		"clusters/some_name/computefleet": computeFleet,
	}

	expectedOutputs := map[string]ComputeFleetStatusResourceModel{
		"invalid/path": {
			Id:          types.StringValue("some_id"),
			ClusterName: types.StringValue("some_name"),
			Region:      types.StringValue("some_region"),
			Status:      types.StringValue(""),
			LastStatusUpdateTime: types.StringValue(
				emptyFleet.GetLastStatusUpdatedTime().String(),
			),
		},
		"clusters/some_name/computefleet": {
			Id:          types.StringValue("some_id"),
			ClusterName: types.StringValue("some_name"),
			Region:      types.StringValue("some_region"),
			Status:      types.StringValue(string(computeFleet.Status)),
			LastStatusUpdateTime: types.StringValue(
				computeFleet.LastStatusUpdatedTime.Round(0).String(),
			),
		},
	}
	for k, test := range tests {
		server, err := mockJsonServer([]string{k}, test)
		if err != nil {
			t.Fatal(err)
		}
		defer server.Close()

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
		r.Read(ctx, req, &resp)

		output := ComputeFleetStatusResourceModel{}
		resp.State.Get(ctx, &output)

		if resp.Diagnostics.HasError() {
			t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics)
		}

		if !reflect.DeepEqual(output, expectedOutputs[k]) {
			t.Fatalf(
				"Epected output did not match actual output: \nO:%v\nE:%v\n",
				output,
				expectedOutputs[k],
			)
		}
	}

	dResp := resource.DeleteResponse{}
	dReq := resource.DeleteRequest{}

	dReq.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.Delete(ctx, dReq, &dResp)

	if !dResp.Diagnostics.HasError() {
		t.Fatal("Expecting delete operation to return error.")
	}

	cResp := resource.CreateResponse{}
	cReq := resource.CreateRequest{}

	cReq.Plan = tfsdk.Plan{
		Schema: mResp.Schema,
	}

	r.Create(ctx, cReq, &cResp)

	if !cResp.Diagnostics.HasError() {
		t.Fatal("Expecting create operation to return error.")
	}

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

	iResp := resource.ImportStateResponse{}
	iReq := resource.ImportStateRequest{ID: "test"}

	iResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.ImportState(ctx, iReq, &iResp)

	if !iResp.Diagnostics.HasError() {
		t.Fatal("Expecting import state operation to return error.")
	}
}
