---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_cifs_local_group_member Data Source - terraform-provider-netapp-ontap"
subcategory: "NAS"
description: |-
  Retrieve CifsLocalGroupMember data source
---

# Data Source cifs_local_group_member

Retrieves cifs local group member configuration

## Supported Platforms

* On-prem ONTAP system 9.10 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_cifs_local_group_member" "protocols_cifs_local_group_member" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  svm_name = "test3"
  group_name = "testme"
  member = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `group_name` (String) Local group name
- `member` (String) Member name
- `svm_name` (String) Svm name
