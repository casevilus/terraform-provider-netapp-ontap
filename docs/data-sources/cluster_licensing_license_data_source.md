---
page_title: "netapp-ontap_cluster_licensing_license Data Source - terraform-provider-netapp-ontap"
subcategory: "Cluster"
description: |-
  ClusterLicensingLicense data source
---

# Data Source cluster licensing license

Retrieves Cluster Licensing License

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_cluster_licensing_license" "cluster_licensing_license" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  name = "snapmirror_sync"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `name` (String) ClusterLicensingLicense name

### Read-Only

- `licenses` (Attributes List) Licenses of the license (see [below for nested schema](#nestedatt--licenses))
- `scope` (String) Scope of the license
- `state` (String) State of the license

<a id="nestedatt--licenses"></a>

### Nested Schema for `licenses`

Read-Only:

- `active` (Boolean) active of the license
- `compliance` (Attributes) compliance of the license (see [below for nested schema](#nestedatt--licenses--compliance))
- `evaluation` (Boolean) evaluation of the license
- `installed_license` (String) installed_license of the license
- `owner` (String) owner of the license
- `serial_number` (String) Serial Number of the license

<a id="nestedatt--licenses--compliance"></a>

### Nested Schema for `licenses.compliance`

Read-Only:

- `state` (String) state of the license
