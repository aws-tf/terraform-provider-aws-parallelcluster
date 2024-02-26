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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewComputeFleetDataSource(t *testing.T) {
	t.Parallel()

	obj := NewComputeFleetDataSource()
	if _, ok := obj.(*ComputeFleetDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ComputeFleetDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitComputeFleetDataSourceMetadata(t *testing.T) {
	d := NewComputeFleetDataSource()
	p := PclusterProvider{}
	resp := datasource.MetadataResponse{}
	req := datasource.MetadataRequest{}
	presp := provider.MetadataResponse{}
	preq := provider.MetadataRequest{}

	p.Metadata(context.TODO(), preq, &presp)
	req.ProviderTypeName = presp.TypeName

	d.Metadata(context.TODO(), req, &resp)

	if !strings.HasPrefix(resp.TypeName, presp.TypeName) {
		t.Fatalf(
			"Error provider typename expected as the prefix for resource or datasource name. \nO: %#v\nE: %#v",
			resp.TypeName,
			presp.TypeName,
		)
	}
}

func TestUnitComputeFleetDataSourceSchema(t *testing.T) {
	d := ComputeFleetDataSource{}
	dModel := ComputeFleetDataSourceModel{}
	resp := datasource.SchemaResponse{}
	req := datasource.SchemaRequest{}

	d.Schema(context.TODO(), req, &resp)

	rDataSource := reflect.TypeOf(dModel)
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

func TestUnitComputeFleetDataSourceConfigure(t *testing.T) {
	d := ComputeFleetDataSource{}

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitComputeFleetDataSourceRead(t *testing.T) {
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

	d := ComputeFleetDataSource{
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
				"status": {},
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

	resp = datasource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}

	someTime := time.Now()
	computeFleet := openapi.DescribeComputeFleetResponseContent{
		Status:                openapi.COMPUTEFLEETSTATUS_ENABLED,
		LastStatusUpdatedTime: &someTime,
	}

	server, err := mockJsonServer(
		mockCfg{path: "clusters/some_name/computefleet", out: computeFleet},
	)
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
		t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics.Errors())
	}

	expectedOutput := ComputeFleetDataSourceModel{
		ClusterName: types.StringValue("some_name"),
		Region:      types.StringValue("some_region"),
		Status:      types.StringValue(string(computeFleet.Status)),
		LastStatusUpdateTime: types.StringValue(
			computeFleet.LastStatusUpdatedTime.Round(0).String(),
		),
	}
	var output ComputeFleetDataSourceModel
	resp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf("Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
}
