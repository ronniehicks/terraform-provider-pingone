package identityproviders

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Flatten(data *schema.ResourceData, identityProvider *models.IdentityProvider, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(identityProvider, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding group",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "environment":
			environment := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "environment_id", environment["id"], diags)
		case "registration":
			registration := flattenRegistration(identityProvider.Registration)
			utils.SetResourceDataWithDiagnostic(data, key, registration, diags)
		case "idp_verification":
			idpVerification := flattenIdpVerification(identityProvider.IdpVerification)
			utils.SetResourceDataWithDiagnostic(data, "idp_verification_certificate_ids", idpVerification, diags)
		case "icon":
			icon := flattenIcon(identityProvider.Icon)
			utils.SetResourceDataWithDiagnostic(data, key, icon, diags)
		case "login_button_icon":
			loginButtonIcon := flattenIcon(identityProvider.LoginButtonIcon)
			utils.SetResourceDataWithDiagnostic(data, key, loginButtonIcon, diags)
		case "sp_signing":
			signingId := flattenSpSigningId(identityProvider.SpSigning)
			utils.SetResourceDataWithDiagnostic(data, "sp_signing_key_id", signingId, diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func FlattenMany(identityProviders *[]models.IdentityProvider) []map[string]interface{} {
	if identityProviders == nil {
		return make([]map[string]interface{}, 0)
	}

	items := make([]map[string]interface{}, 0)

	for _, item := range *identityProviders {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		for key := range target {
			switch key {
			case "registration":
				target["registration"] = flattenRegistration(item.Registration)
			case "idp_verification":
				certs := make([]string, 0)
				for _, certificate := range item.IdpVerification.Certificates {
					certs = append(certs, certificate["id"])
				}
				target["idp_verification_certificate_ids"] = certs
				delete(target, "idp_verification")
			case "icon":
				target["icon"] = flattenIcon(item.Icon)
			case "login_button_icon":
				target["login_button_icon"] = flattenIcon(item.LoginButtonIcon)
			case "sp_signing":
				target["sp_signing_key_id"] = item.SpSigning.Key["id"]
				delete(target, "sp_signing")
			}
		}
		items = append(items, target)
	}

	return items
}

func flattenRegistration(in *models.Registration) *schema.Set {
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

func flattenIcon(in *models.Icon) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	target["id"] = *in.ID
	target["href"] = *in.Href
	delete(target, "ID")
	delete(target, "Href")

	hash := schema.HashResource(resourceIcon())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenSpSigningId(in *models.SpSigning) string {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return ""
	}

	if val, ok := in.Key["id"]; ok {
		return val
	}
	return ""
}

func flattenIdpVerification(in *models.IdpVerification) []string {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	certs := make([]string, 0)
	for _, certificate := range in.Certificates {
		certs = append(certs, certificate["id"])
	}
	target["idp_verification_certificate_ids"] = certs
	delete(target, "certificates")

	return certs
}
