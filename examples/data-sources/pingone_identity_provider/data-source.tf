data "pingone_identity_provider" "openid_idp" {
  environment_id       = local.environment_id
  identity_provider_id = "some_id"
}

data "pingone_identity_provider" "saml_idp" {
  environment_id       = local.environment_id
  identity_provider_id = "some_id"
}

output "openid_idp" {
  value = data.pingone_identity_provider.openid_idp
}

output "saml_idp" {
  value = data.pingone_identity_provider.saml_idp
}

data "pingone_identity_provider" "all" {
  environment_id = local.environment_id
}

output "idps" {
  value = data.pingone_identity_provider.all
}
