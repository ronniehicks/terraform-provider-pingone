package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Resource {
	_, id, _ := ParseId(data.Id())

	resource := models.Resource{
		ID:   &id,
		Name: utils.String(data.Get("name").(string)),
	}

	if val, ok := data.GetOk("description"); ok {
		resource.Description = utils.String(val.(string))
	}
	if val, ok := data.GetOk("type"); ok {
		resource.Type = utils.String(val.(string))
	}
	if val, ok := data.GetOk("audience"); ok {
		resource.Audience = utils.String(val.(string))
	}
	if val, ok := data.GetOk("access_token_validity_seconds"); ok {
		resource.AccessTokenValiditySeconds = utils.Int(val.(int))
	}

	return resource
}
