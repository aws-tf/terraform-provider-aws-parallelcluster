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
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_                   datasource.DataSource = &ClusterListDataSource{}
	metadataObjectTypes                       = map[string]attr.Type{
		"name":    types.StringType,
		"version": types.StringType,
	}
	schedulerObjectTypes = map[string]attr.Type{
		"metadata": types.ObjectType{AttrTypes: metadataObjectTypes},
		"type":     types.StringType,
	}
	clusterObjectTypes = map[string]attr.Type{
		"clusterName":               types.StringType,
		"clusterStatus":             types.StringType,
		"region":                    types.StringType,
		"cloudformationStackArn":    types.StringType,
		"cloudformationStackStatus": types.StringType,
		"scheduler":                 types.ObjectType{AttrTypes: schedulerObjectTypes},
		"version":                   types.StringType,
	}
)

func NewClusterListDataSource() datasource.DataSource {
	return &ClusterListDataSource{}
}

// ClusterListDataSource defines the data source implementation.
type ClusterListDataSource struct {
	client *openapi.APIClient
	awsv4  openapi.AWSv4
}

// ClusterListDataSourceModel describes the data source data model.
type ClusterListDataSourceModel struct {
	ClusterStatus types.List   `tfsdk:"cluster_status"`
	Clusters      types.List   `tfsdk:"clusters"`
	Region        types.String `tfsdk:"region"`
}

func (d *ClusterListDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_list_clusters"
}

func (d *ClusterListDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Retrieve a list of existing clusters.",

		Attributes: map[string]schema.Attribute{
			"cluster_status": schema.ListAttribute{
				MarkdownDescription: "Filter by cluster status. The default is all clusters.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"clusters": schema.ListAttribute{
				MarkdownDescription: "List of clusters.",
				Optional:            true,
				ElementType: types.ObjectType{
					AttrTypes: clusterObjectTypes,
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The AWS Region of the clusters.",
				Optional:            true,
			},
		},
	}
}

func (d *ClusterListDataSource) Configure(
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

func (d *ClusterListDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data ClusterListDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCtx := context.WithValue(context.Background(), openapi.ContextAWSv4, d.awsv4)

	clusterStatuses := make([]openapi.ClusterStatusFilteringOption, 0)
	for _, s := range data.ClusterStatus.Elements() {
		clusterStatuses = append(clusterStatuses, openapi.ClusterStatusFilteringOption(s.String()))
	}

	clusterListReq := d.client.ClusterOperationsAPI.ListClusters(reqCtx)
	clusterListReq = clusterListReq.ClusterStatus(clusterStatuses).Region(data.Region.ValueString())

	clusterList, _, err := clusterListReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error occured while retrieving cluster list",
			fmt.Sprintf("Error: %v", err),
		)
	}

	goClusterList := make([]types.Object, 0)
	if clusters, ok := clusterList.GetClustersOk(); ok {
		sort.Slice(
			clusters,
			func(i, j int) bool { return clusters[i].GetClusterName() > clusters[j].GetClusterName() },
		)
		for _, cluster := range clusters {
			clusterMap, err := cluster.ToMap()
			if err != nil {
				resp.Diagnostics.AddError(
					"Error occured while retrieving cluster list",
					fmt.Sprintf("Error: %v", err),
				)
			}
			delete(clusterMap, "scheduler")

			tfClusterMap, diags := types.MapValueFrom(ctx, types.StringType, clusterMap)
			resp.Diagnostics.Append(diags...)

			tfClusterMapElements := tfClusterMap.Elements()

			var tfMetadataObject types.Object
			var tfSchedulerObject types.Object
			if scheduler, ok := cluster.GetSchedulerOk(); ok {
				schedulerMap, err := scheduler.ToMap()
				if err != nil {
					resp.Diagnostics.AddError(
						"Error occured while retrieving cluster list",
						fmt.Sprintf("Error: %v", err),
					)
				}
				delete(schedulerMap, "metadata")

				tfSchedulerMap, diags := types.MapValueFrom(ctx, types.StringType, schedulerMap)
				resp.Diagnostics.Append(diags...)
				tfSchedulerMapElements := tfSchedulerMap.Elements()

				if metadata, ok := cluster.Scheduler.GetMetadataOk(); ok {
					metadataMap, err := metadata.ToMap()
					if err != nil {
						resp.Diagnostics.AddError(
							"Error occured while retrieving cluster list",
							fmt.Sprintf("Error: %v", err),
						)
					}
					tfMetadataMap, diags := types.MapValueFrom(ctx, types.StringType, metadataMap)
					resp.Diagnostics.Append(diags...)
					tfMetadataObject, diags = types.ObjectValue(
						metadataObjectTypes,
						tfMetadataMap.Elements(),
					)
					resp.Diagnostics.Append(diags...)

				} else {
					tfMetadataObject = types.ObjectNull(metadataObjectTypes)
				}
				tfSchedulerMapElements["metadata"] = tfMetadataObject
				tfSchedulerObject, diags = types.ObjectValue(
					schedulerObjectTypes,
					tfSchedulerMapElements,
				)
				resp.Diagnostics.Append(diags...)

			} else {
				tfSchedulerObject = types.ObjectNull(schedulerObjectTypes)
			}

			tfClusterMapElements["scheduler"] = tfSchedulerObject
			tfClusterObject, diags := types.ObjectValue(clusterObjectTypes, tfClusterMapElements)
			resp.Diagnostics.Append(diags...)

			goClusterList = append(goClusterList, tfClusterObject)
		}
	}

	tfClusterList, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: clusterObjectTypes},
		goClusterList,
	)
	resp.Diagnostics.Append(diags...)
	data.Clusters = tfClusterList

	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
