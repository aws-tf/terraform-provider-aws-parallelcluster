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
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"aws-parallelcluster": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	if _, ok := os.LookupEnv("TEST_REGION"); ok {
		if _, ok := os.LookupEnv("TEST_AVAILABILITY_ZONE"); !ok {
			t.Fatal(`
          Environment 'TEST_AVAILABILITY_ZONE' is not set.
          IF TEST_REGION is set TEST_AVAILABILITY_ZONE must also be set.
        `)
		}
	}
}

func TestUnitNewPclusterProvider(t *testing.T) {
	t.Parallel()

	obj := New("test")()
	if _, ok := obj.(*PclusterProvider); !ok {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: PclusterProvider",
			reflect.TypeOf(obj),
		)
	}
}

func TestUnitPclusterProviderMetadata(t *testing.T) {
	d := New("test")()
	p := PclusterProvider{}
	resp := provider.MetadataResponse{}
	req := provider.MetadataRequest{}
	providerResp := provider.MetadataResponse{}
	providerReq := provider.MetadataRequest{}

	p.Metadata(context.TODO(), providerReq, &providerResp)

	d.Metadata(context.TODO(), req, &resp)

	if !strings.HasPrefix(resp.TypeName, providerResp.TypeName) {
		t.Fatalf(
			"Error provider typename expected as the prefix for provider or provider name. \nO: %#v\nE: %#v",
			resp.TypeName,
			providerResp.TypeName,
		)
	}
}

func TestUnitPclusterProviderSchema(t *testing.T) {
	d := PclusterProvider{}
	dataSourceModel := PclusterProviderModel{}
	resp := provider.SchemaResponse{}
	req := provider.SchemaRequest{}

	d.Schema(context.TODO(), req, &resp)

	rProvider := reflect.TypeOf(dataSourceModel)
	numFields := rProvider.NumField()
	numAttributes := len(resp.Schema.Attributes)

	for i := 0; i < numFields; i++ {
		tag := rProvider.Field(i).Tag
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

// Not parallel: points the AWS SDK at a local fake of CloudFormation and STS through environment variables.
func TestUnitPclusterProviderConfigure(t *testing.T) {
	const (
		stackUrl  = "https://stack.execute-api.us-east-1.amazonaws.com/prod"
		stackRole = "arn:aws:iam::123456789012:role/StackRole"
		knownUrl  = "https://known.execute-api.us-east-1.amazonaws.com/prod"
	)
	stackExists := true
	assumedRole := ""
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		w.Header().Set("Content-Type", "text/xml")
		switch r.Form.Get("Action") {
		case "DescribeStacks":
			if !stackExists {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`<ErrorResponse><Error><Type>Sender</Type><Code>ValidationError</Code>` +
					`<Message>Stack does not exist</Message></Error></ErrorResponse>`))
				return
			}
			_, _ = w.Write([]byte(`<DescribeStacksResponse><DescribeStacksResult><Stacks><member>` +
				`<StackName>MyStack</StackName><Outputs>` +
				`<member><OutputKey>ParallelClusterApiInvokeUrl</OutputKey><OutputValue>` + stackUrl + `</OutputValue></member>` +
				`<member><OutputKey>ParallelClusterApiUserRole</OutputKey><OutputValue>` + stackRole + `</OutputValue></member>` +
				`</Outputs></member></Stacks></DescribeStacksResult></DescribeStacksResponse>`))
		case "AssumeRole":
			assumedRole = r.Form.Get("RoleArn")
			_, _ = w.Write([]byte(`<AssumeRoleResponse><AssumeRoleResult><Credentials>` +
				`<AccessKeyId>assumedKey</AccessKeyId><SecretAccessKey>assumedSecret</SecretAccessKey>` +
				`<SessionToken>assumedToken</SessionToken><Expiration>2030-01-01T00:00:00Z</Expiration>` +
				`</Credentials></AssumeRoleResult></AssumeRoleResponse>`))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}))
	t.Cleanup(server.Close)

	t.Setenv("AWS_ENDPOINT_URL", server.URL)
	t.Setenv("AWS_ACCESS_KEY_ID", "testKey")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "testSecret")
	t.Setenv("AWS_REGION", "us-east-1")
	t.Setenv("AWS_PROFILE", "")
	t.Setenv("AWS_CONFIG_FILE", os.DevNull)
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", os.DevNull)

	unknown := tftypes.NewValue(tftypes.String, tftypes.UnknownValue)
	known := func(v string) tftypes.Value { return tftypes.NewValue(tftypes.String, v) }

	tests := []struct {
		name            string
		attrs           map[string]tftypes.Value
		stackExists     bool
		expectError     bool
		expectUrl       string // Empty when API calls are expected to fail with unknownEndpointTransport.
		expectAssumed   string
		expectAccessKey string
	}{
		{
			name:            "unknown endpoint and role are read from the stack",
			attrs:           map[string]tftypes.Value{"endpoint": unknown, "role_arn": unknown, "api_stack_name": known("MyStack")},
			stackExists:     true,
			expectUrl:       stackUrl,
			expectAssumed:   stackRole,
			expectAccessKey: "assumedKey",
		},
		{
			name:            "unknown endpoint tolerates a missing stack",
			attrs:           map[string]tftypes.Value{"endpoint": unknown, "api_stack_name": known("MyStack")},
			expectAccessKey: "testKey",
		},
		{
			name:        "known endpoint does not tolerate a missing stack",
			attrs:       map[string]tftypes.Value{"endpoint": known(knownUrl), "api_stack_name": known("MyStack")},
			expectError: true,
		},
		{
			name:            "known endpoint is used as is",
			attrs:           map[string]tftypes.Value{"endpoint": known(knownUrl), "api_stack_name": known("MyStack")},
			stackExists:     true,
			expectUrl:       knownUrl,
			expectAccessKey: "testKey",
		},
	}

	p := &PclusterProvider{}
	schemaResp := provider.SchemaResponse{}
	p.Schema(context.TODO(), provider.SchemaRequest{}, &schemaResp)
	objType := schemaResp.Schema.Type().TerraformType(context.TODO()).(tftypes.Object)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackExists, assumedRole = tt.stackExists, ""

			values := map[string]tftypes.Value{}
			for name, attrType := range objType.AttributeTypes {
				values[name] = tftypes.NewValue(attrType, nil)
			}
			for name, v := range tt.attrs {
				values[name] = v
			}
			req := provider.ConfigureRequest{Config: tfsdk.Config{
				Schema: schemaResp.Schema,
				Raw:    tftypes.NewValue(objType, values),
			}}
			resp := provider.ConfigureResponse{}
			p.Configure(context.TODO(), req, &resp)

			if resp.Diagnostics.HasError() != tt.expectError {
				t.Fatalf("Unexpected diagnostics. \nO: %v", resp.Diagnostics)
			}
			if tt.expectError {
				return
			}

			data := resp.ResourceData.(configData)
			cfg := data.client.GetConfig()
			_, failsApiCalls := cfg.HTTPClient.Transport.(unknownEndpointTransport)
			if failsApiCalls != (tt.expectUrl == "") {
				t.Fatalf("Unexpected HTTP transport. \nO: %T", cfg.HTTPClient.Transport)
			}
			if url := cfg.Servers[0].URL; url != tt.expectUrl {
				t.Fatalf("Error matching url. \nO: %v\nE: %v", url, tt.expectUrl)
			}
			if assumedRole != tt.expectAssumed {
				t.Fatalf("Error matching assumed role. \nO: %v\nE: %v", assumedRole, tt.expectAssumed)
			}
			if data.awsv4.AccessKey != tt.expectAccessKey {
				t.Fatalf("Error matching access key. \nO: %v\nE: %v", data.awsv4.AccessKey, tt.expectAccessKey)
			}
		})
	}
}

