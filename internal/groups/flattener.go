package groups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, group *models.Group, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(group, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding group",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "population":
			population := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "population_id", population["id"], diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(groups *[]models.Group) []map[string]interface{} {
	if groups == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *groups {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		items = append(items, target)
	}

	return items
}
