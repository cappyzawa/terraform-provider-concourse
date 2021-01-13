resource "concourse_team" "example" {
  name = "example"
  owner_groups = [
    "oidc:xxxxxxxxxxxxxxxxxxxxxxxx"
  ]
}
