resource "pingone_application_grant" "application_grant" {
  environment_id = local.environment_id
  application_id = "appId"

  scopes = ["resource:read", "resource:create"]
}
