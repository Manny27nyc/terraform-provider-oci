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
	ResolverRequiredOnlyResource = ResolverResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation)

	ResolverResourceConfig = ResolverResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Optional, Update, resolverRepresentation)

	resolverSingularDataSourceRepresentation = map[string]interface{}{
		"resolver_id": Representation{RepType: Required, Create: `${oci_dns_resolver.test_resolver.id}`},
		"scope":       Representation{RepType: Required, Create: `PRIVATE`},
	}

	resolverDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`},
		"id":             Representation{RepType: Optional, Create: `${oci_dns_resolver.test_resolver.id}`},
		"scope":          Representation{RepType: Required, Create: `PRIVATE`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, resolverDataSourceFilterRepresentation}}

	resolverDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_dns_resolver.test_resolver.id}`}},
	}

	resolverRepresentation = map[string]interface{}{
		"resolver_id":    Representation{RepType: Required, Create: `${data.oci_core_vcn_dns_resolver_association.test_vcn_dns_resolver_association.dns_resolver_id}`},
		"attached_views": RepresentationGroup{Optional, resolverAttachedViewsRepresentation},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          Representation{RepType: Required, Create: `PRIVATE`},
	}
	resolverAttachedViewsRepresentation = map[string]interface{}{
		"view_id": Representation{RepType: Required, Create: `${oci_dns_view.test_view.id}`},
	}

	resolverRepresentationRules = RepresentationCopyWithNewProperties(resolverRepresentation, map[string]interface{}{
		"rules": []RepresentationGroup{{Optional, resolverRulesItemsRepresentationClientAddr}, {Optional, resolverRulesItemsRepresentationQname}},
	})

	resolverRulesItemsRepresentationClientAddr = map[string]interface{}{
		"action":                    Representation{RepType: Required, Create: `FORWARD`},
		"destination_addresses":     Representation{RepType: Required, Create: []string{`10.0.0.11`}, Update: []string{`10.0.0.12`}},
		"source_endpoint_name":      Representation{RepType: Required, Create: `endpointName`},
		"client_address_conditions": Representation{RepType: Optional, Create: []string{`192.0.20.0/24`}, Update: []string{`192.0.21.0/24`}},
		"qname_cover_conditions":    Representation{RepType: Optional, Update: []string{}},
	}
	resolverRulesItemsRepresentationQname = map[string]interface{}{
		"action":                    Representation{RepType: Required, Create: `FORWARD`},
		"destination_addresses":     Representation{RepType: Required, Create: []string{`10.0.0.11`}, Update: []string{`10.0.0.12`}},
		"source_endpoint_name":      Representation{RepType: Required, Create: `endpointName`},
		"client_address_conditions": Representation{RepType: Optional, Create: []string{}},
		"qname_cover_conditions":    Representation{RepType: Optional, Create: []string{`internal.example.com`}, Update: []string{`internal2.example.com`}},
	}

	ResolverResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", Required, Create, viewRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsResolverResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsResolverResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_resolver.test_resolver"
	datasourceName := "data.oci_dns_resolvers.test_resolvers"
	singularDatasourceName := "data.oci_dns_resolver.test_resolver"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ResolverResourceDependencies+
		GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation)+
		GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation), "dns", "resolver", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + ResolverResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for 2 minutes for oci_core_vcn resource to get created")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ResolverResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResolverResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation),
		},
		// Create resolver with endpoint
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ResolverResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Optional, Create,
					RepresentationCopyWithNewProperties(resolverRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Create, resolverEndpointRepresentationWithoutNsgId),
		},
		// verify Create with optionals and resolver rules
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ResolverResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Optional, Create,
					RepresentationCopyWithNewProperties(resolverRepresentationRules, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Create, resolverEndpointRepresentationWithoutNsgId),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attached_views.0.view_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.0", "192.0.20.0/24"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.0", "10.0.0.11"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.qname_cover_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source_endpoint_name", "endpointName"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.client_address_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.0", "10.0.0.11"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.0", "internal.example.com"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.source_endpoint_name", "endpointName"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					// Resource discovery is disabled for Resolver
					//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify updates to updatable parameters and add resolver rules
		{
			Config: config + compartmentIdVariableStr + ResolverResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Optional, Update,
					RepresentationCopyWithNewProperties(resolverRepresentationRules, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
					})) +
				GenerateResourceFromRepresentationMap("oci_dns_resolver_endpoint", "test_resolver_endpoint", Optional, Create, resolverEndpointRepresentationWithoutNsgId),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "attached_views.0.view_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(resourceName, "resolver_id"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.forwarding_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_forwarding", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.is_listening", "false"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.client_address_conditions.0", "192.0.21.0/24"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination_addresses.0", "10.0.0.12"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.qname_cover_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source_endpoint_name", "endpointName"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.action", "FORWARD"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.client_address_conditions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.destination_addresses.0", "10.0.0.12"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.qname_cover_conditions.0", "internal2.example.com"),
				resource.TestCheckResourceAttr(resourceName, "rules.1.source_endpoint_name", "endpointName"),

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
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_dns_resolvers", "test_resolvers", Optional, Update, resolverDataSourceRepresentation) +
				compartmentIdVariableStr + ResolverResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Optional, Update, resolverRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.attached_vcn_id"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.default_view_id"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "resolvers.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.is_protected"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "resolvers.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_dns_resolver", "test_resolver", Required, Create, resolverSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResolverResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resolver_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "attached_vcn_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attached_views.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_view_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ResolverResourceConfig + GenerateDataSourceFromRepresentationMap("oci_core_vcn_dns_resolver_association", "test_vcn_dns_resolver_association", Required, Create, vcnDnsResolverAssociationSingularDataSourceRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getDnsResolverImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"scope",
			},
			ResourceName: resourceName,
		},
	})
}

func getDnsResolverImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("resolverId/" + rs.Primary.Attributes["id"] + "/scope/" + rs.Primary.Attributes["scope"]), nil
	}
}
