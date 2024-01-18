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
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	openapi "github.com/aws-tf/terraform-provider-aws-parallelcluster/internal/provider/openapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	testResourceName      = "test"
	configDirectory       = "files/cluster_test"
	configImportDirectory = "files/cluster_import_test"
)

func TestEnd2EndCluster(t *testing.T) {
	clusterName := os.Getenv("TEST_CLUSTER_NAME")
	region := os.Getenv("TEST_REGION")
	os.Setenv("AWS_REGION", region)

	configVariables := config.Variables{
		"cluster_name": config.StringVariable(clusterName),
		"region":       config.StringVariable(region),
		"max_nodes":    config.StringVariable("2"),
		"role_arn":     config.StringVariable(os.Getenv("TEST_ROLE")),
	}
	configUpdateVariables := config.Variables{
		"cluster_name": config.StringVariable(clusterName),
		"region":       config.StringVariable(region),
		"max_nodes":    config.StringVariable("3"),
		"role_arn":     config.StringVariable(os.Getenv("TEST_ROLE")),
	}

	if endpoint, ok := os.LookupEnv("TEST_ENDPOINT"); ok {
		configVariables["endpoint"] = config.StringVariable(endpoint)
		configUpdateVariables["endpoint"] = config.StringVariable(endpoint)
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(configDirectory),
				ConfigVariables:          configVariables,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pcluster_cluster."+testResourceName,
						"cluster_status",
						string(openapi.CLUSTERSTATUS_CREATE_COMPLETE),
					),
					resource.TestCheckResourceAttr(
						"pcluster_cluster."+testResourceName,
						"cloudformation_stack_status",
						string(openapi.CLOUDFORMATIONRESOURCESTATUS_CREATE_COMPLETE),
					),
				),
			},
			// ImportState testing
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(configImportDirectory),
				ConfigVariables:          configVariables,
				ResourceName:             "pcluster_cluster." + testResourceName,
				ImportState:              true,
				ImportStateVerify:        true,
			},
			{
				// Update and Read testing
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(configDirectory),
				ConfigVariables:          configUpdateVariables,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pcluster_cluster."+testResourceName,
						"cluster_status",
						string(openapi.CLUSTERSTATUS_UPDATE_COMPLETE),
					),
					resource.TestCheckResourceAttr(
						"pcluster_cluster."+testResourceName,
						"cloudformation_stack_status",
						string(openapi.CLOUDFORMATIONRESOURCESTATUS_UPDATE_COMPLETE),
					),
				),
			},
			// Cluster Data Source Test
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(configDirectory),
				ConfigVariables:          configUpdateVariables,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.pcluster_cluster."+testResourceName,
						"cluster.clusterName",
						clusterName,
					),
					resource.TestCheckResourceAttrSet(
						"data.pcluster_cluster."+testResourceName,
						"stack_events.#",
					),
					resource.TestCheckResourceAttrSet(
						"data.pcluster_cluster."+testResourceName,
						"log_streams.#",
					),
				),
			},
			// Cluster List Data Source Test
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigDirectory:          config.StaticDirectory(configDirectory),
				ConfigVariables:          configUpdateVariables,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.pcluster_list_clusters."+testResourceName,
						"clusters.#",
					),
				),
			},
		},
	})
}

