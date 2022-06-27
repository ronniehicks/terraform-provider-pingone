package environments

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Description: "`pingone_environments` data source can be used to list all available environments.",
		ReadContext: read,
		Schema: map[string]*schema.Schema{
			"filter": {
				Description: "SCIM filter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": {
				Description: "Environment Id",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
						"license": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pingClient := meta.(*client.Client)

	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}

	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var environments []models.Environment

	if id != nil {
		response, err := client.GetFromEnvironment[models.Environment](pingClient, *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting environment %s", *id),
				Detail:   err.Error(),
			})
		}

		environments = append(environments, *response)
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}

		response, err := client.GetAll[models.Environment](pingClient, params, "environments")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting environments",
				Detail:   err.Error(),
			})
		}

		environments = response.Embedded["environments"]
	}

	if environments == nil {
		return diags
	}

	if err := d.Set("environments", FlattenMany(&environments)); err != nil {
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
