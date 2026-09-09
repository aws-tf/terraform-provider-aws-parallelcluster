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
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
)

func TestUnitRefreshAWSv4(t *testing.T) {
	t.Parallel()

	current := awsv4Test()

	// Expiration well outside the refresh window returns current credentials
	// unchanged.
	notNearExpiry := time.Now().Add(time.Hour)
	out, outExpiration := refreshAWSv4(aws.Config{}, "", current, notNearExpiry)
	if !reflect.DeepEqual(out, current) {
		t.Fatalf("Error matching output and expected. \nO: %#v\nE: %#v", out, current)
	}
	if !outExpiration.Equal(notNearExpiry) {
		t.Fatalf(
			"Error matching expiration and expected. \nO: %#v\nE: %#v",
			outExpiration,
			notNearExpiry,
		)
	}

	// Expiration within the refresh window refreshes the credentials. Static
	// credentials with an empty role let the refresh resolve without a network call.
	staticCfg := aws.Config{
		Region: "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(
			"refreshedKey",
			"refreshedSecret",
			"refreshedToken",
		),
	}
	expectedRefresh := openapi.AWSv4{
		AccessKey:    "refreshedKey",
		SecretKey:    "refreshedSecret",
		SessionToken: "refreshedToken",
		Region:       "us-east-1",
		Service:      "execute-api",
	}
	out, _ = refreshAWSv4(staticCfg, "", current, time.Now().Add(time.Minute))
	if !reflect.DeepEqual(out, expectedRefresh) {
		t.Fatalf(
			"Error matching output and expected. \nO: %#v\nE: %#v",
			out,
			expectedRefresh,
		)
	}

	// When credentials are within the refresh window but re-deriving them fails,
	// the current (stale) credentials and expiration are returned unchanged so
	// callers can still make a best-effort request.
	failingCfg := aws.Config{
		Region: "us-east-1",
		Credentials: aws.CredentialsProviderFunc(
			func(context.Context) (aws.Credentials, error) {
				return aws.Credentials{}, errors.New("credential retrieval failed")
			},
		),
	}
	nearExpiry := time.Now().Add(time.Minute)
	out, outExpiration = refreshAWSv4(failingCfg, "", current, nearExpiry)
	if !reflect.DeepEqual(out, current) {
		t.Fatalf(
			"Error expected current credentials on refresh failure. \nO: %#v\nE: %#v",
			out,
			current,
		)
	}
	if !outExpiration.Equal(nearExpiry) {
		t.Fatalf(
			"Error expected unchanged expiration on refresh failure. \nO: %#v\nE: %#v",
			outExpiration,
			nearExpiry,
		)
	}
}

func TestUnitContextWithAWSv4(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Empty credentials leave the context unchanged so the signer is skipped,
	// preserving pre-refresh unit-test behavior.
	emptyOut := contextWithAWSv4(ctx, openapi.AWSv4{})
	if emptyOut != ctx {
		t.Fatalf("Error expected unchanged context for empty credentials.")
	}
	if value := emptyOut.Value(openapi.ContextAWSv4); value != nil {
		t.Fatalf("Error expected no AWSv4 context value. \nO: %#v", value)
	}

	// Non-empty credentials are attached to the context for signing.
	awsv4 := awsv4Test()
	out := contextWithAWSv4(ctx, awsv4)
	got, ok := out.Value(openapi.ContextAWSv4).(openapi.AWSv4)
	if !ok {
		t.Fatalf("Error expected AWSv4 value in context. \nO: %#v", out.Value(openapi.ContextAWSv4))
	}
	if !reflect.DeepEqual(got, awsv4) {
		t.Fatalf("Error matching output and expected. \nO: %#v\nE: %#v", got, awsv4)
	}
}
