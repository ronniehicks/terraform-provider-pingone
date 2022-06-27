package attributes

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	client "github.com/ronniehicks/terraform-provider-pingone/pingone-client"
	"github.com/ronniehicks/terraform-provider-pingone/pingone-client/models"
)

func Resource() *schema.Resource {
	return &schema.Resource{
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
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Description: "A string that specifies the user attribute, which is unique per provider. The attribute must not be defined as read only from the user schema or of type COMPLEX based on the user schema. Valid examples: username, and name.first. The following attributes may not be used: account, id, created, updated, lifecycle, mfaEnabled, and enabled.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": {
				Description: "A string that specifies a placeholder referring to the attribute (or attributes) from the provider. Placeholders must be valid for the attributes returned by the IdP type and use the ${} syntax (for example, username=\"${email}\"). For SAML, any placeholder is acceptable, and it is mapped against the attributes available in the SAML assertion after authentication. The ${samlAssertion.subject} placeholder is a special reserved placeholder used to refer to the subject name ID in the SAML assertion response.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"mapping_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CORE", "CUSTOM"}, false),
				Description:  "A read-only string that specifies the mapping type. Options are: CORE (This attribute is required by the schema and cannot be removed. The name and update properties cannot be changed.) or CUSTOM (All user-created attributes are of this type.)",
			},
			"update": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"EMPTY_ONLY", "ALWAYS"}, false),
				Description:  "A string that specifies whether to update the user attribute in the directory with the non-empty mapped value from the IdP. Options are: EMPTY_ONLY (only update the user attribute if it has an empty value); ALWAYS (always update the user attribute value).",
			},
		},
	}
}

func ParseId(id string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:identity_provider_id:attribute_id", id)
	}

	return parts[0], parts[1], parts[2], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	identityProviderId := d.Get("identity_provider_id").(string)
	attribute := Expand(d)

	item, err := client.CreateForEnvironment(pingClient, environmentId, attribute, "identityProviders", identityProviderId, "attributes")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating identity provider attribute %s", *attribute.Name),
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s:%s", environmentId, identityProviderId, *item.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, identityProviderId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.ProviderAttribute](pingClient, environmentId, "identityProviders", identityProviderId, "attributes", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting identity provider attribute %s", id),
			Detail:   err.Error(),
		})
	}

	Flatten(d, response, &diags)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, identityProviderId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	scope := Expand(d)
	_, err = client.PutForEnvironment(pingClient, environmentId, scope, "identityProviders", identityProviderId, "attributes", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting identity provider attribute %s", *scope.Name),
			Detail:   err.Error(),
		})
	}

	resourceRead(ctx, d, meta)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, identityProviderId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "identityProviders", identityProviderId, "attributes", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting identity provider attribute %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
