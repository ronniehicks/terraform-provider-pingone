resource "pingone_identity_provider_attributes" "openid_example" {
  environment_id       = local.environment_id
  identity_provider_id = "idpId"
  name                 = "example"
  value                = "$${providerAttributes.example}"
  update               = "EMPTY_ONLY"
}

output "pingone_openid_idp_example_id" {
  value = pingone_identity_provider_attributes.openid_example.id
}
