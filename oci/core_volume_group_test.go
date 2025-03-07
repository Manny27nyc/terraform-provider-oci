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
	VolumeGroupRequiredOnlyResource = VolumeGroupRequiredOnlyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation)

	VolumeGroupResourceConfig = VolumeGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create, volumeGroupRepresentation)

	volumeGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":               Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":              RepresentationGroup{Required, volumeGroupDataSourceFilterRepresentation}}
	volumeGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_volume_group.test_volume_group.id}`}},
	}

	volumeGroupRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"source_details":      RepresentationGroup{Required, volumeGroupSourceDetailsRepresentation},
		"backup_policy_id":    Representation{RepType: Optional, Create: `${data.oci_core_volume_backup_policies.test_volume_user_defined_backup_policies.volume_backup_policies.0.id}`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	volumeGroupSourceDetailsRepresentation = map[string]interface{}{
		"type":       Representation{RepType: Required, Create: `volumeIds`},
		"volume_ids": Representation{RepType: Required, Create: `${oci_core_volume.source_volume_list.*.id}`},
	}
	sourceDetailsJumbledVolumeIdsRepresentation = map[string]interface{}{
		"type":       Representation{RepType: Required, Create: `volumeIds`},
		"volume_ids": Representation{RepType: Required, Create: []string{`${oci_core_volume.source_volume_list.*.id[1]}`, `${oci_core_volume.source_volume_list.*.id[0]}`}},
	}
	sourceDetailsSingleVolumeIdSourceDetailsRepresentation = map[string]interface{}{
		"type":       Representation{RepType: Required, Create: `volumeIds`},
		"volume_ids": Representation{RepType: Required, Create: []string{`${oci_core_volume.source_volume_list.*.id[1]}`}},
	}

	VolumeGroupResourceConfigJumbledVolumeIds = VolumeGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create,
			GetUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, sourceDetailsJumbledVolumeIdsRepresentation}, volumeGroupRepresentation))

	VolumeGroupResourceConfigSingleVolumeId = VolumeGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create,
			GetUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, sourceDetailsSingleVolumeIdSourceDetailsRepresentation}, volumeGroupRepresentation))

	VolumeGroupResourceDependencies = SourceVolumeListDependency +
		`
	data "oci_core_volume_backup_policies" "test_volume_user_defined_backup_policies" {
		compartment_id = var.compartment_id
	}
	` +
		AvailabilityDomainConfig +
		VolumeBackupPolicyRequiredOnlyResource
	VolumeGroupRequiredOnlyResourceDependencies = AvailabilityDomainConfig + SourceVolumeListDependency
	VolumeGroupAsDependency                     = GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation) + SourceVolumeListDependency
	SourceVolumeListDependency                  = `
resource "oci_core_volume" "source_volume_list" {
	count = 2
	display_name = "${format("source-volume-%d", count.index + 1)}"

	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
}
`
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeGroupResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("TestCoreVolumeGroupResource_basic is flaky in replay mode, will skip this test for checkin test")
	}

	httpreplay.SetScenario("TestCoreVolumeGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume_group.test_volume_group"
	datasourceName := "data.oci_core_volume_groups.test_volume_groups"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VolumeGroupResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create, volumeGroupRepresentation), "core", "volumeGroup", t)

	ResourceTest(t, testAccCheckCoreVolumeGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VolumeGroupRequiredOnlyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
				resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// We need to assert that the volume Group created above did cause the source volume to have the volume
		// Group id property populated correctly. Since the TF framework doesn't have a RefreshOnly directive, we are
		// using PlanOnly to trigger a refresh, and then assert on the value
		{
			Config:   config + compartmentIdVariableStr + VolumeGroupRequiredOnlyResource,
			PlanOnly: true,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("oci_core_volume.source_volume_list.0", "volume_group_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create, volumeGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VolumeGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create,
					RepresentationCopyWithNewProperties(volumeGroupRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

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
			Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Update, volumeGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify that the change in order of the volume ids doesn't result in a new resource
		{
			Config:             config + compartmentIdVariableStr + VolumeGroupResourceConfigJumbledVolumeIds,
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// verify that the change in list of volume ids does cause a change in the plan
		{
			Config:             config + compartmentIdVariableStr + VolumeGroupResourceConfigSingleVolumeId,
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_volume_groups", "test_volume_groups", Optional, Update, volumeGroupDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Update, volumeGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "volume_groups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.is_hydrated"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.size_in_mbs"),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.0.type", "volumeIds"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.volume_ids.#"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"backup_policy_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreVolumeGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_group" {
			noResourceFound = false
			request := oci_core.GetVolumeGroupRequest{}

			tmp := rs.Primary.ID
			request.VolumeGroupId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetVolumeGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeGroupLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("CoreVolumeGroup") {
		resource.AddTestSweepers("CoreVolumeGroup", &resource.Sweeper{
			Name:         "CoreVolumeGroup",
			Dependencies: DependencyGraph["volumeGroup"],
			F:            sweepCoreVolumeGroupResource,
		})
	}
}

func sweepCoreVolumeGroupResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()
	volumeGroupIds, err := getVolumeGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeGroupId := range volumeGroupIds {
		if ok := SweeperDefaultResourceId[volumeGroupId]; !ok {
			deleteVolumeGroupRequest := oci_core.DeleteVolumeGroupRequest{}

			deleteVolumeGroupRequest.VolumeGroupId = &volumeGroupId

			deleteVolumeGroupRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeGroup(context.Background(), deleteVolumeGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeGroupId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &volumeGroupId, volumeGroupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeGroupIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "VolumeGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()

	listVolumeGroupsRequest := oci_core.ListVolumeGroupsRequest{}
	listVolumeGroupsRequest.CompartmentId = &compartmentId
	listVolumeGroupsRequest.LifecycleState = oci_core.VolumeGroupLifecycleStateAvailable
	listVolumeGroupsResponse, err := blockstorageClient.ListVolumeGroups(context.Background(), listVolumeGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeGroup := range listVolumeGroupsResponse.Items {
		id := *volumeGroup.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeGroupId", id)
	}
	return resourceIds, nil
}

func volumeGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeGroupResponse, ok := response.Response.(oci_core.GetVolumeGroupResponse); ok {
		return volumeGroupResponse.LifecycleState != oci_core.VolumeGroupLifecycleStateTerminated
	}
	return false
}

func volumeGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient().GetVolumeGroup(context.Background(), oci_core.GetVolumeGroupRequest{
		VolumeGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
