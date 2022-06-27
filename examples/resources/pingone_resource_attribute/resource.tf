resource "pingone_resource_attribute" "resource_attribute" {
  environment_id = local.environment_id
  resource_id    = "resourceId"

  name = "something"
}
