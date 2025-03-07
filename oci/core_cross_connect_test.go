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
	CrossConnectRequiredOnlyResource = CrossConnectResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation)

	CrossConnectResourceConfig = CrossConnectResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectRepresentation)

	crossConnectSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{RepType: Required, Create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	crossConnectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         RepresentationGroup{Required, crossConnectDataSourceFilterRepresentation}}
	crossConnectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_cross_connect.test_cross_connect.id}`}},
	}

	crossConnectRepresentation = map[string]interface{}{
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"location_name":           Representation{RepType: Required, Create: `${data.oci_core_cross_connect_locations.test_cross_connect_locations.cross_connect_locations.0.name}`},
		"port_speed_shape_name":   Representation{RepType: Required, Create: `10 Gbps`},
		"customer_reference_name": Representation{RepType: Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_active":               Representation{RepType: Optional, Create: `true`},
		"macsec_properties":       RepresentationGroup{Optional, crossConnectMacsecPropertiesRepresentation},
	}
	crossConnectMacsecPropertiesRepresentation = map[string]interface{}{
		"state":             Representation{RepType: Required, Create: `ENABLED`, Update: `ENABLED`},
		"encryption_cipher": Representation{RepType: Optional, Create: `AES256_GCM`, Update: `AES256_GCM_XPN`},
		"primary_key":       RepresentationGroup{Optional, crossConnectMacsecPropertiesPrimaryKeyRepresentation},
	}
	crossConnectMacsecPropertiesPrimaryKeyRepresentation = map[string]interface{}{
		"connectivity_association_key_secret_id":  Representation{RepType: Required, Create: `${var.secret_ocid_ckn}`},
		"connectivity_association_name_secret_id": Representation{RepType: Required, Create: `${var.secret_ocid_cak}`},
	}

	CrossConnectResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", Required, Create, crossConnectGroupRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", Required, Create, crossConnectLocationDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreCrossConnectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretIdCKN := getEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := getEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := getEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := getEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	resourceName := "oci_core_cross_connect.test_cross_connect"
	datasourceName := "data.oci_core_cross_connects.test_cross_connects"
	singularDatasourceName := "data.oci_core_cross_connect.test_cross_connect"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+CrossConnectResourceDependencies+secretIdVariableStrCKN+secretIdVariableStrCAK+secretVersionStrCKN+secretVersionStrCAK+
		GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create, crossConnectRepresentation), "core", "crossConnect", t)

	ResourceTest(t, testAccCheckCoreCrossConnectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK +
				GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "location_name"),
				resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "state", "PENDING_CUSTOMER"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK +
				GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create, crossConnectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "location_name"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CrossConnectResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create,
					RepresentationCopyWithNewProperties(crossConnectRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "location_name"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),

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
			Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "location_name"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

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
				GenerateDataSourceFromRepresentationMap("oci_core_cross_connects", "test_cross_connects", Optional, Update, crossConnectDataSourceRepresentation) +
				compartmentIdVariableStr + CrossConnectResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "cross_connects.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.location_name"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.macsec_properties.0.primary_key.0.connectivity_association_key_secret_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.macsec_properties.0.primary_key.0.connectivity_association_name_secret_version"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.port_name"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.port_speed_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.state", "PROVISIONED"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CrossConnectResourceConfig + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "location_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port_speed_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CrossConnectResourceConfig + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify resource import
		// import requires full configuration to handle cross connect dependency on cross connect Group during destroy
		{
			Config:            config + compartmentIdVariableStr + CrossConnectResourceConfig + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cross_connect_id",
				"is_active",
				"far_cross_connect_or_cross_connect_group_id",
				"near_cross_connect_or_cross_connect_group_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreCrossConnectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cross_connect" {
			noResourceFound = false
			request := oci_core.GetCrossConnectRequest{}

			tmp := rs.Primary.ID
			request.CrossConnectId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetCrossConnect(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CrossConnectLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("CoreCrossConnect") {
		resource.AddTestSweepers("CoreCrossConnect", &resource.Sweeper{
			Name:         "CoreCrossConnect",
			Dependencies: DependencyGraph["crossConnect"],
			F:            sweepCoreCrossConnectResource,
		})
	}
}

func sweepCoreCrossConnectResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	crossConnectIds, err := getCrossConnectIds(compartment)
	if err != nil {
		return err
	}
	for _, crossConnectId := range crossConnectIds {
		if ok := SweeperDefaultResourceId[crossConnectId]; !ok {
			deleteCrossConnectRequest := oci_core.DeleteCrossConnectRequest{}

			deleteCrossConnectRequest.CrossConnectId = &crossConnectId

			deleteCrossConnectRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCrossConnect(context.Background(), deleteCrossConnectRequest)
			if error != nil {
				fmt.Printf("Error deleting CrossConnect %s %s, It is possible that the resource is already deleted. Please verify manually \n", crossConnectId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &crossConnectId, crossConnectSweepWaitCondition, time.Duration(3*time.Minute),
				crossConnectSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCrossConnectIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "CrossConnectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listCrossConnectsRequest := oci_core.ListCrossConnectsRequest{}
	listCrossConnectsRequest.CompartmentId = &compartmentId
	listCrossConnectsRequest.LifecycleState = oci_core.CrossConnectLifecycleStateProvisioned
	listCrossConnectsResponse, err := virtualNetworkClient.ListCrossConnects(context.Background(), listCrossConnectsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CrossConnect list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, crossConnect := range listCrossConnectsResponse.Items {
		id := *crossConnect.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "CrossConnectId", id)
	}
	return resourceIds, nil
}

func crossConnectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if crossConnectResponse, ok := response.Response.(oci_core.GetCrossConnectResponse); ok {
		return crossConnectResponse.LifecycleState != oci_core.CrossConnectLifecycleStateTerminated
	}
	return false
}

func crossConnectSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetCrossConnect(context.Background(), oci_core.GetCrossConnectRequest{
		CrossConnectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
