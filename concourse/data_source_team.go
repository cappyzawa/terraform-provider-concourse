package concourse

import (
	"context"
	"fmt"
	"strconv"

	"github.com/concourse/concourse/atc"
	cc "github.com/concourse/concourse/go-concourse/concourse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTeamRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"owner_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"member_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"member_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"viewer_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"viewer_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pipeline_operator_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pipeline_operator_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(cc.Client)
	teamName := d.Get("name").(string)
	teams, err := client.ListTeams()
	if err != nil {
		return diag.FromErr(err)
	}

	var exist bool
	for _, team := range teams {
		if team.Name == teamName {
			exist = true
			d.SetId(strconv.Itoa(team.ID))
			setTeamDataSource(team, d)
		}
	}
	if !exist {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "team does not exist",
		})
		return diags
	}
	return diags
}

func setTeamDataSource(team atc.Team, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	auth := team.Auth
	if auth == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("team: %s does not have auth info", team.Name),
		})
		return diags
	}
	owner, ok := auth["owner"]
	if ok {
		groups, exists := owner["groups"]
		if exists {
			d.Set("owner_groups", groups)
		}
		users, exists := owner["users"]
		if exists {
			d.Set("owner_users", users)
		}
	}
	member, ok := auth["member"]
	if ok {
		groups, exists := member["groups"]
		if exists {
			d.Set("member_groups", groups)
		}
		users, exists := member["users"]
		if exists {
			d.Set("member_users", users)
		}
	}
	viewer, ok := auth["viewer"]
	if ok {
		groups, exists := viewer["groups"]
		if exists {
			d.Set("viewer_groups", groups)
		}
		users, exists := viewer["users"]
		if exists {
			d.Set("viewer_users", users)
		}
	}
	pipelineOperator, ok := auth["pipeline-operator"]
	if ok {
		groups, exists := pipelineOperator["groups"]
		if exists {
			d.Set("pipeline_operator_groups", groups)
		}
		users, exists := pipelineOperator["users"]
		if exists {
			d.Set("pipeline_operator_users", users)
		}
	}
	return diags
}
