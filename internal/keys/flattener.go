package keys

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, key *models.Key, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(key, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding key",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "serial_number":
			newValue := value.(*big.Int).String()
			utils.SetResourceDataWithDiagnostic(data, key, newValue, diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(keys *[]models.Key) []map[string]interface{} {
	if keys == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *keys {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		target["serial_number"] = target["serial_number"].(*big.Int).String()

		items = append(items, target)
	}

	return items
}
