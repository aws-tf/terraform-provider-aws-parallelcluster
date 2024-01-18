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
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var (
	_       provider.Provider = &PclusterProvider{}
	apiName                   = "ParallelCluster"
)

type configData struct {
	client     *openapi.APIClient
	awsv4      openapi.AWSv4
	expiration time.Time
	role       string
}

// PclusterProvider defines the provider implementation.
type PclusterProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PclusterProviderModel describes the provider data model.
type PclusterProviderModel struct {
	Endpoint  types.String `tfsdk:"endpoint"`
	RoleArn   types.String `tfsdk:"role_arn"`
	Region    types.String `tfsdk:"region"`
	Profile   types.String `tfsdk:"profile"`
	AwsKey    types.String `tfsdk:"aws_key"`
	AwsSecret types.String `tfsdk:"aws_secret"`
	ApiName   types.String `tfsdk:"api_name"`
}

func (p *PclusterProvider) Metadata(
	ctx context.Context,
	req provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "pcluster"
	resp.Version = p.version
}

func (p *PclusterProvider) Schema(
	ctx context.Context,
	req provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `
AWS ParallelCluster is an AWS supported open source cluster management tool that helps you to deploy and manage high performance computing (HPC) clusters in the AWS Cloud. It automatically sets up the required compute resources, scheduler, and shared filesystem. You can use AWS ParallelCluster with AWS Batch and Slurm schedulers.

With AWS ParallelCluster, you can quickly build and deploy proof of concept and production HPC compute environments. You can also build and deploy a high level workflow on top of AWS ParallelCluster, such as a genomics portal that automates an entire DNA sequencing workflow.
    `,
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The endpoint of the ParallelCluster API. If unset will be autodetected.",
				Optional:            true,
			},
			"role_arn": schema.StringAttribute{
				MarkdownDescription: "The role used for deploying resources and query data sources.",
				Optional:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The region used for deploying resources and query data sources.",
				Optional:            true,
			},
			"profile": schema.StringAttribute{
				MarkdownDescription: "The aws profile used for deploying resources and query data sources.",
				Optional:            true,
			},
			"aws_key": schema.StringAttribute{
				MarkdownDescription: "The aws key used for deploying resources and query data sources.",
				Optional:            true,
			},
			"aws_secret": schema.StringAttribute{
				MarkdownDescription: "The aws secret used for deploying resources and query data sources.",
				Optional:            true,
				Sensitive:           true,
			},
			"api_name": schema.StringAttribute{
				MarkdownDescription: "The name of the ParallelCluster api. Used to retrieve the api endpoint if not given. Defaults to ParallelCluster.",
				Optional:            true,
			},
		},
	}
}

func GetClusterApiUrl(cfg aws.Config, name string) (string, error) {
	var url string

	svc := apigateway.NewFromConfig(cfg)
	in := &apigateway.GetRestApisInput{}
	apisPaginator := apigateway.NewGetRestApisPaginator(svc, in)

	for apisPaginator.HasMorePages() {
		output, err := apisPaginator.NextPage(context.TODO())
		if err != nil {
			return url, err
		}
		for _, api := range output.Items {
			if *api.Name == name {
				host := fmt.Sprintf(
					"%v.execute-api.%s.amazonaws.com",
					*api.Id,
					cfg.Region,
				)
				return fmt.Sprintf("https://%s/prod", host), err
			}
		}
	}

	if len(url) == 0 {
		return url, fmt.Errorf("API endpoint not found.")
	}

	return url, nil
}

func ConfigureAWSv4(cfg aws.Config, role string) (openapi.AWSv4, time.Time, error) {
	var accessKey string
	var secretKey string
	var sessionToken string
	var expiration time.Time

	if role == "" {
		creds, err := cfg.Credentials.Retrieve(context.TODO())
		if err != nil {
			return openapi.AWSv4{}, time.Now(), fmt.Errorf(
				"Failed to retrieve aws credentials. Error: %w",
				err,
			)
		}
		accessKey = creds.AccessKeyID
		secretKey = creds.SecretAccessKey
		sessionToken = creds.SessionToken
		expiration = creds.Expires
	} else {
		stsClient := sts.NewFromConfig(cfg)
		creds, err := stsClient.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
			RoleArn:         aws.String(role),
			RoleSessionName: aws.String("pcluster-terraform-automation"),
		})
		if err != nil {
			return openapi.AWSv4{}, time.Now(), fmt.Errorf(
				"Failed to assume '%s'. Error: %w",
				role,
				err,
			)
		}
		accessKey = *creds.Credentials.AccessKeyId
		secretKey = *creds.Credentials.SecretAccessKey
		sessionToken = *creds.Credentials.SessionToken
		expiration = *creds.Credentials.Expiration
	}

	awsv4 := openapi.AWSv4{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessionToken,
		Region:       cfg.Region,
		Service:      "execute-api",
	}

	return awsv4, expiration, nil
}

func (p *PclusterProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var data PclusterProviderModel
	var sdata configData

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.

	configuration := openapi.NewConfiguration()

	url := data.Endpoint.ValueString()

	if !data.Profile.IsNull() {
		os.Setenv("AWS_PROFILE", data.Profile.ValueString())
	}

	if !data.AwsKey.IsNull() {
		if data.AwsSecret.IsNull() {
			resp.Diagnostics.AddError(
				"aws_key is set but aws_secret is not.",
				"In order to use a aws key you must allso supply a secret using the aws_secret attribute.",
			)
			return
		}

		os.Setenv("AWS_ACCESS_KEY_ID", data.AwsKey.ValueString())
		os.Setenv("AWS_SECRET_ACCESS_KEY", data.AwsSecret.ValueString())
	}

	if !data.Region.IsNull() {
		os.Setenv("AWS_REGION", data.Region.ValueString())
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to load aws default config.",
			fmt.Sprintf("%v", err),
		)
		return
	}

	if !data.ApiName.IsNull() {
		apiName = data.ApiName.ValueString()
	}

	if data.Endpoint.IsNull() {
		apiUrl, err := GetClusterApiUrl(cfg, apiName)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to retrieve ParallelCluster endpoint url.",
				fmt.Sprintf("%v", err),
			)
			return
		}
		url = apiUrl
	}
	configuration.Servers = openapi.ServerConfigurations{
		{
			URL:         url,
			Description: "Parallel Cluster API",
		},
	}
	sdata.client = openapi.NewAPIClient(configuration)

	if data.RoleArn.IsNull() {
		sdata.role = ""
	} else {
		sdata.role = data.RoleArn.ValueString()
	}

	awsv4, expiration, err := ConfigureAWSv4(cfg, sdata.role)
	if err != nil {
		resp.Diagnostics.AddError("Failed to load AWS credentials.", fmt.Sprintf("%v", err))
		return
	}

	sdata.awsv4 = awsv4
	sdata.expiration = expiration

	resp.DataSourceData = sdata
	resp.ResourceData = sdata
}

func (p *PclusterProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClusterResource,
		NewImageResource,
		NewComputeFleetStatusResource,
	}
}

func (p *PclusterProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewClusterListDataSource,
		NewClusterDataSource,
		NewImageListDataSource,
		NewImageDataSource,
		NewComputeFleetDataSource,
		NewOfficialImageDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PclusterProvider{
			version: version,
		}
	}
}
