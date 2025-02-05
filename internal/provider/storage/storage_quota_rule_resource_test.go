package storage_test

import (
	"fmt"
	"os"
	"testing"

	ntest "github.com/netapp/terraform-provider-netapp-ontap/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccStorageQuotaRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { ntest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ntest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create storage_quota_rule and read
			{
				Config: testAccStorageQuotaRuleResourceBasicConfig("terraform", "terraform", 100, 80),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_quota_rule.example", "qtree.name", ""),
				),
			},
			// Update a option
			{
				Config: testAccStorageQuotaRuleResourceBasicConfig("terraform", "terraform", 100, 70),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_quota_rule.example", "files.hard_limit", "100"),
					resource.TestCheckResourceAttr("netapp-ontap_quota_rule.example", "files.soft_limit", "70"),
				),
			},
			// Import and read
			{
				ResourceName:  "netapp-ontap_quota_rule.example",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s,%s,%s,%s,%s", "terraform_root", "terraform", "tree", "acc_import", "cluster4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_quota_rule.example", "name", "terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccStorageQuotaRuleResourceBasicConfig(volumeName string, svmName string, hardLimit int64, softLimit int64) string {
	host := os.Getenv("TF_ACC_NETAPP_HOST5")
	admin := os.Getenv("TF_ACC_NETAPP_USER")
	password := os.Getenv("TF_ACC_NETAPP_PASS2")
	if host == "" || admin == "" || password == "" {
		fmt.Println("TF_ACC_NETAPP_HOST5, TF_ACC_NETAPP_USER, and TF_ACC_NETAPP_PASS2 must be set for acceptance tests")
		os.Exit(1)
	}
	return fmt.Sprintf(`
provider "netapp-ontap" {
 connection_profiles = [
    {
      name = "cluster4"
      hostname = "%s"
      username = "%s"
      password = "%s"
      validate_certs = false
    },
  ]
}

resource "netapp-ontap_quota_rule" "example" {
	cx_profile_name = "cluster4"
	volume = {
	  name = "%s"
	  }
	svm = {
	  name = "%s"
	  }
	type = "tree"
	qtree = {
	  name = ""
	  }
	files = {
	  hard_limit = %v
	  soft_limit = %v
	  }
  }`, host, admin, password, volumeName, svmName, hardLimit, softLimit)
}
