package attributes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.ResourceAttribute {
	_, _, id, _ := ParseId(data.Id())

	out := models.ResourceAttribute{
		ID:    &id,
		Name:  utils.String(data.Get("name").(string)),
		Value: utils.String(data.Get("value").(string)),
	}

	if val, ok := data.GetOk("type"); ok {
		out.Type = utils.String(val.(string))
	}

	return out
}
