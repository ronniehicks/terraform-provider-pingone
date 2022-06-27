package certificates

import (
	"context"
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
			"key_id": {
				Description: "Key ID",
				Type:        schema.TypeString,
				Required:    true,
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
						"name": {
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
	keyId := d.Get("key_id").(string)
	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}

	var keyApplications []models.KeyApplication

	params := make(map[string]string)
	if filter != nil {
		params["filter"] = *filter
	}

	response, err := client.GetAllFromEnvironment[models.KeyApplication](pingClient, environmentId, params, "keys", keyId, "applications")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure getting key applications",
			Detail:   err.Error(),
		})
	}

	keyApplications = response.Embedded["applications"]

	if keyApplications == nil {
		return diags
	}

	if err := d.Set("applications", FlattenMany(&keyApplications)); err != nil {
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
