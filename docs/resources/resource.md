---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingone_resource Resource - terraform-provider-pingone"
subcategory: ""
description: |-
  pingone_resource is used to manage Resources for an environment.
---

# pingone_resource (Resource)

`pingone_resource` is used to manage Resources for an environment.

## Example Usage

```terraform
resource "pingone_resource" "resource" {
  environment_id = local.environment_id
  name           = "resource"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Environment Id
- `name` (String)

### Optional

- `access_token_validity_seconds` (Number)
- `audience` (String)
- `description` (String)
- `type` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `resource_id` (String)

## Import

Import is supported using the following syntax:

```shell
# import using the environment and resource id from the API
terraform import pingone_resource.resource environment:resource
```
