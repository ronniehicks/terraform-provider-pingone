data "pingone_application_grants" "application_grants" {
  environment_id = local.environment_id
  application_id = "appId"
  # id = "someid"
}
