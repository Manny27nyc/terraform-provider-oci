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
	UserRequiredOnlyResource = UserResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)

	UserResourceConfig = UserResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Update, userRepresentation)

	userSingularDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{RepType: Required, Create: `${oci_identity_user.test_user.id}`},
	}

	userDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"name":           Representation{RepType: Optional, Create: `JohnSmith@example.com`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, userDataSourceFilterRepresentation}}
	userDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_user.test_user.id}`}},
	}

	userRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"description":    Representation{RepType: Required, Create: `John Smith`, Update: `description2`},
		"name":           Representation{RepType: Required, Create: `JohnSmith@example.com`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"email":          Representation{RepType: Optional, Create: `email`, Update: `email2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	UserResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUserResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_user.test_user"
	datasourceName := "data.oci_identity_users.test_users"
	singularDatasourceName := "data.oci_identity_user.test_user"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+UserResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Create, userRepresentation), "identity", "user", t)

	ResourceTest(t, testAccCheckIdentityUserDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "John Smith"),
				resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + UserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + UserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Create, userRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "John Smith"),
				resource.TestCheckResourceAttr(resourceName, "email", "email"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + UserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Update, userRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email", "email2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_users", "test_users", Optional, Update, userDataSourceRepresentation) +
				compartmentIdVariableStr + UserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Update, userRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "users.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "users.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.email", "email2"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.email_verified"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.name", "JohnSmith@example.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.capabilities.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userSingularDataSourceRepresentation) +
				compartmentIdVariableStr + UserResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email", "email2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "email_verified"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "JohnSmith@example.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities.#", "1"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + UserResourceConfig,
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

func testAccCheckIdentityUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user" {
			noResourceFound = false
			request := oci_identity.GetUserRequest{}

			tmp := rs.Primary.ID
			request.UserId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetUser(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.UserLifecycleStateDeleted): true,
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
