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
	oci_mysql "github.com/oracle/oci-go-sdk/v52/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	HeatWaveClusterRequiredOnlyResource = HeatWaveClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Required, Create, heatWaveClusterRepresentation)

	HeatWaveClusterResourceConfig = HeatWaveClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Optional, Update, heatWaveClusterRepresentation)

	heatWaveClusterSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{RepType: Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	mysqlDbSystemHeatWaveRepresentation = map[string]interface{}{
		"admin_password":          Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"admin_username":          Representation{RepType: Required, Create: `adminUser`},
		"availability_domain":     Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"configuration_id":        Representation{RepType: Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              Representation{RepType: Required, Create: `MySQL.HeatWave.VM.Standard.E3`},
		"subnet_id":               Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           RepresentationGroup{Optional, mysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": Representation{RepType: Required, Create: `50`},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":             Representation{RepType: Optional, Create: `MySQL Database Service`},
		"display_name":            Representation{RepType: Optional, Create: `HeatWave-DbSystem`},
		"fault_domain":            Representation{RepType: Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}},
		"hostname_label":          Representation{RepType: Optional, Create: `hostnameLabel`},
		"ip_address":              Representation{RepType: Optional, Create: `10.0.0.3`},
		"maintenance":             RepresentationGroup{Optional, mysqlDbSystemMaintenanceRepresentation},
		"port":                    Representation{RepType: Optional, Create: `3306`},
		"port_x":                  Representation{RepType: Optional, Create: `33306`},
	}

	heatWaveClusterRepresentation = map[string]interface{}{
		"db_system_id": Representation{RepType: Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size": Representation{RepType: Required, Create: `2`, Update: `3`},
		"shape_name":   Representation{RepType: Required, Create: `MySQL.HeatWave.VM.Standard.E3`},
		"state":        Representation{RepType: Optional, Create: `INACTIVE`, Update: `ACTIVE`}, // testing stop & start actions
	}

	HeatWaveClusterResourceDependencies = MysqlConfigurationResourceConfig +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Required, Create, mysqlShapeDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemHeatWaveRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlHeatWaveClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlHeatWaveClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_heat_wave_cluster.test_heat_wave_cluster"

	singularDatasourceName := "data.oci_mysql_heat_wave_cluster.test_heat_wave_cluster"

	var resId, resId2 string

	ResourceTest(t, testAccCheckMysqlHeatWaveClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Required, Create, heatWaveClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies,
		},
		// verify Create & stop
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Optional, Create, heatWaveClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify start & updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Optional, Update, heatWaveClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_mysql_heat_wave_cluster", "test_heat_wave_cluster", Required, Create, heatWaveClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + HeatWaveClusterResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_name", "MySQL.HeatWave.VM.Standard.E3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + HeatWaveClusterResourceConfig,
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

func testAccCheckMysqlHeatWaveClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_heat_wave_cluster" {
			noResourceFound = false
			request := oci_mysql.GetHeatWaveClusterRequest{}

			if value, ok := rs.Primary.Attributes["db_system_id"]; ok {
				request.DbSystemId = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "mysql")

			response, err := client.GetHeatWaveCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.HeatWaveClusterLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("MysqlHeatWaveCluster") {
		resource.AddTestSweepers("MysqlHeatWaveCluster", &resource.Sweeper{
			Name:         "MysqlHeatWaveCluster",
			Dependencies: DependencyGraph["heatWaveCluster"],
			F:            sweepMysqlHeatWaveClusterResource,
		})
	}
}

func sweepMysqlHeatWaveClusterResource(compartment string) error {
	dbSystemClient := GetTestClients(&schema.ResourceData{}).dbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteHeatWaveClusterRequest := oci_mysql.DeleteHeatWaveClusterRequest{}
			deleteHeatWaveClusterRequest.DbSystemId = &mysqlDbSystemId

			deleteHeatWaveClusterRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteHeatWaveCluster(context.Background(), deleteHeatWaveClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting HeatWaveCluster of DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &mysqlDbSystemId, heatWaveClusterSweepWaitCondition, time.Duration(3*time.Minute),
				heatWaveClusterSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func heatWaveClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if heatWaveClusterResponse, ok := response.Response.(oci_mysql.GetHeatWaveClusterResponse); ok {
		return heatWaveClusterResponse.LifecycleState != oci_mysql.HeatWaveClusterLifecycleStateDeleted
	}
	return false
}

func heatWaveClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dbSystemClient().GetHeatWaveCluster(context.Background(), oci_mysql.GetHeatWaveClusterRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
