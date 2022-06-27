data "pingone_certificate_applications" "cert_apps" {
  environment_id = local.environment_id
  certificate_id = "someKeyId"
}
