package signonpolicyassignments

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
		Description: "`pingone_application_sign_on_policy_assignments` data source can be used to list all available" +
			" Sign On Policy Assignments (Authentication Policies) for an environment/application.",
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
			"sign_on_policy_assignments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assignment_id": {
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
						"application": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sign_on_policy": {
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

	var signOnPolicyAssignments []models.ApplicationSignOnPolicyAssignment
	if id != nil {
		response, err := client.GetFromEnvironment[models.ApplicationSignOnPolicyAssignment](pingClient, environmentId, "applications", applicationId, "signOnPolicyAssignments", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting signOnPolicyAssignment %s", *id),
				Detail:   err.Error(),
			})
		}

		signOnPolicyAssignments = []models.ApplicationSignOnPolicyAssignment{*response}
	} else {
		response, err := client.GetAllFromEnvironment[models.ApplicationSignOnPolicyAssignment](pingClient, environmentId, nil, "applications", applicationId, "signOnPolicyAssignments")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting signOnPolicyAssignments",
				Detail:   err.Error(),
			})
		}

		signOnPolicyAssignments = response.Embedded["signOnPolicyAssignments"]
	}

	if signOnPolicyAssignments == nil {
		return diags
	}

	if err := d.Set("sign_on_policy_assignments", FlattenMany(&signOnPolicyAssignments)); err != nil {
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
