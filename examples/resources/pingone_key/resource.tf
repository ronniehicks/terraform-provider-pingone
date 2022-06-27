# Create a signing key
resource "pingone_key" "key" {
  environment_id = local.environment_id
  name = "Doc test cert"
  algorithm = "RSA"
  issuer_dn = "CN=Doc test cert, OU=Ping Identity, O=Ping Identity, L=, ST=, C=US"
  usage_type = "SIGNING"
  default = false
  key_length = 2048
  signature_algorithm = "SHA256withRSA"
  subject_dn = "CN=some.fake.com"
  validity_period = 42
}
