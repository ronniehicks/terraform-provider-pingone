data "pingone_application_attributes" "application_attributes" {
  environment_id = local.environment_id
  application_id = "appId"
  # id = "someid"
}
