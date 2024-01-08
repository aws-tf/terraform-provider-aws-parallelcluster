output "images" {
  description = "All images."
  value       = [for c in data.pcluster_list_images.example.images : { name = c.imageId, status = c.imageBuildStatus }]
}

output "image" {
  value = try(data.pcluster_image.example[0].image, null)
}

resource "local_file" "logs" {
  count    = length(data.pcluster_image.example)
  content  = jsonencode(data.pcluster_image.example[count.index].log_streams)
  filename = "logs/logs.json"
}

resource "local_file" "stack_events" {
  count    = length(data.pcluster_image.example)
  content  = jsonencode(data.pcluster_image.example[count.index].stack_events)
  filename = "logs/stack_events.json"
}

output "images_pending" {
  description = "All images."
  value       = [for c in data.pcluster_list_images.example_pending.images : { name = c.imageId, status = c.imageBuildStatus }]
}

output "image_pending" {
  value = try(data.pcluster_image.example_pending[0].image, null)
}

resource "local_file" "logs_pending" {
  count    = length(data.pcluster_image.example_pending)
  content  = jsonencode(data.pcluster_image.example_pending[count.index].log_streams)
  filename = "logs/logs_pending.json"
}

resource "local_file" "stack_events_pending" {
  count    = length(data.pcluster_image.example_pending)
  content  = jsonencode(data.pcluster_image.example_pending[count.index].stack_events)
  filename = "logs/stack_events_pending.json"
}
