resource "pingone_application_attribute" "application_attribute" {
  environment_id = local.environment_id
  application_id = "appId"

  name     = "firstName"
  required = true
  value    = "$${user.name.given}"
}
