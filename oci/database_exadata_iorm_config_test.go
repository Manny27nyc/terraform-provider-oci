// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataIormConfigRequiredOnlyResource = ExadataIormConfigResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Required, Create, exadataIormConfigRepresentation)

	ExadataIormConfigResourceConfig = ExadataIormConfigResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Optional, Update, exadataIormConfigRepresentation)

	exadataIormConfigSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{RepType: Required, Create: `${oci_database_db_system.t.id}`},
	}

	exadataIormConfigRepresentation = map[string]interface{}{
		"db_system_id": Representation{RepType: Required, Create: `${oci_database_db_system.t.id}`},
		"objective":    Representation{RepType: Optional, Create: `AUTO`, Update: `BALANCED`},
		"db_plans":     RepresentationGroup{Required, dbPlanRepresentation},
	}

	dbPlanRepresentation = map[string]interface{}{
		"db_name": Representation{RepType: Required, Create: `default`, Update: `default`},
		"share":   Representation{RepType: Required, Create: `1`, Update: `2`},
	}

	ExadataIormConfigResourceDependencies = DefinedTagsDependencies + `

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}
	data "oci_identity_availability_domain" "ad" {
		compartment_id 		= "${var.compartment_id}"
		ad_number      		= 3
	}
	resource "oci_core_subnet" "exadata_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.22.0/24"
		display_name        = "ExadataSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}", "${oci_core_security_list.exadata_shapes_security_list.id}"]
		dns_label           = "subnetexadata1"
	}
	resource "oci_core_subnet" "exadata_backup_subnet" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		cidr_block          = "10.1.23.0/24"
		display_name        = "ExadataBackupSubnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "subnetexadata2"
	}

	resource "oci_core_security_list" "exadata_shapes_security_list" {
		compartment_id = "${var.compartment_id}"
		vcn_id         = "${oci_core_virtual_network.t.id}"
		display_name   = "ExadataSecurityList"

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "6"
		}

		ingress_security_rules {
			source    = "10.1.22.0/24"
			protocol  = "1"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "6"
		}

		egress_security_rules {
			destination = "10.1.22.0/24"
			protocol    = "1"
		}
	}

	resource "oci_database_db_system" "t" {
		availability_domain = "${data.oci_identity_availability_domain.ad.name}"
		compartment_id = "${var.compartment_id}"
		subnet_id = "${oci_core_subnet.exadata_subnet.id}"
		backup_subnet_id = "${oci_core_subnet.exadata_backup_subnet.id}"
		database_edition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
		disk_redundancy = "HIGH"
		shape = "Exadata.Quarter1.84"
		cpu_core_count = "22"
		ssh_public_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV"]
		domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
		hostname = "myOracleDB"
		data_storage_size_in_gb = "256"
		license_model = "LICENSE_INCLUDED"
		node_count = "1"
		time_zone = "US/Pacific"
		db_home {
			db_version = "12.1.0.2"
			database {
				admin_password = "BEstrO0ng_#11"
				db_name = "aTFdb"
			}
		}
	}
	`
)

// issue-routing-tag: database/default
func TestDatabaseExadataIormConfigResource_basic(t *testing.T) {
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_Exadata") {
		t.Skip("Skipping suppressed DBSystem_Exadata")
	}

	httpreplay.SetScenario("TestDatabaseExadataIormConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadata_iorm_config.test_exadata_iorm_config"

	singularDatasourceName := "data.oci_database_exadata_iorm_config.test_exadata_iorm_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ExadataIormConfigResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Optional, Create, exadataIormConfigRepresentation), "database", "exadataIormConfig", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExadataIormConfigResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Required, Create, exadataIormConfigRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "db_plans.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExadataIormConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExadataIormConfigResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Optional, Create, exadataIormConfigRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "objective", "AUTO"),

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
			Config: config + compartmentIdVariableStr + ExadataIormConfigResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Optional, Update, exadataIormConfigRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "objective", "BALANCED"),
				resource.TestCheckResourceAttr(resourceName, "db_plans.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_exadata_iorm_config", "test_exadata_iorm_config", Required, Create, exadataIormConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExadataIormConfigResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "db_plans.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "objective", "BALANCED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
