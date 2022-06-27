data "pingone_key_export" "key_text_export" {
  environment_id = local.environment_id
  id = "someKeyId"
  # export_format = "p7b"
}
