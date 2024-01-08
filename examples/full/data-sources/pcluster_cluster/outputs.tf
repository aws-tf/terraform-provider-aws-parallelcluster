output "clusters" {
  description = "All clusters."
  value       = [for c in data.pcluster_list_clusters.example.clusters : { name = c.clusterName, status = c.clusterStatus }]
}

output "compute_fleet" {
  description = "Compute fleet of the first cluster"
  value       = data.pcluster_compute_fleet_status.example
}

output "cluster" {
  value = data.pcluster_cluster.example.cluster
}

resource "local_file" "logs" {
  content  = jsonencode(data.pcluster_cluster.example.log_streams)
  filename = "logs/logs.json"
}

resource "local_file" "stack_events" {
  content  = jsonencode(data.pcluster_cluster.example.stack_events)
  filename = "logs/stack_events.json"
}
