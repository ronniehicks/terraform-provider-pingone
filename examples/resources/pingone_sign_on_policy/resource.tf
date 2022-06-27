resource "pingone_sign_on_policy" "sign_on_policy" {
  environment_id = local.environment_id

  name = "something"
}
