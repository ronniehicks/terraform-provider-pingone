package customdomains

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/ronniehicks/terraform-provider-pingone/internal/utils"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "`pingone_custom_domain` is used to manage Custom Domains for an environment.",
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
			"custom_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"canonical_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_expiration": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
		},
	}
}

func ParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:custom_domain_id", id)
	}

	return parts[0], parts[1], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	item := Expand(d)

	response, err := client.CreateForEnvironment(pingClient, environmentId, item, "customDomains")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating custom domain %s", *item.DomainName),
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s", environmentId, *response.ID))

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

	response, err := client.GetFromEnvironment[models.CustomDomain](pingClient, environmentId, "customDomains", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting custom domain %s", id),
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

	customDomain := Expand(d)

	_, err = client.PutForEnvironment(pingClient, environmentId, customDomain, "customDomains", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting custom domain %s", *customDomain.DomainName),
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

	err = client.DeleteForEnvironment(pingClient, environmentId, "customDomains", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting custom domain %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
