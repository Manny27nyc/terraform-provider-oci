// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	operatorActionSingularDataSourceRepresentation = map[string]interface{}{
		"operator_action_id": Representation{RepType: Required, Create: `${data.oci_operator_access_control_operator_actions.test_operator_actions.operator_action_collection.0.items.0.id}`},
	}

	operatorActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `name`},
		"resource_type":  Representation{RepType: Optional, Create: `EXADATAINFRASTRUCTURE`},
	}

	OperatorActionResourceConfig = ""
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlOperatorActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOperatorAccessControlOperatorActionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_operator_actions.test_operator_actions"
	singularDatasourceName := "data.oci_operator_access_control_operator_action.test_operator_action"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_actions", "test_operator_actions", Required, Create, operatorActionDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorActionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "operator_action_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "operator_action_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_actions", "test_operator_actions", Required, Create, operatorActionDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_action", "test_operator_action", Required, Create, operatorActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorActionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "properties.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
			),
		},
	})
}
