package networking_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	ntest "github.com/netapp/terraform-provider-netapp-ontap/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkIpRouteResourceAlias(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { ntest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ntest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Missing Required argument
			{
				Config:      testAccNetworkIPIRouteResourceConfigAliasMissingVars("non-existent"),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
			// Non existent SVM
			{
				Config:      testAccNetworkIPIRouteResourceConfigAlias("non-existent"),
				ExpectError: regexp.MustCompile("2621462"),
			},
			// Test create with no gateway
			{
				Config: testAccNetworkIPIRouteResourceConfigAlias("ansibleSVM"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "svm_name", "ansibleSVM"),
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "destination.address", "0.0.0.0"),
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "destination.netmask", "0"),
				),
			},
			// test create with a gateway
			{
				Config: testAccNetworkIPIRouteResourceWithGatewayConfigAlias("ansibleSVM", "10.10.10.254", 20),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "svm_name", "ansibleSVM"),
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "destination.address", "10.10.10.254"),
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "destination.netmask", "20"),
				),
			},
			// Import and read
			{
				ResourceName:  "netapp-ontap_networking_ip_route_resource.example",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s,%s,%s", "carchi-test", "10.10.10.254", "cluster4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netapp-ontap_networking_ip_route_resource.example", "svm_name", "carchi-test"),
				),
			},
		},
	})
}

func testAccNetworkIPIRouteResourceConfigAlias(svmName string) string {
	host := os.Getenv("TF_ACC_NETAPP_HOST")
	admin := os.Getenv("TF_ACC_NETAPP_USER")
	password := os.Getenv("TF_ACC_NETAPP_PASS")
	if host == "" || admin == "" || password == "" {
		fmt.Println("TF_ACC_NETAPP_HOST, TF_ACC_NETAPP_USER, and TF_ACC_NETAPP_PASS must be set for acceptance tests")
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

resource "netapp-ontap_networking_ip_route_resource" "example" {
  cx_profile_name = "cluster4"
  svm_name = "%s"
  gateway = "10.10.10.1"
}
`, host, admin, password, svmName)
}

func testAccNetworkIPIRouteResourceWithGatewayConfigAlias(svmName string, address string, netmask int) string {
	host := os.Getenv("TF_ACC_NETAPP_HOST")
	admin := os.Getenv("TF_ACC_NETAPP_USER")
	password := os.Getenv("TF_ACC_NETAPP_PASS")
	if host == "" || admin == "" || password == "" {
		fmt.Println("TF_ACC_NETAPP_HOST, TF_ACC_NETAPP_USER, and TF_ACC_NETAPP_PASS must be set for acceptance tests")
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

resource "netapp-ontap_networking_ip_route_resource" "example" {
  cx_profile_name = "cluster4"
  svm_name = "%s"
  gateway = "10.10.10.1"
  destination = {
    address = "%s"
    netmask = %d
    }
}
`, host, admin, password, svmName, address, netmask)
}

func testAccNetworkIPIRouteResourceConfigAliasMissingVars(svmName string) string {
	host := os.Getenv("TF_ACC_NETAPP_HOST")
	admin := os.Getenv("TF_ACC_NETAPP_USER")
	password := os.Getenv("TF_ACC_NETAPP_PASS")
	if host == "" || admin == "" || password == "" {
		fmt.Println("TF_ACC_NETAPP_HOST, TF_ACC_NETAPP_USER, and TF_ACC_NETAPP_PASS must be set for acceptance tests")
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

resource "netapp-ontap_networking_ip_route_resource" "example" {
  cx_profile_name = "cluster4"
  svm_name = "%s"
}
`, host, admin, password, svmName)
}
