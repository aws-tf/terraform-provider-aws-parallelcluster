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

variable "profile" {
  type        = string
  description = "The aws profile to use."
  default     = null
}

variable "region" {
  type        = string
  description = "The region to create the cluster in"
  default     = "us-west-2"
}
