# concourse_team

## Example Usage

```hcl
data "concourse_team" "main" {
  name = "main"
}
```

## Argument References

The following arguments are supported:

* `name` - (Required) The name of the team.

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
