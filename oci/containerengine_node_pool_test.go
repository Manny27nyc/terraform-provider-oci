// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v52/containerengine"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NodePoolRequiredOnlyResource = NodePoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation)

	NodePoolResourceConfig = NodePoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation)

	nodePoolSingularDataSourceRepresentation = map[string]interface{}{
		"node_pool_id": Representation{RepType: Required, Create: `${oci_containerengine_node_pool.test_node_pool.id}`},
	}

	nodePoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cluster_id":     Representation{RepType: Optional, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"name":           Representation{RepType: Optional, Create: `name`, Update: `name2`},
		"filter":         RepresentationGroup{Required, nodePoolDataSourceFilterRepresentation}}
	nodePoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_containerengine_node_pool.test_node_pool.id}`}},
	}

	nodePoolRepresentation = map[string]interface{}{
		"cluster_id":          Representation{RepType: Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  Representation{RepType: Required, Create: `${oci_containerengine_cluster.test_cluster.kubernetes_version}`},
		"name":                Representation{RepType: Required, Create: `name`, Update: `name2`},
		"node_image_name":     Representation{RepType: Required, Create: `Oracle-Linux-7.6`},
		"node_shape":          Representation{RepType: Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		"subnet_ids":          Representation{RepType: Required, Create: []string{`${oci_core_subnet.nodePool_Subnet_1.id}`, `${oci_core_subnet.nodePool_Subnet_2.id}`}},
		"initial_node_labels": RepresentationGroup{Optional, nodePoolInitialNodeLabelsRepresentation},
		"node_metadata":       Representation{RepType: Optional, Create: map[string]string{"nodeMetadata": "nodeMetadata"}, Update: map[string]string{"nodeMetadata2": "nodeMetadata2"}},
		"quantity_per_subnet": Representation{RepType: Optional, Create: `1`, Update: `2`},
		"ssh_public_key":      Representation{RepType: Optional, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}}
	nodePoolInitialNodeLabelsRepresentation = map[string]interface{}{
		"key":   Representation{RepType: Optional, Create: `key`, Update: `key2`},
		"value": Representation{RepType: Optional, Create: `value`, Update: `value2`},
	}

	NodePoolResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool_option", "test_node_pool_option", Required, Create, nodePoolOptionSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_1", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.22.0/24`}, "dns_label": Representation{RepType: Required, Create: `nodepool1`}})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "nodePool_Subnet_2", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.23.0/24`}, "dns_label": Representation{RepType: Required, Create: `nodepool2`}})) +
		GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.20.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster1`}})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.0.21.0/24`}, "dns_label": Representation{RepType: Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		}))
)

// issue-routing-tag: containerengine/default
func TestContainerengineNodePoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineNodePoolResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_node_pool.test_node_pool"
	datasourceName := "data.oci_containerengine_node_pools.test_node_pools"
	singularDatasourceName := "data.oci_containerengine_node_pool.test_node_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+NodePoolResourceDependencies+nodePoolResourceConfigForVMStandard+
		GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Create, nodePoolRepresentation), "containerengine", "nodePool", t)

	ResourceTest(t, testAccCheckContainerengineNodePoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Required, Create, nodePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Create, nodePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

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
			Config: config + compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard + GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, GetUpdatedRepresentationCopy("node_metadata", Representation{RepType: Optional, Update: map[string]string{"nodeMetadata": "nodeMetadata"}}, nodePoolRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "node_image_name"),
				resource.TestCheckResourceAttr(resourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "node_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

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
				GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pools", "test_node_pools", Optional, Update, nodePoolDataSourceRepresentation) +
				compartmentIdVariableStr + NodePoolResourceDependencies + nodePoolResourceConfigForVMStandard +
				GenerateResourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool", Optional, Update, nodePoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

				resource.TestCheckResourceAttr(datasourceName, "node_pools.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.kubernetes_version"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_pools.0.node_image_name"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.node_source.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(datasourceName, "node_pools.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_node_pool", "test_node_pool",
					Required, Create,
					nodePoolSingularDataSourceRepresentation) + nodePoolResourceConfigForVMStandard + compartmentIdVariableStr + NodePoolResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_node_labels.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_image_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "quantity_per_subnet", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subnet_ids.#", "2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + NodePoolResourceConfig,
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

func testAccCheckContainerengineNodePoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).containerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_node_pool" {
			noResourceFound = false
			request := oci_containerengine.GetNodePoolRequest{}

			tmp := rs.Primary.ID
			request.NodePoolId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "containerengine")

			_, err := client.GetNodePool(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !InSweeperExcludeList("ContainerengineNodePool") {
		resource.AddTestSweepers("ContainerengineNodePool", &resource.Sweeper{
			Name:         "ContainerengineNodePool",
			Dependencies: DependencyGraph["nodePool"],
			F:            sweepContainerengineNodePoolResource,
		})
	}
}

func sweepContainerengineNodePoolResource(compartment string) error {
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()
	nodePoolIds, err := getNodePoolIds(compartment)
	if err != nil {
		return err
	}
	for _, nodePoolId := range nodePoolIds {
		if ok := SweeperDefaultResourceId[nodePoolId]; !ok {
			deleteNodePoolRequest := oci_containerengine.DeleteNodePoolRequest{}

			deleteNodePoolRequest.NodePoolId = &nodePoolId

			deleteNodePoolRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteNodePool(context.Background(), deleteNodePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting NodePool %s %s, It is possible that the resource is already deleted. Please verify manually \n", nodePoolId, error)
				continue
			}
		}
	}
	return nil
}

func getNodePoolIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "NodePoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()

	listNodePoolsRequest := oci_containerengine.ListNodePoolsRequest{}
	listNodePoolsRequest.CompartmentId = &compartmentId
	listNodePoolsResponse, err := containerEngineClient.ListNodePools(context.Background(), listNodePoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NodePool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, nodePool := range listNodePoolsResponse.Items {
		id := *nodePool.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "NodePoolId", id)
	}
	return resourceIds, nil
}
