package attributes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, scope *models.ProviderAttribute, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(scope, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding provider attribute",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "identity_provider":
			identityProvider := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "identity_provider_id", identityProvider["id"], diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(attributes *[]models.ProviderAttribute) []map[string]interface{} {
	if attributes == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *attributes {
		target := make(map[string]interface{})

		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		for key := range target {
			switch key {
			case "identity_provider":
				target["identity_provider_id"] = item.IdentityProvider["id"]
				delete(target, "identity_provider")
			}
		}

		items = append(items, target)
	}

	return items
}
