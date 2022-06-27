# - export TF_VAR_P1_CLIENT_SECRET="secretstuff"

variable "P1_CLIENT_SECRET" {
  type        = string
  description = "PingOne Client Secret Key"
  sensitive   = true
}

resource "pingone_identity_provider" "openid" {
  environment_id       = local.environment_id
  name                 = "Justatest"
  description          = "This is just a test description change."
  type                 = "OPENID_CONNECT"
  enabled              = false
  authn_request_signed = false
  icon {
    href = "https://picsum.photos/200/300"
    id   = "iconId"
  }
  login_button_icon {
    href = "https://picsum.photos/200/300"
    id   = "iconId"
  }
  registration {
    population_id = "populationId"
  }
  scopes                     = ["openid"]
  client_id                  = "testing"
  client_secret              = var.P1_CLIENT_SECRET
  authorization_endpoint     = "https://pingone.com/as/auth"
  issuer                     = "https://pingone.com/as/issuer"
  jwks_endpoint              = "https://pingone.com/as/jwks"
  token_endpoint             = "https://pingone.com/as/token"
  token_endpoint_auth_method = "CLIENT_SECRET_BASIC"
  user_info_endpoint         = "https://pingone.com/as/user"
}

resource "pingone_identity_provider" "saml" {
  environment_id       = local.environment_id
  name                 = "JustaSamlTest"
  description          = "This is just a SAML test"
  type                 = "SAML"
  enabled              = false
  authn_request_signed = false
  icon {
    href = "https://picsum.photos/200/300"
    id   = "iconId"
  }
  login_button_icon {
    href = "https://picsum.photos/200/300"
    id   = "iconId"
  }
  idp_entity_id                    = "testingthissaml"
  idp_verification_certificate_ids = ["idpVerificationCertId"]
  sp_entity_id                     = "https://auth.pingone.com/entityId"
  sp_signing_key_id                = "signingKeyId"
  sso_binding                      = "HTTP_POST"
  sso_endpoint                     = "https://pingone.com/saml20/sp/acs/test"
}
