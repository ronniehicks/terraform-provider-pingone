package signonpolicyactions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Expand(data *schema.ResourceData) models.SignOnPolicyAction {
	_, _, id, _ := ParseId(data.Id())

	action := models.SignOnPolicyAction{
		ID:       &id,
		Priority: utils.Int(data.Get("priority").(int)),
		Type:     utils.String(data.Get("type").(string)),
	}

	if val, ok := data.GetOk("confirm_identity_provider_attributes"); ok {
		action.ConfirmIdentityProviderAttributes = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("enforce_lockout_for_identity_providers"); ok {
		action.EnforceLockoutForIdentityProviders = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("recovery"); ok {
		item := make(map[string]bool)
		err := mapstructure.Decode(val.(map[string]interface{}), &item)
		if err == nil {
			action.Recovery = item
		}
	}
	if val, ok := data.GetOk("registration"); ok {
		action.Registration = expandRegistration(val.(*schema.Set).List())
	}
	if val, ok := data.GetOk("social_providers"); ok {
		action.SocialProviders = val.([]map[string]string)
	}
	if val, ok := data.GetOk("authenticator"); ok {
		action.Authenticator = val.(map[string]bool)
	}
	if val, ok := data.GetOk("bound_biometrics"); ok {
		action.BoundBiometrics = val.(map[string]bool)
	}
	if val, ok := data.GetOk("email"); ok {
		action.Email = val.(map[string]bool)
	}
	if val, ok := data.GetOk("security_key"); ok {
		action.SecurityKey = val.(map[string]bool)
	}
	if val, ok := data.GetOk("sms"); ok {
		action.Sms = val.(map[string]bool)
	}
	if val, ok := data.GetOk("voice"); ok {
		action.Voice = val.(map[string]bool)
	}
	if val, ok := data.GetOk("applications"); ok {
		items := utils.ExpandList[models.MultiFactorAuthenticationActionApplication](val.([]map[string]interface{}))
		action.Applications = items
	}
	if val, ok := data.GetOk("no_device_mode"); ok {
		action.NoDeviceMode = utils.String(val.(string))
	}
	if val, ok := data.GetOk("agreement"); ok {
		action.Agreement = val.(map[string]string)
	}
	if val, ok := data.GetOk("discovery_rules"); ok {
		action.DiscoveryRules = expandDiscoveryRules(val.(*schema.Set).List())
	}
	if val, ok := data.GetOk("identity_provider"); ok {
		action.IdentityProvider = val.(map[string]string)
	}
	if val, ok := data.GetOk("acr_values"); ok {
		action.AcrValues = utils.String(val.(string))
	}
	if val, ok := data.GetOk("pass_user_context"); ok {
		action.PassUserContext = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("prevent_multiple_prompts_per_flow"); ok {
		action.PreventMultiplePromptsPerFlow = utils.Bool(val.(bool))
	}
	if val, ok := data.GetOk("prompt_interval_seconds"); ok {
		action.PromptIntervalSeconds = utils.Int(val.(int))
	}
	if val, ok := data.GetOk("prompt_text"); ok {
		action.PromptText = utils.String(val.(string))
	}
	if val, ok := data.GetOk("attributes"); ok {
		items := utils.ExpandSet[map[string]interface{}](val.(*schema.Set).List())
		action.Attributes = items
	}

	if val, ok := data.GetOk("condition"); ok {
		items := utils.ExpandSet[models.SignOnPolicyCondition](val.(*schema.Set).List())
		action.Condition = &items[0]
	}

	return action
}

func expandRegistration(in []interface{}) *models.SignOnPolicyActionRegistration {
	item := in[0]
	var registration models.SignOnPolicyActionRegistration
	if err := mapstructure.Decode(item, &registration); err != nil {
		return nil
	}

	for _, raw := range in {
		top := raw.(map[string]interface{})
		if val, ok := top["population_id"]; ok {
			population := make(map[string]string)
			population["id"] = val.(string)
			registration.Population = population
		}
	}

	return &registration
}

func expandDiscoveryRules(in []interface{}) []models.DiscoveryRule {
	discoveryRules := make([]models.DiscoveryRule, 0)

	for _, raw := range in {
		var target models.DiscoveryRule
		if err := mapstructure.Decode(raw, &target); err != nil {
			continue
		}

		top := raw.(map[string]interface{})
		if val, ok := top["identity_provider_id"]; ok {
			identityProvider := make(map[string]string)
			identityProvider["id"] = val.(string)
			target.IdentityProvider = identityProvider
		}

		discoveryRules = append(discoveryRules, target)
	}

	return discoveryRules
}
