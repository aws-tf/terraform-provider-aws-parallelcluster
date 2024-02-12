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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	metadataObjectTfTypes = map[string]tftypes.Type{
		"name":    tftypes.String,
		"version": tftypes.String,
	}
	schedulerObjectTfTypes = map[string]tftypes.Type{
		"metadata": tftypes.Object{
			AttributeTypes: metadataObjectTfTypes,
		},
		"type": tftypes.String,
	}
	clusterObjectTfTypes = map[string]tftypes.Type{
		"clusterName":               tftypes.String,
		"clusterStatus":             tftypes.String,
		"region":                    tftypes.String,
		"cloudformationStackArn":    tftypes.String,
		"cloudformationStackStatus": tftypes.String,
		"scheduler": tftypes.Object{
			AttributeTypes: schedulerObjectTfTypes,
		},
		"version": tftypes.String,
	}
)

func TestUnitNewClusterListDataSource(t *testing.T) {
	t.Parallel()

	obj := NewClusterListDataSource()
	if _, ok := obj.(*ClusterListDataSource); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: ClusterListDataSource",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitClusterListDataSourceMetadata(t *testing.T) {
	d := NewClusterListDataSource()
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

func TestUnitClusterListDataSourceSchema(t *testing.T) {
	d := ClusterListDataSource{}
	dataSourceModel := ClusterListDataSourceModel{}
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

func TestUnitClusterListDataSourceConfigure(t *testing.T) {
	d := ClusterListDataSource{}
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

func TestUnitClusterListDataSourceRead(t *testing.T) {
	t.Parallel()
	resp := datasource.ReadResponse{}
	req := datasource.ReadRequest{}
	mResp := datasource.SchemaResponse{}
	mReq := datasource.SchemaRequest{}
	ctx := context.TODO()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	d := ClusterListDataSource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	d.Schema(ctx, mReq, &mResp)

	req.Config = tfsdk.Config{
		Schema: mResp.Schema,
	}
	resp.State = tfsdk.State{
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
				"clusters": {},
				"cluster_status": tftypes.NewValue(
					tftypes.List{ElementType: tftypes.String},
					[]tftypes.Value{
						tftypes.NewValue(
							tftypes.String,
							string(openapi.CLUSTERSTATUS_CREATE_COMPLETE),
						),
					},
				),
				"region": {},
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

	name := "test"
	version := "some_version"

	scheduler := openapi.NewScheduler(name)
	scheduler.Metadata = &openapi.Metadata{
		Name:    &name,
		Version: &version,
	}

	expectedOutput := openapi.ClusterInfoSummary{
		ClusterName:               name,
		CloudformationStackArn:    "some_stack",
		CloudformationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
		Scheduler:                 scheduler,
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respJSON, err := openapi.ListClustersResponseContent{
			Clusters: []openapi.ClusterInfoSummary{expectedOutput},
		}.MarshalJSON()
		if err != nil {
			t.Fatal("Failed to marshal list clusters response JSON.")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respJSON))
	}))
	defer server.Close()

	cfg = openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: server.URL,
		},
	}

	d = ClusterListDataSource{
		awsv4:  awsv4Test(),
		client: openapi.NewAPIClient(cfg),
	}

	req.Config = tfsdk.Config{
		Raw: tftypes.NewValue(
			tftypes.Object{},
			map[string]tftypes.Value{
				"clusters": tftypes.NewValue(
					tftypes.List{
						ElementType: tftypes.Object{AttributeTypes: clusterObjectTfTypes},
					},
					[]tftypes.Value{
						tftypes.NewValue(
							tftypes.Object{AttributeTypes: clusterObjectTfTypes},
							map[string]tftypes.Value{
								"clusterName": tftypes.NewValue(tftypes.String, "test"),
								"clusterStatus": tftypes.NewValue(
									tftypes.String,
									string(openapi.CLUSTERSTATUS_CREATE_COMPLETE),
								),
								"region": tftypes.NewValue(
									tftypes.String,
									"us-east-1",
								),
								"cloudformationStackArn": tftypes.NewValue(
									tftypes.String,
									"some_arn",
								),
								"cloudformationStackStatus": tftypes.NewValue(
									tftypes.String,
									string(openapi.CLOUDFORMATIONRESOURCESTATUS_CREATE_COMPLETE),
								),
								"scheduler": tftypes.NewValue(
									tftypes.Object{
										AttributeTypes: schedulerObjectTfTypes,
									},
									map[string]tftypes.Value{
										"metadata": tftypes.NewValue(
											tftypes.Object{
												AttributeTypes: metadataObjectTfTypes,
											},
											map[string]tftypes.Value{
												"name":    tftypes.NewValue(tftypes.String, ""),
												"version": tftypes.NewValue(tftypes.String, ""),
											},
										),
										"type": tftypes.NewValue(tftypes.String, ""),
									},
								),
								"version": tftypes.NewValue(tftypes.String, "some_version"),
							},
						),
					},
				),
				"cluster_status": tftypes.NewValue(
					tftypes.List{ElementType: tftypes.String},
					[]tftypes.Value{},
				),
				"region": tftypes.NewValue(tftypes.String, ""),
			},
		),
		Schema: mResp.Schema,
	}

	resp = datasource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics)
	}

	output := ClusterListDataSourceModel{}
	diags := resp.State.Get(ctx, &output)
	if diags.HasError() {
		t.Fatalf("Read operation returned unexpected error while retrieving state data: %v", diags)
	}

	expectedBytOut, err := expectedOutput.MarshalJSON()
	if err != nil {
		t.Fatalf("Unexpected error while converting expected output to json: %v", err.Error())
	}

	if output.Clusters.Elements()[0].String() != string(expectedBytOut) {
		t.Fatalf(
			"Expected output did not match output from read operation. E: %v\nO: %v\n",
			string(expectedBytOut),
			output.Clusters.Elements()[0].String(),
		)
	}
}
