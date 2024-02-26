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

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
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

	input := openapi.ListClustersResponseContent{
		Clusters: []openapi.ClusterInfoSummary{
			{
				ClusterName:               name,
				CloudformationStackArn:    "some_stack",
				CloudformationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
				Scheduler:                 scheduler,
			},
		},
	}
	server, err := mockJsonServer(mockCfg{path: "clusters", out: input})
	if err != nil {
		t.Fatal(err)
	}
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

	resp = datasource.ReadResponse{}
	resp.State = tfsdk.State{
		Schema: mResp.Schema,
	}
	d.Read(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Read operation returned unexpected error: %v", resp.Diagnostics)
	}

	expectedOutput := ClusterListDataSourceModel{
		ClusterStatus: types.ListValueMust(types.StringType, []attr.Value{}),
		Region:        types.StringValue(""),
		Clusters: types.ListValueMust(
			types.ObjectType{AttrTypes: clusterObjectTypes},
			[]attr.Value{types.ObjectValueMust(clusterObjectTypes, map[string]attr.Value{
				"clusterName": types.StringValue(input.Clusters[0].ClusterName),
				"clusterStatus": types.StringValue(
					string(input.Clusters[0].ClusterStatus),
				),
				"region": types.StringValue(input.Clusters[0].Region),
				"cloudformationStackArn": types.StringValue(
					input.Clusters[0].CloudformationStackArn,
				),
				"cloudformationStackStatus": types.StringValue(
					string(input.Clusters[0].CloudformationStackStatus),
				),
				"scheduler": types.ObjectValueMust(
					schedulerObjectTypes,
					map[string]attr.Value{
						"metadata": types.ObjectValueMust(
							metadataObjectTypes,
							map[string]attr.Value{
								"name": types.StringValue(
									*input.Clusters[0].Scheduler.Metadata.Name,
								),
								"version": types.StringValue(
									*input.Clusters[0].Scheduler.Metadata.Version,
								),
							},
						),
						"type": types.StringValue(input.Clusters[0].Scheduler.Type),
					},
				),
				"version": types.StringValue(input.Clusters[0].Version),
			})},
		),
	}

	output := ClusterListDataSourceModel{}
	diags := resp.State.Get(ctx, &output)
	if diags.HasError() {
		t.Fatalf("Read operation returned unexpected error while retrieving state data: %v", diags)
	}

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Fatalf(
			"Expected output did not match output from read operation. \nE: %v\nO: %v\n",
			expectedOutput,
			output,
		)
	}
}
