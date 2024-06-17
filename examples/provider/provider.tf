provider "aws-parallelcluster" {
  role_arn = var.role_arn
  endpoint = var.endpoint
  profile  = var.profile
}

terraform {
  required_version = ">= 1.5.7"
  required_providers {
    aws-parallelcluster = {
      source  = "aws-tf/aws-parallelcluster"
      version = "1.0.0"
    }
  }
}
