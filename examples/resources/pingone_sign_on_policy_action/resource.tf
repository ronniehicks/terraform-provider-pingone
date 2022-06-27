resource "pingone_sign_on_policy_action" "sign_on_policy_action" {
  environment_id = local.environment_id
  policy_id      = "policyId"
  priority       = 1
  type           = "IDENTIFIER_FIRST"
  recovery = {
    enabled = true
  }
  condition {
    seconds_since = "$${session.lastSignOn.withAuthenticator.pwd.at}"
    greater       = 1800
  }
  registration {
    enabled                              = true
    population_id                        = "33f3394f-7b44-4468-b310-2c27d627ac20"
    confirm_identity_provider_attributes = true
  }
  discovery_rules {
    identity_provider_id = "0f39b692-7c1c-4f20-92c8-ee5ca07ef5dc"

    condition = {
      value    = "$${identifier}"
      contains = "@lumeris.com"
    }
  }
}
