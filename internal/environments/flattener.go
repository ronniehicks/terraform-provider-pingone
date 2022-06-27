package environments

import (
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func FlattenMany(environments *[]models.Environment) []map[string]interface{} {
	if environments == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *environments {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		items = append(items, target)
	}

	return items
}
