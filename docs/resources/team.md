# concourse_team

This resource can operates concourse teams.

## Example Usage

```hcl
resource "concourse_team" "main" {
  name = "main"
  owner_groups = [
    "oidc:123456789"
  ]
}
```

## Argument References

* `name` - (Required) The name of the team.
* `owner_groups` - (Optional) The owner group list.
* `owner_users` - (Optional) The owner user list.
* `member_groups` - (Optional) The member group list.
* `member_users` - (Optional) The member user list.
* `viewer_groups` - (Optional) The viewer group list.
* `viewer_users` - (Optional) The viewer user list.
* `pipeline_operator_groups` - (Optional) The pipeline operator group list.
* `pipeline_operator_users` - (Optional) The pipeline operator user list.

## Attributes References

* `id` - The ID of the team
* `name` - The name of the team.
* `owner_groups` - The owner group list.
* `owner_users` - The owner user list.
* `member_groups` - The member group list.
* `member_users` - The member user list.
* `viewer_groups` - The viewer group list.
* `viewer_users` - The viewer user list.
* `pipeline_operator_groups` - The pipeline operator group list.
* `pipeline_operator_users` - The pipeline operator user list.
