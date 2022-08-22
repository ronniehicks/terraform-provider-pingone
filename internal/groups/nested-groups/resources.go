package nestedgroups

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
		Description:   "`pingone_nested_group` is used to manage Group Nesting for an environment.",
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		DeleteContext: resourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Description: "Environment Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"group_id": {
				Description: "Parent Group Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"nested_group_id": {
				Description: "Child Group Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
	}
}

func ParseId(id string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("unexpected format of ID (%s), expected environment_id:group_id:nested_group_id", id)
	}

	return parts[0], parts[1], parts[2], nil
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId := d.Get("environment_id").(string)
	groupId := d.Get("group_id").(string)
	nestedGroupId := d.Get("nested_group_id").(string)
	group := models.Group{
		ID: &nestedGroupId,
	}
	thing, err := client.CreateForEnvironment(pingClient, environmentId, group, "groups", groupId, "memberOfGroups")
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure creating nested group %s", groupId),
			Detail:   err.Error(),
		})
	}

	// "ID" returned by this call is the nested group
	d.SetId(fmt.Sprintf("%s:%s:%s", environmentId, groupId, *thing.ID))

	resourceRead(ctx, d, meta)

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)
	var diags diag.Diagnostics

	environmentId, groupId, nestedGroupId, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	response, err := client.GetFromEnvironment[models.Group](pingClient, environmentId, "groups", groupId, "memberOfGroups", nestedGroupId)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure getting group %s-%s", groupId, nestedGroupId),
			Detail:   err.Error(),
		})
	}

	utils.SetResourceDataWithDiagnostic(d, "environment_id", environmentId, &diags)
	utils.SetResourceDataWithDiagnostic(d, "group_id", groupId, &diags)
	Flatten(d, response, &diags)

	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pingClient := meta.(*client.Client)

	var diags diag.Diagnostics

	environmentId, groupId, nestedGroupId, err := ParseId(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure parsing id %s", d.Id()),
			Detail:   err.Error(),
		})
	}

	err = client.DeleteForEnvironment(pingClient, environmentId, "groups", groupId, "memberOfGroups", nestedGroupId)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failure deleting group %s-%s", groupId, nestedGroupId),
			Detail:   err.Error(),
		})
	}

	return diags
}
