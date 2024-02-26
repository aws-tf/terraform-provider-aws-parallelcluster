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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnitNewOfficialImageDataSource(t *testing.T) {
	t.Parallel()

	obj := NewOfficialImageDataSource()
	if _, ok := obj.(*OfficialImageDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: OfficialImageDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitOfficialImageDataSourceMetadata(t *testing.T) {
	d := NewOfficialImageDataSource()
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

func TestUnitOfficialImageDataSourceSchema(t *testing.T) {
	d := OfficialImageDataSource{}
	dataSourceModel := OfficialImageDataSourceModel{}
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

func TestUnitOfficialImageDataSourceConfigure(t *testing.T) {
	d := OfficialImageDataSource{}

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitOfficialImageDataSourceRead(t *testing.T) {
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

	d := OfficialImageDataSource{
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
				"architecture": tftypes.NewValue(
					tftypes.String,
					"some_architecture",
				),
				"os": tftypes.NewValue(
					tftypes.String,
					"some_os",
				),
				"region": tftypes.NewValue(
					tftypes.String,
					"some_region",
				),
				"official_images": {},
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

	officialImageList := openapi.ListOfficialImagesResponseContent{
		Images: []openapi.AmiInfo{{
			Architecture: "some_architecture",
			AmiId:        "some_id",
			Name:         "some_name",
			Os:           "some_os",
			Version:      "some_version",
		}},
	}

	server, err := mockJsonServer(mockCfg{path: "images/official", out: officialImageList})
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

	if resp.Diagnostics.HasError() {
		t.Fatalf("Unexpected error during read operation: %v", resp.Diagnostics.Errors())
	}

	expectedOutput := OfficialImageDataSourceModel{
		Architecture: types.StringValue("some_architecture"),
		Os:           types.StringValue("some_os"),
		Region:       types.StringValue("some_region"),
		OfficialImages: types.ListValueMust(
			types.MapType{ElemType: types.StringType},
			[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{
				"architecture": types.StringValue(officialImageList.Images[0].Architecture),
				"amiId":        types.StringValue(officialImageList.Images[0].AmiId),
				"name":         types.StringValue(officialImageList.Images[0].Name),
				"os":           types.StringValue(officialImageList.Images[0].Os),
				"version":      types.StringValue(officialImageList.Images[0].Version),
			})},
		),
	}

	var output OfficialImageDataSourceModel
	resp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match actual output: \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
}
