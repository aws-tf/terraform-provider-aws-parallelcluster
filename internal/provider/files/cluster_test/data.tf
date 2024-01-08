data "pcluster_list_clusters" "test" {
  region     = var.region
  depends_on = [pcluster_cluster.test]
}

data "pcluster_cluster" "test" {
  cluster_name = var.cluster_name
  depends_on   = [pcluster_cluster.test]
}
