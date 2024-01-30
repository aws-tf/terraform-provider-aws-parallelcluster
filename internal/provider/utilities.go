package provider

import (
	"context"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type AttributeValidator struct {
	description         string
	markdownDescription string
	validatorFunction   func(context.Context, validator.StringRequest, *validator.StringResponse)
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
