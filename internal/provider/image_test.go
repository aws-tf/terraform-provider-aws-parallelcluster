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
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"reflect"
	"testing"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type imageTestConfig struct {
	role         string
	useUserRole  string
	apiStackName string
	imageId      string
	resourceName string
	endpoint     string
	region       string
	az           string
}

func TestEnd2EndImage(t *testing.T) {
	testConfig := &imageTestConfig{
		role:         "null",
		useUserRole:  "false",
		apiStackName: "null",
		endpoint:     "null",
		resourceName: "test",
	}

	if region, ok := os.LookupEnv("TEST_REGION"); ok {
		os.Setenv("AWS_REGION", region)
		testConfig.region = region
	} else {
		os.Setenv("AWS_REGION", defaultRegion)
		testConfig.region = defaultRegion
	}

	if az, ok := os.LookupEnv("TEST_AVAILABILITY_ZONE"); ok {
		testConfig.az = az
	} else {
		testConfig.az = defaultAz
	}

	if name, ok := os.LookupEnv("TEST_IMAGE_NAME"); ok {
		testConfig.imageId = name
	} else {
		testConfig.imageId = "pcluster-image-build-test-01"
	}

	if endpoint, ok := os.LookupEnv("TEST_ENDPOINT"); ok {
		testConfig.endpoint = fmt.Sprintf("%q", endpoint)
	}

	if role, ok := os.LookupEnv("TEST_ROLE"); ok {
		testConfig.role = fmt.Sprintf("%q", role)
	}

	if _, ok := os.LookupEnv("TEST_USE_USER_ROLE"); ok {
		testConfig.useUserRole = "true"
	}

	if name, ok := os.LookupEnv("TEST_PCAPI_STACK_NAME"); ok {
		testConfig.apiStackName = fmt.Sprintf("%q", name)
	}

	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"aws":   {Source: "hashicorp/aws", VersionConstraint: "~> 5.0"},
			"tls":   {Source: "hashicorp/tls", VersionConstraint: ">= 3.4"},
			"local": {Source: "hashicorp/local"},
			"null":  {Source: "hashicorp/null"},
		},
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testConfig.imageResourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"pcluster_image."+testConfig.resourceName,
						"image_configuration",
						"data.null_data_source.values",
						"inputs.image_config",
					),
					resource.TestCheckResourceAttr(
						"pcluster_image."+testConfig.resourceName,
						"image_build_status",
						string(openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE),
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "pcluster_image." + testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"cloudformation_stack_arn",
					"cloudformation_stack_status",
					"rollback_on_failure",
				},
			},
			// Image Data Source Test
			{
				Config: testConfig.imageDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.pcluster_image."+testConfig.resourceName,
						"image.imageId",
						testConfig.imageId,
					),
					resource.TestCheckResourceAttr(
						"pcluster_image."+testConfig.resourceName,
						"image_build_status",
						string(openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE),
					),
					resource.TestCheckResourceAttrPair(
						"data.pcluster_image."+testConfig.resourceName,
						"image.imageConfiguration",
						"pcluster_image."+testConfig.resourceName,
						"image_configuration",
					),
					resource.TestCheckResourceAttrSet(
						"data.pcluster_image."+testConfig.resourceName,
						"stack_events.#",
					),
					resource.TestCheckResourceAttrSet(
						"data.pcluster_image."+testConfig.resourceName,
						"log_streams.#",
					),
				),
			},
			// Image List Data Source Test
			{
				Config: testConfig.imageListDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.pcluster_list_images."+testConfig.resourceName,
						"images.#",
					),
				),
			},
		},
	})
}

func (c *imageTestConfig) providerConfig() string {
	return fmt.Sprintf(`
provider "pcluster" {
  role_arn = %s
  use_user_role = %s
  api_stack_name = %s
  endpoint = %s
  region   = "%s"
}

`, c.role, c.useUserRole, c.apiStackName, c.endpoint, c.region)
}

func (c *imageTestConfig) imageListDataSourceConfig() string {
	dataSource := fmt.Sprintf(`
data "pcluster_list_images" "%s" {
  image_status = "AVAILABLE"
  region = "%s"
}

`, c.resourceName, c.region)

	return dataSource + c.imageResourceConfig()
}

func (c *imageTestConfig) imageDataSourceConfig() string {
	dataSource := fmt.Sprintf(`
data "pcluster_image" "%s" {
  image_id = "%s"
}

`, c.resourceName, c.imageId)

	return dataSource + c.imageResourceConfig()
}

