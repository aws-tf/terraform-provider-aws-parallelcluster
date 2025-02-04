resource "aws-parallelcluster_image" "custom_image_1" {
  image_id            = "CustomImage1-${random_id.suffix.id}"
  image_configuration = try(templatefile("files/image-build-demo.yaml", local.config_vars), "{}")
  rollback_on_failure = false
}

resource "aws-parallelcluster_image" "custom_image_2" {
  image_id = "CustomImage2-${random_id.suffix.id}"
  image_configuration = yamlencode({
    "Build" : {
      "InstanceType" : local.config_vars.instanceType,
      "ParentImage" : local.config_vars.parentImage,
      "UpdateOsPackages" : {
        "Enabled" : false
      },
      "Iam" : {
        "CleanupLambdaRole" : local.config_vars.cleanupLambdaRole
      }
    }
  })
  rollback_on_failure = false
}