resource "pcluster_list_clusters" "example" {
  region         = var.region
  cluster_status = "CREATE_COMPLETE"
}
