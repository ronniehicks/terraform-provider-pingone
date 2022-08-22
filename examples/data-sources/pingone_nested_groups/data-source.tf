data "pingone_nested_groups" "groups" {
  environment_id = local.environment_id
  group_id       = "someid"
  # nested_group_id = "someid"
}
