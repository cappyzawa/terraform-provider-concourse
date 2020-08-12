# Concourse Provider

This provider enables to operate concourse.

## Example Usage

```hcl
provider "concourse" {
  url      = "https://example.concourse.com"
  username = "foo"
  password = "bar"
}

resource "concourse_team" "main" {
  name = "main"
  owner_groups = [
    "oidc:123456789"
  ]
}
```

## Argument References

The following arguments are supported:

* `url` - (Required) The url of your concourse.
* `username` - (Required) The username of a local user.
* `password` - (Required) The password of a local user.
