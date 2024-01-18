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

variable "use_user_role" {
  type        = bool
  description = "If set to true retrieve the user role exported from the ParallelCluster cloudformation stack."
  default     = null
}

variable "api_stack_name" {
  type        = string
  description = "Retrieve the api stack endpoint from the given cloudformation stack name."
  default     = null
}
