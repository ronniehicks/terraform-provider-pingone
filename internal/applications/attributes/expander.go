package attributes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.ApplicationAttribute {
	_, applicationId, id, _ := ParseId(data.Id())

	application := make(map[string]string)
	application["id"] = applicationId
	attribute := models.ApplicationAttribute{
		ID:          &id,
		Name:        utils.String(data.Get("name").(string)),
		Required:    utils.Bool(data.Get("required").(bool)),
		Value:       utils.String(data.Get("value").(string)),
		Application: application,
	}

	if val, ok := data.GetOk("mapping_type"); ok {
		attribute.MappingType = utils.String(val.(string))
	}

	return attribute
}
