data "pingone_application_policy_assignments" "application_policy_assignments" {
  environment_id = local.environment_id
  application_id = "appId"
  # id = "someid"
}
