---
page_title: "ONTAP: NFS_Service "
subcategory: "NAS"
description: |-
  Retrieves the NFS configuration of SVMs.
---

# Data Source nfs_service

Retrieves the NFS configuration of SVMs.

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_nfs_service" "protocols_nfs_services" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  svm_name = "ansibleSVM"
}
```

<!-- schema generated by tfplugindocs -->
## Argument Reference

### Required

- `cx_profile_name` (String) Connection profile name
- `svm_name` (String) IPInterface svm name

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `enabled` (Boolean) NFS should be enabled or disabled
- `protocol` (Attributes) Protocol (see [below for nested schema](#nestedatt--protocol))
- `root` (Attributes) Specific Root user options (see [below for nested schema](#nestedatt--root))
- `security` (Attributes) NFS Security options (see [below for nested schema](#nestedatt--security))
- `showmount_enabled` (Boolean) Whether SVM allows showmount
- `transport` (Attributes) (see [below for nested schema](#nestedatt--transport))
- `vstorage_enabled` (Boolean) Whether Vstorage is enabled
- `windows` (Attributes) (see [below for nested schema](#nestedatt--windows))

~> **NOTE:** `root`, `security`, `windows` requires ONTAP 9.11 or higher

<a id="nestedatt--protocol"></a>

### Nested Schema for `protocol`

Read-Only:

- `v3_enabled` (Boolean) NFSv3 enabled
- `v40_enabled` (Boolean) NFSv4.0 enabled
- `v40_features` (Attributes) NFSv4.0 features (see [below for nested schema](#nestedatt--protocol--v40_features))
- `v41_enabled` (Boolean) NFSv4.1 enabled
- `v41_features` (Attributes) NFSv4.1 features (see [below for nested schema](#nestedatt--protocol--v41_features))
- `v4_id_domain` (String) User ID domain for NFSv4

<a id="nestedatt--protocol--v40_features"></a>

### Nested Schema for `protocol.v40_features`

Read-Only:

- `acl_enabled` (Boolean) Enable ACL for NFSv4.0
- `read_delegation_enabled` (Boolean) Enable Read File Delegation for NFSv4.0
- `write_delegation_enabled` (Boolean) Enable Write File Delegation for NFSv4.0


<a id="nestedatt--protocol--v41_features"></a>

### Nested Schema for `protocol.v41_features`

Read-Only:

- `acl_enabled` (Boolean) Enable ACL for NFSv4.1
- `pnfs_enabled` (Boolean) Enabled pNFS (parallel NFS) for NFSv4.1
- `read_delegation_enabled` (Boolean) Enable Read File Delegation for NFSv4.1
- `write_delegation_enabled` (Boolean) Enable Write File Delegation for NFSv4.1


<a id="nestedatt--root"></a>

### Nested Schema for `root`

Read-Only:

- `ignore_nt_acl` (Boolean) Ignore NTFS ACL for root user
- `skip_write_permission_check` (Boolean) Skip write permissions check for root user

<a id="nestedatt--security"></a>

### Nested Schema for `security`

Read-Only:

- `chown_mode` (String) Specifies whether file ownership can be changed only by the superuser, or if a non-root user can also change file ownership
- `nt_acl_display_permission` (Boolean) Controls the permissions that are displayed to NFSv3 and NFSv4 clients on a file or directory that has an NT ACL set
- `ntfs_unix_security` (String) Specifies how NFSv3 security changes affect NTFS volumes
- `permitted_encryption_types` (List of String) Specifies the permitted encryption types for Kerberos over NFS.
- `rpcsec_context_idle` (Number) Specifies, in seconds, the amount of time a RPCSEC_GSS context is permitted to remain unused before it is deleted

<a id="nestedatt--transport"></a>

### Nested Schema for `transport`

Read-Only:

- `tcp_enabled` (Boolean) tcp enabled
- `tcp_max_transfer_size` (Number) Max tcp transfer size
- `udp_enabled` (Boolean) udp enabled

~> **NOTE:** `tcp_max_transfer_size` requires ONTAP 9.11 or higher

<a id="nestedatt--windows"></a>

### Nested Schema for `windows`

Read-Only:

- `default_user` (String) default Windows user for the NFS server
- `map_unknown_uid_to_default_user` (Boolean) whether or not the mapping of an unknown UID to the default Windows user is enabled
- `v3_ms_dos_client_enabled` (Boolean) if permission checks are to be skipped for NFS WRITE calls from root/owner.
