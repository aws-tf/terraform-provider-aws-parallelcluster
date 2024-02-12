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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestUnitNewImageListDataSource(t *testing.T) {
	t.Parallel()

	obj := NewImageListDataSource()
	if _, ok := obj.(*ImageListDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ImageListDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitImageListDataSourceMetadata(t *testing.T) {
	d := NewImageListDataSource()
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

func TestUnitImageListDataSourceSchema(t *testing.T) {
	d := ImageListDataSource{}
	dataSourceModel := ImageListDataSourceModel{}
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

func TestUnitImageListDataSourceConfigure(t *testing.T) {
	d := ImageListDataSource{}
	resp := datasource.ConfigureResponse{}
	req := datasource.ConfigureRequest{}

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

	d.Configure(context.TODO(), req, &resp)

	if d.client == nil {
		t.Fatal("Error client expected to be set.")
	}

	if d.awsv4 != awsv4 {
		t.Fatalf("Error matching output expected. O: %#v\nE: %#v",
			d.awsv4,
			awsv4,
		)
	}
}

func TestUnitImageListDataSourceRead(t *testing.T) {
	t.Parallel()
	resp := datasource.ReadResponse{}
	req := datasource.ReadRequest{}
	mResp := datasource.SchemaResponse{}
	mReq := datasource.SchemaRequest{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	d := ImageListDataSource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	d.Schema(context.TODO(), mReq, &mResp)

	req.Config = tfsdk.Config{
		Schema: mResp.Schema,
	}

	d.Read(context.TODO(), req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("Expecting read operation to return error.")
	}
}
