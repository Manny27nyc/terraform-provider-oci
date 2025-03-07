// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	DatabaseSoftwareImageResourceConfigForExaccShape = DatabaseSoftwareImageResourceDependenciesForExaShape +
		GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape)

	DatabaseSoftwareImageResourceConfigForExadataShape = DatabaseSoftwareImageResourceDependenciesForExaShape +
		GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExadataShape)

	databaseSoftwareImageSingularDataSourceRepresentationForExaccShape = map[string]interface{}{
		"database_software_image_id": Representation{RepType: Required, Create: `${oci_database_database_software_image.test_database_software_image.id}`},
	}

	databaseSoftwareImageDataSourceRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":       Representation{RepType: Optional, Create: `DB_Image_V19`},
		"image_shape_family": Representation{RepType: Optional, Create: `EXACC_SHAPE`},
		"image_type":         Representation{RepType: Optional, Create: `DATABASE_IMAGE`},
		"state":              Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":             RepresentationGroup{Required, databaseSoftwareImageDataSourceFilterRepresentationForExaccShape}}
	databaseSoftwareImageDataSourceFilterRepresentationForExaccShape = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_database_software_image.test_database_software_image.id}`}},
	}

	databaseSoftwareImageDataSourceRepresentationForExadataShape = GetUpdatedRepresentationCopy("image_shape_family",
		Representation{RepType: Required, Create: `EXADATA_SHAPE`},
		databaseSoftwareImageDataSourceRepresentationForExaccShape)

	databaseSoftwareImageDataSourceRepresentationForVmBmShape = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":       Representation{RepType: Required, Create: `VMBM_DB_Image_V19`},
		"database_version":   Representation{RepType: Required, Create: `19.0.0.0`},
		"patch_set":          Representation{RepType: Required, Create: `19.7.0.0`},
		"image_shape_family": Representation{RepType: Optional, Create: `VMBM_SHAPE`},
		"image_type":         Representation{RepType: Optional, Create: `DATABASE_IMAGE`},
	}

	databaseSoftwareImageRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Create: `${var.compartment_id}`},
		"database_version":   Representation{RepType: Required, Create: `19.0.0.0`},
		"display_name":       Representation{RepType: Required, Create: `DB_Image_V19`, Update: `DB_Image_V19_U1`},
		"patch_set":          Representation{RepType: Required, Create: `19.7.0.0`},
		"image_shape_family": Representation{RepType: Optional, Create: `EXACC_SHAPE`},
		"database_software_image_one_off_patches": Representation{RepType: Optional, Create: []string{"29910218", "31113249"}},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"Department": "Exacc_Finance"}, Update: map[string]string{"Department": "Exacc_Accounting"}},

		"image_type": Representation{RepType: Optional, Create: `DATABASE_IMAGE`},
	}

	databaseSoftwareImageRepresentationForExadataShape = GetUpdatedRepresentationCopy("image_shape_family",
		Representation{RepType: Required, Create: `EXADATA_SHAPE`},
		databaseSoftwareImageRepresentationForExaccShape)

	DatabaseSoftwareImageResourceDependenciesForExaShape = DefinedTagsDependencies

	DatabaseSoftwareImageResourceDependenciesForSourceDbHome = GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentation) +
		ExaBaseDependencies + AvailabilityDomainConfig +
		DatabaseSoftwareImageResourceDependenciesForExaShape

	databaseSoftwareImageRepresentationForSourceDbHome = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"database_version": Representation{RepType: Optional, Create: `12.1.0.2`},
		"display_name":     Representation{RepType: Required, Create: `image1`, Update: `displayName2`},
		"patch_set":        Representation{RepType: Optional, Create: `12.1.0.2.210119`},
		"database_software_image_one_off_patches": Representation{RepType: Optional, Create: []string{"31113249", "27929509"}},
		"defined_tags":       Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_shape_family": Representation{RepType: Optional, Create: `EXADATA_SHAPE`},
		"image_type":         Representation{RepType: Optional, Create: `DATABASE_IMAGE`},
		"ls_inventory":       Representation{RepType: Optional, Create: nil},
		"source_db_home_id":  Representation{RepType: Required, Create: `${oci_database_db_home.test_db_home.id}`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseSoftwareImageResourceForExaccShape(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResourceForExaccShape") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	}

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	ResourceTest(t, testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape, []resource.TestStep{

		// verify creation of Database Software Image with EXACC_SHAPE
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageRepresentationForExaccShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
		// verify Update of Database Software Image created for shape EXACC_SHAPE
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", Optional, Update, databaseSoftwareImageDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.is_upgrade_supported"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageSingularDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExaccShape,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database_software_image" {
			noResourceFound = false
			request := oci_database.GetDatabaseSoftwareImageRequest{}

			tmp := rs.Primary.ID
			request.DatabaseSoftwareImageId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetDatabaseSoftwareImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseSoftwareImageLifecycleStateDeleted): true, string(oci_database.DatabaseSoftwareImageLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("DatabaseDatabaseSoftwareImageForExaccShape") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImageForExaccShape", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImageForExaccShape",
			Dependencies: DependencyGraph["databaseSoftwareImageForExaccShape"],
			F:            sweepDatabaseDatabaseSoftwareImageResourceForExaccShape,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResourceForExaccShape(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIdsForExaccShape(compartment)
	if err != nil {
		return err
	}
	for _, databaseSoftwareImageId := range databaseSoftwareImageIds {
		if ok := SweeperDefaultResourceId[databaseSoftwareImageId]; !ok {
			deleteDatabaseSoftwareImageRequest := oci_database.DeleteDatabaseSoftwareImageRequest{}

			deleteDatabaseSoftwareImageRequest.DatabaseSoftwareImageId = &databaseSoftwareImageId

			deleteDatabaseSoftwareImageRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabaseSoftwareImage(context.Background(), deleteDatabaseSoftwareImageRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseSoftwareImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseSoftwareImageId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitConditionForExacc, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationForExacc, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIdsForExaccShape(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DatabaseSoftwareImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listDatabaseSoftwareImagesRequest := oci_database.ListDatabaseSoftwareImagesRequest{}
	listDatabaseSoftwareImagesRequest.CompartmentId = &compartmentId
	listDatabaseSoftwareImagesRequest.LifecycleState = oci_database.DatabaseSoftwareImageSummaryLifecycleStateAvailable
	listDatabaseSoftwareImagesRequest.ImageShapeFamily = oci_database.DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape
	listDatabaseSoftwareImagesResponse, err := databaseClient.ListDatabaseSoftwareImages(context.Background(), listDatabaseSoftwareImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseSoftwareImage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseSoftwareImage := range listDatabaseSoftwareImagesResponse.Items {
		id := *databaseSoftwareImage.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseSoftwareImageId", id)
	}
	return resourceIds, nil
}

func databaseSoftwareImageSweepWaitConditionForExacc(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperationForExacc(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// Exadata test function

// issue-routing-tag: database/default
func TestDatabaseDatabaseSoftwareImageResourceExadata_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResourceExadata_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResourceExadata_basic") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResourceExadata_basic")
	}

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DatabaseSoftwareImageResourceDependenciesForExaShape+
		GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageRepresentationForExadataShape), "database", "databaseSoftwareImage", t)

	ResourceTest(t, testAccCheckDatabaseDatabaseSoftwareImageDestroyExa, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageRepresentationForExadataShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape,
		},

		//verify Create with source_db_home
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForSourceDbHome +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageRepresentationForSourceDbHome),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_home_id"),
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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageRepresentationForExadataShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create,
					RepresentationCopyWithNewProperties(databaseSoftwareImageRepresentationForExadataShape, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
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
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExadataShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
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
			PreConfig: WaitTillCondition(testAccProvider, &resId, databaseSoftwareImageWaitTillAvailableConditionExa, time.Duration(20*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationExa, "database", true),
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", Optional, Update, databaseSoftwareImageDataSourceRepresentationForExadataShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExadataShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.is_upgrade_supported"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageSingularDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExadataShape,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExadataShape,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{
				//	"source_db_home_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroyExa(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database_software_image" {
			noResourceFound = false
			request := oci_database.GetDatabaseSoftwareImageRequest{}

			tmp := rs.Primary.ID
			request.DatabaseSoftwareImageId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

			response, err := client.GetDatabaseSoftwareImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseSoftwareImageLifecycleStateDeleted): true, string(oci_database.DatabaseSoftwareImageLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("DatabaseDatabaseSoftwareImageExa") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImageExa", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImageExa",
			Dependencies: DependencyGraph["databaseSoftwareImageExa"],
			F:            sweepDatabaseDatabaseSoftwareImageResourceExa,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResourceExa(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIdsExa(compartment)
	if err != nil {
		return err
	}
	for _, databaseSoftwareImageId := range databaseSoftwareImageIds {
		if ok := SweeperDefaultResourceId[databaseSoftwareImageId]; !ok {
			deleteDatabaseSoftwareImageRequest := oci_database.DeleteDatabaseSoftwareImageRequest{}

			deleteDatabaseSoftwareImageRequest.DatabaseSoftwareImageId = &databaseSoftwareImageId

			deleteDatabaseSoftwareImageRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabaseSoftwareImage(context.Background(), deleteDatabaseSoftwareImageRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseSoftwareImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseSoftwareImageId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitConditionExa, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationExa, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIdsExa(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DatabaseSoftwareImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listDatabaseSoftwareImagesRequest := oci_database.ListDatabaseSoftwareImagesRequest{}
	listDatabaseSoftwareImagesRequest.CompartmentId = &compartmentId
	listDatabaseSoftwareImagesRequest.LifecycleState = oci_database.DatabaseSoftwareImageSummaryLifecycleStateAvailable
	listDatabaseSoftwareImagesResponse, err := databaseClient.ListDatabaseSoftwareImages(context.Background(), listDatabaseSoftwareImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseSoftwareImage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseSoftwareImage := range listDatabaseSoftwareImagesResponse.Items {
		id := *databaseSoftwareImage.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseSoftwareImageId", id)
	}
	return resourceIds, nil
}

func databaseSoftwareImageSweepWaitConditionExa(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageWaitTillAvailableConditionExa(response common.OCIOperationResponse) bool {
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		fmt.Print("Checking whether the state of resource is Available: ", databaseSoftwareImageResponse.LifecycleState == oci_database.DatabaseSoftwareImageLifecycleStateAvailable, "\n")
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateAvailable)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperationExa(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
