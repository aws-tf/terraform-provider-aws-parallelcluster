######################
#### CLUSTER VARS ####
######################

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
  default     = "pcluster-example"
}

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
