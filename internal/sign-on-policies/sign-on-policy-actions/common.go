package signonpolicyactions

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCondition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"seconds_since": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"greater": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"and": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceConditionBlock(),
			},
			"or": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceConditionBlock(),
			},
			"not": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     resourceConditionBlock(),
			},
		},
	}
}

func resourceConditionBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"equals": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contains": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_range": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"greater": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"seconds_since": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRegistration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"confirm_identity_provider_attributes": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"population_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAttributes() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceDiscoveryRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"condition": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
