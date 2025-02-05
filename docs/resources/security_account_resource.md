---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_security_account Resource - terraform-provider-netapp-ontap"
subcategory: "Security"
description: |-
  SecurityAccount resource
---

# Resource Security Account

Create/Modify/Delete a ONTAP user account

## Related ONTAP commands

```commandline
* security login create
* security login delete
* security login modify
* security login password
* security login lock
* security login unlock
```

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
resource "netapp-ontap_security_account" "security_account" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  name = "carchitest"
  applications = [{
    application = "http"
    authentication_methods = ["password"]
  }]
  password = "P@ssw0rd"
}

```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `name` (String) SecurityAccount name

### Optional

- `applications` (Attributes List) List of applications (see [below for nested schema](#nestedatt--applications))
- `comment` (String) Account comment
- `locked` (Boolean) Account locked
- `owner` (Attributes) Account owner (see [below for nested schema](#nestedatt--owner))
- `password` (String, Sensitive) Account password
- `role` (Attributes) Account role (see [below for nested schema](#nestedatt--role))
- `second_authentication_method` (String) Second authentication method

### Read-Only

- `id` (String) SecurityAccount id
- `owner_id` (String) Account owner uuid

<a id="nestedatt--applications"></a>

### Nested Schema for `applications`

Required:

- `application` (String) Application name

Optional:

- `authentication_methods` (List of String) List of authentication methods
- `second_authentication_method` (String) Second authentication method

<a id="nestedatt--owner"></a>

### Nested Schema for `owner`

Optional:

- `name` (String) Account owner name

<a id="nestedatt--role"></a>

### Nested Schema for `role`

Optional:

- `name` (String) Account role name

## Import

This resource supports import, which allows you to import existing security account into the state of this resource.
Import require a unique ID composed of the security account name and connection profile, separated by a comma.

id = `name`, `cx_profile_name`

### Terraform Import

 For example

 ```shell
  terraform import netapp-ontap_security_account.act_import acc_user,cluster4
 ```

### Terraform Import Block

This requires Terraform 1.5 or higher, and will auto create the configuration for you

First create the block

```terraform
import {
  to = netapp-ontap_security_account.act_import
  id = "acc_user,cluster4"
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
# __generated__ by Terraform from "acc_user,cluster4"
resource "netapp-ontap_security_account" "act_import" {
  cx_profile_name = "cluster4"
  name       = "acc_user"
  applications = [
    {
      application = "http"
      authentication_methods = ["password"]
      second_authentication_method = "none"
    }
  ]
  comment         = null
  locked          = false
  owner = {
    name = "abccluster-1"
  }
  password = null # sensitive
  role = {
    name = "admin"
  }
  second_authentication_method = null
}
```
