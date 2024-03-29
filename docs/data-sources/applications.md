---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_applications Data Source - terraform-provider-pingone"
subcategory: ""
description: |-
  pingone_applications data source can be used to list all available applications for an environment.
---

# pingone_applications (Data Source)

`pingone_applications` data source can be used to list all available applications for an environment.

## Example Usage

```terraform
data "pingone_applications" "applications" {
  environment_id = local.environment_id
  # id = "someid"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Environment ID

### Read-Only

- `applications` (List of Object) (see [below for nested schema](#nestedatt--applications))
- `filter` (String) SCIM filter
- `id` (String) The ID of this resource.

<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `access_control` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--access_control))
- `acs_urls` (List of String)
- `application_id` (String)
- `assertion_duration` (Number)
- `assertion_signed` (Boolean)
- `assign_actor_roles` (Boolean)
- `bundle_id` (String)
- `default_target_url` (String)
- `description` (String)
- `enabled` (Boolean)
- `environment` (Map of String)
- `grant_types` (List of String)
- `home_page_url` (String)
- `icon` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--icon))
- `idp_signing` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--idp_signing))
- `login_page_url` (String)
- `mobile` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--mobile))
- `name` (String)
- `name_id_format` (String)
- `package_name` (String)
- `pkce_enforcement` (String)
- `post_logout_redirect_uris` (List of String)
- `protocol` (String)
- `redirect_uris` (List of String)
- `refresh_token_duration` (Number)
- `refresh_token_rolling_duration` (Number)
- `response_signed` (Boolean)
- `response_types` (List of String)
- `slo_binding` (String)
- `slo_endpoint` (String)
- `slo_response_endpoint` (String)
- `sp_encryption` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--sp_encryption))
- `sp_entity_id` (String)
- `sp_verification` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--sp_verification))
- `support_unsigned_request_object` (Boolean)
- `token_endpoint_auth_method` (String)
- `type` (String)

<a id="nestedobjatt--applications--access_control"></a>
### Nested Schema for `applications.access_control`

Read-Only:

- `group` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--access_control--group))
- `role` (Map of String)

<a id="nestedobjatt--applications--access_control--group"></a>
### Nested Schema for `applications.access_control.group`

Read-Only:

- `groups` (List of String)
- `type` (String)



<a id="nestedobjatt--applications--icon"></a>
### Nested Schema for `applications.icon`

Read-Only:

- `href` (String)
- `id` (String)


<a id="nestedobjatt--applications--idp_signing"></a>
### Nested Schema for `applications.idp_signing`

Read-Only:

- `algorithm` (String)
- `key` (String)


<a id="nestedobjatt--applications--mobile"></a>
### Nested Schema for `applications.mobile`

Read-Only:

- `bundle_id` (String)
- `integrity_detection` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--mobile--integrity_detection))
- `package_name` (String)
- `passcode_refresh_duration` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--mobile--passcode_refresh_duration))

<a id="nestedobjatt--applications--mobile--integrity_detection"></a>
### Nested Schema for `applications.mobile.integrity_detection`

Read-Only:

- `cache_duration` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--mobile--integrity_detection--cache_duration))
- `mode` (String)

<a id="nestedobjatt--applications--mobile--integrity_detection--cache_duration"></a>
### Nested Schema for `applications.mobile.integrity_detection.mode`

Read-Only:

- `amount` (Number)
- `units` (String)



<a id="nestedobjatt--applications--mobile--passcode_refresh_duration"></a>
### Nested Schema for `applications.mobile.passcode_refresh_duration`

Read-Only:

- `duration` (Number)
- `time_unit` (String)



<a id="nestedobjatt--applications--sp_encryption"></a>
### Nested Schema for `applications.sp_encryption`

Read-Only:

- `algorithm` (String)
- `certificates` (String)


<a id="nestedobjatt--applications--sp_verification"></a>
### Nested Schema for `applications.sp_verification`

Read-Only:

- `authn_request_signed` (String)
- `certificates` (List of String)


