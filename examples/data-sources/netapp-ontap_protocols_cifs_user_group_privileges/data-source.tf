data "netapp-ontap_protocols_cifs_user_group_privileges_data_source" "protocols_cifs_user_group_privileges" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  filter = {
    # name = "user1"
    svm_name = "ansibleSVM"
  }
}
