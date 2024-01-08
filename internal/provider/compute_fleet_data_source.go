// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource = &ComputeFleetDataSource{}
)

func NewComputeFleetDataSource() datasource.DataSource {
	return &ComputeFleetDataSource{}
}

// ComputeFleetDataSource defines the data source implementation.
type ComputeFleetDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ComputeFleetDataSourceModel describes the data source data model.
type ComputeFleetDataSourceModel struct {
	ClusterName          types.String `tfsdk:"cluster_name"`
	Region               types.String `tfsdk:"region"`
	Status               types.String `tfsdk:"status"`
	LastStatusUpdateTime types.String `tfsdk:"last_status_update_time"`
}

func (d *ComputeFleetDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_compute_fleet_status"
}

func (d *ComputeFleetDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Describe the status of the compute fleet.",

		Attributes: map[string]schema.Attribute{
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Filter by architecture. The default is no filtering.",
				Optional:            false,
				Required:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region that official images are listed in.",
				Optional:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The status of the compute fleet.",
				Optional:            false,
				Required:            false,
				Computed:            true,
			},
			"last_status_update_time": schema.StringAttribute{
				MarkdownDescription: "The timestamp representing the last status update time.",
				Optional:            false,
				Required:            false,
				Computed:            true,
			},
		},
	}
}

func (d *ComputeFleetDataSource) Configure(
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

func (d *ComputeFleetDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ComputeFleetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, d.awsv4)

	computeFleet, rawHttp, err := d.client.ClusterComputeFleetAPI.DescribeComputeFleet(
		reqCtx,
		data.ClusterName.ValueString(),
	).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("%v", err.Error()),
			fmt.Sprintf("%v", rawHttp.Body),
		)
	}

	if computeFleet != nil {
		data.Status = types.StringValue(string(computeFleet.GetStatus()))
		data.LastStatusUpdateTime = types.StringValue(
			computeFleet.GetLastStatusUpdatedTime().String(),
		)
	}

	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
