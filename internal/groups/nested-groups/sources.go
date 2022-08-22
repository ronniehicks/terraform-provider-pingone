package nestedgroups

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
		Description: "`pingone_nested_groups` data source can be used to list all available Nested Groups for the provided User Group.",
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Description: "SCIM filter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment_id": {
				Description: "Environment ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"group_id": {
				Description: "Group ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"nested_group_id": {
				Description: "Nested Group ID",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"group_memberships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nested_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
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
	groupId := d.Get("group_id").(string)

	var filter *string
	if val, ok := d.GetOk("filter"); ok {
		filter = utils.String(val.(string))
	}
	var nestedGroupId *string
	if val, ok := d.GetOk("nested_group_id"); ok {
		nestedGroupId = utils.String(val.(string))
	}

	var groups []models.Group
	if nestedGroupId != nil {
		response, err := client.GetFromEnvironment[models.Group](pingClient, environmentId, "groups", groupId, "memberOfGroups", *nestedGroupId)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting group %s-%s", groupId, *nestedGroupId),
				Detail:   err.Error(),
			})
		}

		groups = []models.Group{*response}
	} else {
		params := make(map[string]string)
		if filter != nil {
			params["filter"] = *filter
		}
		response, err := client.GetAllFromEnvironment[models.Group](pingClient, environmentId, params, "groups", groupId, "memberOfGroups")
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failure getting groups for %s", groupId),
				Detail:   err.Error(),
			})
		}

		groups = response.Embedded["groupMemberships"]
	}

	if groups == nil {
		return diags
	}

	if err := d.Set("group_memberships", FlattenMany(&groups)); err != nil {
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
