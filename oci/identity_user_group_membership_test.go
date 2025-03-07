// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_identity "github.com/oracle/oci-go-sdk/v52/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	userGroupMembershipDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"group_id":       Representation{RepType: Optional, Create: `${oci_identity_group.test_group.id}`},
		"user_id":        Representation{RepType: Optional, Create: `${oci_identity_user.test_user.id}`},
		"filter":         RepresentationGroup{Required, userGroupMembershipDataSourceFilterRepresentation}}
	userGroupMembershipDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_user_group_membership.test_user_group_membership.id}`}},
	}

	userGroupMembershipRepresentation = map[string]interface{}{
		"group_id": Representation{RepType: Required, Create: `${oci_identity_group.test_group.id}`},
		"user_id":  Representation{RepType: Required, Create: `${oci_identity_user.test_user.id}`},
	}

	UserGroupMembershipResourceDependencies = GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUserGroupMembershipResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"
	datasourceName := "data.oci_identity_user_group_memberships.test_user_group_memberships"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+UserGroupMembershipResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", Required, Create, userGroupMembershipRepresentation), "identity", "userGroupMembership", t)

	ResourceTest(t, testAccCheckIdentityUserGroupMembershipDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UserGroupMembershipResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", Required, Create, userGroupMembershipRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_user_group_memberships", "test_user_group_memberships", Optional, Update, userGroupMembershipDataSourceRepresentation) +
				compartmentIdVariableStr + UserGroupMembershipResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", Optional, Update, userGroupMembershipRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "memberships.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.user_id"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityUserGroupMembershipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user_group_membership" {
			noResourceFound = false
			request := oci_identity.GetUserGroupMembershipRequest{}

			tmp := rs.Primary.ID
			request.UserGroupMembershipId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetUserGroupMembership(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.UserGroupMembershipLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
