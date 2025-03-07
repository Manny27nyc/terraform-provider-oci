// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkLoadBalancersPolicyDataSourceRepresentation = map[string]interface{}{}

	NetworkLoadBalancersPolicyResourceConfig = ""
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancersPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancersPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	datasourceName := "data.oci_network_load_balancer_network_load_balancers_policies.test_network_load_balancers_policies"

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_policies", "test_network_load_balancers_policies", Required, Create, networkLoadBalancersPolicyDataSourceRepresentation) +
				NetworkLoadBalancersPolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancers_policy_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "network_load_balancers_policy_collection.0.items.#", "3"),
			),
		},
	})
}
