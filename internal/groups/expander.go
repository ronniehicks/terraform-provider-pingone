package groups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Group {
	_, id, _ := ParseId(data.Id())

	group := models.Group{
		ID:   &id,
		Name: utils.String(data.Get("name").(string)),
	}

	if val, ok := data.GetOk("description"); ok {
		group.Description = utils.String(val.(string))
	}
	if val, ok := data.GetOk("external_id"); ok {
		group.ExternalId = utils.String(val.(string))
	}
	if val, ok := data.GetOk("user_filter"); ok {
		group.UserFilter = utils.String(val.(string))
	}
	if val, ok := data.GetOk("population_id"); ok {
		populationId := val.(string)
		population := make(map[string]string)
		population["id"] = populationId
		group.Population = population
	}
	if val, ok := data.GetOk("custom_data"); ok {
		group.CustomData = val.(map[string]interface{})
	}

	return group
}
