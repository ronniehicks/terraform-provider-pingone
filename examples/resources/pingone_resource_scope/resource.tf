resource "pingone_resource_scope" "resource_scope" {
  environment_id = local.environment_id
  resource_id    = "resourceId"

  name = "something"
}
