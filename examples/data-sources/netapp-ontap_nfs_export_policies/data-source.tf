data "netapp-ontap_nfs_export_policies" "export_policies" {
  cx_profile_name = "cluster4"
  filter = {
    #name = "default"
    svm_name = "svm*"
  }
}
