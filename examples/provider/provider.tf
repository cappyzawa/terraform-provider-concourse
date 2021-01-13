variable "concourse_url" {
  type = string
}

variable "concourse_username" {
  type = string
}

variable "concourse_password" {
  type      = string
  sensitive = true
}

provider "concourse" {
  url      = var.concourse_url
  username = var.concourse_username
  password = var.concourse_password
}
