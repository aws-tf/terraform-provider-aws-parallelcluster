package provider

import (
	"context"

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
