package svm_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	ntest "github.com/netapp/terraform-provider-netapp-ontap/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSvmResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { ntest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ntest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSvmResourceConfig("tfsvm4", "test", "default"),
				Check: resource.ComposeTestCheckFunc(
					// Check to see the svm name is correct,
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "tfsvm4"),
					// Check to see if Ipspace is set correctly
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "ipspace", "Default"),
					// Check that a ID has been set (we don't know what the vaule is as it changes
					resource.TestCheckResourceAttrSet("netapp-ontap_svm.example", "id"),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "comment", "test")),
			},
			// Update a comment
			{
				Config: testAccSvmResourceConfig("tfsvm4", "carchi8py was here", "default"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "comment", "carchi8py was here"),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "tfsvm4")),
			},
			// Update a comment with an empty string
			{
				Config: testAccSvmResourceConfig("tfsvm4", "", "default"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "comment", ""),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "tfsvm4")),
			},
			// Update snapshot policy default-1weekly and comment "carchi8py was here"
			{
				Config: testAccSvmResourceConfig("tfsvm4", "carchi8py was here", "default-1weekly"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "comment", "carchi8py was here"),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "snapshot_policy", "default-1weekly"),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "tfsvm4")),
			},
			// Update snapshot policy with empty string
			{
				Config:      testAccSvmResourceConfig("tfsvm4", "carchi8py was here", ""),
				ExpectError: regexp.MustCompile("cannot be updated with empty string"),
			},
			// change SVM name
			{
				Config: testAccSvmResourceConfig("tfsvm3", "carchi8py was here", "default"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "comment", "carchi8py was here"),
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "tfsvm3")),
			},
			// Fail if the name already exist
			{
				Config:      testAccSvmResourceConfig("terraform", "carchi8py was here", "default"),
				ExpectError: regexp.MustCompile("13434908"),
			},
			// Import and read
			{
				ResourceName:  "netapp-ontap_svm.example",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s,%s", "terraform", "cluster4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_svm.example", "name", "terraform"),
				),
			},
		},
	})
}
func testAccSvmResourceConfig(svm, comment string, snapshotPolicy string) string {
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

resource "netapp-ontap_svm" "example" {
  cx_profile_name = "cluster4"
  name = "%s"
  ipspace = "Default"
  comment = "%s"
  snapshot_policy = "%s"
  subtype = "default"
  language = "en_us.utf_8"
  aggregates = [
    {
      name = "terraform"
    },
  ]
  max_volumes = "unlimited"
}`, host, admin, password, svm, comment, snapshotPolicy)
}
