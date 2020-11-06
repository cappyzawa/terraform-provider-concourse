variable "concourse_url" {
  type = string
}

variable "concourse_username" {
  type = string
}

variable "concourse_password" {
  type = string
}

provider "concourse" {
  url      = var.concourse_url
  username = var.concourse_username
  password = var.concourse_password
}

resource "concourse_team" "provider-test" {
  name = "provider-test"
  owner_groups = [
    "oidc:12222222222222"
  ]
}
