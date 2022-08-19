package certificates

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
			"id": {
				Description: "Certificate ID",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"expires_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuer_dn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"serial_number": {
							Type:        schema.TypeString,
							Description: "String representation of a potentially very large number",
							Computed:    true,
						},
						"signature_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"starts_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject_dn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"usage_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validity_period": {
							Type:     schema.TypeInt,
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
	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var certificates []models.Certificate
	if id != nil {
		response, err := client.GetFromEnvironment[models.Certificate](pingClient, environmentId, "certificates", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting certificate %s", *id),
				Detail:   err.Error(),
			})
		}

		certificates = []models.Certificate{*response}
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}
		response, err := client.GetAllFromEnvironment[models.Certificate](pingClient, environmentId, params, "certificates")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting certificates",
				Detail:   err.Error(),
			})
		}

		certificates = response.Embedded["certificates"]
	}

	if certificates == nil {
		return diags
	}

	if err := d.Set("certificates", FlattenMany(&certificates)); err != nil {
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
