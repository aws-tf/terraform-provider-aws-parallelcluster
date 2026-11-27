// Copyright 2026 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// dropConnection simulates a network failure by closing the connection without a response.
func dropConnection(w http.ResponseWriter, _ *http.Request) {
	if conn, _, err := w.(http.Hijacker).Hijack(); err == nil {
		_ = conn.Close()
	}
}

func httpStatus(code int) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_, _ = fmt.Fprintf(w, `{"message":"%s"}`, http.StatusText(code))
	}
}

func TestUnitPCAPIHTTPClient(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		method    string
		fail      http.HandlerFunc
		wantCode  int
		wantCalls int32
	}{
		{"GET retried on dropped connection", http.MethodGet, dropConnection, http.StatusOK, 2},
		{"GET retried on 503", http.MethodGet, httpStatus(http.StatusServiceUnavailable), http.StatusOK, 2},
		{"POST retried on 429, body re-sent", http.MethodPost, httpStatus(http.StatusTooManyRequests), http.StatusOK, 2},
		{"404 not retried", http.MethodGet, httpStatus(http.StatusNotFound), http.StatusNotFound, 1},
		{"400 not retried", http.MethodPost, httpStatus(http.StatusBadRequest), http.StatusBadRequest, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if calls.Add(1) == 1 {
					c.fail(w, r)
					return
				}
				b, _ := io.ReadAll(r.Body)
				_, _ = w.Write(b)
			}))
			defer server.Close()

			req, _ := http.NewRequest(c.method, server.URL, strings.NewReader("payload"))
			resp, err := newPCAPIHTTPClient().Do(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			body, _ := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if resp.StatusCode != c.wantCode {
				t.Fatalf("expected status %d, got %d", c.wantCode, resp.StatusCode)
			}
			if c.wantCode == http.StatusOK && string(body) != "payload" {
				t.Fatalf("expected request body to be re-sent, got %q", body)
			}
			if calls.Load() != c.wantCalls {
				t.Fatalf("expected %d calls, got %d", c.wantCalls, calls.Load())
			}
		})
	}
}

// statusServer drops the connection of the first 3 GETs to `path` (more than the single silent retry of Go's
// transport), then returns `inProgress` once and `final` afterwards.
// POST to `createPath` returns `created`.
func statusServer(path string, describe func(status string) []byte, inProgress, final, createPath string, created []byte) *httptest.Server {
	var mu sync.Mutex
	gets := 0
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost && r.URL.Path == createPath {
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write(created)
			return
		}
		if r.Method != http.MethodGet || r.URL.Path != path {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		mu.Lock()
		gets++
		g := gets
		mu.Unlock()
		switch {
		case g <= 3:
			dropConnection(w, r)
		case g == 4:
			_, _ = w.Write(describe(inProgress))
		default:
			_, _ = w.Write(describe(final))
		}
	}))
}

func pcapiClient(server *httptest.Server) *openapi.APIClient {
	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{{URL: server.URL}}
	cfg.HTTPClient = newPCAPIHTTPClient()
	return openapi.NewAPIClient(cfg)
}

// nullPlan returns a plan for s with all attributes null except the given ones.
func nullPlan(ctx context.Context, s schema.Schema, values map[string]tftypes.Value) tfsdk.Plan {
	objType := s.Type().TerraformType(ctx).(tftypes.Object)
	attrs := map[string]tftypes.Value{}
	for name, typ := range objType.AttributeTypes {
		if v, ok := values[name]; ok {
			attrs[name] = v
		} else {
			attrs[name] = tftypes.NewValue(typ, nil)
		}
	}
	return tfsdk.Plan{Raw: tftypes.NewValue(objType, attrs), Schema: s}
}

// Reproduces V2412031283: status polls fail while CloudFormation keeps creating the cluster.
// Create must not return an error, otherwise Terraform taints the healthy cluster.
func TestUnitClusterCreateNotTaintedOnNetworkErrors(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	name := "RecursiveAICluster3"
	created, _ := openapi.CreateClusterResponseContent{Cluster: openapi.ClusterInfoSummary{
		ClusterName:               name,
		ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_IN_PROGRESS,
		CloudformationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_IN_PROGRESS,
	}}.MarshalJSON()
	describe := func(status string) []byte {
		b, _ := openapi.DescribeClusterResponseContent{
			ClusterName:               name,
			ClusterStatus:             openapi.ClusterStatus(status),
			ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
			CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		}.MarshalJSON()
		return b
	}
	server := statusServer("/v3/clusters/"+name, describe,
		string(openapi.CLUSTERSTATUS_CREATE_IN_PROGRESS), string(openapi.CLUSTERSTATUS_CREATE_COMPLETE),
		"/v3/clusters", created)
	defer server.Close()
	r := ClusterResource{awsv4: awsv4Test(), client: pcapiClient(server)}

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	req := resource.CreateRequest{Plan: nullPlan(ctx, schemaResp.Schema, map[string]tftypes.Value{
		"cluster_name":          tftypes.NewValue(tftypes.String, name),
		"cluster_configuration": tftypes.NewValue(tftypes.String, "some_config"),
	})}
	resp := resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Create(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Create returned an error, so Terraform would taint the cluster: %v", resp.Diagnostics.Errors())
	}
	var state ClusterResourceModel
	resp.State.Get(ctx, &state)
	if state.ClusterStatus != types.StringValue(string(openapi.CLUSTERSTATUS_CREATE_COMPLETE)) {
		t.Fatalf("expected cluster_status CREATE_COMPLETE, got %v", state.ClusterStatus)
	}
}

// Same as above, for images.
func TestUnitImageCreateNotTaintedOnNetworkErrors(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := "flaky-image"
	created, _ := openapi.BuildImageResponseContent{Image: openapi.ImageInfoSummary{
		ImageId:          id,
		ImageBuildStatus: openapi.IMAGEBUILDSTATUS_BUILD_IN_PROGRESS,
	}}.MarshalJSON()
	describe := func(status string) []byte {
		b, _ := openapi.DescribeImageResponseContent{ImageId: id, ImageBuildStatus: openapi.ImageBuildStatus(status)}.MarshalJSON()
		return b
	}
	server := statusServer("/v3/images/custom/"+id, describe,
		string(openapi.IMAGEBUILDSTATUS_BUILD_IN_PROGRESS), string(openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE),
		"/v3/images/custom", created)
	defer server.Close()
	r := ImageResource{awsv4: awsv4Test(), client: pcapiClient(server)}

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	req := resource.CreateRequest{Plan: nullPlan(ctx, schemaResp.Schema, map[string]tftypes.Value{
		"image_id":            tftypes.NewValue(tftypes.String, id),
		"image_configuration": tftypes.NewValue(tftypes.String, "some_config"),
	})}
	resp := resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}

	r.Create(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Create returned an error, so Terraform would taint the image: %v", resp.Diagnostics.Errors())
	}
	var state ImageResourceModel
	resp.State.Get(ctx, &state)
	if state.ImageBuildStatus != types.StringValue(string(openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE)) {
		t.Fatalf("expected image_build_status BUILD_COMPLETE, got %v", state.ImageBuildStatus)
	}
}
