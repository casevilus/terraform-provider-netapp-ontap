resource "netapp-ontap_quota_rule" "storage_quota_rule" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  volume = {
    name = "lunTest"
    }
  svm = {
    name = "carchi-test"
    }
  type = "tree"
  qtree = {
    name = ""
    }
  users = [{
    name = ""
    }]
  files = {
    hard_limit = 100
    soft_limit = 80
    }
}
