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
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource = &OfficialImageDataSource{}
)

func NewOfficialImageDataSource() datasource.DataSource {
	return &OfficialImageDataSource{}
}

// OfficialImageDataSource defines the data source implementation.
type OfficialImageDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// OfficialImageDataSourceModel describes the data source data model.
type OfficialImageDataSourceModel struct {
	Architecture   types.String `tfsdk:"architecture"`
	Os             types.String `tfsdk:"os"`
	Region         types.String `tfsdk:"region"`
	OfficialImages types.List   `tfsdk:"official_images"`
}

func (d *OfficialImageDataSource) getClient() *openapi.APIClient {
	return d.client
}

func (d *OfficialImageDataSource) getAWSv4() openapi.AWSv4 {
	return d.awsv4
}

func (d *OfficialImageDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_list_official_images"
}

func (d *OfficialImageDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Retrieve the list of AWS ParallelCluster official images.",

		Attributes: map[string]schema.Attribute{
			"architecture": schema.StringAttribute{
				MarkdownDescription: "Filter by architecture. The default is no filtering.",
				Optional:            true,
			},
			"official_images": schema.ListAttribute{
				MarkdownDescription: "List of official images.",
				Optional:            false,
				Required:            false,
				Computed:            true,
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
			"os": schema.StringAttribute{
				MarkdownDescription: "Filter by OS distribution. The default is no filtering.",
				Optional:            true,
			},

			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that official images are listed in.",
				Optional:            true,
			},
		},
	}
}

func (d *OfficialImageDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(configData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf(
				"Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)

		return
	}

	d.client = client.client
	d.awsv4 = client.awsv4
}

func (d *OfficialImageDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data OfficialImageDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, d.awsv4)

	imageListReq := d.client.ImageOperationsAPI.ListOfficialImages(reqCtx)
	if !data.Architecture.IsNull() {
		imageListReq = imageListReq.Architecture(data.Architecture.ValueString())
	}
	if !data.Os.IsNull() {
		imageListReq = imageListReq.Os(data.Os.ValueString())
	}
	if !data.Region.IsNull() {
		imageListReq = imageListReq.Region(data.Region.ValueString())
	}

	imageList, _, err := imageListReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error occured while retrieving image list",
			fmt.Sprintf("Error: %v", err),
		)
	}

	goOfficialImageList := make([]types.Map, 0)
	if images, ok := imageList.GetImagesOk(); ok {
		sort.Slice(
			images,
			func(i, j int) bool { return images[i].GetName() > images[j].GetName() },
		)
		for _, image := range images {
			imageMap, err := image.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving image list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			tfOfficialImageMap, diags := types.MapValueFrom(ctx, types.StringType, imageMap)
			resp.Diagnostics.Append(diags...)

			goOfficialImageList = append(goOfficialImageList, tfOfficialImageMap)
		}
	}

	tfOfficialImageList, diags := types.ListValueFrom(
		ctx,
		types.MapType{ElemType: types.StringType},
		goOfficialImageList,
	)
	resp.Diagnostics.Append(diags...)
	data.OfficialImages = tfOfficialImageList

	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
