package populations

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Population {
	_, id, _ := ParseId(data.Id())

	population := models.Population{
		ID:   &id,
		Name: utils.String(data.Get("name").(string)),
	}

	if val, ok := data.GetOk("description"); ok {
		population.Description = utils.String(val.(string))
	}
	if val, ok := data.GetOk("password_policy_id"); ok {
		passwordPolicyId := val.(string)
		passwordPolicy := make(map[string]string)
		passwordPolicy["id"] = passwordPolicyId
		population.PasswordPolicy = passwordPolicy
	}

	return population
}
