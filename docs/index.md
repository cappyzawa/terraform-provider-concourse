---
page_title: "concourse Provider"
subcategory: ""
description: |-
  
---

# concourse Provider



## Example Usage

```terraform
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
```

## Schema

### Required

- **password** (String, Sensitive)
- **url** (String)
- **username** (String)
