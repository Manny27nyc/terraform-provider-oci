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
	networkSecurityGroupVnicDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
	}

	NetworkSecurityGroupVnicResourceConfig = VnicAttachmentResourceConfig
)

// issue-routing-tag: core/virtualNetwork
func TestCoreNetworkSecurityGroupVnicResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupVnicResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_network_security_group_vnics.test_network_security_group_vnics"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupVnicResourceConfig,
		},
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_network_security_group_vnics", "test_network_security_group_vnics", Required, Create, networkSecurityGroupVnicDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkSecurityGroupVnicResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.time_associated"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_vnics.0.vnic_id"),
			),
		},
	})
}
