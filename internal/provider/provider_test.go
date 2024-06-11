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
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

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

func TestUnitPclusterProviderConfigure(t *testing.T) {
	t.Skip()
}
