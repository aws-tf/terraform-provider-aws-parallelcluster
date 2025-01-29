data "aws-parallelcluster_list_official_images" "parent_image" {
  region       = var.region
  os           = var.os
  architecture = var.architecture
}