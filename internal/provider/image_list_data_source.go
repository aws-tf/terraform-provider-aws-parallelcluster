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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_                datasource.DataSource = &ImageListDataSource{}
	imageObjectTypes                       = map[string]attr.Type{
		"imageId":                   types.StringType,
		"ec2AmiInfo":                types.MapType{ElemType: types.StringType},
		"region":                    types.StringType,
		"version":                   types.StringType,
		"cloudformationStackArn":    types.StringType,
		"imageBuildStatus":          types.StringType,
		"cloudformationStackStatus": types.StringType,
	}
)

func NewImageListDataSource() datasource.DataSource {
	return &ImageListDataSource{}
}

// ImageListDataSource defines the data source implementation.
type ImageListDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ImageListDataSourceModel describes the data source data model.
type ImageListDataSourceModel struct {
	ImageStatus types.String `tfsdk:"image_status"`
	Images      types.List   `tfsdk:"images"`
	Region      types.String `tfsdk:"region"`
}

func ValidateImageStatusFilter(
	ctx context.Context,
	req validator.StringRequest,
	resp *validator.StringResponse,
) {
	_, err := openapi.NewImageStatusFilteringOptionFromValue(req.ConfigValue.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Validation Failure.",
			fmt.Sprintf("Must be one of %v", openapi.AllowedImageStatusFilteringOptionEnumValues),
		)
	}
}

func (d *ImageListDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_list_images"
}

func (d *ImageListDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Retrieve the list of existing custom images.",

		Attributes: map[string]schema.Attribute{
			"image_status": schema.StringAttribute{
				MarkdownDescription: "Filter by image status.",
				Optional:            false,
				Required:            true,
				Validators: []validator.String{
					&AttributeValidator{
						description: fmt.Sprintf(
							"Valid values: %v",
							openapi.AllowedImageStatusFilteringOptionEnumValues,
						),
						markdownDescription: fmt.Sprintf(
							"Valid values: `%v`",
							openapi.AllowedImageStatusFilteringOptionEnumValues,
						),
						validatorFunction: ValidateImageStatusFilter,
					},
				},
			},
			"images": schema.ListAttribute{
				MarkdownDescription: "List of images.",
				Optional:            false,
				Required:            false,
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: imageObjectTypes,
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region of the images.",
				Optional:            true,
			},
		},
	}
}

func (d *ImageListDataSource) Configure(
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

func (d *ImageListDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ImageListDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, d.awsv4)

	imageStatus := openapi.ImageStatusFilteringOption(data.ImageStatus.ValueString())

	imageListReq := d.client.ImageOperationsAPI.ListImages(reqCtx)
	imageListReq = imageListReq.ImageStatus(imageStatus).Region(data.Region.ValueString())

	imageList, _, err := imageListReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error occured while retrieving image list",
			fmt.Sprintf("Error: %v", err),
		)
	}

	goImageList := make([]types.Object, 0)
	if images, ok := imageList.GetImagesOk(); ok {
		sort.Slice(
			images,
			func(i, j int) bool { return images[i].GetImageId() > images[j].GetImageId() },
		)

		for _, image := range images {
			if !image.HasCloudformationStackArn() {
				image.SetCloudformationStackArn("")
			}
			if !image.HasCloudformationStackStatus() {
				image.SetCloudformationStackStatus("")
			}
			imageMap, err := image.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving image list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			delete(imageMap, "ec2AmiInfo")

			tfImageMap, diags := types.MapValueFrom(ctx, types.StringType, imageMap)

			resp.Diagnostics.Append(diags...)

			tfImageMapElements := tfImageMap.Elements()

			var tfEc2AmiInfoMap types.Map
			if ec2AmiInfo, ok := image.GetEc2AmiInfoOk(); ok {
				ec2AmiInfoMap, err := ec2AmiInfo.ToMap()
				if err != nil {
					resp.Diagnostics.AddError(
						"Error occured while retrieving image list",
						fmt.Sprintf("Error: %v", err),
					)
				}

				tfEc2AmiInfoMap, diags = types.MapValueFrom(ctx, types.StringType, ec2AmiInfoMap)
				resp.Diagnostics.Append(diags...)

			} else {
				tfEc2AmiInfoMap = types.MapNull(types.StringType)
			}

			tfImageMapElements["ec2AmiInfo"] = tfEc2AmiInfoMap
			tfImageObject, diags := types.ObjectValue(
				imageObjectTypes,
				tfImageMapElements,
			)
			resp.Diagnostics.Append(diags...)

			goImageList = append(goImageList, tfImageObject)
		}
	}

	tfImageList, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: imageObjectTypes},
		goImageList,
	)
	resp.Diagnostics.Append(diags...)
	data.Images = tfImageList

	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
