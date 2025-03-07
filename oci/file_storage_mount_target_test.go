// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v52/filestorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MountTargetRequiredOnlyResource = MountTargetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Required, Create, mountTargetRepresentation)

	mountTargetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Optional, Create: `mount-target-5`, Update: `displayName2`},
		"id":                  Representation{RepType: Optional, Create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"state":               Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, mountTargetDataSourceFilterRepresentation}}
	mountTargetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_file_storage_mount_target.test_mount_target.id}`}},
	}

	mountTargetRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"subnet_id":           Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{RepType: Optional, Create: `mount-target-5`, Update: `displayName2`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":      Representation{RepType: Optional, Create: `hostnamelabel`},
		"ip_address":          Representation{RepType: Optional, Create: `10.0.0.5`},
		"nsg_ids":             Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
	}

	MountTargetResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
			"dns_label":           Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: file_storage/default
func TestFileStorageMountTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageMountTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_mount_target.test_mount_target"
	datasourceName := "data.oci_file_storage_mount_targets.test_mount_targets"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+MountTargetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Create, mountTargetRepresentation), "filestorage", "mountTarget", t)

	ResourceTest(t, testAccCheckFileStorageMountTargetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Required, Create, mountTargetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MountTargetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Create, mountTargetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids.0"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Create,
					RepresentationCopyWithNewProperties(mountTargetRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
			Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Update, mountTargetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
				GenerateDataSourceFromRepresentationMap("oci_file_storage_mount_targets", "test_mount_targets", Optional, Update, mountTargetDataSourceRepresentation) +
				compartmentIdVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", Optional, Update, mountTargetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.export_set_id"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.private_ip_ids.#"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.time_created"),
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

// issue-routing-tag: file_storage/default
func TestFileStorageMountTargetResource_failedWorkRequest(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageMountTargetResource_failedWorkRequest")
	defer httpreplay.SaveScenario()
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_mount_target.test_mount_target2"

	ResourceTest(t, testAccCheckFileStorageMountTargetDestroy, []resource.TestStep{
		// verify resource creation fails for the second mount target with the same ip_address
		{
			Config: config + compartmentIdVariableStr + MountTargetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target1", Optional, Update, mountTargetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target2", Optional, Update, mountTargetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
			),
			ExpectError: regexp.MustCompile("Resource creation failed"),
		},
	})
}

func testAccCheckFileStorageMountTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).fileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_mount_target" {
			noResourceFound = false
			request := oci_file_storage.GetMountTargetRequest{}

			tmp := rs.Primary.ID
			request.MountTargetId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "file_storage")

			response, err := client.GetMountTarget(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.MountTargetLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("FileStorageMountTarget") {
		resource.AddTestSweepers("FileStorageMountTarget", &resource.Sweeper{
			Name:         "FileStorageMountTarget",
			Dependencies: DependencyGraph["mountTarget"],
			F:            sweepFileStorageMountTargetResource,
		})
	}
}

func sweepFileStorageMountTargetResource(compartment string) error {
	fileStorageClient := GetTestClients(&schema.ResourceData{}).fileStorageClient()
	mountTargetIds, err := getMountTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, mountTargetId := range mountTargetIds {
		if ok := SweeperDefaultResourceId[mountTargetId]; !ok {
			deleteMountTargetRequest := oci_file_storage.DeleteMountTargetRequest{}

			deleteMountTargetRequest.MountTargetId = &mountTargetId

			deleteMountTargetRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteMountTarget(context.Background(), deleteMountTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting MountTarget %s %s, It is possible that the resource is already deleted. Please verify manually \n", mountTargetId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &mountTargetId, mountTargetSweepWaitCondition, time.Duration(3*time.Minute),
				mountTargetSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getMountTargetIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "MountTargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := GetTestClients(&schema.ResourceData{}).fileStorageClient()

	listMountTargetsRequest := oci_file_storage.ListMountTargetsRequest{}
	listMountTargetsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for MountTarget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listMountTargetsRequest.AvailabilityDomain = &availabilityDomainName

		listMountTargetsRequest.LifecycleState = oci_file_storage.ListMountTargetsLifecycleStateActive
		listMountTargetsResponse, err := fileStorageClient.ListMountTargets(context.Background(), listMountTargetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting MountTarget list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, mountTarget := range listMountTargetsResponse.Items {
			id := *mountTarget.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "MountTargetId", id)
		}

	}
	return resourceIds, nil
}

func mountTargetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mountTargetResponse, ok := response.Response.(oci_file_storage.GetMountTargetResponse); ok {
		return mountTargetResponse.LifecycleState != oci_file_storage.MountTargetLifecycleStateDeleted
	}
	return false
}

func mountTargetSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.fileStorageClient().GetMountTarget(context.Background(), oci_file_storage.GetMountTargetRequest{
		MountTargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
