package identityproviders

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment Id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string that specifies the name of the IdP. This is a required property. ",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the description of the IdP.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string that specifies the IdP type. This is a required property. Options are FACEBOOK, GOOGLE, LINKEDIN, OPENID_CONNECT, APPLE, AMAZON, TWITTER, YAHOO,and SAML.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "A boolean that specifies the current enabled state of the IdP.",
			},
			"icon": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        resourceIcon(),
				MaxItems:    1,
				Description: "The ID and HREF for the IdP icon.",
			},
			"login_button_icon": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        resourceIcon(),
				MaxItems:    1,
				Description: "The image ID and HREF for the IdP login button icon. For Facebook, Google, and LinkedIn IdPs, updates to the login button are ignored to preserve the IdP branding rules.",
			},
			"registration": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        resourceRegistration(),
				Description: "The optional registration object designates an external IdP as authoritative. Setting this attribute gives management of linked users to the IdP and also triggers just-in-time provisioning of new users. These users are created in the population indicated with registration.population.id.",
			},
			"authorization_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the the OIDC identity provider's authorization endpoint. This value must be a URL that uses https. This is a required property.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the application ID from the OIDC identity provider. This is a required property.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the application secret from the OIDC identity provider. This is a required property.",
				Sensitive:   true,
			},
			"discovery_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the OIDC identity provider's discovery endpoint. This value must be a URL that uses https.",
			},
			"issuer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the issuer to which the authentication is sent for the OIDC identity provider. This value must be a URL that uses https. This is a required property.",
			},
			"jwks_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the OIDC identity provider's jwks endpoint. This value must be a URL that uses https. This is a required property.",
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "An array that specifies the scopes to include in the authentication request to the OIDC identity provider. This is a required property.",
			},
			"token_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the OIDC identity provider's token endpoint. This is a required property.",
			},
			"token_endpoint_auth_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the OIDC identity provider's token endpoint authentication method. Options are CLIENT_SECRET_BASIC (default), CLIENT_SECRET_POST, and NONE. This is a required property.",
			},
			"user_info_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the OIDC identity provider's userInfo endpoint.",
			},
			"authn_request_signed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs.",
			},
			"idp_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response.",
			},
			"idp_verification_certificate_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A array that specifies the identity provider's certificate IDs used to verify the signature on the signed assertion from the identity provider. Signing is done with a private key and verified with a public key.",
			},
			"sp_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the service provider's entity ID, used to look up the application.",
			},
			"sp_signing_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the service provider's signing key ID.",
			},
			"sso_binding": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the binding for the authentication request. Options are HTTP_POST and HTTP_REDIRECT.",
			},
			"sso_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the SSO endpoint for the authentication request.",
			},
		},
	}
}

func resourceIcon() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"href": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRegistration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"population_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:identity_provider_id", id)
	}

	return parts[0], parts[1], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	identityProvider := Expand(d)

	app, err := client.CreateForEnvironment(pingClient, environmentId, identityProvider, "identityProviders")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating identityProvider %s", *identityProvider.Name),
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s", environmentId, *app.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.IdentityProvider](pingClient, environmentId, "identityProviders", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting identityProvider %s", id),
			Detail:   err.Error(),
		})
	}

	utils.SetResourceDataWithDiagnostic(d, "environment_id", environmentId, &diags)
	Flatten(d, response, &diags)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}
	identityProvider := Expand(d)

	_, err = client.PutForEnvironment(pingClient, environmentId, identityProvider, "identityProviders", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting identityProvider %s", *identityProvider.Name),
			Detail:   err.Error(),
		})
	}

	resourceRead(ctx, d, meta)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "identityProviders", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting identityProvider %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
