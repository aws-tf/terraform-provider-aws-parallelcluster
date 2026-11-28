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

locals {
  # Some attributes of the `aws-parallelcluster_cluster` data source are inherently time varying:
  # they change on their own while the cluster is alive, with no cluster update involved.
  # Dropping them from the output is the only way to tolerate this with terraform-plugin-testing.
  #   - `log_streams` grows as CloudWatch creates a stream the first time a node writes to a log
  #      file (e.g. `pcluster-check-update` appearing minutes after the other streams).
  #   - `stack_events` grows as CloudFormation keeps appending events.
  #   - `cluster.computeFleetStatus` transitions on its own (START_REQUESTED -> STARTING -> RUNNING).
  time_varying_data_source_attributes = ["log_streams", "stack_events"]
  time_varying_cluster_attributes     = ["computeFleetStatus"]
}

output "cluster_data_source" {
  value = merge(
    {
      for k, v in data.aws-parallelcluster_cluster.test : k => v
      if !contains(local.time_varying_data_source_attributes, k)
    },
    {
      cluster = {
        for k, v in data.aws-parallelcluster_cluster.test.cluster : k => v
        if !contains(local.time_varying_cluster_attributes, k)
      }
    },
  )
}
