terraform {
  required_providers {
    pingone = {
      version = "0.1"
      source  = "registry.terraform.io/ronniehicks/pingone"
    }
  }
}

variable "P1_CLIENT_SECRET" {
  type        = string
  description = "PingOne Client Secret Key"
  sensitive   = true
}
variable "P1_CLIENT_ID" {
  type        = string
  description = "PingOne Client Id"
}
variable "P1_ENV_ID" {
  type        = string
  description = "PingOne Environment Id"
}
provider "pingone" {
  client_id     = var.P1_CLIENT_ID
  client_secret = var.P1_CLIENT_SECRET
}

locals {
  environment_id = var.P1_ENV_ID
}
