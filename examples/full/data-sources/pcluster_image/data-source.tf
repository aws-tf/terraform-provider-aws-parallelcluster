data "pcluster_list_images" "example_pending" {
  image_status = "PENDING"
}

data "pcluster_image" "example_pending" {
  count    = length(data.pcluster_list_images.example_pending.images)
  image_id = data.pcluster_list_images.example_pending.images[count.index].imageId
}

data "pcluster_list_images" "example" {
  image_status = "AVAILABLE"
}

data "pcluster_image" "example" {
  count    = length(data.pcluster_list_images.example.images)
  image_id = data.pcluster_list_images.example.images[count.index].imageId
}
