package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// awsv4RefreshWindow is how long before expiration we proactively refresh
// credentials, so that requests are always signed with credentials that have
// enough remaining lifetime to complete, rather than ones about to expire.
const awsv4RefreshWindow = 5 * time.Minute

// refreshAWSv4 returns freshly-derived SigV4 credentials when the current ones
// are at or near expiration, otherwise it returns the current credentials
// unchanged.
//
// Refresh is skipped when expiration is the zero value (credentials supplied
// directly, e.g. in unit tests) so existing behavior is preserved. When a
// refresh is attempted but fails, the current (possibly stale) credentials and
// expiration are returned so callers can still make a best-effort request and
// surface the original error.
func refreshAWSv4(
	cfg aws.Config,
	role string,
	current openapi.AWSv4,
	expiration time.Time,
) (openapi.AWSv4, time.Time) {
	if expiration.IsZero() {
		return current, expiration
	}
	if time.Now().Add(awsv4RefreshWindow).Before(expiration) {
		return current, expiration
	}

	awsv4, newExpiration, err := ConfigureAWSv4(cfg, role)
	if err != nil {
		return current, expiration
	}
	return awsv4, newExpiration
}

// contextWithAWSv4 returns ctx with the given SigV4 credentials attached for
// request signing. When the credentials are empty (e.g. unit tests that do not
// configure credentials), ctx is returned unchanged so the signer is skipped,
// preserving the pre-refresh behavior.
func contextWithAWSv4(ctx context.Context, awsv4 openapi.AWSv4) context.Context {
	if awsv4.AccessKey == "" && awsv4.SecretKey == "" && awsv4.SessionToken == "" {
		return ctx
	}
	return context.WithValue(ctx, openapi.ContextAWSv4, awsv4)
}

type mockCfg struct {
	out         jsonable
	outText     string
	path        string
	method      string
	useJsonable bool
	httpError   int
}

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

func mockJsonServer(mocks ...mockCfg) (*httptest.Server, error) {
	for _, m := range mocks {
		var err error
		if m.useJsonable || m.out != nil {
			_, err = m.out.MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed to marshal list clusters response JSON")
			}
		}
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mocks {
			if m.method == "" {
				m.method = http.MethodGet
			}
			if r.URL.Path == "/v3/"+m.path && r.Method == m.method {
				w.Header().Set("Content-Type", "application/json")
				if m.httpError != 0 {
					w.WriteHeader(m.httpError)
				}
				if m.useJsonable || m.out != nil {
					j, err := m.out.MarshalJSON()
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					_, _ = w.Write(j)
					return
				} else {
					_, _ = w.Write([]byte(m.outText))
				}
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
		return fmt.Errorf("not expecting error when configuring without provider data")
	}

	if d.getClient() != nil {
		return fmt.Errorf("client should not be set when provider data is not set")
	}

	req.ProviderData = configData{
		awsv4:  awsv4,
		client: openapi.NewAPIClient(cfg),
	}

	d.Configure(context.TODO(), req, &resp)

	if d.getClient() == nil {
		return fmt.Errorf("client expected to be set")
	}

	if d.getAWSv4() != awsv4 {
		return fmt.Errorf("error matching output expected. O: %#v\nE: %#v",
			d.getAWSv4(),
			awsv4,
		)
	}

	req.ProviderData = "Some invalid data"
	d.Configure(context.TODO(), req, &resp)
	if !resp.Diagnostics.HasError() {
		return fmt.Errorf("expecting error when configuring with invalid data")
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
		return fmt.Errorf("not expecting error when configuring without provider data")
	}

	if d.getClient() != nil {
		return fmt.Errorf("client should not be set when provider data is not set")
	}

	req.ProviderData = configData{
		awsv4:  awsv4,
		client: openapi.NewAPIClient(cfg),
	}

	d.Configure(context.TODO(), req, &resp)

	if d.getClient() == nil {
		return fmt.Errorf("client expected to be set")
	}

	if d.getAWSv4() != awsv4 {
		return fmt.Errorf("error matching output expected. O: %#v\nE: %#v",
			d.getAWSv4(),
			awsv4,
		)
	}

	req.ProviderData = "Some invalid data"
	d.Configure(context.TODO(), req, &resp)
	if !resp.Diagnostics.HasError() {
		return fmt.Errorf("expecting error when configuring with invalid data")
	}

	return nil
}
