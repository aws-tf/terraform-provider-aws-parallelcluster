resource "pcluster_image" "build-demo2" {
  image_id            = "Demo2"
  image_configuration = file("files/image-build-demo.yaml")
}

resource "pcluster_image" "build-demo3" {
  image_id            = "Demo3"
  image_configuration = file("files/image-build-demo.yaml")
}
