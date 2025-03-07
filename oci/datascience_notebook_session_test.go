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
	oci_datascience "github.com/oracle/oci-go-sdk/v52/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NotebookSessionRequiredOnlyResource = NotebookSessionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Required, Create, notebookSessionRepresentation)

	NotebookSessionResourceConfig = NotebookSessionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionRepresentation)

	notebookSessionSingularDataSourceRepresentation = map[string]interface{}{
		"notebook_session_id": Representation{RepType: Required, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
	}

	notebookSessionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
		"project_id":     Representation{RepType: Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, notebookSessionDataSourceFilterRepresentation}}
	notebookSessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_datascience_notebook_session.test_notebook_session.id}`}},
	}

	notebookSessionRepresentation = map[string]interface{}{
		"compartment_id":                         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"notebook_session_configuration_details": RepresentationGroup{Required, notebookSessionNotebookSessionConfigurationDetailsRepresentation},
		"project_id":                             Representation{RepType: Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                           Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                           Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	notebookSessionNotebookSessionConfigurationDetailsRepresentation = map[string]interface{}{
		"shape":                     Representation{RepType: Required, Create: `VM.Standard.E2.2`, Update: `VM.Standard.E2.4`},
		"subnet_id":                 Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_storage_size_in_gbs": Representation{RepType: Optional, Create: `50`, Update: `51`},
	}

	NotebookSessionResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	datasourceName := "data.oci_datascience_notebook_sessions.test_notebook_sessions"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+NotebookSessionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create, notebookSessionRepresentation), "datascience", "notebookSession", t)

	ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Required, Create, notebookSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create, notebookSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create,
					RepresentationCopyWithNewProperties(notebookSessionRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_sessions", "test_notebook_sessions", Optional, Update, notebookSessionDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Required, Create, notebookSessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceConfig,
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

func testAccCheckDatascienceNotebookSessionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_notebook_session" {
			noResourceFound = false
			request := oci_datascience.GetNotebookSessionRequest{}

			tmp := rs.Primary.ID
			request.NotebookSessionId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datascience")

			response, err := client.GetNotebookSession(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.NotebookSessionLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("DatascienceNotebookSession") {
		resource.AddTestSweepers("DatascienceNotebookSession", &resource.Sweeper{
			Name:         "DatascienceNotebookSession",
			Dependencies: DependencyGraph["notebookSession"],
			F:            sweepDatascienceNotebookSessionResource,
		})
	}
}

func sweepDatascienceNotebookSessionResource(compartment string) error {
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()
	notebookSessionIds, err := getNotebookSessionIds(compartment)
	if err != nil {
		return err
	}
	for _, notebookSessionId := range notebookSessionIds {
		if ok := SweeperDefaultResourceId[notebookSessionId]; !ok {
			deleteNotebookSessionRequest := oci_datascience.DeleteNotebookSessionRequest{}

			deleteNotebookSessionRequest.NotebookSessionId = &notebookSessionId

			deleteNotebookSessionRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteNotebookSession(context.Background(), deleteNotebookSessionRequest)
			if error != nil {
				fmt.Printf("Error deleting NotebookSession %s %s, It is possible that the resource is already deleted. Please verify manually \n", notebookSessionId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &notebookSessionId, notebookSessionSweepWaitCondition, time.Duration(3*time.Minute),
				notebookSessionSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getNotebookSessionIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "NotebookSessionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()

	listNotebookSessionsRequest := oci_datascience.ListNotebookSessionsRequest{}
	listNotebookSessionsRequest.CompartmentId = &compartmentId
	listNotebookSessionsRequest.LifecycleState = oci_datascience.ListNotebookSessionsLifecycleStateActive
	listNotebookSessionsResponse, err := dataScienceClient.ListNotebookSessions(context.Background(), listNotebookSessionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NotebookSession list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, notebookSession := range listNotebookSessionsResponse.Items {
		id := *notebookSession.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "NotebookSessionId", id)
	}
	return resourceIds, nil
}

func notebookSessionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if notebookSessionResponse, ok := response.Response.(oci_datascience.GetNotebookSessionResponse); ok {
		return notebookSessionResponse.LifecycleState != oci_datascience.NotebookSessionLifecycleStateDeleted
	}
	return false
}

func notebookSessionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataScienceClient().GetNotebookSession(context.Background(), oci_datascience.GetNotebookSessionRequest{
		NotebookSessionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
