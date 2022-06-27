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
  client_id     = "3c789e2c-4900-4772-8d79-1148ef3baab1"
  client_secret = var.P1_CLIENT_SECRET
}

locals {
  environment_id = "5c3a5b10-ce82-4f4b-80de-6fc187022f35"
}

locals {
  environment_id = "envId"
}
