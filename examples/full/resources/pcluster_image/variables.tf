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

#######################
#### PROVIDER VARS ####
#######################

variable "role_arn" {
  type        = string
  description = "The role used by the pcluster provider."
  default     = null
}

variable "endpoint" {
  type        = string
  description = "The endpoint used by the pcluster provider."
  default     = null
}

variable "api_stack_name" {
  type        = string
  description = "The name of the cloudformation stack used for the ParallelCluster API."
  default     = null
}

variable "parent_image" {
  type        = string
  description = "The image id to use as the base AMI of the new image."
  default     = "ami-033a1ef04d047e1ca"
}

variable "profile" {
  type        = string
  description = "The aws profile to use."
  default     = null
}

variable "region" {
  type        = string
  description = "The region to create the image in"
  default     = "us-east-1"
}

variable "default_az" {
  type        = string
  description = "The default availability zone to create the image in"
  default     = "us-east-1a"
}
