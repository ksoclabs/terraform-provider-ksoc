---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ksoc Provider"
subcategory: ""
description: |-
  
---

# ksoc Provider



## Example Usage

```terraform
provider "ksoc" {
  access_key_id = "ksoc_access_key"
  secret_key    = "ksoc_secret_key"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_key_id` (String, Sensitive) Ksoc Customer Access ID
- `secret_key` (String, Sensitive) Ksoc Customer Secret Key

### Optional

- `ksoc_api_url` (String) Ksoc API to target. Defaults to https://api.ksoc.com
