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
	oci_nosql "github.com/oracle/oci-go-sdk/v52/nosql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TableRequiredOnlyResource = TableResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Required, Create, tableRepresentation)

	TableResourceConfig = TableResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Update, tableRepresentation)

	tableSingularDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": Representation{RepType: Required, Create: `${oci_nosql_table.test_table.id}`},
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
	}
	ddlStatement = "CREATE TABLE IF NOT EXISTS test_table(id INTEGER, name STRING, age STRING, info JSON, PRIMARY KEY(SHARD(id)))"

	tableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `test_table`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, tableDataSourceFilterRepresentation}}
	tableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_nosql_table.test_table.id}`}},
	}

	tableRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"ddl_statement":       Representation{RepType: Required, Create: ddlStatement},
		"name":                Representation{RepType: Required, Create: `test_table`},
		"table_limits":        RepresentationGroup{Required, tableTableLimitsRepresentation},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_reclaimable": Representation{RepType: Optional, Create: `false`},
	}
	tableTableLimitsRepresentation = map[string]interface{}{
		"max_read_units":     Representation{RepType: Required, Create: `10`, Update: `11`},
		"max_storage_in_gbs": Representation{RepType: Required, Create: `10`, Update: `11`},
		"max_write_units":    Representation{RepType: Required, Create: `10`, Update: `11`},
	}

	TableResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: nosql/default
func TestNosqlTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_nosql_table.test_table"

	datasourceName := "data.oci_nosql_tables.test_tables"
	singularDatasourceName := "data.oci_nosql_table.test_table"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+TableResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Create, tableRepresentation), "nosql", "table", t)

	ResourceTest(t, testAccCheckNosqlTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TableResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Required, Create, tableRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TableResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Create, tableRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TableResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Create,
					RepresentationCopyWithNewProperties(tableRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + TableResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Update, tableRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "11"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "11"),

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
				GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_tables", Optional, Update, tableDataSourceRepresentation) +
				compartmentIdVariableStr + TableResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", Optional, Update, tableRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "table_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "table_collection.0.id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_table", Required, Create, tableSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TableResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name_or_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ddl_statement", ddlStatement),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_reclaimable", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_table"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_read_units", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_storage_in_gbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_write_units", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TableResourceConfig,
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

func testAccCheckNosqlTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).nosqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_nosql_table" {
			noResourceFound = false
			request := oci_nosql.GetTableRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["table_name_or_id"]; ok {
				request.TableNameOrId = &value
			} else if rs.Primary.ID != "" {
				tmp := rs.Primary.ID
				request.TableNameOrId = &tmp
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "nosql")

			response, err := client.GetTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_nosql.TableLifecycleStateDeleted): true,
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("NosqlTable") {
		resource.AddTestSweepers("NosqlTable", &resource.Sweeper{
			Name:         "NosqlTable",
			Dependencies: DependencyGraph["table"],
			F:            sweepNosqlTableResource,
		})
	}
}

func sweepNosqlTableResource(compartment string) error {
	nosqlClient := GetTestClients(&schema.ResourceData{}).nosqlClient()
	tableIds, err := getTableIds(compartment)
	if err != nil {
		return err
	}
	for _, tableId := range tableIds {
		if ok := SweeperDefaultResourceId[tableId]; !ok {
			deleteTableRequest := oci_nosql.DeleteTableRequest{}

			deleteTableRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "nosql")
			_, error := nosqlClient.DeleteTable(context.Background(), deleteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting Table %s %s, It is possible that the resource is already deleted. Please verify manually \n", tableId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &tableId, tableSweepWaitCondition, time.Duration(3*time.Minute),
				tableSweepResponseFetchOperation, "nosql", true)
		}
	}
	return nil
}

func getTableIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "TableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	nosqlClient := GetTestClients(&schema.ResourceData{}).nosqlClient()

	listTablesRequest := oci_nosql.ListTablesRequest{}
	listTablesRequest.CompartmentId = &compartmentId
	listTablesRequest.LifecycleState = oci_nosql.ListTablesLifecycleStateActive
	listTablesResponse, err := nosqlClient.ListTables(context.Background(), listTablesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Table list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, table := range listTablesResponse.Items {
		id := *table.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "TableId", id)
	}
	return resourceIds, nil
}

func tableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if tableResponse, ok := response.Response.(oci_nosql.GetTableResponse); ok {
		return tableResponse.LifecycleState != oci_nosql.TableLifecycleStateDeleted
	}
	return false
}

func tableSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.nosqlClient().GetTable(context.Background(), oci_nosql.GetTableRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
