/**
 *  Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"). You may not
 *  use this file except in compliance with the License. A copy of the License is
 *  located at
 *
 *  http://aws.amazon.com/apache2.0/
 *
 *  or in the "LICENSE.txt" file accompanying this file. This file is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
 *  implied. See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  # Required Infrastructure Submodule
 *  The required infra submodule deploys a vpc, subnets, routes, gateways, and creates a
 *  key pair. These are necessary resources for the API to deploy and manage clusters.
 */

output "cluster_resource" {
  value = resource.aws-parallelcluster_cluster.test
}

output "cluster_compute_fleet_resource" {
  value = resource.aws-parallelcluster_compute_fleet_status.test.status
}

output "cluster_data_source" {
  value = data.aws-parallelcluster_cluster.test
}
