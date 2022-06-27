package signonpolicyassignments

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.ApplicationSignOnPolicyAssignment {
	_, _, id, _ := ParseId(data.Id())

	signOnPolicy := make(map[string]string)
	signOnPolicy["id"] = data.Get("sign_on_policy_id").(string)

	grant := models.ApplicationSignOnPolicyAssignment{
		ID:           &id,
		Priority:     utils.Int(data.Get("priority").(int)),
		SignOnPolicy: signOnPolicy,
	}

	return grant
}
