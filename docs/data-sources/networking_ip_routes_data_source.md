---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_network_ip_routes Data Source - terraform-provider-netapp-ontap"
subcategory: "Networking"
description: |-
  Retrieves the IP routes of SVMs.
---

# Data Source network_ip_routes

Retrieves the IP routes of SVMs.

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_network_ip_routes" "network_ip_routes" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  gateway = "10.10.10.254"
  filter = {
    svm_name = "*a*"
    destination = {
      address = "0.0.0.0",
      netmask = "24",
    }
    gateway = "10.*"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `gateway` (String) The IP address of the gateway router leading to the destination.

### Optional

- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `ip_routes` (Attributes List) (see [below for nested schema](#nestedatt--ip_routes))

<a id="nestedatt--filter"></a>

### Nested Schema for `filter`

Optional:

- `destination` (Attributes) destination IP address information (see [below for nested schema](#nestedatt--filter--destination))
- `gateway` (String) The IP address of the gateway router leading to the destination.
- `svm_name` (String) IP Route svm name

<a id="nestedatt--filter--destination"></a>

### Nested Schema for `filter.destination`

Optional:

- `address` (String) IPv4 or IPv6 address
- `netmask` (String) netmask length (16) or IPv4 mask (255.255.0.0). For IPv6, valid range is 1 to 127.



<a id="nestedatt--ip_routes"></a>

### Nested Schema for `ip_routes`

Required:

- `cx_profile_name` (String) Connection profile name
- `destination` (Attributes) destination IP address information (see [below for nested schema](#nestedatt--ip_routes--destination))
- `gateway` (String) The IP address of the gateway router leading to the destination.

Optional:

- `svm_name` (String) IPInterface svm name

Read-Only:

- `metric` (Number) Indicates a preference order between several routes to the same destination.

<a id="nestedatt--ip_routes--destination"></a>

### Nested Schema for `ip_routes.destination`

Required:

- `address` (String) IPv4 or IPv6 address

Read-Only:

- `netmask` (String) netmask length (16) or IPv4 mask (255.255.0.0). For IPv6, valid range is 1 to 127.
