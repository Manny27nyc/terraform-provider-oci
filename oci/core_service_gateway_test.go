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
	oci_core "github.com/oracle/oci-go-sdk/v52/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ServiceGatewayRequiredOnlyResource = ServiceGatewayResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Required, Create, serviceGatewayRepresentation)

	serviceGatewayDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, serviceGatewayDataSourceFilterRepresentation}}
	serviceGatewayDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_service_gateway.test_service_gateway.id}`}},
	}

	serviceGatewayRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"services":       RepresentationGroup{Required, serviceGatewayServicesRepresentation},
		"vcn_id":         Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `MyServiceGateway`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"route_table_id": Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
	}
	serviceGatewayServicesRepresentation = map[string]interface{}{
		"service_id": Representation{RepType: Required, Create: `${lookup(data.oci_core_services.test_services.services[0], "id")}`},
	}

	ServiceGatewayResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/serviceGateway
func TestCoreServiceGatewayResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreServiceGatewayResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_service_gateway.test_service_gateway"
	datasourceName := "data.oci_core_service_gateways.test_service_gateways"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ServiceGatewayResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Create, serviceGatewayRepresentation), "core", "serviceGateway", t)

	ResourceTest(t, testAccCheckCoreServiceGatewayDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Required, Create, serviceGatewayRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
					[]string{
						"service_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Create, serviceGatewayRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyServiceGateway"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
					[]string{
						"service_id",
						"service_name",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceGatewayResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Create,
					RepresentationCopyWithNewProperties(serviceGatewayRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyServiceGateway"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
					[]string{
						"service_id",
						"service_name",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Update, serviceGatewayRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
					[]string{
						"service_id",
						"service_name",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_core_service_gateways", "test_service_gateways", Optional, Update, serviceGatewayDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceGatewayResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Optional, Update, serviceGatewayRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "service_gateways.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.block_traffic"),
				resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.route_table_id"),
				resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.services.#", "1"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "service_gateways.0.services", map[string]string{},
					[]string{
						"service_id",
						"service_name",
					}),
				resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.vcn_id"),
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

func testAccCheckCoreServiceGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_service_gateway" {
			noResourceFound = false
			request := oci_core.GetServiceGatewayRequest{}

			tmp := rs.Primary.ID
			request.ServiceGatewayId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetServiceGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ServiceGatewayLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("CoreServiceGateway") {
		resource.AddTestSweepers("CoreServiceGateway", &resource.Sweeper{
			Name:         "CoreServiceGateway",
			Dependencies: DependencyGraph["serviceGateway"],
			F:            sweepCoreServiceGatewayResource,
		})
	}
}

func sweepCoreServiceGatewayResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	serviceGatewayIds, err := getServiceGatewayIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceGatewayId := range serviceGatewayIds {
		if ok := SweeperDefaultResourceId[serviceGatewayId]; !ok {
			deleteServiceGatewayRequest := oci_core.DeleteServiceGatewayRequest{}

			deleteServiceGatewayRequest.ServiceGatewayId = &serviceGatewayId

			deleteServiceGatewayRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteServiceGateway(context.Background(), deleteServiceGatewayRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceGateway %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceGatewayId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &serviceGatewayId, serviceGatewaySweepWaitCondition, time.Duration(3*time.Minute),
				serviceGatewaySweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getServiceGatewayIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ServiceGatewayId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listServiceGatewaysRequest := oci_core.ListServiceGatewaysRequest{}
	listServiceGatewaysRequest.CompartmentId = &compartmentId
	listServiceGatewaysRequest.LifecycleState = oci_core.ServiceGatewayLifecycleStateAvailable
	listServiceGatewaysResponse, err := virtualNetworkClient.ListServiceGateways(context.Background(), listServiceGatewaysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceGateway list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceGateway := range listServiceGatewaysResponse.Items {
		id := *serviceGateway.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ServiceGatewayId", id)
	}
	return resourceIds, nil
}

func serviceGatewaySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceGatewayResponse, ok := response.Response.(oci_core.GetServiceGatewayResponse); ok {
		return serviceGatewayResponse.LifecycleState != oci_core.ServiceGatewayLifecycleStateTerminated
	}
	return false
}

func serviceGatewaySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetServiceGateway(context.Background(), oci_core.GetServiceGatewayRequest{
		ServiceGatewayId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
