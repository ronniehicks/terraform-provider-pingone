package keys

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Key {
	_, id, _ := ParseId(data.Id())

	key := models.Key{
		ID:   &id,
		Name: utils.String(data.Get("name").(string)),
	}

	if val, ok := data.GetOk("algorithm"); ok {
		key.Algorithm = utils.String(val.(string))
	}
	if val, ok := data.GetOk("created_at"); ok {
		key.CreatedAt = utils.String(val.(string))
	}
	if val, ok := data.GetOk("default"); ok {
		key.Default = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("expires_at"); ok {
		key.ExpiresAt = utils.String(val.(string))
	}
	if val, ok := data.GetOk("key_id"); ok {
		key.ID = utils.String(val.(string))
	}
	if val, ok := data.GetOk("issuer_dn"); ok {
		key.IssuerDN = utils.String(val.(string))
	}
	if val, ok := data.GetOk("key_length"); ok {
		key.KeyLength = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("name"); ok {
		key.Name = utils.String(val.(string))
	}
	if val, ok := data.GetOk("serial_number"); ok {
		key.SerialNumber = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("signature_algorithm"); ok {
		key.SignatureAlgorithm = utils.String(val.(string))
	}
	if val, ok := data.GetOk("starts_at"); ok {
		key.StartsAt = utils.String(val.(string))
	}
	if val, ok := data.GetOk("status"); ok {
		key.Status = utils.String(val.(string))
	}
	if val, ok := data.GetOk("subject_dn"); ok {
		key.SubjectDN = utils.String(val.(string))
	}
	if val, ok := data.GetOk("usage_type"); ok {
		key.UsageType = utils.String(val.(string))
	}
	if val, ok := data.GetOk("validity_period"); ok {
		key.ValidityPeriod = utils.Int(val.(int))
	}
	return key
}
