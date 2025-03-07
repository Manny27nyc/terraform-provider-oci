// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VirtualCircuitRequiredOnlyResource = VirtualCircuitResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitRequiredOnlyRepresentation)

	virtualCircuitSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_circuit_id": Representation{RepType: Required, Create: `${oci_core_virtual_circuit.test_virtual_circuit.id}`},
	}

	virtualCircuitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `PROVISIONED`},
		"filter":         RepresentationGroup{Required, virtualCircuitDataSourceFilterRepresentation}}
	virtualCircuitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit.id}`}},
	}

	virtualCircuitPublicRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":                   Representation{RepType: Required, Create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsPublicRequiredOnlyRepresentation},
		"customer_bgp_asn":       Representation{RepType: Required, Create: `10`, Update: `11`},
		"public_prefixes":        RepresentationGroup{Required, virtualCircuitPublicPrefixesRepresentation},
	}
	virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation = RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(virtualCircuitPublicRequiredOnlyRepresentation, []string{"customer_bgp_asn"}), map[string]interface{}{
			"customer_asn": Representation{RepType: Required, Create: `10`, Update: `11`},
		})

	virtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":                   Representation{RepType: Required, Create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           Representation{RepType: Required, Create: `10`, Update: `11`},
		"gateway_id":             Representation{RepType: Required, Create: `${oci_core_drg.test_drg.id}`},
	}
	virtualCircuitRepresentation = map[string]interface{}{
		"ip_mtu":                 Representation{RepType: Optional, Create: `MTU_1500`, Update: `MTU_9000`},
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":                   Representation{RepType: Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{RepType: Optional, Create: `10 Gbps`, Update: `20 Gbps`},
		"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           Representation{RepType: Required, Create: `10`, Update: `11`},
		"defined_tags":           Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"gateway_id":             Representation{RepType: Optional, Create: `${oci_core_drg.test_drg.id}`},
		"region":                 Representation{RepType: Optional, Create: `us-phoenix-1`},
		"routing_policy":         Representation{RepType: Optional, Create: []string{`REGIONAL`}, Update: []string{`GLOBAL`}},
	}

	virtualCircuitWithProviderRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"type":                   Representation{RepType: Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   Representation{RepType: Optional, Create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes.0.name}"},
		"cross_connect_mappings": RepresentationGroup{Required, virtualCircuitCrossConnectMappingsRepresentation},
		"customer_asn":           Representation{RepType: Required, Create: `10`, Update: `11`},
		"display_name":           Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":             Representation{RepType: Optional, Create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    Representation{RepType: Optional, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		// provider_service_key_name can only be updated by a Fast Connect Service Provider
		// "provider_service_key_name": Representation{RepType: Optional, Create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
		"region": Representation{RepType: Optional, Create: `us-phoenix-1`},
	}

	crossConnectMappingsPublicRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": Representation{RepType: Required, Create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"vlan": Representation{RepType: Required, Create: `200`, Update: `300`},
	}
	crossConnectMappingsRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": Representation{RepType: Required, Create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"customer_bgp_peering_ip":                 Representation{RepType: Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   Representation{RepType: Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
		"vlan":                                    Representation{RepType: Required, Create: `200`, Update: `300`},
	}
	virtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"customer_bgp_peering_ip": Representation{RepType: Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":   Representation{RepType: Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
	}
	virtualCircuitPublicPrefixesRepresentation = map[string]interface{}{
		"cidr_block": Representation{RepType: Required, Create: `11.0.0.0/24`, Update: `11.0.1.0/24`},
	}

	VirtualCircuitWithProviderResourceConfigFilter = `
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"

	filter {
		name = "type"
		values = [ "LAYER2" ]
	}

	filter {
		name = "private_peering_bgp_management"
		values = [ "CUSTOMER_MANAGED" ]
	}

	filter {
		name = "supported_virtual_circuit_types"
		values = [ "${var.virtual_circuit_type}" ]
	}

	filter {
		name = "public_peering_bgp_management"
		values = [ "ORACLE_MANAGED" ]
	}

	filter {
		name = "provider_service_key_management"
		values = ["PROVIDER_MANAGED"]
	}
}

data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes" {
  #Required
  provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}"
}
`

	VirtualCircuitPrivatePropertyVariables = `
variable "virtual_circuit_type" { default = "PRIVATE" }

`

	VirtualCircuitPublicPropertyVariables = `
variable "virtual_circuit_type" { default = "PUBLIC" }

