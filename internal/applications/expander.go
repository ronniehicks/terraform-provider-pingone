package applications

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.Application {
	_, id, _ := ParseId(data.Id())

	app := models.Application{
		ID:       &id,
		Name:     utils.String(data.Get("name").(string)),
		Type:     utils.String(data.Get("type").(string)),
		Enabled:  utils.Bool(data.Get("enabled").(bool)),
		Protocol: utils.String(data.Get("protocol").(string)),
	}

	// TODO: Reflect this?

	if val, ok := data.GetOk("description"); ok {
		app.Description = utils.String(val.(string))
	}
	if val, ok := data.GetOk("home_page_url"); ok {
		app.HomePageUrl = utils.String(val.(string))
	}
	if val, ok := data.GetOk("login_page_url"); ok {
		app.LoginPageUrl = utils.String(val.(string))
	}
	if val, ok := data.GetOk("assign_actor_roles"); ok {
		app.AssignActorRoles = utils.Bool(val.(bool))
	}

	// OIDC
	if val, ok := data.GetOk("token_endpoint_auth_method"); ok {
		app.TokenEndpointAuthMethod = utils.String(val.(string))
	}
	if val, ok := data.GetOk("pkce_enforcement"); ok {
		app.PkceEnforcement = utils.String(val.(string))
	}
	if val, ok := data.GetOk("refresh_token_duration"); ok {
		app.RefreshTokenDuration = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("refresh_token_rolling_duration"); ok {
		app.RefreshTokenRollingDuration = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("support_unsigned_request_object"); ok {
		app.SupportUnsignedRequestObject = utils.Bool(val.(bool))
	}

	// Native App
	if val, ok := data.GetOk("bundle_id"); ok {
		app.BundleID = utils.String(val.(string))
	}
	if val, ok := data.GetOk("package_name"); ok {
		app.PackageName = utils.String(val.(string))
	}

	// SAML
	if val, ok := data.GetOk("sp_entity_id"); ok {
		app.SpEntityId = utils.String(val.(string))
	}
	if val, ok := data.GetOk("assertion_duration"); ok {
		app.AssertionDuration = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("default_target_url"); ok {
		app.DefaultTargetUrl = utils.String(val.(string))
	}
	if val, ok := data.GetOk("slo_binding"); ok {
		app.SloBinding = utils.String(val.(string))
	}
	if val, ok := data.GetOk("slo_endpoint"); ok {
		app.SloEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("slo_response_endpoint"); ok {
		app.SloResponseEndpoint = utils.String(val.(string))
	}
	if val, ok := data.GetOk("response_signed"); ok {
		app.ResponseSigned = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("assertion_signed"); ok {
		app.AssertionSigned = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("name_id_format"); ok {
		app.NameIdFormat = utils.String(val.(string))
	}

	if val, ok := data.GetOk("grant_types"); ok {
		strs := utils.ExpandStringList(val.([]interface{}))
		app.GrantTypes = strs
	}

	if val, ok := data.GetOk("acs_urls"); ok {
		strs := utils.ExpandStringList(val.([]interface{}))
		app.AcsUrls = strs
	}

	if val, ok := data.GetOk("post_logout_redirect_uris"); ok {
		strs := utils.ExpandStringList(val.([]interface{}))
		app.PostLogoutRedirectUris = strs
	}

	if val, ok := data.GetOk("redirect_uris"); ok {
		strs := utils.ExpandStringList(val.([]interface{}))
		app.RedirectUris = strs
	}

	if val, ok := data.GetOk("response_types"); ok {
		strs := utils.ExpandStringList(val.([]interface{}))
		app.ResponseTypes = strs
	}

	if val, ok := data.GetOk("access_control"); ok {
		app.AccessControl = expandAccessControl(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("icon"); ok {
		app.Icon = expandIcon(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("mobile"); ok {
		app.Mobile = expandMobile(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("idp_signing"); ok {
		app.IdpSigning = expandIdpSigning(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("sp_encryption"); ok {
		app.SpEncryption = expandSpEncryption(val.(*schema.Set).List())
	}

	if val, ok := data.GetOk("sp_verification"); ok {
		app.SpVerification = expandSpVerification(val.(*schema.Set).List())
	}

	return app
}

func expandAccessControl(in []interface{}) *models.ApplicationAccessControl {
	result := models.ApplicationAccessControl{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["group"]; ok {
			result.Group = expandAccessControlGroup(val.(*schema.Set).List())
		}

		if val, ok := top["role"]; ok {
			role := make(map[string]string)
			for key, v := range val.(map[string]interface{}) {
				role[key] = v.(string)
			}
			result.Role = role
		}
	}

	return &result
}

func expandAccessControlGroup(in []interface{}) *models.ApplicationAccessControlGroup {
	if len(in) < 1 {
		return nil
	}

	result := models.ApplicationAccessControlGroup{}

	for _, raw := range in {
		sub := raw.(map[string]interface{})
		if val, ok := sub["type"]; ok {
			result.Type = utils.String(val.(string))
		}
		if val, ok := sub["groups"]; ok {
			g := utils.ExpandStringList(val.([]interface{}))
			var groups []map[string]string
			for _, group := range g {
				gb := make(map[string]string)
				gb["id"] = *group
				groups = append(groups, gb)
			}
			result.Groups = groups
		}
	}

	return &result
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

func expandMobile(in []interface{}) *models.ApplicationMobile {
	result := models.ApplicationMobile{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["bundle_id"]; ok {
			result.BundleID = utils.String(val.(string))
		}

		if val, ok := top["package_name"]; ok {
			result.PackageName = utils.String(val.(string))
		}

		if val, ok := top["integrity_detection"]; ok {
			result.IntegrationDetection = expandMobileIntegrityDetection(val.(*schema.Set).List())
		}

		if val, ok := top["passcode_refresh_duration"]; ok {
			result.PasscodeRefreshDuration = expandMobilePasscodeRefreshDuration(val.(*schema.Set).List())
		}
	}

	return &result
}

func expandMobileIntegrityDetection(in []interface{}) *models.ApplicationMobileIntegrityDetection {
	result := models.ApplicationMobileIntegrityDetection{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["mode"]; ok {
			result.Mode = utils.String(val.(string))
		}

		if val, ok := top["cache_duration"]; ok {
			result.CacheDuration = expandMobileIntegrityDetectionCacheDuration(val.(*schema.Set).List())
		}
	}

	return &result
}

func expandMobileIntegrityDetectionCacheDuration(in []interface{}) *models.ApplicationMobileIntegrityDetectionCacheDuration {
	result := models.ApplicationMobileIntegrityDetectionCacheDuration{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["amount"]; ok {
			result.Amount = utils.Int(val.(int))
		}

		if val, ok := top["units"]; ok {
			result.Units = utils.String(val.(string))
		}
	}

	return &result
}

func expandMobilePasscodeRefreshDuration(in []interface{}) *models.ApplicationMobilePasscodeRefreshDuration {
	result := models.ApplicationMobilePasscodeRefreshDuration{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["time_unit"]; ok {
			result.TimeUnit = utils.String(val.(string))
		}

		if val, ok := top["duration"]; ok {
			result.Duration = utils.Int(val.(int))
		}
	}

	return &result
}

func expandIdpSigning(in []interface{}) *models.ApplicationIdpSigning {
	result := models.ApplicationIdpSigning{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["key"]; ok {
			key := make(map[string]string)
			key["id"] = val.(string)
			result.Key = key
		}

		if val, ok := top["algorithm"]; ok {
			result.Algorithm = utils.String(val.(string))
		}
	}

	return &result
}

func expandSpEncryption(in []interface{}) *models.ApplicationSpEncryption {
	result := models.ApplicationSpEncryption{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["algorithm"]; ok {
			result.Algorithm = utils.String(val.(string))
		}
		if val, ok := top["certificates"]; ok {
			certId := utils.String(val.(string))
			result.Certificates = &models.ApplicationCertificates{ID: certId}
		}
	}

	return &result
}

func expandSpVerification(in []interface{}) *models.ApplicationSpVerification {
	result := models.ApplicationSpVerification{}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["authn_request_signed"]; ok {
			result.AuthnRequestSigned = utils.Bool(val.(bool))
		}
		if val, ok := top["groups"]; ok {
			g := utils.ExpandStringList(val.([]interface{}))
			var certificates []*models.ApplicationCertificates
			for _, id := range g {
				cert := models.ApplicationCertificates{ID: id}
				certificates = append(certificates, &cert)
			}
			result.Certificates = certificates
		}
	}

	return &result
}
