package attributes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.ProviderAttribute {
	_, idendityProviderId, id, _ := ParseId(data.Id())

	identityProvider := make(map[string]string)
	identityProvider["id"] = idendityProviderId
	attribute := models.ProviderAttribute{
		ID:               &id,
		Name:             utils.String(data.Get("name").(string)),
		Value:            utils.String(data.Get("value").(string)),
		IdentityProvider: identityProvider,
		Update:           utils.String(data.Get("update").(string)),
	}

	if val, ok := data.GetOk("mapping_type"); ok {
		attribute.MappingType = utils.String(val.(string))
	}

	return attribute
}
