data "pingone_resource_scopes" "resource_scopes" {
  environment_id = local.environment_id
  resource_id    = "resourceId"
  # id = "someid"
}
