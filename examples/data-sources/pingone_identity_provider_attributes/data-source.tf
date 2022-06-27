
data "pingone_identity_provider_attributes" "openid_idp" {
  environment_id       = local.environment_id
  identity_provider_id = "some_id"
}

data "pingone_identity_provider_attributes" "saml_idp" {
  environment_id       = local.environment_id
  identity_provider_id = "some_id"
}

output "openid_idp" {
  value = data.pingone_identity_provider_attributes.openid_idp
}

output "saml_idp" {
  value = data.pingone_identity_provider_attributes.saml_idp
}
