package populations

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "`pingone_population` is used to manage User Populations for an environment.",
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
			"population_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:        schema.TypeString,
				Description: "A user-defined identifier for the group. Use this propertry to syncronize a group in PingOne with the same group in an external system. PingOne does not directly use this property.",
				Optional:    true,
			},
			"password_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:group_id", id)
	}

	return parts[0], parts[1], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	population := Expand(d)

	app, err := client.CreateForEnvironment(pingClient, environmentId, population, "populations")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating population %s", *population.Name),
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s", environmentId, *app.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.Population](pingClient, environmentId, "populations", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting population %s", id),
			Detail:   err.Error(),
		})
	}

	utils.SetResourceDataWithDiagnostic(d, "environment_id", environmentId, &diags)
	Flatten(d, response, &diags)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	population := Expand(d)

	_, err = client.PutForEnvironment(pingClient, environmentId, population, "populations", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting population %s", *population.Name),
			Detail:   err.Error(),
		})
	}

	resourceRead(ctx, d, meta)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "populations", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting population %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
