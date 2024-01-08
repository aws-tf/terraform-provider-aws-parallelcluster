resource "pcluster_list_images" "example" {
  cluster_status = "AVAILABLE"
  region         = var.region
}
