// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	DataSafeConfigurationRequiredOnlyResource = DataSafeConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Required, Create, dataSafeConfigurationRepresentation)

	DataSafeConfigurationResourceConfig = DataSafeConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Update, dataSafeConfigurationRepresentation)

	dataSafeConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
	}

	dataSafeConfigurationRepresentation = map[string]interface{}{
		"is_enabled":     Representation{RepType: Required, Create: `true`},
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
	}

	DataSafeConfigurationResourceDependencies = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeDataSafeConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDataSafeConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_data_safe_configuration.test_data_safe_configuration"

	singularDatasourceName := "data.oci_data_safe_data_safe_configuration.test_data_safe_configuration"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DataSafeConfigurationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Create, dataSafeConfigurationRepresentation), "datasafe", "dataSafeConfiguration", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Required, Create, dataSafeConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Create, dataSafeConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),

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
			Config: config + compartmentIdVariableStr + DataSafeConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Update, dataSafeConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),

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
				GenerateDataSourceFromRepresentationMap("oci_data_safe_data_safe_configuration", "test_data_safe_configuration", Optional, Create, dataSafeConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeConfigurationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "url"),
			),
		},
	})
}
