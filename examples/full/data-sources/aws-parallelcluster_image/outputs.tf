/** Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License"). You may not
* use this file except in compliance with the License. A copy of the License is
* located at
*
* http://aws.amazon.com/apache2.0/
*
* or in the "LICENSE.txt" file accompanying this file. This file is distributed
* on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
* implied. See the License for the specific language governing permissions and
* limitations under the License.
*/

output "images" {
  description = "All images."
  value       = [for c in data.aws-parallelcluster_list_images.example.images : { name = c.imageId, status = c.imageBuildStatus }]
}

output "image" {
  value = try(data.aws-parallelcluster_image.example[0].image, null)
}

resource "local_file" "logs" {
  count    = length(data.aws-parallelcluster_image.example)
  content  = jsonencode(data.aws-parallelcluster_image.example[count.index].log_streams)
  filename = "logs/logs.json"
}

resource "local_file" "stack_events" {
  count    = length(data.aws-parallelcluster_image.example)
  content  = jsonencode(data.aws-parallelcluster_image.example[count.index].stack_events)
  filename = "logs/stack_events.json"
}

output "images_pending" {
  description = "All images."
  value       = [for c in data.aws-parallelcluster_list_images.example_pending.images : { name = c.imageId, status = c.imageBuildStatus }]
}

output "image_pending" {
  value = try(data.aws-parallelcluster_image.example_pending[0].image, null)
}

resource "local_file" "logs_pending" {
  count    = length(data.aws-parallelcluster_image.example_pending)
  content  = jsonencode(data.aws-parallelcluster_image.example_pending[count.index].log_streams)
  filename = "logs/logs_pending.json"
}

resource "local_file" "stack_events_pending" {
  count    = length(data.aws-parallelcluster_image.example_pending)
  content  = jsonencode(data.aws-parallelcluster_image.example_pending[count.index].stack_events)
  filename = "logs/stack_events_pending.json"
}
