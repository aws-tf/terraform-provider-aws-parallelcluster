data "pcluster_list_clusters" "example" {}

data "pcluster_compute_fleet_status" "example" {
  cluster_name = data.pcluster_list_clusters.example.clusters[0].clusterName
}

data "pcluster_cluster" "example" {
  cluster_name = data.pcluster_list_clusters.example.clusters[0].clusterName
}
