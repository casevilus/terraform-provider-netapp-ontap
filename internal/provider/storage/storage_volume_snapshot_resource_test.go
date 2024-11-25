package storage_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	ntest "github.com/netapp/terraform-provider-netapp-ontap/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccStorageVolumeSnapshotResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { ntest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ntest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// non-existant SVM return code 2621462. Must happen before create/read
			{
				Config:      testAccStorageVolumeSnapshotResourceConfig("non-existant", "my comment"),
				ExpectError: regexp.MustCompile("svm non-existant not found"),
			},
			// Create and read testing
			{
				Config: testAccStorageVolumeSnapshotResourceConfig("terraform", "my comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "volume_name", "terraform"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "name", "snaptest"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "svm_name", "terraform"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "comment", "my comment"),
				),
			},
			// Update and read testing
			{
				Config: testAccStorageVolumeSnapshotResourceConfig("terraform", "new comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "volume_name", "terraform"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "name", "snaptest"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "svm_name", "terraform"),
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "comment", "new comment"),
				),
			},
			// Test importing a resource
			{
				ResourceName:  "netapp-ontap_volume_snapshot.example",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s,%s,%s,%s", "snaptest", "terraform", "terraform", "cluster4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_volume_snapshot.example", "name", "snaptest"),
				),
			},
		},
	})
}

func testAccStorageVolumeSnapshotResourceConfig(svmName string, comment string) string {
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

resource "netapp-ontap_volume_snapshot" "example" {
  cx_profile_name = "cluster4"
  name = "snaptest"
  volume_name = "terraform"
  svm_name = "%s"
  comment = "%s"
}`, host, admin, password, svmName, comment)
}
