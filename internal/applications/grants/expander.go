package grants

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.ApplicationGrant {
	_, applicationId, id, _ := ParseId(data.Id())

	application := make(map[string]string)
	application["id"] = applicationId
	resource := make(map[string]string)
	resource["id"] = data.Get("resource_id").(string)

	scopes := make([]map[string]string, 0)

	dScopes := utils.ExpandStringList(data.Get("scopes").([]interface{}))
	for _, dScope := range dScopes {
		scope := make(map[string]string)
		scope["id"] = *dScope
		scopes = append(scopes, scope)
	}

	grant := models.ApplicationGrant{
		ID:          &id,
		Application: application,
		Resource:    resource,
		Scopes:      scopes,
	}

	return grant
}
