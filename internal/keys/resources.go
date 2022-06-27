package keys

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
			"key_id": {
				Description: "Key Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"algorithm": {
				Type:         schema.TypeString,
				Description:  "The key algorithm.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"RSA", "EC", "UNKNOWN"}, false),
			},
			"default": {
				Type:        schema.TypeBool,
				Description: "A boolean that specifies whether this is the default key for the specified environment.",
				Optional:    true,
				Default:     false,
			},
			"created_at": {
				Type:        schema.TypeString,
				Description: "The time the resource was created.",
				Computed:    true,
			},
			"expires_at": {
				Type:        schema.TypeString,
				Description: "The time the key resource expires.",
				Computed:    true,
			},
			"issuer_dn": {
				Type:        schema.TypeString,
				Description: "A string that specifies the distinguished name of the certificate issuer.",
				Required:    true,
			},
			"key_length": {
				Type:         schema.TypeInt,
				Description:  "An integer that specifies the key length. For RSA keys, options are 2048, 3072, and 7680. For elliptical curve (EC) keys, options are 224, 256, and 384.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{2048, 3072, 7680, 224, 256, 384}),
			},
			"serial_number": {
				Type:        schema.TypeInt,
				Description: "An integer that specifies the serial number of the key or certificate.",
				Computed:    true,
			},
			"signature_algorithm": {
				Type:        schema.TypeString,
				Description: "A string that specifies the signature algorithm of the key. Examples: 'SHA256withRSA', 'SHA512withECDSA'.",
				Required:    true,
				ForceNew:    true,
			},
			"starts_at": {
				Type:        schema.TypeString,
				Description: "The time the validity period starts.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "A string that specifies the status of the key. Options are VALID, EXPIRED, NOT_YET_VALID, and REVOKED.",
				Computed:    true,
			},
			"subject_dn": {
				Type:        schema.TypeString,
				Description: "A string that specifies the distinguished name of the subject being secured.",
				Required:    true,
				ForceNew:    true,
			},
			"usage_type": {
				Type:         schema.TypeString,
				Description:  "A string that specifies how the certificate is used. Options are ENCRYPTION and SIGNING.",
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"ENCRYPTION", "SIGNING"}, false),
			},
			"validity_period": {
				Type:        schema.TypeInt,
				Description: "An integer that specifies the number of days the key is valid.",
				Optional:    true,
				Default:     365,
				ForceNew:    true,
			},
		},
	}
}

func ParseId(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:key_id", id)
	}

	return parts[0], parts[1], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	key := Expand(d)

	app, err := client.CreateForEnvironment(pingClient, environmentId, key, "keys")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating key %s", *key.Name),
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

	response, err := client.GetFromEnvironment[models.Key](pingClient, environmentId, "keys", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting key %s", id),
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
	key := Expand(d)

	_, err = client.PutForEnvironment(pingClient, environmentId, key, "keys", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting key %s", *key.Name),
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

	err = client.DeleteForEnvironment(pingClient, environmentId, "keys", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting key %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
