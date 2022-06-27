data "pingone_key_applications" "key_apps" {
  environment_id = local.environment_id
  key_id = "someKeyId"
}
