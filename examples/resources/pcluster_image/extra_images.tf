resource "pcluster_image" "build-demo2" {
  image_id            = "imageBuilderDemo2"
  image_configuration = file("files/image-build-demo.yaml")
}

resource "pcluster_image" "build-demo3" {
  image_id            = "imageBUilderDemo3"
  image_configuration = file("files/image-build-demo.yaml")
}
