/**
 *  Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"). You may not
 *  use this file except in compliance with the License. A copy of the License is
 *  located at
 *
 *  http://aws.amazon.com/apache2.0/
 *
 *  or in the "LICENSE.txt" file accompanying this file. This file is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
 *  implied. See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  # Required Infrastructure Submodule
 *  The required infra submodule deploys a vpc, subnets, routes, gateways, and creates a
 *  key pair. These are necessary resources for the API to deploy and manage clusters.
 */

######################
#### CLUSTER VARS ####
######################

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
  default     = "pcluster-example"
}

variable "max_nodes" {
  type        = string
  description = "The maximum number of compute nodes for the cluster."
  default     = "4"
}

variable "min_nodes" {
  type        = string
  description = "The minimum number of compute nodes for the cluster."
  default     = "1"
}

variable "subnet" {
  type        = string
  description = "The subnet to deploy a cluster to."
  default     = null
}

variable "key_pair" {
  type        = string
  description = "The keypair used to deploy a cluster."
  default     = null
}

variable "region" {
  type        = string
  description = "The region to create the cluster in"
  default     = "us-east-1"
}

variable "compute_node_type" {
  type        = string
  description = "The type of instance for compute nodes."
  default     = "t3.small"
}

variable "head_node_type" {
  type        = string
  description = "The type of instance for head nodes."
  default     = "t3.small"
}

#######################
#### PROVIDER VARS ####
#######################

variable "role_arn" {
  type        = string
  description = "The role used by the ParallelCluster provider."
  default     = null
}

variable "endpoint" {
  type        = string
  description = "The endpoint used by the ParallelCluster provider."
  default     = null
}

variable "use_user_role" {
  type        = bool
  description = "If set to true retrieve the user role exported from the ParallelCluster CloudFormation stack."
  default     = null
}

variable "api_stack_name" {
  type        = string
  description = "Retrieve the ParallelCluster API stack endpoint from the given CloudFormation stack name."
  default     = null
}
