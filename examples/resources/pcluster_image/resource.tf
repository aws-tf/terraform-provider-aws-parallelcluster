resource "pcluster_image" "build-demo" {
  image_id            = "imageBuilderDemo"
  rollback_on_failure = false
  image_configuration = file("files/image-build-demo.yaml")
}
