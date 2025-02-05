---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_quota_rule Data Source - terraform-provider-netapp-ontap"
subcategory: "Storage"
description: |-
  QuotaRule data source
---

# Data Source quota_rule

Retrieves QuotaRule data

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_quota_rule" "storage_quota_rule" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  volume = {
    name = "lunTest"
    }
  svm = {
    name = "carchi-test"
    }
  type = "tree"
  qtree = {
    name = "testacc"
    }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `qtree` (Attributes) Qtree name for the rule (see [below for nested schema](#nestedatt--qtree))
- `svm` (Attributes) Existing SVM (see [below for nested schema](#nestedatt--svm))
- `type` (String) Quota type for the rule. This type can be user, group, or tree
- `volume` (Attributes) Existing volume (see [below for nested schema](#nestedatt--volume))

### Read-Only

- `files` (Attributes) (see [below for nested schema](#nestedatt--files))
- `group` (Attributes) group to which the group quota policy rule applies (see [below for nested schema](#nestedatt--group))
- `id` (String) The ID of this resource.
- `user_mapping` (Boolean) user mapping for user quota policy rules
- `users` (Attributes Set) user to which the user quota policy rule applies (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--qtree"></a>

### Nested Schema for `qtree`

Required:

- `name` (String) name of the qtree

<a id="nestedatt--svm"></a>

### Nested Schema for `svm`

Required:

- `name` (String) name of the SVM

<a id="nestedatt--volume"></a>

### Nested Schema for `volume`

Required:

- `name` (String) name of the volume

<a id="nestedatt--files"></a>

### Nested Schema for `files`

Read-Only:

- `hard_limit` (Number) Specifies the hard limit for files
- `soft_limit` (Number) Specifies the soft limit for files

<a id="nestedatt--group"></a>

### Nested Schema for `group`

Read-Only:

- `name` (String) name of the group

<a id="nestedatt--users"></a>

### Nested Schema for `users`

Read-Only:

- `name` (String) name of the user
