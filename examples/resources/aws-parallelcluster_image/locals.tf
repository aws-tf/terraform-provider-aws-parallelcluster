locals {
  config_vars = {
    instanceType      = var.instance_type
    parentImage       = data.aws-parallelcluster_list_official_images.parent_image.official_images[0].amiId
    cleanupLambdaRole = aws_iam_role.cleanup_lambda_role.arn
  }
  partition  = data.aws_partition.current.partition
  region     = data.aws_region.current.name
  account_id = data.aws_caller_identity.current.account_id
}