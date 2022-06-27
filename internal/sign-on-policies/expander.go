package signonpolicies

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.SignOnPolicy {
	_, id, _ := ParseId(data.Id())

	signOnPolicy := models.SignOnPolicy{
		ID:   &id,
		Name: utils.String(data.Get("name").(string)),
	}

	if val, ok := data.GetOk("description"); ok {
		signOnPolicy.Description = utils.String(val.(string))
	}
	if val, ok := data.GetOk("default"); ok {
		signOnPolicy.Default = utils.Bool(val.(bool))
	}

	return signOnPolicy
}
