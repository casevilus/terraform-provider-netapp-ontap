---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_quota_rule Resource - terraform-provider-netapp-ontap"
subcategory: "Storage"
description: |-
  StorageQuotaRule resource
---

# Resource Quota Rule

Create/Modify/Delete a quota rule resource

## Related ONTAP commands

```commandline
* quota policy rule create
* quota policy rule modify
* quota policy rule delete
```

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
resource "netapp-ontap_quota_rule" "storage_quota_rule" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  volume = {
    name = "lunTest"
    }
  svm = {
    name = "test"
    }
  type = "tree"
  qtree = {
    name = "test"
    }
  files = {
    hard_limit = 100
    soft_limit = 70
    }
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `qtree` (Attributes) Qtree for which to create the rule. For default tree rules, the qtree name must be specified as "" (see [below for nested schema](#nestedatt--qtree))
- `svm` (Attributes) Existing SVM in which to create the qtree (see [below for nested schema](#nestedatt--svm))
- `type` (String) Quota type for the rule. This type can be user, group, or tree
- `volume` (Attributes) Existing volume in which to create the qtree (see [below for nested schema](#nestedatt--volume))

### Optional

- `files` (Attributes) (see [below for nested schema](#nestedatt--files))
- `group` (Attributes) If the quota type is group, this property takes the group name. For default group quota rules, the group name must be specified as "" (see [below for nested schema](#nestedatt--group))
- `users` (Attributes Set) If the quota type is user, this property takes the user name. For default user quota rules, the user name must be specified as "" (see [below for nested schema](#nestedatt--users))

### Read-Only

- `id` (String) The ID of this resource.

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

Optional:

- `hard_limit` (Number) Specifies the hard limit for files
- `soft_limit` (Number) Specifies the soft limit for files

<a id="nestedatt--group"></a>

### Nested Schema for `group`

Required:

- `name` (String) name of the group

<a id="nestedatt--users"></a>

### Nested Schema for `users`

Optional:

- `name` (String) name of the user

## Import

This Resource supports import, which allows you to import existing quota rules into the state of this resource.
Import require a unique ID composed of the volume name, svm name, type, qtree and cx_profile_name, separated by a comma.

id = `name`,`volume_name`, `svm_name`, `cx_profile_name`

### Terraform Import

For example

 ```shell
  terraform import netapp-ontap_quota_rule.example vol,svm,tree,test,cluster4
 ```

!> The terraform import CLI command can only import resources into the state. Importing via the CLI does not generate configuration. If you want to generate the accompanying configuration for imported resources, use the import block instead.

### Terraform Import Block

This requires Terraform 1.5 or higher, and will auto create the configuration for you

First create the block

```terraform
import {
  to = netapp-ontap_quota_rule.lun_import
  id = "vol,svm,tree,test,cluster4"
}
```

Next run, this will auto create the configuration for you

```shell
terraform plan -generate-config-out=generated.tf
```

This will generate a file called generated.tf, which will contain the configuration for the imported resource

```terraform
# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "vol,svm,tree,test,cluster4"
resource "netapp-ontap_quota_rule" "storage_quota_rule" {
  cx_profile_name = "cluster4"
  files = {
    hard_limit = 100
    soft_limit = 80
  }
  group = null
  id = "abcd"
  qtree = {
    name = "testacc"
  }
  svm = {
    name = "carchi-test"
  }
  type = "tree"
  users = null
  volume = {
    name = "lunTest"
  }
}
```
