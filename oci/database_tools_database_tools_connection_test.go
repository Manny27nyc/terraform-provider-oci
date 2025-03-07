// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v52/databasetools"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseToolsConnectionRequiredOnlyResource = DatabaseToolsConnectionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Required, Create, databaseToolsConnectionRepresentation)

	DatabaseToolsConnectionResourceConfig = DatabaseToolsConnectionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Update, databaseToolsConnectionRepresentation)

	databaseToolsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": Representation{RepType: Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
	}

	databaseToolsConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `tf_connection_name`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"type":           Representation{RepType: Optional, Create: []string{`ORACLE_DATABASE`}},
		"filter":         RepresentationGroup{Required, databaseToolsConnectionDataSourceFilterRepresentation}}
	databaseToolsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`}},
	}

	databaseToolsConnectionRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Required, Create: `tf_connection_name`, Update: `displayName2`},
		"type":                Representation{RepType: Required, Create: `ORACLE_DATABASE`},
		"advanced_properties": Representation{RepType: Optional, Create: map[string]string{"oracle.jdbc.loginTimeout": "0"}, Update: map[string]string{"oracle.net.CONNECT_TIMEOUT": "0"}},
		"connection_string":   Representation{RepType: Required, Create: `tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com`, Update: `connectionString2`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          RepresentationGroup{Optional, databaseToolsConnectionKeyStoresRepresentation},
		"private_endpoint_id": Representation{RepType: Optional, Create: `${oci_database_tools_database_tools_private_endpoint.test_private_endpoint.id}`},
		"related_resource":    RepresentationGroup{Optional, databaseToolsConnectionRelatedResourceRepresentation},
		"user_name":           Representation{RepType: Required, Create: `${oci_identity_user.test_user.name}`},
		"user_password":       RepresentationGroup{Required, databaseToolsConnectionUserPasswordRepresentation},
		"lifecycle":           RepresentationGroup{Required, ignoreChangesDatabaseToolsConnectionRepresentation},
	}
	databaseToolsConnectionKeyStoresRepresentation = map[string]interface{}{
		"key_store_content":  RepresentationGroup{Optional, databaseToolsConnectionKeyStoresKeyStoreContentRepresentation},
		"key_store_password": RepresentationGroup{Optional, databaseToolsConnectionKeyStoresKeyStorePasswordRepresentation},
		"key_store_type":     Representation{RepType: Optional, Create: `JAVA_KEY_STORE`, Update: `JAVA_TRUST_STORE`},
	}
	databaseToolsConnectionRelatedResourceRepresentation = map[string]interface{}{
		"entity_type": Representation{RepType: Required, Create: `DATABASE`},
		"identifier":  Representation{RepType: Required, Create: `ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q`, Update: `identifier2`},
	}
	databaseToolsConnectionUserPasswordRepresentation = map[string]interface{}{
		"value_type": Representation{RepType: Required, Create: `SECRETID`},
		"secret_id":  Representation{RepType: Required, Create: `ocid1.vaultsecret.region1.sea.amaaaaaazlynb3aahrylxtg7peotj6yybjblsqocjumsg5fp6g1111111111`}, // ${oci_vault_secret.test_secret.id}
	}
	databaseToolsConnectionKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": Representation{RepType: Required, Create: `SECRETID`},
		"secret_id":  Representation{RepType: Optional, Create: `ocid1.vaultsecret.region1.sea.amaaaaaazlynb3aahrylxtg7peotj6yybjblsqocjumsg5fp6g1111111111`}, // `${oci_vault_secret.test_secret.id}`},
	}
	databaseToolsConnectionKeyStoresKeyStorePasswordRepresentation = map[string]interface{}{
		"value_type": Representation{RepType: Required, Create: `SECRETID`},
		"secret_id":  Representation{RepType: Optional, Create: `ocid1.vaultsecret.region1.sea.amaaaaaazlynb3aahrylxtg7peotj6yybjblsqocjumsg5fp6g1111111111`}, //`${oci_vault_secret.test_secret.id}`},
	}

	ignoreChangesDatabaseToolsConnectionRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsConnectionResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_private_endpoint", Required, Create, databaseToolsPrivateEndpointRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestDatabaseToolsDatabaseToolsConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_connection"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DatabaseToolsConnectionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Create, databaseToolsConnectionRepresentation), "databasetools", "databaseToolsConnection", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// 0. verify create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Required, Create, databaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// 1. delete before next create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation),
			},
			// 2. verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Create, databaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // On R1: "3": "1" + "2" for Operators = "3"
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

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

			// 3. verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Create,
						RepresentationCopyWithNewProperties(databaseToolsConnectionRepresentation, map[string]interface{}{
							"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// 4. verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Update, databaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "identifier2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// 5. verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_connections", Optional, Update, databaseToolsConnectionDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Optional, Update, databaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.#", "1"),
				),
			},
			// 6. verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", Required, Create, databaseToolsConnectionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "connectionString2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.identifier", "identifier2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
				),
			},
			// 7. remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceConfig +
					GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", Required, Create, databaseToolsEndpointServiceDataSourceRepresentation),
			},
			// 8. verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_connection" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsConnectionRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.DatabaseToolsConnection.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.DatabaseToolsConnection.GetLifecycleState())
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DatabaseToolsDatabaseToolsConnection") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsConnection", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsConnection",
			Dependencies: DependencyGraph["databaseToolsConnection"],
			F:            sweepDatabaseToolsDatabaseToolsConnectionResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsConnectionResource(compartment string) error {
	databaseToolsClient := GetTestClients(&schema.ResourceData{}).databaseToolsClient()
	databaseToolsConnectionIds, err := getDatabaseToolsConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
		if ok := SweeperDefaultResourceId[databaseToolsConnectionId]; !ok {
			deleteDatabaseToolsConnectionRequest := oci_database_tools.DeleteDatabaseToolsConnectionRequest{}

			deleteDatabaseToolsConnectionRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

			deleteDatabaseToolsConnectionRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsConnection(context.Background(), deleteDatabaseToolsConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &databaseToolsConnectionId, databaseToolsConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				databaseToolsConnectionSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsConnectionIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := GetTestClients(&schema.ResourceData{}).databaseToolsClient()

	listDatabaseToolsConnectionsRequest := oci_database_tools.ListDatabaseToolsConnectionsRequest{}
	listDatabaseToolsConnectionsRequest.CompartmentId = &compartmentId
	listDatabaseToolsConnectionsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsConnectionsLifecycleStateActive
	listDatabaseToolsConnectionsResponse, err := databaseToolsClient.ListDatabaseToolsConnections(context.Background(), listDatabaseToolsConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsConnection := range listDatabaseToolsConnectionsResponse.Items {
		id := *databaseToolsConnection.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionId", id)
	}
	return resourceIds, nil
}

func databaseToolsConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsConnectionResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsConnectionResponse); ok {
		return databaseToolsConnectionResponse.DatabaseToolsConnection.GetLifecycleState() != oci_database_tools.LifecycleStateDeleted
	}
	return false
}

func databaseToolsConnectionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseToolsClient().GetDatabaseToolsConnection(context.Background(), oci_database_tools.GetDatabaseToolsConnectionRequest{
		DatabaseToolsConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
