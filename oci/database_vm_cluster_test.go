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
	oci_database "github.com/oracle/oci-go-sdk/v52/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VmClusterRequiredOnlyResource = VmClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)

	VmClusterResourceConfig = VmClusterResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation)

	vmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":              Representation{RepType: Optional, Create: `vmCluster`},
		"exadata_infrastructure_id": Representation{RepType: Optional, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                    RepresentationGroup{Required, vmClusterDataSourceFilterRepresentation}}
	vmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_vm_cluster.test_vm_cluster.id}`}},
	}

	vmClusterRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              Representation{RepType: Required, Create: `4`, Update: `6`},
		"display_name":                Representation{RepType: Required, Create: `vmCluster`},
		"exadata_infrastructure_id":   Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  Representation{RepType: Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":             Representation{RepType: Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       Representation{RepType: Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"data_storage_size_in_tbs":    Representation{RepType: Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs": Representation{RepType: Optional, Create: `120`, Update: `160`},
		"db_servers":                  Representation{RepType: Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     Representation{RepType: Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": Representation{RepType: Optional, Create: `false`},
		"license_model":               Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs":          Representation{RepType: Optional, Create: `60`, Update: `90`},
		"time_zone":                   Representation{RepType: Optional, Create: `US/Pacific`},
	}

	VmClusterResourceDependencies = VmClusterNetworkValidatedResourceConfig
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_vm_cluster.test_vm_cluster"
	datasourceName := "data.oci_database_vm_clusters.test_vm_clusters"
	singularDatasourceName := "data.oci_database_vm_cluster.test_vm_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VmClusterResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Create, vmClusterRepresentation), "database", "vmCluster", t)

	ResourceTest(t, testAccCheckDatabaseVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Create, vmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Create,
					RepresentationCopyWithNewProperties(vmClusterRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "6"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_database_vm_clusters", "test_vm_clusters", Optional, Update, vmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.cpus_enabled"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_servers.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.vm_cluster_network_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr +
				GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
				VmClusterResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cpu_core_count",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetVmClusterRequest{}

			tmp := rs.Primary.ID
			request.VmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.VmClusterLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("DatabaseVmCluster") {
		resource.AddTestSweepers("DatabaseVmCluster", &resource.Sweeper{
			Name:         "DatabaseVmCluster",
			Dependencies: DependencyGraph["vmCluster"],
			F:            sweepDatabaseVmClusterResource,
		})
	}
}

func sweepDatabaseVmClusterResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	vmClusterIds, err := getVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterId := range vmClusterIds {
		if ok := SweeperDefaultResourceId[vmClusterId]; !ok {
			deleteVmClusterRequest := oci_database.DeleteVmClusterRequest{}

			deleteVmClusterRequest.VmClusterId = &vmClusterId

			deleteVmClusterRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmCluster(context.Background(), deleteVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting VmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &vmClusterId, vmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getVmClusterIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "VmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listVmClustersRequest := oci_database.ListVmClustersRequest{}
	listVmClustersRequest.CompartmentId = &compartmentId
	listVmClustersRequest.LifecycleState = oci_database.VmClusterSummaryLifecycleStateAvailable
	listVmClustersResponse, err := databaseClient.ListVmClusters(context.Background(), listVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vmCluster := range listVmClustersResponse.Items {
		id := *vmCluster.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterId", id)
	}
	return resourceIds, nil
}

func vmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterResponse, ok := response.Response.(oci_database.GetVmClusterResponse); ok {
		return vmClusterResponse.LifecycleState != oci_database.VmClusterLifecycleStateTerminated
	}
	return false
}

func vmClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetVmCluster(context.Background(), oci_database.GetVmClusterRequest{
		VmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