func TestUnitPopulateClusterResourceInfo(t *testing.T) {
	t.Parallel()
	contents := []openapi.ClusterInfoSummary{
		{
			ClusterName:               "Case01",
			Region:                    "us-east-1",
			Version:                   "3.7.0",
			CloudformationStackArn:    "some_arn",
			ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
			CloudformationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		},
		{
			ClusterName: "Case02",
		},
	}

	cases := []struct {
		content  openapi.ClusterInfoSummary
		expected ClusterResourceModel
	}{
		{
			content: contents[0],
			expected: ClusterResourceModel{
				ClusterName:            types.StringValue(contents[0].ClusterName),
				Id:                     types.StringValue(contents[0].ClusterName),
				Region:                 types.StringValue(contents[0].Region),
				Version:                types.StringValue(contents[0].Version),
				CloudformationStackArn: types.StringValue(contents[0].CloudformationStackArn),
				ClusterStatus:          types.StringValue(string(contents[0].ClusterStatus)),
				CloudformationStackStatus: types.StringValue(
					string(contents[0].CloudformationStackStatus),
				),
			},
		},
		{
			content: contents[1],
			expected: ClusterResourceModel{
				ClusterName: types.StringValue(contents[1].ClusterName),
				Id:          types.StringValue(contents[1].ClusterName),
			},
		},
	}

	for _, c := range cases {
		data := ClusterResourceModel{}
		populateClusterResourceInfo(&c.content, &data)
		if !reflect.DeepEqual(data, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v",
				data,
				c.expected,
			)
		}
	}
}

func mockHttpServer(configPath string, config string, t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != configPath {
			t.Errorf("Expected to request '%s', got: %s", configPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(config))
	}))

	return server
}

