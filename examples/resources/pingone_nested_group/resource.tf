resource "pingone_nested_group" "group" {
  environment_id  = local.environment_id
  group_id        = "someid"
  nested_group_id = "nestedid"
}
