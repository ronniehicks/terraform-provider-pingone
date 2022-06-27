data "pingone_environments" "environments" {
  filter = "name sw \"something\""
  # filter = "organization.id eq \"orgId\""
  # id = "someid"
}