func TestUnitWaitClusterReady(t *testing.T) {
	names := map[string]string{
		"failed":         "cluster01",
		"complete":       "cluster02",
		"updateFailed":   "cluster03",
		"updateComplete": "cluster04",
		"deleteFailed":   "cluster05",
		"deleteComplete": "clsuter06",
		"wait":           "cluster07",
		"waitUpdate":     "cluster08",
		"waitDelete":     "cluster09",
	}

	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var status openapi.ClusterStatus
		switch name := path.Base(r.URL.Path); name {
		case names["failed"]:
			status = openapi.CLUSTERSTATUS_CREATE_FAILED
		case names["complete"]:
			status = openapi.CLUSTERSTATUS_CREATE_COMPLETE
		case names["updateFailed"]:
			status = openapi.CLUSTERSTATUS_UPDATE_FAILED
		case names["updateComplete"]:
			status = openapi.CLUSTERSTATUS_UPDATE_COMPLETE
		case names["deleteFailed"]:
			status = openapi.CLUSTERSTATUS_DELETE_FAILED
		case names["deleteComplete"]:
			status = openapi.CLUSTERSTATUS_DELETE_COMPLETE
		case names["wait"]:
			if count > 2 {
				status = openapi.CLUSTERSTATUS_CREATE_COMPLETE
				count = 0
			} else {
				status = openapi.CLUSTERSTATUS_CREATE_IN_PROGRESS
			}
			count++
		case names["waitUpdate"]:
			if count > 2 {
				status = openapi.CLUSTERSTATUS_UPDATE_COMPLETE
				count = 0
			} else {
				status = openapi.CLUSTERSTATUS_UPDATE_IN_PROGRESS
			}
			count++
		case names["waitDelete"]:
			if count > 2 {
				status = openapi.CLUSTERSTATUS_DELETE_COMPLETE
				count = 0
			} else {
				status = openapi.CLUSTERSTATUS_DELETE_IN_PROGRESS
			}
			count++
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		clusterJson, _ := openapi.DescribeClusterResponseContent{
			ClusterName:               path.Base(r.URL.Path),
			ClusterStatus:             status,
			ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
			CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		}.MarshalJSON()
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(clusterJson))
	}))

	defer server.Close()
	cR := ClusterResource{}
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
		expected *openapi.DescribeClusterResponseContent
	}{
		{
			name: names["failed"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["failed"],
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_FAILED,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["complete"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["complete"],
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["updateFailed"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["updateFailed"],
				ClusterStatus:             openapi.CLUSTERSTATUS_UPDATE_FAILED,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["updateComplete"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["updateComplete"],
				ClusterStatus:             openapi.CLUSTERSTATUS_UPDATE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["deleteFailed"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["deleteFailed"],
				ClusterStatus:             openapi.CLUSTERSTATUS_DELETE_FAILED,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["deleteComplete"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["deleteComplete"],
				ClusterStatus:             openapi.CLUSTERSTATUS_DELETE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["wait"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["wait"],
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["waitUpdate"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["waitUpdate"],
				ClusterStatus:             openapi.CLUSTERSTATUS_UPDATE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name: names["waitDelete"],
			expected: &openapi.DescribeClusterResponseContent{
				ClusterName:               names["waitDelete"],
				ClusterStatus:             openapi.CLUSTERSTATUS_DELETE_COMPLETE,
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			name:     names["notFound"],
			expected: nil,
		},
	}

	for _, c := range cases {
		out, err := cR.waitClusterReady(context.TODO(), c.name)
		if err != nil {
			if c.name == names["NotFound"] {
				if err.Error() != failedToFindClusterErr {
					t.Fatalf("Expected to see: %v", failedToFindClusterErr)
				}
				continue
			} else {
				t.Fatalf("Failure while getting cluster: %v", err)
			}
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v",
				out,
				c.expected,
			)
		}
	}
}

func TestUnitGetCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clusterJson, _ := openapi.DescribeClusterResponseContent{
			ClusterName:               path.Base(r.URL.Path),
			ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
			ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
			CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
		}.MarshalJSON()
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(clusterJson))
	}))
	defer server.Close()
	cR := ClusterResource{}
	c := openapi.NewConfiguration()

	c.Servers = openapi.ServerConfigurations{
		{
			URL:         server.URL,
			Description: "Parallel Cluster API",
		},
	}
	cR.client = openapi.NewAPIClient(c)

	contents := []string{"cluster01", "cluster02"}

	cases := []struct {
		content  string
		expected openapi.DescribeClusterResponseContent
	}{
		{
			content: contents[0],
			expected: openapi.DescribeClusterResponseContent{
				ClusterName:               contents[0],
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
		{
			content: contents[1],
			expected: openapi.DescribeClusterResponseContent{
				ClusterName:               contents[1],
				ComputeFleetStatus:        openapi.COMPUTEFLEETSTATUS_DISABLED,
				ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
				CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			},
		},
	}

	for _, c := range cases {
		out, err := cR.getCluster(context.TODO(), c.content)
		if err != nil {
			t.Fatalf("Failure while getting cluster: %v", err)
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v",
				out,
				c.expected,
			)
		}
	}
}

func TestUnitPopulateClusterDataDesc(t *testing.T) {
	configPath := "/config"
	config := `{"some":"json"}`

	server := mockHttpServer(configPath, config, t)

	clusterConfigUrl := server.URL + configPath
	defer server.Close()

	contents := []*openapi.DescribeClusterResponseContent{
		{
			ClusterName:               "Case01",
			Region:                    "us-east-1",
			Version:                   "3.7.0",
			CloudformationStackArn:    "some_arn",
			ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
			CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			ClusterConfiguration: openapi.ClusterConfigurationStructure{
				Url: &clusterConfigUrl,
			},
		},
		{},
	}
	emptyClusterObject, err := types.ObjectValue(
		clusterDescriptionObjectTypes,
		map[string]attr.Value{
			"version":                   types.StringValue(""),
			"cloudformationStackArn":    types.StringValue(""),
			"clusterStatus":             types.StringValue(""),
			"cloudFormationStackStatus": types.StringValue(""),
			"clusterConfiguration":      types.StringValue(""),
			"clusterName":               types.StringValue(""),
			"region":                    types.StringValue(""),
			"scheduler":                 types.ObjectNull(schedulerObjectTypes),
			"creationTime": types.StringValue(
				time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			),
			"lastUpdatedTime": types.StringValue(
				time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			),
			"computeFleetStatus": types.StringValue(""),
			"tags":               types.ListNull(types.MapType{ElemType: types.StringType}),
			"headNode":           types.MapNull(types.StringType),
			"failures":           types.ListNull(types.MapType{ElemType: types.StringType}),
			"loginNodes":         types.ObjectNull(loginNodesObjectTypes),
		},
	)

	clusterObject, err := types.ObjectValue(clusterDescriptionObjectTypes, map[string]attr.Value{
		"version": types.StringValue(contents[0].Version),
		"cloudformationStackArn": types.StringValue(
			contents[0].CloudformationStackArn,
		),
		"clusterStatus": types.StringValue(
			string(contents[0].ClusterStatus),
		),
		"cloudFormationStackStatus": types.StringValue(
			string(contents[0].CloudFormationStackStatus),
		),
		"clusterConfiguration": types.StringValue(config),
		"clusterName":          types.StringValue(contents[0].ClusterName),
		"region":               types.StringValue(contents[0].Region),
		"scheduler":            types.ObjectNull(schedulerObjectTypes),
		"creationTime": types.StringValue(
			time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
		),
		"lastUpdatedTime": types.StringValue(
			time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
		),
		"computeFleetStatus": types.StringValue(""),
		"tags":               types.ListNull(types.MapType{ElemType: types.StringType}),
		"headNode":           types.MapNull(types.StringType),
		"failures":           types.ListNull(types.MapType{ElemType: types.StringType}),
		"loginNodes":         types.ObjectNull(loginNodesObjectTypes),
	})
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		content  *openapi.DescribeClusterResponseContent
		expected *ClusterDataSourceModel
	}{
		{
			content: contents[0],
			expected: &ClusterDataSourceModel{
				Cluster: clusterObject,
			},
		},
		{
			content: contents[1],
			expected: &ClusterDataSourceModel{
				Cluster: emptyClusterObject,
			},
		},
	}

	for _, c := range cases {
		data := &ClusterDataSourceModel{}
		resp := &datasource.ReadResponse{}
		populateClusterDataSource(c.content, data, resp)
		if !reflect.DeepEqual(data, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v",
				data,
				c.expected,
			)
		}
	}
}

func TestUnitPopulateClusterResourceDesc(t *testing.T) {
	t.Parallel()
	configPath := "/config"
	config := `{"some":"json"}`

	server := mockHttpServer(configPath, config, t)

	clusterConfigUrl := server.URL + configPath
	defer server.Close()

	contents := []*openapi.DescribeClusterResponseContent{
		{
			ClusterName:               "Case01",
			Region:                    "us-east-1",
			Version:                   "3.7.0",
			CloudformationStackArn:    "some_arn",
			ClusterStatus:             openapi.CLUSTERSTATUS_CREATE_COMPLETE,
			CloudFormationStackStatus: openapi.CLOUDFORMATIONSTACKSTATUS_CREATE_COMPLETE,
			ClusterConfiguration: openapi.ClusterConfigurationStructure{
				Url: &clusterConfigUrl,
			},
		},
		{
			ClusterName: "Case02",
		},
	}

	cases := []struct {
		content  *openapi.DescribeClusterResponseContent
		expected *ClusterResourceModel
	}{
		{
			content: contents[0],
			expected: &ClusterResourceModel{
				ClusterName:            types.StringValue(contents[0].ClusterName),
				Id:                     types.StringValue(contents[0].ClusterName),
				Region:                 types.StringValue(contents[0].Region),
				Version:                types.StringValue(contents[0].Version),
				CloudformationStackArn: types.StringValue(contents[0].CloudformationStackArn),
				ClusterStatus:          types.StringValue(string(contents[0].ClusterStatus)),
				CloudformationStackStatus: types.StringValue(
					string(contents[0].CloudFormationStackStatus),
				),
				ClusterConfiguration: types.StringValue(config),
			},
		},
		{
			content: contents[1],
			expected: &ClusterResourceModel{
				ClusterName: types.StringValue(contents[1].ClusterName),
				Id:          types.StringValue(contents[1].ClusterName),
			},
		},
	}

	for _, c := range cases {
		data := &ClusterResourceModel{ClusterName: types.StringValue("Make all Values Known.")}
		populateClusterResourceDesc(c.content, data)
		if !reflect.DeepEqual(data, c.expected) {
			t.Fatalf(
				"Error matching output and expected. \nO: %#v\nE: %#v",
				data,
				c.expected,
			)
		}
	}
}
