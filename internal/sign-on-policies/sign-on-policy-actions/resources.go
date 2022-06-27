package signonpolicyactions

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
		Description: "`pingone_sign_on_policy_action` is used to manage Sign On (Authentication) Policy " +
			" actions for an environment/sign on policy.",
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
			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"LOGIN", "MULTI_FACTOR_AUTHENTICATION", "IDENTIFIER_FIRST", "PROGRESSIVE_PROFILING", "AGREEMENT", "IDENTITY_PROVIDER",
				}, false),
			},
			"condition": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceCondition(),
			},
			// Login action
			"confirm_identity_provider_attributes": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enforce_lockout_for_identity_providers": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"recovery": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"registration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceRegistration(),
				MaxItems: 1,
			},
			"social_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			// MFA action
			"authenticator": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"bound_biometrics": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"email": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"security_key": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"sms": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"voice": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"applications": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"auto_entrollment": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"device_authorization": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
			},
			"no_device_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Agreement action
			"agreement": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// Identifier First action
			"discovery_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceDiscoveryRules(),
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Identity Provider action
			"acr_values": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pass_user_context": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Progressive Profiling action
			"prevent_multiple_prompts_per_flow": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"prompt_interval_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prompt_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func ParseId(id string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:policy_id:action_id", id)
	}

	return parts[0], parts[1], parts[2], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	policyId := d.Get("policy_id").(string)
	action := Expand(d)

	item, err := client.CreateForEnvironment(pingClient, environmentId, action, "signOnPolicies", policyId, "actions")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating sign on policy action %d", *action.Priority),
			Detail:   err.Error(),
		})
	}

	d.SetId(fmt.Sprintf("%s:%s:%s", environmentId, policyId, *item.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, policyId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.SignOnPolicyAction](pingClient, environmentId, "signOnPolicies", policyId, "actions", id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting sign on policy actions %s", id),
			Detail:   err.Error(),
		})
	}

	Flatten(d, response, &diags)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, policyId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	action := Expand(d)
	_, err = client.PutForEnvironment(pingClient, environmentId, action, "signOnPolicies", policyId, "actions", id)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure putting sign on policy action %d", *action.Priority),
			Detail:   err.Error(),
		})
	}

	resourceRead(ctx, d, meta)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, policyId, id, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "signOnPolicies", policyId, "actions", id)
	if err != nil {
		if strings.Contains(err.Error(), "Cannot delete last action") {
			return nil
		}

		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting sign on policy action %s", id),
			Detail:   err.Error(),
		})
	}

	return diags
}
