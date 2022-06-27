resource "pingone_custom_domain" "custom_domain" {
  environment_id = local.environment_id
  domain_name    = "fake.com"
}
