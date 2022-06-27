package signonpolicyactions

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
		Description: "`pingone_sign_on_policy_actions` data source can be used to list all available Sign On (Authentication)" +
			" Policy Actions for an environment/sign on policy.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_id": {
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
						"sign_on_policy": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"condition": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     resourceCondition(),
						},
						// Login action
						"confirm_identity_provider_attributes": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enforce_lockout_for_identity_providers": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"recovery": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"registration": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceRegistration()),
						},
						"social_providers": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						// MFA action
						"authenticator": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"bound_biometrics": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"email": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"security_key": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"sms": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"voice": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"applications": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"auto_entrollment": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeBool,
										},
									},
									"device_authorization": {
										Type:     schema.TypeMap,
										Computed: true,
									},
								},
							},
						},
						"no_device_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// Agreement action
						"agreement": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						// Identifier First action
						"discovery_rules": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     utils.ResourceToDataSource(resourceDiscoveryRules()),
						},
						"identity_provider": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						// Identity Provider action
						"acr_values": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pass_user_context": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						// Progressive Profiling action
						"prevent_multiple_prompts_per_flow": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"prompt_interval_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prompt_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem:     resourceAttributes(),
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
	policyId := d.Get("policy_id").(string)
	var id *string
	if val, ok := d.GetOk("id"); ok {
		id = utils.String(val.(string))
	}

	var actions []models.SignOnPolicyAction
	if id != nil {
		response, err := client.GetFromEnvironment[models.SignOnPolicyAction](pingClient, environmentId, "signOnPolicies", policyId, "actions", *id)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting sign on policy action %s", *id),
				Detail:   err.Error(),
			})
		}

		actions = []models.SignOnPolicyAction{*response}
	} else {
		params := make(map[string]string)
		response, err := client.GetAllFromEnvironment[models.SignOnPolicyAction](pingClient, environmentId, params, "signOnPolicies", policyId, "actions")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failure getting sign on policy actions",
				Detail:   err.Error(),
			})
		}

		actions = response.Embedded["actions"]
	}

	if actions == nil {
		return diags
	}

	if err := d.Set("actions", FlattenMany(&actions)); err != nil {
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
