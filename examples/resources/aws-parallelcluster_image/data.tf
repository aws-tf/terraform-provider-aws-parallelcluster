data "aws_caller_identity" "current" {}
data "aws_partition" "current" {}
data "aws_region" "current" {}

data "aws-parallelcluster_list_official_images" "parent_image" {
  region       = var.region
  os           = var.os
  architecture = var.architecture
}