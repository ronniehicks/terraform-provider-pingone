package applications

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func FlattenMany(in *[]models.Application) []map[string]interface{} {
	if in == nil {
		return make([]map[string]interface{}, 0)
	}

	out := make([]map[string]interface{}, 0)
	for _, item := range *in {
		target := make(map[string]interface{})
		if err := mapstructure.Decode(item, &target); err != nil {
			continue
		}

		for key := range target {
			switch key {
			case "icon":
				icon := flattenIcon(item.Icon)
				target[key] = icon
			case "access_control":
				thing := flattenAccessControl(item.AccessControl)
				target[key] = thing
			case "mobile":
				thing := flattenMobile(item.Mobile)
				target[key] = thing
			case "idp_signing":
				thing := flattenIdpSigning(item.IdpSigning)
				target[key] = thing
			case "sp_encryption":
				thing := flattenSpEncryption(item.SpEncryption)
				target[key] = thing
			case "sp_verification":
				thing := flattenSpVerification(item.SpVerification)
				target[key] = thing
			}
		}

		out = append(out, target)
	}

	return out
}

func Flatten(data *schema.ResourceData, in *models.Application, diags *diag.Diagnostics) diag.Diagnostics {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure decoding application",
			Detail:   err.Error(),
		})
	}

	for key, value := range target {
		switch key {
		case "environment":
			environment := value.(map[string]string)
			utils.SetResourceDataWithDiagnostic(data, "environment_id", environment["id"], diags)
		case "icon":
			icon := flattenIcon(in.Icon)
			utils.SetResourceDataWithDiagnostic(data, key, icon, diags)
		case "access_control":
			thing := flattenAccessControl(in.AccessControl)
			utils.SetResourceDataWithDiagnostic(data, key, thing, diags)
		case "mobile":
			thing := flattenMobile(in.Mobile)
			utils.SetResourceDataWithDiagnostic(data, key, thing, diags)
		case "idp_signing":
			thing := flattenIdpSigning(in.IdpSigning)
			utils.SetResourceDataWithDiagnostic(data, key, thing, diags)
		case "sp_encryption":
			thing := flattenSpEncryption(in.SpEncryption)
			utils.SetResourceDataWithDiagnostic(data, key, thing, diags)
		case "sp_verification":
			thing := flattenSpVerification(in.SpVerification)
			utils.SetResourceDataWithDiagnostic(data, key, thing, diags)
		default:
			utils.SetResourceDataWithDiagnostic(data, key, value, diags)
		}
	}

	return *diags
}

func flattenIcon(in *models.Icon) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	target["id"] = *in.ID
	target["href"] = *in.Href

	hash := schema.HashResource(resourceIcon())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenAccessControl(in *models.ApplicationAccessControl) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	delete(target, "group")
	delete(target, "role")

	if in.Group != nil {
		target["group"] = flattenAccessControlGroup(in.Group)
	}

	if in.Role != nil {
		role := make(map[string]interface{})
		if err := mapstructure.Decode(in.Role, &role); err == nil {
			target["role"] = role
		}
	}

	hash := schema.HashResource(resourceAccessControl())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenAccessControlGroup(in *models.ApplicationAccessControlGroup) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	groups := make([]interface{}, 0)
	for _, group := range in.Groups {
		groups = append(groups, group["id"])
	}

	target["type"] = *in.Type
	target["groups"] = groups

	hash := schema.HashResource(resourceAccessControlGroup())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenMobile(in *models.ApplicationMobile) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	delete(target, "integrity_detection")
	delete(target, "passcode_refresh_duration")

	target["bundle_id"] = *in.BundleID
	target["package_name"] = *in.PackageName

	if in.IntegrationDetection != nil {
		target["integrity_detection"] = flattenIntegrityDetection(in.IntegrationDetection)
	}

	if in.PasscodeRefreshDuration != nil {
		target["passcode_refresh_duration"] = flattenPasscodeRefreshDuration(in.PasscodeRefreshDuration)
	}

	hash := schema.HashResource(resourceMobile())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenPasscodeRefreshDuration(in *models.ApplicationMobilePasscodeRefreshDuration) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	target["duration"] = *in.Duration
	target["time_unit"] = *in.TimeUnit

	hash := schema.HashResource(resourcePasscodeRefreshDuration())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenIntegrityDetection(in *models.ApplicationMobileIntegrityDetection) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	delete(target, "cache_duration")

	if in.CacheDuration != nil {
		target["cache_duration"] = flattenCacheDuration(in.CacheDuration)
	}

	hash := schema.HashResource(resourceIntegrityDetection())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenCacheDuration(in *models.ApplicationMobileIntegrityDetectionCacheDuration) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	hash := schema.HashResource(resourceCacheDuration())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenIdpSigning(in *models.ApplicationIdpSigning) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	target["key"] = in.Key["id"]
	target["algorithm"] = *in.Algorithm

	hash := schema.HashResource(resourceIdpSigning())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenSpEncryption(in *models.ApplicationSpEncryption) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	delete(target, "certificates")

	if in.Certificates != nil {
		target["certificates"] = in.Certificates.ID
	}

	hash := schema.HashResource(resourceMobile())
	return schema.NewSet(hash, []interface{}{target})
}

func flattenSpVerification(in *models.ApplicationSpVerification) *schema.Set {
	target := make(map[string]interface{})
	if err := mapstructure.Decode(in, &target); err != nil {
		return nil
	}

	certificates := make([]string, 0)
	for _, cert := range in.Certificates {
		certificates = append(certificates, *cert.ID)
	}

	target["certificates"] = certificates

	hash := schema.HashResource(resourceSpVerification())
	return schema.NewSet(hash, []interface{}{target})
}
