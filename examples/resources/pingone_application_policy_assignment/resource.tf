resource "pingone_application_policy_assignment" "application_policy_assignment" {
  environment_id    = local.environment_id
  application_id    = "appId"
  sign_on_policy_id = "policyId"

  priority = 1
}
