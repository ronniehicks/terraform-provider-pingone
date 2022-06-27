data "pingone_sign_on_policy_actions" "sign_on_policy_actions" {
  environment_id = local.environment_id
  policy_id      = "policyId"
  # id = "someid"
}
