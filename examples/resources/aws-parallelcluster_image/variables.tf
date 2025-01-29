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

variable "region" {
  description = "The region the ParallelCluster API is deployed in."
  type        = string
}

variable "profile" {
  type        = string
  description = "The AWS profile used to deploy the clusters."
  default     = null
}

variable "api_stack_name" {
  type        = string
  description = "The name of the CloudFormation stack used to deploy the ParallelCluster API."
}

variable "api_version" {
  type        = string
  description = "The version of the ParallelCluster API."
}

variable "os" {
  type        = string
  description = "The OS of the ParallelCluster image."
}

variable "architecture" {
  type        = string
  description = "The architecture of the ParallelCluster image."
}

variable "instance_type" {
  type        = string
  description = "The build instance type used to build the ParallelCluster image."
}
