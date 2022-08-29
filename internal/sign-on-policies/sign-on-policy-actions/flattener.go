package signonpolicyactions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, action *models.SignOnPolicyAction, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(action, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding sign on policy action",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "environment":
			environment := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "environment_id", environment["id"], diags)
		case "registration":
			registration := flattenRegistration(action.Registration)
			utils.SetResourceDataWithDiagnostic(data, key, registration, diags)
		case "attributes":
			attributes := flattenAttributes(action.Attributes)
			utils.SetResourceDataWithDiagnostic(data, key, attributes, diags)
		case "sign_on_policy":
			val := action.SignOnPolicy["id"]
			utils.SetResourceDataWithDiagnostic(data, "policy_id", val, diags)
		case "identity_provider":
			identity_provider := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "identity_provider_id", identity_provider["id"], diags)
		case "condition":
			condition := flattenCondition(action.Condition)
			utils.SetResourceDataWithDiagnostic(data, key, condition, diags)
		case "discovery_rules":
			rules := flattenDiscoveryRules(action.DiscoveryRules)
			utils.SetResourceDataWithDiagnostic(data, key, rules, diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(actions *[]models.SignOnPolicyAction) []map[string]interface{} {
	if actions == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *actions {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		// TODO: Make this better
		delete(target, "condition")
		delete(target, "discovery_rules")

		if item.Attributes != nil {
			target["attributes"] = flattenAttributes(item.Attributes)
		}
		if item.Registration != nil {
			target["registration"] = flattenRegistration(item.Registration)
		}
		if item.DiscoveryRules != nil {
			target["discovery_rules"] = flattenDiscoveryRules(item.DiscoveryRules)
		}
		if item.Condition != nil {
			target["condition"] = flattenCondition(item.Condition)
		}

		items = append(items, target)
	}

	return items
}

func flattenAttributes(in []map[string]interface{}) *schema.Set {
	items := make([]interface{}, 0)

	for _, attribute := range in {
		items = append(items, attribute)
	}

	hash := schema.HashResource(resourceAttributes())
	return schema.NewSet(hash, items)
}

func flattenRegistration(in *models.SignOnPolicyActionRegistration) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	if val, ok := in.Population["id"]; ok {
		target["population_id"] = val
		delete(target, "population")
	}

	hash := schema.HashResource(resourceRegistration())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenDiscoveryRules(in []models.DiscoveryRule) []map[string]interface{} {
	items := make([]map[string]interface{}, 0)

	for _, rule := range in {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(rule, &target); err != nil {
			continue
		}

		if val, ok := rule.IdentityProvider["id"]; ok {
			target["identity_provider_id"] = val
			delete(target, "identity_provider")
		}

		items = append(items, target)
	}

	return items
}

func flattenCondition(in *models.SignOnPolicyCondition) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	if in.Greater != nil {
		target["greater"] = *in.Greater
	}
	if in.SecondsSince != nil {
		target["seconds_since"] = *in.SecondsSince
	}
	if in.Value != nil {
		target["value"] = *in.Value
	}
	if in.Equals != nil {
		target["equals"] = *in.Equals
	}
	if in.Contains != nil {
		target["contains"] = *in.Contains
	}

	hash := schema.HashResource(resourceCondition())
	return schema.NewSet(hash, []interface{}{target})
}