func (c *imageTestConfig) imageResourceConfig() string {
	return c.providerConfig() + fmt.Sprintf(`
locals {
	default_region = "%s"
	default_az = "%s"
}

resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}

resource "aws_default_subnet" "public_az1" {
  availability_zone = local.default_az
}

resource "aws_default_security_group" "default" {
  vpc_id = aws_default_vpc.default.id

  ingress {
    from_port        = 0
    to_port          = 0
    protocol         = -1
    self			 = true
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = -1
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

data "pcluster_list_official_images" "parent_image" {
        region = local.default_region
        os = "alinux2"
        architecture = "x86_64"
}

// Null resource allows us to use built-in test functions above
data "null_data_source" "values" {
  inputs = {
    image_config = yamlencode({
      "Build":{
              "InstanceType": "c5.2xlarge",
              "ParentImage": data.pcluster_list_official_images.parent_image.official_images[0].amiId,
			  "SubnetId": aws_default_subnet.public_az1.id,
              "SecurityGroupIds": [aws_default_security_group.default.id],
              "UpdateOsPackages": {"Enabled": false}
      }
    })
	}
}

resource "pcluster_image" "%s" {
  image_id = "%s"
  image_configuration = data.null_data_source.values.inputs.image_config
  rollback_on_failure = false
}
`, c.region, c.az, c.resourceName, c.imageId)
}

func TestUnitWaitImageReady(t *testing.T) {
	names := map[string]string{
		"failed":         "image01",
		"complete":       "image02",
		"deleteFailed":   "image03",
		"deleteComplete": "image04",
		"wait":           "image05",
		"waitDelete":     "image06",
		"InvalidStatus":  "image07",
	}

	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var status openapi.ImageBuildStatus
		switch name := path.Base(r.URL.Path); name {
		case names["failed"]:
			status = openapi.IMAGEBUILDSTATUS_BUILD_FAILED
		case names["complete"]:
			status = openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE
		case names["deleteFailed"]:
			status = openapi.IMAGEBUILDSTATUS_DELETE_FAILED
		case names["deleteComplete"]:
			status = openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE
		case names["wait"]:
			if count > 2 {
				status = openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE
				count = 0
			} else {
				status = openapi.IMAGEBUILDSTATUS_BUILD_IN_PROGRESS
			}
			count++
		case names["waitDelete"]:
			if count > 2 {
				status = openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE
				count = 0
			} else {
				status = openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE
			}
			count++
		case names["InvalidStatus"]:
			status = "InvalidStatus"
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		clusterJson, _ := openapi.DescribeImageResponseContent{
			ImageId:          path.Base(r.URL.Path),
			ImageBuildStatus: status,
		}.MarshalJSON()
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(clusterJson))
		if err != nil {
			t.Fatalf("Failed to mock http request %v", err)
		}
	}))

	defer server.Close()
	cR := ImageResource{}
	c := openapi.NewConfiguration()

	c.Servers = openapi.ServerConfigurations{
		{
			URL:         server.URL,
			Description: "Parallel Cluster API",
		},
	}
	cR.client = openapi.NewAPIClient(c)

	cases := []struct {
		name     string
		expected *openapi.DescribeImageResponseContent
	}{
		{
			name: names["failed"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["failed"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_BUILD_FAILED,
			},
		},
		{
			name: names["complete"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["complete"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
			},
		},
		{
			name: names["deleteFailed"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["deleteFailed"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_DELETE_FAILED,
			},
		},
		{
			name: names["deleteComplete"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["deleteComplete"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE,
			},
		},
		{
			name: names["wait"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["wait"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_BUILD_COMPLETE,
			},
		},
		{
			name: names["waitDelete"],
			expected: &openapi.DescribeImageResponseContent{
				ImageId:          names["waitDelete"],
				ImageBuildStatus: openapi.IMAGEBUILDSTATUS_DELETE_COMPLETE,
			},
		},
		{
			name:     names["notFound"],
			expected: nil,
		},
		{
			name:     names["InvalidStatus"],
			expected: nil,
		},
	}

	for _, c := range cases {
		out, err := cR.waitImageReady(context.TODO(), c.name)
		if err != nil {
			if c.name == names["NotFound"] {
				if err.Error() != failedToFindImageErr {
					t.Fatalf("Expected to see: %v", failedToFindClusterErr)
				}
				continue
			}
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v\nErr:%v",
				out,
				c.expected,
				err,
			)
		}
	}
}
