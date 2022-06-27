# Web App
resource "pingone_application" "web" {
  environment_id = local.environment_id
  name           = "tf_web"
  enabled        = true
  type           = "WEB_APP"
  protocol       = "OPENID_CONNECT"

  grant_types                    = ["AUTHORIZATION_CODE"]
  post_logout_redirect_uris      = ["https://fake.com/logout"]
  redirect_uris                  = ["https://fake.com/"]
  response_types                 = ["CODE"]
  token_endpoint_auth_method     = "CLIENT_SECRET_BASIC"
  pkce_enforcement               = "REQUIRED"
  refresh_token_duration         = 86400
  refresh_token_rolling_duration = 86400

  icon {
    id   = "1d39eadb-ee72-41a1-a460-f5a5fd2b0a27"
    href = "https://picsum.photos/200/300"
  }

  access_control {
    group {
      type   = "ANY_GROUP"
      groups = ["cd39fd84-955c-4742-b72d-1f46f3341c3b"]
    }
  }
}

# Native App
resource "pingone_application" "native" {
  environment_id = local.environment_id
  name           = "tf_native"
  enabled        = true
  type           = "NATIVE_APP"
  protocol       = "OPENID_CONNECT"

  grant_types                    = ["AUTHORIZATION_CODE", "IMPLICIT"]
  post_logout_redirect_uris      = ["https://fake.com/logout"]
  redirect_uris                  = ["https://fake.com"]
  response_types                 = ["CODE", "TOKEN", "ID_TOKEN"]
  token_endpoint_auth_method     = "NONE"
  pkce_enforcement               = "REQUIRED"
  refresh_token_duration         = 86400
  refresh_token_rolling_duration = 86400

  mobile {
    bundle_id    = "myAndroidPackage.myApp"
    package_name = "myAndroidPackage.myApp"
    integrity_detection {
      mode = "ENABLED"
      cache_duration {
        amount = 10
        units  = "MINUTES"
      }
    }

    passcode_refresh_duration {
      duration  = 45
      time_unit = "SECONDS"
    }
  }
}

# SPA App
resource "pingone_application" "spa" {
  environment_id = local.environment_id
  name           = "tf_spa"
  enabled        = true
  type           = "SINGLE_PAGE_APP"
  protocol       = "OPENID_CONNECT"

  grant_types                    = ["IMPLICIT"]
  post_logout_redirect_uris      = ["https://fake.com/logout"]
  redirect_uris                  = ["https://fake.com/"]
  response_types                 = ["TOKEN", "ID_TOKEN"]
  token_endpoint_auth_method     = "NONE"
  pkce_enforcement               = "REQUIRED"
  refresh_token_duration         = 86400
  refresh_token_rolling_duration = 86400
}

# Service App
resource "pingone_application" "service" {
  environment_id = local.environment_id
  name           = "tf_service"
  enabled        = true
  type           = "SERVICE"
  protocol       = "OPENID_CONNECT"

  grant_types                    = ["AUTHORIZATION_CODE", "CLIENT_CREDENTIALS", "IMPLICIT", "REFRESH_TOKEN"]
  post_logout_redirect_uris      = ["https://fake.com/logout"]
  redirect_uris                  = ["https://fake.com/"]
  response_types                 = ["TOKEN", "ID_TOKEN", "CODE"]
  token_endpoint_auth_method     = "CLIENT_SECRET_POST"
  pkce_enforcement               = "REQUIRED"
  refresh_token_duration         = 86400
  refresh_token_rolling_duration = 86400
}

# Worker App
resource "pingone_application" "worker" {
  environment_id = local.environment_id
  name           = "tf_worker"
  enabled        = true
  type           = "WORKER"
  protocol       = "OPENID_CONNECT"

  grant_types                    = ["CLIENT_CREDENTIALS"]
  post_logout_redirect_uris      = ["https://fake.com/logout"]
  redirect_uris                  = ["https://fake.com/"]
  token_endpoint_auth_method     = "CLIENT_SECRET_BASIC"
  pkce_enforcement               = "REQUIRED"
  refresh_token_duration         = 86400
  refresh_token_rolling_duration = 86400
}

# SAML App
resource "pingone_application" "saml" {
  environment_id = local.environment_id
  name           = "tf_saml"
  enabled        = true
  type           = "WEB_APP"
  protocol       = "SAML"

  assertion_duration    = 60
  acs_urls              = ["https://localhost:1337/acs"]
  slo_response_endpoint = "https://localhost:1337/slo"
  sp_entity_id          = "tf_saml"
  name_id_format        = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
}
