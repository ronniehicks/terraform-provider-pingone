package scopes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Scope {
	_, resourceId, id, _ := ParseId(data.Id())

	resource := make(map[string]string)
	resource["id"] = resourceId
	scope := models.Scope{
		ID:       &id,
		Name:     utils.String(data.Get("name").(string)),
		Resource: resource,
	}

	if val, ok := data.GetOk("description"); ok {
		scope.Description = utils.String(val.(string))
	}

	return scope
}
