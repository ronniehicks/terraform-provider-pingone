package signonpolicyassignments

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description: "`pingone_application_sign_on_policy_assignment` is used to manage Sign On Policy Assignments" +
			" (Authentication Policies) for an environment/application.",
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment Id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assignment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"sign_on_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ParseId(id string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:application_id:assignment_id", id)
	}

	return parts[0], parts[1], parts[2], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	applicationId := d.Get("application_id").(string)
	assignment := Expand(d)

	item, err := client.CreateForEnvironment(pingClient, environmentId, assignment, "applications", applicationId, "signOnPolicyAssignments")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure creating application sign on policy assignment",
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s:%s", environmentId, applicationId, *item.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, applicationId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.ApplicationSignOnPolicyAssignment](pingClient, environmentId, "applications", applicationId, "signOnPolicyAssignments", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting application sign on policy assignment %s", id),
			Detail:   err.Error(),
		})
	}
	Flatten(d, response, &diags)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, applicationId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	assignment := Expand(d)
	_, err = client.PutForEnvironment(pingClient, environmentId, assignment, "applications", applicationId, "signOnPolicyAssignments", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failure putting application sign on policy assignment",
			Detail:   err.Error(),
		})
	}

	resourceRead(ctx, d, meta)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, applicationId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "applications", applicationId, "signOnPolicyAssignments", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting application sign on policy assignment %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
