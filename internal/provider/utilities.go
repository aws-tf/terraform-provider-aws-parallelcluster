package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type AttributeValidator struct {
	description         string
	markdownDescription string
	validatorFunction   func(context.Context, validator.StringRequest, *validator.StringResponse)
}

type resourceConfigurable interface {
	getClient() *openapi.APIClient
	getAWSv4() openapi.AWSv4
	Configure(context.Context, resource.ConfigureRequest, *resource.ConfigureResponse)
}

type dataConfigurable interface {
	getClient() *openapi.APIClient
	getAWSv4() openapi.AWSv4
	Configure(context.Context, datasource.ConfigureRequest, *datasource.ConfigureResponse)
}

type jsonable interface {
	MarshalJSON() ([]byte, error)
}

func (m *AttributeValidator) Description(ctx context.Context) string {
	return m.description
}

func (m *AttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.markdownDescription
}

func (f *AttributeValidator) ValidateString(
	ctx context.Context,
	req validator.StringRequest,
	resp *validator.StringResponse,
) {
	f.validatorFunction(ctx, req, resp)
}

func awsv4Test() openapi.AWSv4 {
	return openapi.AWSv4{
		AccessKey:    "testKey",
		SecretKey:    "testSecret",
		SessionToken: "testToken",
		Service:      "testService",
	}
}

func mockJsonWithTextServer(paths []string, outputs ...string) (*httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, b := range outputs {
			if r.URL.Path == "/"+paths[i] {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(b))
				return
			}
		}
	},
	))

	return server, nil
}

func mockJsonServer(paths []string, jsons ...jsonable) (*httptest.Server, error) {
	var respJSONs [][]byte

	for _, j := range jsons {
		respJSON, err := j.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal list clusters response JSON.")
		}
		respJSONs = append(respJSONs, respJSON)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, j := range respJSONs {
			if r.URL.Path == "/v3/"+paths[i] {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(j)
				return
			}
		}
	},
	))

	return server, nil
}

func standardResourceConfigureTests(d resourceConfigurable) error {
	resp := resource.ConfigureResponse{}
	req := resource.ConfigureRequest{}

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: "testURL",
		},
	}

	awsv4 := awsv4Test()

	d.Configure(context.TODO(), req, &resp)
	if resp.Diagnostics.HasError() {
		return fmt.Errorf("Not expecting error when configuring without provider data.")
	}

	if d.getClient() != nil {
		return fmt.Errorf("Error: Client should not be set when provider data is not set.")
	}

	req.ProviderData = configData{
		awsv4:  awsv4,
		client: openapi.NewAPIClient(cfg),
	}

	d.Configure(context.TODO(), req, &resp)

	if d.getClient() == nil {
		return fmt.Errorf("Error client expected to be set.")
	}

	if d.getAWSv4() != awsv4 {
		return fmt.Errorf("Error matching output expected. O: %#v\nE: %#v",
			d.getAWSv4(),
			awsv4,
		)
	}

	req.ProviderData = "Some invalid data"
	d.Configure(context.TODO(), req, &resp)
	if !resp.Diagnostics.HasError() {
		return fmt.Errorf("Expecting error when configuring with invalid data.")
	}

	return nil
}

func standardDataConfigureTests(d dataConfigurable) error {
	resp := datasource.ConfigureResponse{}
	req := datasource.ConfigureRequest{}

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		openapi.ServerConfiguration{
			URL: "testURL",
		},
	}

	awsv4 := awsv4Test()

	d.Configure(context.TODO(), req, &resp)
	if resp.Diagnostics.HasError() {
		return fmt.Errorf("Not expecting error when configuring without provider data.")
	}

	if d.getClient() != nil {
		return fmt.Errorf("Error: Client should not be set when provider data is not set.")
	}

	req.ProviderData = configData{
		awsv4:  awsv4,
		client: openapi.NewAPIClient(cfg),
	}

	d.Configure(context.TODO(), req, &resp)

	if d.getClient() == nil {
		return fmt.Errorf("Error client expected to be set.")
	}

	if d.getAWSv4() != awsv4 {
		return fmt.Errorf("Error matching output expected. O: %#v\nE: %#v",
			d.getAWSv4(),
			awsv4,
		)
	}

	req.ProviderData = "Some invalid data"
	d.Configure(context.TODO(), req, &resp)
	if !resp.Diagnostics.HasError() {
		return fmt.Errorf("Expecting error when configuring with invalid data.")
	}

	return nil
}
