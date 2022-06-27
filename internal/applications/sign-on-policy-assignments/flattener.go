package signonpolicyassignments

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, assignment *models.ApplicationSignOnPolicyAssignment, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(assignment, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding application sign on policy assignment",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "environment":
			environment := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "environment_id", environment["id"], diags)
		case "application":
			application := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "application_id", application["id"], diags)
		case "sign_on_policy":
			signOnPolicy := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "sign_on_policy_id", signOnPolicy["id"], diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(grants *[]models.ApplicationSignOnPolicyAssignment) []map[string]interface{} {
	if grants == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *grants {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		items = append(items, target)
	}

	return items
}
