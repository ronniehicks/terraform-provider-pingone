package attributes

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
		Description: "`pingone_resource_attributes` data source can be used to list all available Resource Attributes for an environment/resource.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Description: "SCIM filter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_id": {
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
						"resource": {
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
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
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
	resourceId := d.Get("resource_id").(string)
	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var attributes []models.ResourceAttribute
	if id != nil {
		response, err := client.GetFromEnvironment[models.ResourceAttribute](pingClient, environmentId, "resources", resourceId, "attributes", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting attribute %s", *id),
				Detail:   err.Error(),
			})
		}

		attributes = []models.ResourceAttribute{*response}
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}
		response, err := client.GetAllFromEnvironment[models.ResourceAttribute](pingClient, environmentId, params, "resources", resourceId, "attributes")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting attributes",
				Detail:   err.Error(),
			})
		}

		attributes = response.Embedded["attributes"]
	}

	if attributes == nil {
		return diags
	}

	if err := d.Set("attributes", FlattenMany(&attributes)); err != nil {
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
