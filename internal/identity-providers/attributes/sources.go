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
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identity_provider_id": {
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mapping_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_provider_id": {
							Type:     schema.TypeString,
							Required: true,
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
	identityProviderId := d.Get("identity_provider_id").(string)
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var attributes []models.ProviderAttribute
	if id != nil {
		response, err := client.GetFromEnvironment[models.ProviderAttribute](pingClient, environmentId, "identityProviders", identityProviderId, "attributes", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting attribute %s", *id),
				Detail:   err.Error(),
			})
		}

		attributes = []models.ProviderAttribute{*response}
	} else {
		response, err := client.GetAllFromEnvironment[models.ProviderAttribute](pingClient, environmentId, nil, "identityProviders", identityProviderId, "attributes")
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
