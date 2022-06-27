# - export TF_VAR_P1_CLIENT_SECRET="secretstuff"

variable "P1_CLIENT_SECRET" {
  type        = string
  description = "PingOne Client Secret Key"
  sensitive   = true
}

provider "pingone" {
  client_id     = "client_id"
  client_secret = var.P1_CLIENT_SECRET
}
