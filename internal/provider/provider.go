package provider

import (
	"context"

	"github.com/ronniehicks/terraform-provider-pingone/internal/applications"
	appAttributes "github.com/ronniehicks/terraform-provider-pingone/internal/applications/attributes"
	appGrants "github.com/ronniehicks/terraform-provider-pingone/internal/applications/grants"
	appPolicies "github.com/ronniehicks/terraform-provider-pingone/internal/applications/sign-on-policy-assignments"
	"github.com/ronniehicks/terraform-provider-pingone/internal/certificates"
	certApps "github.com/ronniehicks/terraform-provider-pingone/internal/certificates/applications"
	customdomains "github.com/ronniehicks/terraform-provider-pingone/internal/custom-domains"
	"github.com/ronniehicks/terraform-provider-pingone/internal/environments"
	"github.com/ronniehicks/terraform-provider-pingone/internal/groups"
	identityproviders "github.com/ronniehicks/terraform-provider-pingone/internal/identity-providers"
	idpattributes "github.com/ronniehicks/terraform-provider-pingone/internal/identity-providers/attributes"
	keys "github.com/ronniehicks/terraform-provider-pingone/internal/keys"
	keyApps "github.com/ronniehicks/terraform-provider-pingone/internal/keys/applications"
	keyExport "github.com/ronniehicks/terraform-provider-pingone/internal/keys/export"
	keyCsr "github.com/ronniehicks/terraform-provider-pingone/internal/keys/signing-request"
	"github.com/ronniehicks/terraform-provider-pingone/internal/populations"
	"github.com/ronniehicks/terraform-provider-pingone/internal/resources"
	"github.com/ronniehicks/terraform-provider-pingone/internal/resources/attributes"
	"github.com/ronniehicks/terraform-provider-pingone/internal/resources/scopes"
	signonpolicies "github.com/ronniehicks/terraform-provider-pingone/internal/sign-on-policies"
	signonpolicyactions "github.com/ronniehicks/terraform-provider-pingone/internal/sign-on-policies/sign-on-policy-actions"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"client_secret": {
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
				},
				"token_url": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"api_url": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"pingone_application":                   applications.Resource(),
				"pingone_application_attribute":         appAttributes.Resource(),
				"pingone_application_grant":             appGrants.Resource(),
				"pingone_application_policy_assignment": appPolicies.Resource(),
				"pingone_custom_domain":                 customdomains.Resource(),
				"pingone_group":                         groups.Resource(),
				"pingone_population":                    populations.Resource(),
				"pingone_resource":                      resources.Resource(),
				"pingone_resource_scope":                scopes.Resource(),
				"pingone_resource_attribute":            attributes.Resource(),
				"pingone_sign_on_policy":                signonpolicies.Resource(),
				"pingone_sign_on_policy_action":         signonpolicyactions.Resource(),
				"pingone_identity_provider":             identityproviders.Resource(),
				"pingone_identity_provider_attributes":  idpattributes.Resource(),
				"pingone_key":                           keys.Resource(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"pingone_environments":                   environments.DataSource(),
				"pingone_applications":                   applications.DataSource(),
				"pingone_application_attributes":         appAttributes.DataSource(),
				"pingone_application_grants":             appGrants.DataSource(),
				"pingone_application_policy_assignments": appPolicies.DataSource(),
				"pingone_custom_domains":                 customdomains.DataSource(),
				"pingone_groups":                         groups.DataSource(),
				"pingone_populations":                    populations.DataSource(),
				"pingone_resources":                      resources.DataSource(),
				"pingone_resource_scopes":                scopes.DataSource(),
				"pingone_resource_attributes":            attributes.DataSource(),
				"pingone_sign_on_policies":               signonpolicies.DataSource(),
				"pingone_sign_on_policy_actions":         signonpolicyactions.DataSource(),
				"pingone_identity_provider":              identityproviders.DataSource(),
				"pingone_identity_provider_attributes":   idpattributes.DataSource(),
				"pingone_certificates":                   certificates.DataSource(),
				"pingone_certificate_applications":       certApps.DataSource(),
				"pingone_keys":                           keys.DataSource(),
				"pingone_key_applications":               keyApps.DataSource(),
				"pingone_key_export":                     keyExport.DataSource(),
				"pingone_key_signing_request":            keyCsr.DataSource(),
			},
			ConfigureContextFunc: providerConfigure,
		}
	}
}

// Resources needed: Certificates, Custom Domains, External IdPs,
// Nice to have: Web hooks, Group nesting?

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	var tokenUrl, apiUrl *string
	tVal, ok := d.GetOk("token_url")
	if ok {
		placeholder := tVal.(string)
		tokenUrl = &placeholder
	}
	aVal, ok := d.GetOk("api_url")
	if ok {
		placeholder := aVal.(string)
		apiUrl = &placeholder
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, err := client.NewClient(ctx, &clientId, &clientSecret, tokenUrl, apiUrl)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error creating client",
			Detail:   err.Error(),
		})
	}

	return client, diags
}
