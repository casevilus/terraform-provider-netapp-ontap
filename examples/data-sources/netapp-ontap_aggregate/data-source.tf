# data "netapp-ontap_aggregate" "storage_aggregate" {
#   # required to know which system to interface with
#   cx_profile_name = "cluster4"
#   name = "aggr1"
# }
data "netapp-ontap_aggregate" "storage_aggregate" {
  # required to know which system to interface with
  cx_profile_name = "fsx"
  name = "aggr1"
}