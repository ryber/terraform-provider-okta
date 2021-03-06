package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAuthServer(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager("okta_auth_server")
	config := mgr.GetFixtures("datasource.tf", ri, t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("okta_auth_server.test", "id"),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "name", fmt.Sprintf("testAcc_%d", ri)),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "status", "ACTIVE"),
				),
			},
		},
	})
}
