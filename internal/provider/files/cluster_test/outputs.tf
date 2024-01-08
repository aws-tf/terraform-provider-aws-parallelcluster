output "cluster_resource" {
  value = resource.pcluster_cluster.test
}

output "cluster_compute_fleet_resource" {
  value = resource.pcluster_compute_fleet_status.test.status
}

output "cluster_data_source" {
  value = data.pcluster_cluster.test
}

output "cluster_list_data_source" {
  value = data.pcluster_list_clusters.test
}