func TestUnitConfigureAWSv4(t *testing.T) {
	t.Parallel()

	// When the initial credential configuration fails (e.g. credentials cannot
	// be retrieved), ConfigureAWSv4 surfaces the error so the provider's
	// Configure can report it instead of proceeding with invalid credentials.
	failingCfg := aws.Config{
		Region: "us-east-1",
		Credentials: aws.CredentialsProviderFunc(
			func(context.Context) (aws.Credentials, error) {
				return aws.Credentials{}, errors.New("credential retrieval failed")
			},
		),
	}
	_, _, err := ConfigureAWSv4(failingCfg, "")
	if err == nil {
		t.Fatalf("Error expected when initial credential configuration fails.")
	}
	if !strings.Contains(err.Error(), "failed to retrieve aws credentials") {
		t.Fatalf("Error message did not describe the failure. \nO: %v", err)
	}

	// When credentials resolve successfully, ConfigureAWSv4 returns populated
	// SigV4 credentials and no error.
	staticCfg := aws.Config{
		Region: "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(
			"testKey",
			"testSecret",
			"testToken",
		),
	}
	awsv4, _, err := ConfigureAWSv4(staticCfg, "")
	if err != nil {
		t.Fatalf("Not expecting error on successful credential configuration. \nO: %v", err)
	}
	expected := openapi.AWSv4{
		AccessKey:    "testKey",
		SecretKey:    "testSecret",
		SessionToken: "testToken",
		Region:       "us-east-1",
		Service:      "execute-api",
	}
	if !reflect.DeepEqual(awsv4, expected) {
		t.Fatalf("Error matching output and expected. \nO: %#v\nE: %#v", awsv4, expected)
	}
}

func TestUnitUnknownEndpointTransport(t *testing.T) {
	t.Parallel()

	_, err := (&http.Client{Transport: unknownEndpointTransport{}}).Get("/v3/clusters/test")
	if err == nil || !strings.Contains(err.Error(), "its endpoint is not known yet") {
		t.Fatalf("Error message did not explain the unknown endpoint. \nO: %v", err)
	}
}
