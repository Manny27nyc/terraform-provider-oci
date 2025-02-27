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
	oci_core "github.com/oracle/oci-go-sdk/v52/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConsoleHistoryRequiredOnlyResource = ConsoleHistoryResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Required, Create, consoleHistoryRepresentation)

	consoleHistoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         Representation{RepType: Optional, Create: `${oci_core_instance.test_instance.id}`},
		"state":               Representation{RepType: Optional, Create: `SUCCEEDED`},
		"filter":              RepresentationGroup{Required, consoleHistoryDataSourceFilterRepresentation}}
	consoleHistoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_console_history.test_console_history.id}`}},
	}

	consoleHistoryRepresentation = map[string]interface{}{
		"instance_id":   Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"defined_tags":  Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ConsoleHistoryResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreConsoleHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreConsoleHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_console_history.test_console_history"
	datasourceName := "data.oci_core_console_histories.test_console_histories"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ConsoleHistoryResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Optional, Create, consoleHistoryRepresentation), "core", "consoleHistory", t)

	ResourceTest(t, testAccCheckCoreConsoleHistoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ConsoleHistoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Required, Create, consoleHistoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ConsoleHistoryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ConsoleHistoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Optional, Create, consoleHistoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ConsoleHistoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Optional, Update, consoleHistoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				GenerateDataSourceFromRepresentationMap("oci_core_console_histories", "test_console_histories", Optional, Update, consoleHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + ConsoleHistoryResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Optional, Update, consoleHistoryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "console_histories.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "console_histories.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "console_histories.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "console_histories.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.time_created"),
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

func testAccCheckCoreConsoleHistoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_console_history" {
			noResourceFound = false
			request := oci_core.GetConsoleHistoryRequest{}

			tmp := rs.Primary.ID
			request.InstanceConsoleHistoryId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			_, err := client.GetConsoleHistory(context.Background(), request)

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
	if !InSweeperExcludeList("CoreConsoleHistory") {
		resource.AddTestSweepers("CoreConsoleHistory", &resource.Sweeper{
			Name:         "CoreConsoleHistory",
			Dependencies: DependencyGraph["consoleHistory"],
			F:            sweepCoreConsoleHistoryResource,
		})
	}
}

func sweepCoreConsoleHistoryResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	consoleHistoryIds, err := getConsoleHistoryIds(compartment)
	if err != nil {
		return err
	}
	for _, consoleHistoryId := range consoleHistoryIds {
		if ok := SweeperDefaultResourceId[consoleHistoryId]; !ok {
			deleteConsoleHistoryRequest := oci_core.DeleteConsoleHistoryRequest{}

			deleteConsoleHistoryRequest.InstanceConsoleHistoryId = &consoleHistoryId

			deleteConsoleHistoryRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteConsoleHistory(context.Background(), deleteConsoleHistoryRequest)
			if error != nil {
				fmt.Printf("Error deleting ConsoleHistory %s %s, It is possible that the resource is already deleted. Please verify manually \n", consoleHistoryId, error)
				continue
			}
		}
	}
	return nil
}

func getConsoleHistoryIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ConsoleHistoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

	listConsoleHistoriesRequest := oci_core.ListConsoleHistoriesRequest{}
	listConsoleHistoriesRequest.CompartmentId = &compartmentId
	listConsoleHistoriesResponse, err := computeClient.ListConsoleHistories(context.Background(), listConsoleHistoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ConsoleHistory list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, consoleHistory := range listConsoleHistoriesResponse.Items {
		id := *consoleHistory.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ConsoleHistoryId", id)
	}
	return resourceIds, nil
}
