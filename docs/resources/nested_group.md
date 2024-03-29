---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_nested_group Resource - terraform-provider-pingone"
subcategory: ""
description: |-
  pingone_nested_group is used to manage Group Nesting for an environment.
---

# pingone_nested_group (Resource)

`pingone_nested_group` is used to manage Group Nesting for an environment.

## Example Usage

```terraform
resource "pingone_nested_group" "group" {
  environment_id  = local.environment_id
  group_id        = "someid"
  nested_group_id = "nestedid"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Environment Id
- `group_id` (String) Parent Group Id
- `nested_group_id` (String) Child Group Id

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
# import using the environment, group id and nested group id from the API
terraform import pingone_nested_group.group environment:group:nested
```
