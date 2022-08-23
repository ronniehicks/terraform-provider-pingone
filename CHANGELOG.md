## 1.1.0 (August 22, 2022)

FEATURES:

* **New Resource**: `pingone_nested_group`

* **New Data Source**: `pingone_nested_groups`

## 1.0.0 (August 19, 2022)

BREAKING CHANGES:

* Data Source `pingone_identity_provider` has been renamed to `pingone_identity_providers`

IMPROVEMENTS:

* resource/pingone_key: add support to provide either `expires_at` or `validity_period`

BUG FIXES:

* data/pingone_certificates/keys: treat `serial_number` as string because value can be _very_ large number

## 0.0.6 (August 15, 2022)

BUG FIXES:

* data/pingone_applications and resource/pingone_application
  * treat `authn_request_signed` as string because of lack of tri-state support in TF SDK
  * remove default value for `pkce_encforcement`
* data/pingone_applications: fix application id filter property

## 0.0.3 (June 28, 2022)

FEATURES:

* **New Resource**: `pingone_identity_provider`
* **New Resource**: `pingone_identity_provider_attribute`
* **New Resource**: `pingone_key`

* **New Data Source**: `pingone_identity_providers`
* **New Data Source**: `pingone_identity_provider_attributes`
* **New Data Source**: `pingone_certificates`
* **New Data Source**: `pingone_certificate_applications`
* **New Data Source**: `pingone_keys`
* **New Data Source**: `pingone_key_applications`
* **New Data Source**: `pingone_key_export`
* **New Data Source**: `pingone_key_signing_request`

## 0.0.1 (June 09, 2022)

FEATURES:

* **New Resource**: `pingone_application`
* **New Resource**: `pingone_application_attribute`
* **New Resource**: `pingone_application_grant`
* **New Resource**: `pingone_application_policy_assignment`
* **New Resource**: `pingone_custom_domain`
* **New Resource**: `pingone_group`
* **New Resource**: `pingone_population`
* **New Resource**: `pingone_resource`
* **New Resource**: `pingone_resource_scope`
* **New Resource**: `pingone_resource_attribute`
* **New Resource**: `pingone_sign_on_policy`
* **New Resource**: `pingone_sign_on_policy_action`

* **New Data Source**: `pingone_environments`
* **New Data Source**: `pingone_applications`
* **New Data Source**: `pingone_application_attributes`
* **New Data Source**: `pingone_application_grants`
* **New Data Source**: `pingone_application_policy_assignments`
* **New Data Source**: `pingone_custom_domains`
* **New Data Source**: `pingone_groups`
* **New Data Source**: `pingone_populations`
* **New Data Source**: `pingone_resources`
* **New Data Source**: `pingone_resource_scopes`
* **New Data Source**: `pingone_resource_attributes`
* **New Data Source**: `pingone_sign_on_policies`
* **New Data Source**: `pingone_sign_on_policie_actions`
