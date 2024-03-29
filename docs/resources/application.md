---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_application Resource - terraform-provider-pingone"
subcategory: ""
description: |-
  pingone_application is used to managed an application for an environment.
---

# pingone_application (Resource)

`pingone_application` is used to managed an application for an environment.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean)
- `environment_id` (String) Environment Id
- `name` (String)
- `protocol` (String)
- `type` (String)

### Optional

- `access_control` (Block Set) (see [below for nested schema](#nestedblock--access_control))
- `acs_urls` (List of String)
- `assertion_duration` (Number)
- `assertion_signed` (Boolean)
- `assign_actor_roles` (Boolean)
- `bundle_id` (String)
- `default_target_url` (String)
- `description` (String)
- `grant_types` (List of String)
- `home_page_url` (String)
- `icon` (Block Set) (see [below for nested schema](#nestedblock--icon))
- `idp_signing` (Block Set) (see [below for nested schema](#nestedblock--idp_signing))
- `login_page_url` (String)
- `mobile` (Block Set) (see [below for nested schema](#nestedblock--mobile))
- `name_id_format` (String)
- `package_name` (String)
- `pkce_enforcement` (String)
- `post_logout_redirect_uris` (List of String)
- `redirect_uris` (List of String)
- `refresh_token_duration` (Number)
- `refresh_token_rolling_duration` (Number)
- `response_signed` (Boolean)
- `response_types` (List of String)
- `slo_binding` (String)
- `slo_endpoint` (String)
- `slo_response_endpoint` (String)
- `sp_encryption` (Block Set) (see [below for nested schema](#nestedblock--sp_encryption))
- `sp_entity_id` (String)
- `sp_verification` (Block Set) (see [below for nested schema](#nestedblock--sp_verification))
- `support_unsigned_request_object` (Boolean)
- `token_endpoint_auth_method` (String)

### Read-Only

- `application_id` (String) Application Id
- `id` (String) The ID of this resource.

<a id="nestedblock--access_control"></a>
### Nested Schema for `access_control`

Optional:

- `group` (Block Set) (see [below for nested schema](#nestedblock--access_control--group))
- `role` (Map of String)

<a id="nestedblock--access_control--group"></a>
### Nested Schema for `access_control.group`

Required:

- `groups` (List of String)
- `type` (String)



<a id="nestedblock--icon"></a>
### Nested Schema for `icon`

Required:

- `href` (String)

Read-Only:

- `id` (String) The ID of this resource.


<a id="nestedblock--idp_signing"></a>
### Nested Schema for `idp_signing`

Optional:

- `algorithm` (String)
- `key` (String)


<a id="nestedblock--mobile"></a>
### Nested Schema for `mobile`

Optional:

- `bundle_id` (String)
- `integrity_detection` (Block Set) (see [below for nested schema](#nestedblock--mobile--integrity_detection))
- `package_name` (String)
- `passcode_refresh_duration` (Block Set) (see [below for nested schema](#nestedblock--mobile--passcode_refresh_duration))

<a id="nestedblock--mobile--integrity_detection"></a>
### Nested Schema for `mobile.integrity_detection`

Optional:

- `cache_duration` (Block Set) (see [below for nested schema](#nestedblock--mobile--integrity_detection--cache_duration))
- `mode` (String)

<a id="nestedblock--mobile--integrity_detection--cache_duration"></a>
### Nested Schema for `mobile.integrity_detection.cache_duration`

Optional:

- `amount` (Number)
- `units` (String)



<a id="nestedblock--mobile--passcode_refresh_duration"></a>
### Nested Schema for `mobile.passcode_refresh_duration`

Optional:

- `duration` (Number)
- `time_unit` (String)



<a id="nestedblock--sp_encryption"></a>
### Nested Schema for `sp_encryption`

Required:

- `certificates` (String)

Optional:

- `algorithm` (String)


<a id="nestedblock--sp_verification"></a>
### Nested Schema for `sp_verification`

Required:

- `certificates` (List of String)

Optional:

- `authn_request_signed` (String) String representation of a bool so we can handle tristate

## Import

Import is supported using the following syntax:

```shell
# import using the environment and application id from the API
terraform import pingone_application.app environment:application
```
