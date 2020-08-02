package concourse

import (
	"context"
	"strconv"

	"github.com/concourse/concourse/atc"
	cc "github.com/concourse/concourse/go-concourse/concourse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
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

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(cc.Client)
	teamName := d.Get("name").(string)
	ownerGroups := ifs2strs(d.Get("owner_groups").([]interface{}))
	ownerUsers := ifs2strs(d.Get("owner_users").([]interface{}))
	team := atc.Team{
		Name: teamName,
		Auth: atc.TeamAuth{
			"owner": {
				"groups": ownerGroups,
				"users":  ownerUsers,
			},
		},
	}
	createdTeam, _, _, err := client.Team(teamName).CreateOrUpdate(team)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to create team",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(strconv.Itoa(createdTeam.ID))
	t := flattenTeamAuthData(createdTeam.Auth)
	for k, v := range t {
		d.Set(k, v)
	}
	return diags
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(cc.Client)
	teamName := d.Get("name").(string)
	teams, err := client.ListTeams()
	if err != nil {
		return diag.FromErr(err)
	}
	for _, team := range teams {
		if team.Name == teamName {
			d.SetId(strconv.Itoa(team.ID))
			t := flattenTeamAuthData(team.Auth)
			for k, v := range t {
				d.Set(k, v)
			}
		}
	}
	return diags
}

func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(cc.Client)
	if d.HasChange("name") {
		before, _ := d.GetChange("name")
		beforeName := before.(string)
		if err := client.Team(beforeName).DestroyTeam(beforeName); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "failed to create team",
			})
			return diags
		}
	}
	teamName := d.Get("name").(string)
	ownerGroups := ifs2strs(d.Get("owner_groups").([]interface{}))
	ownerUsers := ifs2strs(d.Get("owner_users").([]interface{}))
	team := atc.Team{
		Name: teamName,
		Auth: atc.TeamAuth{
			"owner": {
				"groups": ownerGroups,
				"users":  ownerUsers,
			},
		},
	}
	createdTeam, _, _, err := client.Team(teamName).CreateOrUpdate(team)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to create team",
		})
		return diags
	}
	d.SetId(strconv.Itoa(createdTeam.ID))
	t := flattenTeamAuthData(createdTeam.Auth)
	for k, v := range t {
		d.Set(k, v)
	}
	return diags
}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(cc.Client)
	teamName := d.Get("name").(string)

	if err := client.Team(teamName).DestroyTeam(teamName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to destory team",
		})
		return diags
	}
	d.SetId("")
	return diags
}

func ifs2strs(ifs []interface{}) []string {
	strs := []string{}
	for _, f := range ifs {
		strs = append(strs, f.(string))
	}
	return strs
}
