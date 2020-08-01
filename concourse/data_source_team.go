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
			t := flattenTeamAuthData(team.Auth)
			for k, v := range t {
				d.Set(k, v)
			}
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

func flattenTeamAuthData(auth atc.TeamAuth) map[string][]string {
	m := make(map[string][]string)
	for k, v := range auth {
		// e.g. k="owner", k="pipeline-operator", etc.
		if k == "pipeline-operator" {
			k = "pipeline_operator"
		}
		groups, ok := v["groups"]
		if ok {
			m[fmt.Sprintf("%s_groups", k)] = groups
		}
		users, ok := v["users"]
		if ok {
			m[fmt.Sprintf("%s_users", k)] = users
		}
	}
	return m
}
