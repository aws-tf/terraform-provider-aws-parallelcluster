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

	err := standardDataConfigureTests(&d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnitImageListDataSourceRead(t *testing.T) {
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

	req.Config = tfsdk.Config{
		Raw: tftypes.NewValue(
			tftypes.Object{},
			map[string]tftypes.Value{
				"image_status": tftypes.NewValue(
					tftypes.String,
					string(openapi.IMAGESTATUSFILTERINGOPTION_AVAILABLE),
				),
				"images": {},
				"region": tftypes.NewValue(tftypes.String, "some_region"),
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
	someStatus := openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE
	imageList := openapi.ListImagesResponseContent{
		Images: []openapi.ImageInfoSummary{{
			ImageId:                   "some_id",
			Region:                    "some_region",
			Version:                   "some_version",
			ImageBuildStatus:          openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
			CloudformationStackArn:    &someArn,
			CloudformationStackStatus: &someStatus,
			Ec2AmiInfo: &openapi.Ec2AmiInfoSummary{
				AmiId: "some_id",
			},
		}},
	}
	server, err := mockJsonServer([]string{"images/custom"}, imageList)
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

	expectedOutput := ImageListDataSourceModel{
		ImageStatus: types.StringValue(string(openapi.IMAGESTATUSFILTERINGOPTION_AVAILABLE)),
		Region:      types.StringValue("some_region"),
		Images: types.ListValueMust(types.ObjectType{AttrTypes: imageObjectTypes}, []attr.Value{
			types.ObjectValueMust(imageObjectTypes, map[string]attr.Value{
				"imageId": types.StringValue(imageList.Images[0].ImageId),
				"ec2AmiInfo": types.MapValueMust(
					types.StringType,
					map[string]attr.Value{
						"amiId": types.StringValue(imageList.Images[0].Ec2AmiInfo.AmiId),
					},
				),
				"region":  types.StringValue(imageList.Images[0].Region),
				"version": types.StringValue(imageList.Images[0].Version),
				"cloudformationStackArn": types.StringValue(
					*imageList.Images[0].CloudformationStackArn,
				),
				"imageBuildStatus": types.StringValue(
					string(imageList.Images[0].ImageBuildStatus),
				),
				"cloudformationStackStatus": types.StringValue(
					string(*imageList.Images[0].CloudformationStackStatus),
				),
			}),
		}),
	}

	var output ImageListDataSourceModel
	resp.State.Get(ctx, &output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf(
			"Expected output did not match output from read operation. \nO: %v\nE: %v\n",
			output,
			expectedOutput,
		)
	}
}
