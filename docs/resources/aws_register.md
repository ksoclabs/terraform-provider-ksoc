---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ksoc_aws_register Resource - terraform-provider-ksoc"
subcategory: ""
description: |-
  Register service with Ksoc
---

# ksoc_aws_register (Resource)

Register service with Ksoc



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `aws_account_id` (String, Sensitive) Ksoc Customer AWS account ID
- `ksoc_assumed_role_arn` (String) Ksoc Role to Trust

### Read-Only

- `api_path` (String) Target of the API path
- `id` (String) The ID of this resource.
- `ksoc_registered` (Boolean) Target of the API path