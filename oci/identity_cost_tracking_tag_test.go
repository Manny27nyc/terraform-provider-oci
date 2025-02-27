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
	costTrackingTagDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	CostTrackingTagResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityCostTrackingTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCostTrackingTagResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_cost_tracking_tags.test_cost_tracking_tags"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_cost_tracking_tags", "test_cost_tracking_tags", Required, Create, costTrackingTagDataSourceRepresentation) +
				compartmentIdVariableStr + CostTrackingTagResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "tags.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_cost_tracking"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_retired"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
			),
		},
	})
}
