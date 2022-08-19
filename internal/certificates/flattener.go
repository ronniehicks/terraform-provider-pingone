package certificates

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, certificate *models.Certificate, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(certificate, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding certificate",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		utils.SetResourceDataWithDiagnostic(data, key, value, diags)
	}

	return *diags
}

func FlattenMany(certificates *[]models.Certificate) []map[string]interface{} {
	if certificates == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *certificates {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		target["serial_number"] = target["serial_number"].(*big.Int).String()

		items = append(items, target)
	}

	return items
}
