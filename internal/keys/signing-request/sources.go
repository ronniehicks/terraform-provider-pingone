package keys

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
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
				Description: "Key ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"export_format": {
				Description: "Export format requested",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"export_text": {
				Description: "Exported key as a string",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pingClient := meta.(*client.Client)

	environmentId := d.Get("environment_id").(string)
	id := d.Get("id").(string)
	export_format := d.Get("export_format").(string)

	var acceptHeader *string
	if export_format == "" {
		acceptHeader = utils.String("application/x-x509-ca-cert")
	} else {
		acceptHeader = utils.String("application/" + export_format)
	}

	response, err := client.GetStringFromEnvironment(pingClient, environmentId, *acceptHeader, "keys", id, "csr")

	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting key export %s", id),
			Detail:   err.Error(),
		})
	}

	if response == nil {
		return diags
	}

	if err := d.Set("export_text", *response); err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error setting key text.",
			Detail:   err.Error(),
		})
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
