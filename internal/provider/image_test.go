// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	imageId      string
	parentImage  string
	resourceName string
	endpoint     string
	region       string
}

func TestEnd2EndImage(t *testing.T) {
	testConfig := &imageTestConfig{
		role:         os.Getenv("TEST_ROLE"),
		endpoint:     os.Getenv("TEST_ENDPOINT"),
		imageId:      "pcluster-image-build-test-01",
		parentImage:  "ami-0872c164f38dcc49f",
		resourceName: "test",
		region:       os.Getenv("TEST_REGION"),
	}

	if name, ok := os.LookupEnv("TEST_IMAGE_NAME"); ok {
		testConfig.imageId = name
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
				Config: testConfig.testAccImageResourceConfig(),
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
				},
			},
			// Image Data Source Test
			{
				Config: testConfig.testAccImageDataSourceConfig(),
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
				Config: testConfig.testAccImageListDataSourceConfig(),
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

func (c *imageTestConfig) testAccProviderConfig() string {
	var endpoint string
	if c.endpoint == "" {
		endpoint = "null"
	} else {
		endpoint = fmt.Sprintf("%q", c.endpoint)
	}
	return fmt.Sprintf(`
provider "pcluster" {
  role_arn = %q
  endpoint = %v
  region   = %q
}

`, c.role, endpoint, c.region)
}

func (c *imageTestConfig) testAccImageListDataSourceConfig() string {
	dataSource := fmt.Sprintf(`
data "pcluster_list_images" "%v" {
  image_status = "AVAILABLE"
  region = "%v"
}

`, c.resourceName, c.region)

	return dataSource + c.testAccImageResourceConfig()
}

func (c *imageTestConfig) testAccImageDataSourceConfig() string {
	dataSource := fmt.Sprintf(`
data "pcluster_image" "%v" {
  image_id = "%v"
}

`, c.resourceName, c.imageId)

	return dataSource + c.testAccImageResourceConfig()
}

func (c *imageTestConfig) testAccImageResourceConfig() string {
	return c.testAccProviderConfig() + fmt.Sprintf(`
resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}

// Null resource allows us to use built-in test functions above
data "null_data_source" "values" {
  inputs = {
    image_config = yamlencode({
      "Build":{
              "InstanceType": "c5.2xlarge",
              "ParentImage": "%v"
      }
    })
	}
}

resource "pcluster_image" "%v" {
  image_id = "%s"
  image_configuration = data.null_data_source.values.inputs.image_config
}
`, c.parentImage, c.resourceName, c.imageId)
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
		w.Write([]byte(clusterJson))
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
