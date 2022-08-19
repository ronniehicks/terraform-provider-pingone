package grants

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
		Description: "`pingone_application_grants` data source can be used to list all available resource grants for an environment/application.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"grants": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"grant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scopes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"application": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"environment": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
	applicationId := d.Get("application_id").(string)
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var grants []models.ApplicationGrant
	if id != nil {
		response, err := client.GetFromEnvironment[models.ApplicationGrant](pingClient, environmentId, "applications", applicationId, "grants", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting grant %s", *id),
				Detail:   err.Error(),
			})
		}

		grants = []models.ApplicationGrant{*response}
	} else {
		response, err := client.GetAllFromEnvironment[models.ApplicationGrant](pingClient, environmentId, nil, "applications", applicationId, "grants")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting grants",
				Detail:   err.Error(),
			})
		}

		grants = response.Embedded["grants"]
	}

	if grants == nil {
		return diags
	}

	if err := d.Set("grants", FlattenMany(&grants)); err != nil {
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
