package signonpolicies

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, scope *models.SignOnPolicy, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(scope, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding sign on policy",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "environment":
			environment := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "environment_id", environment["id"], diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(signOnPolicies *[]models.SignOnPolicy) []map[string]interface{} {
	if signOnPolicies == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *signOnPolicies {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		delete(target, "environment")

		items = append(items, target)
	}

	return items
}
