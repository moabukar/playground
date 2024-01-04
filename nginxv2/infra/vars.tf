variable "environment" {
  type        = string
  description = "The environment to deploy to."
}

variable "execution_role_arn" {
  type        = string
  description = "The role which should be assumed by ecs services, lambdas , ec2 instances."
}

variable "spacelift_commit_sha" {
  type        = string
  default     = null
  description = "The current commit sha"
}

variable "instance" {
  type        = string
  default     = ""
  description = "The specific name of this instance of the component"
}

variable "environment_vars" {
  type        = list(any)
  default     = []
  description = "Any environment variables to pass"
}

variable "task_desired_count" {
  type        = number
  default     = 1
  description = "The number of tasks required"
}

variable "target_group_arns" {
  type        = list(string)
  default     = []
  description = "Any target groups to attach to for custom routing"
}
