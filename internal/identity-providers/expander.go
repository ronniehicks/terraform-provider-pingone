package identityproviders

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.IdentityProvider {
	_, id, _ := ParseId(data.Id())

	identityProvider := models.IdentityProvider{
		ID:      &id,
		Name:    utils.String(data.Get("name").(string)),
		Enabled: utils.Bool((data.Get("enabled").(bool))),
		Type:    utils.String(data.Get("type").(string)),
	}

	if val, ok := data.GetOk("description"); ok {
		identityProvider.Description = utils.String(val.(string))
	}

	if val, ok := data.GetOk("icon"); ok {
		identityProvider.Icon = expandIcon(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("login_button_icon"); ok {
		identityProvider.LoginButtonIcon = expandIcon(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("registration"); ok {
		identityProvider.Registration = expandRegistration(val.(*schema.Set).List())
	}

	// OpenId Connect Request Model
	if val, ok := data.GetOk("authorization_endpoint"); ok {
		identityProvider.AuthorizationEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("client_id"); ok {
		identityProvider.ClientId = utils.String(val.(string))
	}
	if val, ok := data.GetOk("client_secret"); ok {
		identityProvider.ClientSecret = utils.String(val.(string))
	}
	if val, ok := data.GetOk("discovery_endpoint"); ok {
		identityProvider.DiscoveryEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("issuer"); ok {
		identityProvider.Issuer = utils.String(val.(string))
	}
	if val, ok := data.GetOk("jwks_endpoint"); ok {
		identityProvider.JwksEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("scopes"); ok {
		scopes := make([]string, 0)

		dScopes := utils.ExpandStringList(val.([]interface{}))
		for _, dScope := range dScopes {
			scopes = append(scopes, *dScope)
		}
		identityProvider.Scopes = scopes
	}
	if val, ok := data.GetOk("token_endpoint"); ok {
		identityProvider.TokenEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("token_endpoint_auth_method"); ok {
		identityProvider.TokenEndpointAuthMethod = utils.String(val.(string))
	}
	if val, ok := data.GetOk("user_info_endpoint"); ok {
		identityProvider.UserInfoEndpoint = utils.String(val.(string))
	}

	// SAML Request Model
	identityProvider.AuthnRequestSigned = utils.Bool(data.Get("authn_request_signed").(bool))
	// We cannot use data.GetOk for bool when false - ok get set to false since false is the default bool value

	if val, ok := data.GetOk("idp_entity_id"); ok {
		identityProvider.IdpEntityId = utils.String(val.(string))
	}
	if val, ok := data.GetOk("idp_verification_certificate_ids"); ok {
		certificateIds := make([]string, 0)

		dCertificateIds := utils.ExpandStringList(val.([]interface{}))
		for _, dCertificateId := range dCertificateIds {
			certificateIds = append(certificateIds, *dCertificateId)
		}

		identityProvider.IdpVerification = expandIdpVerification(certificateIds)
	}
	if val, ok := data.GetOk("sp_entity_id"); ok {
		identityProvider.SpEntityId = utils.String(val.(string))
	}
	if val, ok := data.GetOk("sp_signing_key_id"); ok {
		identityProvider.SpSigning = expandSpSigning(val.(string))
	}
	if val, ok := data.GetOk("sso_binding"); ok {
		identityProvider.SsoBinding = utils.String(val.(string))
	}
	if val, ok := data.GetOk("sso_endpoint"); ok {
		identityProvider.SsoEndpoint = utils.String(val.(string))
	}

	return identityProvider
}

func expandIcon(in []interface{}) *models.Icon {
	result := models.Icon{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["id"]; ok {
			result.ID = utils.String(val.(string))
		}

		if val, ok := top["href"]; ok {
			result.Href = utils.String(val.(string))
		}
	}

	return &result
}

func expandRegistration(in []interface{}) *models.Registration {
	result := models.Registration{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["population_id"]; ok {
			population := map[string]string{}
			population["id"] = val.(string)
			result.Population = population
		}
	}

	return &result
}

func expandIdpVerification(in []string) *models.IdpVerification {
	result := models.IdpVerification{}
	certificates := []map[string]string{}

	for _, raw := range in {
		certificate := make(map[string]string)

		certificate["id"] = raw
		certificates = append(certificates, certificate)
	}

	result.Certificates = certificates
	return &result
}

func expandSpSigning(in string) *models.SpSigning {
	result := models.SpSigning{}

	key := make(map[string]string)
	key["id"] = in
	result.Key = key

	return &result
}
