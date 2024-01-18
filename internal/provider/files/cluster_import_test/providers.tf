provider "pcluster" {
  endpoint       = var.endpoint
  role_arn       = var.role_arn
  use_user_role  = var.use_user_role
  api_stack_name = var.api_stack_name
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    tls = {
      source  = "hashicorp/tls"
      version = ">= 3.4"
    }
  }
}
