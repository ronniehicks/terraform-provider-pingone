package identityproviders

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identity_provider_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"icon": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceIcon()),
						},
						"login_button_icon": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceIcon()),
						},
						"registration": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceRegistration()),
						},
						"authorization_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_secret": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"discovery_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"jwks_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"scopes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"token_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"token_endpoint_auth_method": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_info_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"authn_request_signed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"idp_entity_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_verification_certificate_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sp_entity_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_signing_key_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_binding": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pingClient := meta.(*client.Client)

	environmentId := d.Get("environment_id").(string)
	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}
	var id *string
	if val, ok := d.GetOk("identity_provider_id"); ok {
		id = utils.String(val.(string))
	}

	var identityProviders []models.IdentityProvider
	if id != nil {
		response, err := client.GetFromEnvironment[models.IdentityProvider](pingClient, environmentId, "identityProviders", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting identityProviders %s", *id),
				Detail:   err.Error(),
			})
		}

		identityProviders = []models.IdentityProvider{*response}
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}
		response, err := client.GetAllFromEnvironment[models.IdentityProvider](pingClient, environmentId, params, "identityProviders")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting identityProviders",
				Detail:   err.Error(),
			})
		}

		identityProviders = response.Embedded["identityProviders"]
	}

	if identityProviders == nil {
		return diags
	}

	if err := d.Set("identity_providers", FlattenMany(&identityProviders)); err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error setting key",
			Detail:   err.Error(),
		})
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
