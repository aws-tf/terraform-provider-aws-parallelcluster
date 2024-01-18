/** Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
