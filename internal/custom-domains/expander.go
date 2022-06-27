package customdomains

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.CustomDomain {
	_, id, _ := ParseId(data.Id())

	out := models.CustomDomain{
		ID:         &id,
		DomainName: utils.String(data.Get("domain_name").(string)),
	}

	if val, ok := data.GetOk("canonical_name"); ok {
		out.CanonicalName = utils.String(val.(string))
	}
	if val, ok := data.GetOk("status"); ok {
		out.Status = utils.String(val.(string))
	}
	if val, ok := data.GetOk("certificate_expiration"); ok {
		expiration := val.(string)
		certificate := make(map[string]string)
		certificate["expiresAt"] = expiration
		out.Certificate = certificate
	}

	return out
}
