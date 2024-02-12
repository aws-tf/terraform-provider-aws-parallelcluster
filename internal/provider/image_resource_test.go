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
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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
	resp := resource.ConfigureResponse{}
	req := resource.ConfigureRequest{}

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: "testURL",
		},
	}

	awsv4 := awsv4Test()

	req.ProviderData = configData{
		awsv4:  awsv4,
		client: openapi.NewAPIClient(cfg),
	}

	r.Configure(context.TODO(), req, &resp)

	if r.client == nil {
		t.Fatal("Error client expected to be set.")
	}

	if r.awsv4 != awsv4 {
		t.Fatalf("Error matching output expected. O: %#v\nE: %#v",
			r.awsv4,
			awsv4,
		)
	}
}

func TestUnitImageResourceRead(t *testing.T) {
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

	r := ImageResource{
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
	iReq := resource.ImportStateRequest{ID: "Test"}

	iResp.State = tfsdk.State{
		Schema: mResp.Schema,
	}

	r.ImportState(ctx, iReq, &iResp)

	if !iResp.Diagnostics.HasError() {
		t.Fatal("Expecting import state operation to return error.")
	}
}
