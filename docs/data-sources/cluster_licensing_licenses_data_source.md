---
page_title: "netapp-ontap_cluster_licensing_licenses Data Source - terraform-provider-netapp-ontap"
subcategory: "Cluster"
description: |-
  ClusterLicensingLicenses data source
---

# Data Source cluster licensing licenses

Retrieves Cluster Licensing Licenses

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_cluster_licensing_licenses" "cluster_licensing_licenses" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  filter = {
    name = "snapmirror_sy*"
  }
}

```

## Schema

### Required

- `cx_profile_name` (String) Connection profile name

### Optional

- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `cluster_licensing_licenses` (Attributes List) (see [below for nested schema](#nestedatt--cluster_licensing_licenses))

<a id="nestedatt--filter"></a>

### Nested Schema for `filter`

Optional:

- `name` (String) ClusterLicensingLicense name

<a id="nestedatt--cluster_licensing_licenses"></a>

### Nested Schema for `cluster_licensing_licenses`

Required:

- `cx_profile_name` (String) Connection profile name
- `name` (String) ClusterLicensingLicense name

Read-Only:

- `licenses` (Attributes List) Licenses of the license (see [below for nested schema](#nestedatt--cluster_licensing_licenses--licenses))
- `scope` (String) Scope of the license
- `state` (String) State of the license

<a id="nestedatt--cluster_licensing_licenses--licenses"></a>

### Nested Schema for `cluster_licensing_licenses.licenses`

Read-Only:

- `active` (Boolean) active of the license
- `compliance` (Attributes) compliance of the license (see [below for nested schema](#nestedatt--cluster_licensing_licenses--licenses--compliance))
- `evaluation` (Boolean) evaluation of the license
- `installed_license` (String) installed_license of the license
- `owner` (String) owner of the license
- `serial_number` (String) Serial Number of the license

<a id="nestedatt--cluster_licensing_licenses--licenses--compliance"></a>

### Nested Schema for `cluster_licensing_licenses.licenses.compliance`

Read-Only:

- `state` (String) state of the license
