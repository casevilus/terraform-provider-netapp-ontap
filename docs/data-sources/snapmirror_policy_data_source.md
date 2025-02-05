---
page_title: "netapp-ontap_snapmirror_policy Data Source - terraform-provider-netapp-ontap"
subcategory: "SnapMirror"
description: |-
  Retrieves SnapMirror policy of SVMs.
---

# Data Source snapmirror_policy

Retrieves SnapMirror policy of SVMs.

## Supported Platforms

* On-prem ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```terraform
data "netapp-ontap_snapmirror_policy" "snapmirror_policy" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  name = "Asynchronous"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `name` (String) SnapmirrorPolicy name

### Read-Only

- `comment` (String) Comment associated with the policy.
- `copy_all_source_snapshots` (Boolean) Specifies that all the source Snapshot copies (including the one created by SnapMirror before the transfer begins) should be copied to the destination on a transfer.
- `copy_latest_source_snapshot` (Boolean) Specifies that the latest source Snapshot copy (created by SnapMirror before the transfer begins) should be copied to the destination on a transfer. 'Retention' properties cannot be specified along with this property. This is applicable only to async policies. Property can only be set to 'true'.
- `create_snapshot_on_source` (Boolean) Specifies that all the source Snapshot copies (including the one created by SnapMirror before the transfer begins) should be copied to the destination on a transfer.
- `identity_preservation` (String) Specifies which configuration of the source SVM is replicated to the destination SVM.
- `network_compression_enabled` (Boolean) Specifies whether network compression is enabled for transfers
- `retention` (Attributes List) Rules for Snapshot copy retention. (see [below for nested schema](#nestedatt--retention))
- `svm_name` (String) SnapmirrorPolicy svm name
- `sync_type` (String) SnapmirrorPolicy sync type. [sync, strict_sync, automated_failover]
- `transfer_schedule_name` (String) The schedule used to update asynchronous relationships
- `type` (String) SnapmirrorPolicy type. [async, sync, continuous]
- `id` (String) SnapmirrorPolicy uuid

<a id="nestedatt--retention"></a>

### Nested Schema for `retention`

Read-Only:

- `count` (Number) Number of Snapshot copies to be kept for retention.
- `creation_schedule_name` (Attributes) Schedule used to create Snapshot copies on the destination for long term retention. 
- `label` (String) Snapshot copy label
- `prefix` (String) Specifies the prefix for the Snapshot copy name to be created as per the schedule
