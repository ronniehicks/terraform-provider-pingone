package signonpolicies

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
		Description: "`pingone_sign_on_policies` data source can be used to list all available Sign On (Authentication)" +
			" Policies for an environment.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sign_on_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": {
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
						"default": {
							Type:     schema.TypeBool,
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
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var signOnPolicies []models.SignOnPolicy
	if id != nil {
		response, err := client.GetFromEnvironment[models.SignOnPolicy](pingClient, environmentId, "signOnPolicies", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting sign on policy %s", *id),
				Detail:   err.Error(),
			})
		}

		signOnPolicies = []models.SignOnPolicy{*response}
	} else {
		params := make(map[string]string)
		response, err := client.GetAllFromEnvironment[models.SignOnPolicy](pingClient, environmentId, params, "signOnPolicies")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting sign on policies",
				Detail:   err.Error(),
			})
		}

		signOnPolicies = response.Embedded["signOnPolicies"]
	}

	if signOnPolicies == nil {
		return diags
	}

	if err := d.Set("sign_on_policies", FlattenMany(&signOnPolicies)); err != nil {
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
