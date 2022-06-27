data "pingone_certificates" "certificates" {
  environment_id = local.environment_id
  id = "someid"
}
