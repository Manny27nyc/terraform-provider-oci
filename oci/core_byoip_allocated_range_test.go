// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	byoipAllocatedRangeDataSourceRepresentation = map[string]interface{}{
		"byoip_range_id": Representation{RepType: Required, Create: `${var.byoip_range_id}`},
	}

	ByoipAllocatedRangeResourceConfig = PublicIpPoolAddCapacityResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation)
)

// issue-routing-tag: core/vcnip
func TestCoreByoipAllocatedRangeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreByoipAllocatedRangeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_byoip_allocated_ranges.test_byoip_allocated_ranges"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + ByoipAllocatedRangeResourceConfig,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for oci_core_public_ip_pool and oci_core_public_ip_pool_capacity resource to get created first")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_byoip_allocated_ranges", "test_byoip_allocated_ranges", Required, Create, byoipAllocatedRangeDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipAllocatedRangeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "byoip_allocated_range_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "byoip_allocated_range_collection.0.items.#", "1"),
			),
		},
	})
}