`
	VirtualCircuitResourceDependencies          = DrgRequiredOnlyResource + CrossConnectWithGroupResourceConfig
	VirtualCircuitResourceDependenciesCopyForVC = DrgRequiredOnlyResource + CrossConnectWithGroupResourceConfigCopyForVC
)

// issue-routing-tag: core/default
func TestCoreVirtualCircuitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretIdCKN := getEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := getEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := getEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := getEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VirtualCircuitResourceDependencies+VirtualCircuitPublicPropertyVariables+secretIdVariableStrCKN+secretIdVariableStrCAK+secretVersionStrCAK+secretVersionStrCKN+
		GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation), "core", "virtualCircuit", t)

	ResourceTest(t, testAccCheckCoreVirtualCircuitDestroy, []resource.TestStep{
		// verify Create - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create,
					RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(virtualCircuitRepresentation, []string{"gateway_id", "cross_connect_mappings", "customer_asn"}),
						map[string]interface{}{
							"cross_connect_mappings": RepresentationGroup{Required, crossConnectMappingsPublicRequiredOnlyRepresentation},
							"customer_bgp_asn":       Representation{RepType: Required, Create: `10`, Update: `11`},
						})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update from customer_bgp_asn to customer_asn
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Update, virtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create - PRIVATE Virtual Circuit with Provider
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitWithProviderRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update - PRIVATE Virtual Circuit with Provider
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitWithProviderRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},

		// verify Create - PRIVATE Virtual Circuit
		{
			Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + VirtualCircuitRequiredOnlyResource + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + VirtualCircuitWithProviderResourceConfigFilter +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create, virtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_1500"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Create,
					RepresentationCopyWithNewProperties(virtualCircuitRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_1500"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuits", "test_virtual_circuits", Optional, Update, virtualCircuitDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_ipv6session_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_management"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_session_state"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.customer_asn", "11"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.gateway_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.oracle_bgp_asn"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.routing_policy.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.service_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.state", "PROVISIONED"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.type", "PRIVATE"),
			),
		},
		// verify singular datasource
		{
			Config: config + GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Required, Create, virtualCircuitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_ipv6session_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_management"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_session_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_bgp_asn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oracle_bgp_asn", "31898"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "COLOCATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PRIVATE"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependencies + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", Optional, Update, virtualCircuitRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"region",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreVirtualCircuitDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_virtual_circuit" {
			noResourceFound = false
			request := oci_core.GetVirtualCircuitRequest{}

			tmp := rs.Primary.ID
			request.VirtualCircuitId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetVirtualCircuit(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VirtualCircuitLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("CoreVirtualCircuit") {
		resource.AddTestSweepers("CoreVirtualCircuit", &resource.Sweeper{
			Name:         "CoreVirtualCircuit",
			Dependencies: DependencyGraph["virtualCircuit"],
			F:            sweepCoreVirtualCircuitResource,
		})
	}
}

func sweepCoreVirtualCircuitResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	virtualCircuitIds, err := getVirtualCircuitIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualCircuitId := range virtualCircuitIds {
		if ok := SweeperDefaultResourceId[virtualCircuitId]; !ok {
			deleteVirtualCircuitRequest := oci_core.DeleteVirtualCircuitRequest{}

			deleteVirtualCircuitRequest.VirtualCircuitId = &virtualCircuitId

			deleteVirtualCircuitRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVirtualCircuit(context.Background(), deleteVirtualCircuitRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualCircuit %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualCircuitId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &virtualCircuitId, virtualCircuitSweepWaitCondition, time.Duration(3*time.Minute),
				virtualCircuitSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVirtualCircuitIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "VirtualCircuitId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listVirtualCircuitsRequest := oci_core.ListVirtualCircuitsRequest{}
	listVirtualCircuitsRequest.CompartmentId = &compartmentId
	listVirtualCircuitsRequest.LifecycleState = oci_core.VirtualCircuitLifecycleStateProvisioned
	listVirtualCircuitsResponse, err := virtualNetworkClient.ListVirtualCircuits(context.Background(), listVirtualCircuitsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualCircuit list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualCircuit := range listVirtualCircuitsResponse.Items {
		id := *virtualCircuit.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "VirtualCircuitId", id)
	}
	return resourceIds, nil
}

func virtualCircuitSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualCircuitResponse, ok := response.Response.(oci_core.GetVirtualCircuitResponse); ok {
		return virtualCircuitResponse.LifecycleState != oci_core.VirtualCircuitLifecycleStateTerminated
	}
	return false
}

func virtualCircuitSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetVirtualCircuit(context.Background(), oci_core.GetVirtualCircuitRequest{
		VirtualCircuitId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
