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
	oci_load_balancer "github.com/oracle/oci-go-sdk/v52/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BackendRequiredOnlyResource = BackendResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Required, Create, backendRepresentation)

	backendDataSourceRepresentation = map[string]interface{}{
		"backendset_name":  Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, backendDataSourceFilterRepresentation}}
	backendDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`${oci_load_balancer_backend.test_backend.name}`}},
	}

	backendRepresentation = map[string]interface{}{
		"backendset_name":  Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"ip_address":       Representation{RepType: Required, Create: `10.0.0.3`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             Representation{RepType: Required, Create: `10`},
		"backup":           Representation{RepType: Optional, Create: `false`, Update: `true`},
		"drain":            Representation{RepType: Optional, Create: `false`, Update: `true`},
		"offline":          Representation{RepType: Optional, Create: `false`, Update: `true`},
		"weight":           Representation{RepType: Optional, Create: `10`, Update: `11`},
	}

	BackendResourceDependencies = GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerBackendResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backend.test_backend"
	datasourceName := "data.oci_load_balancer_backends.test_backends"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BackendResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Create, backendRepresentation), "loadbalancer", "backend", t)

	ResourceTest(t, testAccCheckLoadBalancerBackendDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Required, Create, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Create, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "drain", "false"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "offline", "false"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "weight", "10"),

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
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "drain", "true"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "offline", "true"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "weight", "11"),

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
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_backends", "test_backends", Optional, Update, backendDataSourceRepresentation) +
				compartmentIdVariableStr + BackendResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Optional, Update, backendRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "backendset_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.backup", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.drain", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(datasourceName, "backends.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.offline", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.port", "10"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.weight", "11"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"backendset_name",
				"state",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerBackendDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_backend" {
			noResourceFound = false
			request := oci_load_balancer.GetBackendRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendName = &value
			}

			if value, ok := rs.Primary.Attributes["backendset_name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")

			_, err := client.GetBackend(context.Background(), request)

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
	if !InSweeperExcludeList("LoadBalancerBackend") {
		resource.AddTestSweepers("LoadBalancerBackend", &resource.Sweeper{
			Name:         "LoadBalancerBackend",
			Dependencies: DependencyGraph["backend"],
			F:            sweepLoadBalancerBackendResource,
		})
	}
}

func sweepLoadBalancerBackendResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	backendIds, err := getBackendIds(compartment)
	if err != nil {
		return err
	}
	for _, backendId := range backendIds {
		if ok := SweeperDefaultResourceId[backendId]; !ok {
			deleteBackendRequest := oci_load_balancer.DeleteBackendRequest{}

			deleteBackendRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteBackend(context.Background(), deleteBackendRequest)
			if error != nil {
				fmt.Printf("Error deleting Backend %s %s, It is possible that the resource is already deleted. Please verify manually \n", backendId, error)
				continue
			}
		}
	}
	return nil
}

func getBackendIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "BackendId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listBackendsRequest := oci_load_balancer.ListBackendsRequest{}

	backendsetNames, error := getBackendSetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting backendsetName required for Backend resource requests \n")
	}
	for _, backendsetName := range backendsetNames {
		listBackendsRequest.BackendSetName = &backendsetName

		loadBalancerIds, error := getLoadBalancerIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting loadBalancerId required for Backend resource requests \n")
		}
		for _, loadBalancerId := range loadBalancerIds {
			listBackendsRequest.LoadBalancerId = &loadBalancerId

			listBackendsResponse, err := loadBalancerClient.ListBackends(context.Background(), listBackendsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting Backend list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, backend := range listBackendsResponse.Items {
				id := *backend.Name
				resourceIds = append(resourceIds, id)
				AddResourceIdToSweeperResourceIdMap(compartmentId, "BackendId", id)
			}

		}
	}
	return resourceIds, nil
}
