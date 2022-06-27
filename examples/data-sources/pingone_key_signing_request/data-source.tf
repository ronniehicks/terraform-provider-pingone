data "pingone_key_signing_request" "key_csr_export" {
  environment_id = local.environment_id
  id = "43ee922d-7093-4487-a963-f973f11d3076"
  export_format = "pkcs10"
}
