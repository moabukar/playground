
variable "provider_region" {
  type    = string
  default = "us-east-1"

}

variable "rds_password" {
  type      = string
  sensitive = true
  default   = "rizk123456"
}

variable "rds_username" {
  type      = string
  sensitive = true
  default   = "rizk"
}

variable "alb-name" {
  type      = string
  sensitive = true
  default   = "rizk"
}
