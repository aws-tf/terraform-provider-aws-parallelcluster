/** Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License"). You may not
* use this file except in compliance with the License. A copy of the License is
* located at
*
* http://aws.amazon.com/apache2.0/
*
* or in the "LICENSE.txt" file accompanying this file. This file is distributed
* on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
* implied. See the License for the specific language governing permissions and
* limitations under the License.
*/

output "clusters" {
  description = "All clusters."
  value       = [for c in data.aws-parallelcluster_list_clusters.example.clusters : { name = c.clusterName, status = c.clusterStatus }]
}

output "compute_fleet" {
  description = "Compute fleet of the first cluster"
  value       = data.aws-parallelcluster_compute_fleet_status.example
}

output "cluster" {
  value = data.aws-parallelcluster_cluster.example.cluster
}

resource "local_file" "logs" {
  content  = jsonencode(data.aws-parallelcluster_cluster.example.log_streams)
  filename = "logs/logs.json"
}

resource "local_file" "stack_events" {
  content  = jsonencode(data.aws-parallelcluster_cluster.example.stack_events)
  filename = "logs/stack_events.json"
}
