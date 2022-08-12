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
provider "pingone" {
  client_id     = "client_id"
  client_secret = var.P1_CLIENT_SECRET
}

locals {
  environment_id = "env_id"
}

