package applications

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
		Description: "`pingone_applications` data source can be used to list all available applications for an environment.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Description: "SCIM filter",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_id": {
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"home_page_url": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"login_page_url": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"icon": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceIcon()),
						},
						"assign_actor_roles": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"access_control": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceAccessControl()),
						},
						// OIDC only settings
						"grant_types": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"post_logout_redirect_uris": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"redirect_uris": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"response_types": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"token_endpoint_auth_method": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"pkce_enforcement": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"refresh_token_duration": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"refresh_token_rolling_duration": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"support_unsigned_request_object": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},

						// Native App only settings
						"bundle_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"package_name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"mobile": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceMobile()),
						},

						// SAML App only settings
						"sp_entity_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"acs_urls": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"assertion_duration": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"default_target_url": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"slo_binding": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"slo_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"slo_response_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"response_signed": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"assertion_signed": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
						},
						"idp_signing": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceIdpSigning()),
						},
						"sp_encryption": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceSpEncryption()),
						},
						"sp_verification": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceSpVerification()),
						},
						"name_id_format": {
							Type:     schema.TypeString,
							Computed: true,
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
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var items []models.Application
	if id != nil {
		response, err := client.GetFromEnvironment[models.Application](pingClient, environmentId, "applications", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting application %s", *id),
				Detail:   err.Error(),
			})
		}

		items = []models.Application{*response}
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}
		response, err := client.GetAllFromEnvironment[models.Application](pingClient, environmentId, params, "applications")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting applications",
				Detail:   err.Error(),
			})
		}

		items = response.Embedded["applications"]
	}

	if items == nil {
		return diags
	}

	if err := d.Set("applications", FlattenMany(&items)); err != nil {
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
