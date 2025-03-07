// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	acbDBName = RandomString(1, charsetWithoutDigits) + RandomString(13, charset)

	ACDatabaseResourceConfig = ACDatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, ACDatabaseRepresentation)

	ACDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           Representation{RepType: Required, Create: `${var.compartment_id}`},
		"autonomous_vm_cluster_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"display_name":             Representation{RepType: Optional, Create: `containerdatabases2`},
		"infrastructure_type":      Representation{RepType: Optional, Create: `CLOUD_AT_CUSTOMER`},
		"state":                    Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                   RepresentationGroup{Required, autonomousContainerDatabaseDataSourceFilterRepresentation},
	}

	ACDatabaseRepresentation = map[string]interface{}{
		"display_name":                 Representation{RepType: Required, Create: `containerdatabases2`},
		"patch_model":                  Representation{RepType: Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     Representation{RepType: Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                RepresentationGroup{Required, autonomousContainerDatabaseBackupConfigRepresentation},
		"key_store_id":                 Representation{RepType: Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"compartment_id":               Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               Representation{RepType: Optional, Create: acbDBName},
		"defined_tags":                 Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": Representation{RepType: Optional, Create: `STANDARD`},
	}

	ACDatabaseBackupConfigRepresentation = map[string]interface{}{
		"recovery_window_in_days": Representation{RepType: Optional, Create: `10`, Update: `11`},
	}

	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           Representation{RepType: Required, Create: `RECOVERY_APPLIANCE`},
		"id":             Representation{RepType: Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"internet_proxy": Representation{RepType: Optional, Create: `internetProxy`},
		"vpc_password":   Representation{RepType: Optional, Create: `vpcPassword`, Update: `vpcPassword2`},
		"vpc_user":       Representation{RepType: Optional, Create: `bkupUser1`},
	}

	acdBackupConfigLocalRepresentation = map[string]interface{}{
		"backup_destination_details": RepresentationGroup{Optional, map[string]interface{}{
			"type": Representation{RepType: Required, Create: `LOCAL`}}},
		"recovery_window_in_days": Representation{RepType: Optional, Create: `7`},
	}

	ACDatabaseResourceDependencies = AutonomousExadataInfrastructureResourceConfig +
		GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create,
			RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": Representation{RepType: Required, Create: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create,
			RepresentationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": Representation{RepType: Required, Create: "true"}})) +
		GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", Optional, Create, keyStoreRepresentation) + KmsVaultIdVariableStr + OkvSecretVariableStr

	dgDbUniqueName = RandomString(10, charsetWithoutDigits)
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabase_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabase_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ACDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
					GetUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + ACDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, ACDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", Optional, Create, ACDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ACDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ACDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ACDatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"maintenance_window_details",
				"backup_config.0.backup_destination_details.0.vpc_password",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabase_rotateDatabase(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabase_rotateDatabase")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals and rotate key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
					RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify no rotation of key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `false`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify rotate key
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update,
					RepresentationCopyWithNewProperties(autonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": Representation{RepType: Optional, Create: `true`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", Optional, Update, autonomousContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update, autonomousContainerDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Update,
					GetUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, autonomousContainerDatabaseRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousContainerDatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"rotate_key_trigger",
				"maintenance_window_details",
			},
			ResourceName: resourceName,
		},
	})
}
