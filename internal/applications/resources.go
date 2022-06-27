package applications

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "`pingone_application` is used to managed an application for an environment.",
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
			"application_id": {
				Description: "Application Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"NATIVE_APP", "WEB_APP", "SINGLE_PAGE_APP", "SERVICE", "WORKER", "CUSTOM_APP", "PORTAL_LINK_APP",
				}, false),
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"OPENID_CONNECT", "SAML", "EXTERNAL_LINK",
				}, false),
			},
			"home_page_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_page_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"icon": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceIcon(),
			},
			"assign_actor_roles": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// tags - Ping Fed only, do we need?
			"access_control": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceAccessControl(),
			},

			// OIDC only settings
			"grant_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"AUTHORIZATION_CODE", "IMPLICIT", "REFRESH_TOKEN", "CLIENT_CREDENTIALS",
					}, false),
				},
			},
			"post_logout_redirect_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"redirect_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"response_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"TOKEN", "ID_TOKEN", "CODE",
					}, false),
				},
			},
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"NONE", "CLIENT_SECRET_BASIC", "CLIENT_SECRET_POST",
				}, false),
			},
			"pkce_enforcement": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPTIONAL",
				ValidateFunc: validation.StringInSlice([]string{
					"OPTIONAL", "REQUIRED", "S256_REQUIRED",
				}, false),
			},
			"refresh_token_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(60),
			},
			"refresh_token_rolling_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(60),
			},
			"support_unsigned_request_object": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Native App only settings
			"bundle_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mobile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceMobile(),
			},

			// SAML App only settings
			"sp_entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"acs_urls": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"assertion_duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_target_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slo_binding": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"HTTP_REDIRECT", "HTTP_POST",
				}, false),
			},
			"slo_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slo_response_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_signed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"assertion_signed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"idp_signing": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceIdpSigning(),
			},
			"sp_encryption": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceSpEncryption(),
			},
			"sp_verification": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceSpVerification(),
			},
			"name_id_format": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
					"urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
					"urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
					"urn:oasis:names:tc:SAML:2.0:nameid-format:transient",
				}, false),
			},
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	application := Expand(d)

	app, err := client.CreateForEnvironment(pingClient, environmentId, application, "applications")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating application %s", *application.Name),
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

	response, err := client.GetFromEnvironment[models.Application](pingClient, environmentId, "applications", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting application %s", id),
			Detail:   err.Error(),
		})
	}

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
	application := Expand(d)

	_, err = client.PutForEnvironment(pingClient, environmentId, application, "applications", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting application %s", *application.Name),
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

	err = client.DeleteForEnvironment(pingClient, environmentId, "applications", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting application %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}

func ParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:application_id", id)
	}

	return parts[0], parts[1], nil
}
