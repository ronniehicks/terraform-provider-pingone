package applications

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcon() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"href": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAccessControl() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"group": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceAccessControlGroup(),
			},
		},
	}
}

func resourceAccessControlGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ALL_GROUPS", "ANY_GROUP",
				}, false),
			},
			"groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceCacheDuration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"amount": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"units": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"HOURS", "MINUTES",
				}, false),
			},
		},
	}
}

func resourceIntegrityDetection() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"EVALUATE", "DISABLED", "ENABLED",
				}, false),
			},
			"cache_duration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceCacheDuration(),
			},
		},
	}
}

func resourcePasscodeRefreshDuration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"time_unit": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"SECONDS",
				}, false),
			},
		},
	}
}

func resourceMobile() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"integrity_detection": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceIntegrityDetection(),
			},
			"passcode_refresh_duration": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourcePasscodeRefreshDuration(),
			},
		},
	}
}

func resourceIdpSigning() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSpEncryption() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificates": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSpVerification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authn_request_signed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "String representation of a bool so we can handle tristate",
			},
			"certificates": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
