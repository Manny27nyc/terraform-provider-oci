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
	protectionCapabilityGroupTagDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `Java`},
		"type":           Representation{RepType: Optional, Create: `REQUEST_PROTECTION_CAPABILITY`},
	}

	ProtectionCapabilityGroupTagResourceConfig = ""
)

// issue-routing-tag: waf/default
func TestWafProtectionCapabilityGroupTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafProtectionCapabilityGroupTagResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_waf_protection_capability_group_tags.test_protection_capability_group_tags"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waf_protection_capability_group_tags", "test_protection_capability_group_tags", Optional, Create, protectionCapabilityGroupTagDataSourceRepresentation) +
				compartmentIdVariableStr + ProtectionCapabilityGroupTagResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "Java"),
				resource.TestCheckResourceAttr(datasourceName, "type", "REQUEST_PROTECTION_CAPABILITY"),

				resource.TestCheckResourceAttrSet(datasourceName, "protection_capability_group_tag_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "protection_capability_group_tag_collection.0.items.#", "1"),
			),
		},
	})
}
