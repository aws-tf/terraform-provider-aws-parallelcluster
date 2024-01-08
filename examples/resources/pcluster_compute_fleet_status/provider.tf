provider "pcluster" {
  endpoint = var.endpoint
  role_arn = var.role_arn
}

terraform {
  required_version = ">= 1.5.0"
  required_providers {
    pcluster = {
      source  = "terraform.local/local/pcluster"
      version = "0.0.1"
    }
  }
}
