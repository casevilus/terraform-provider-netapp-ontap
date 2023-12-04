data "netapp-ontap_security_accounts_data_source" "security_accounts" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  filter = {
    name = "vsadmin"
    svm_name = "testImport"
  }
}
