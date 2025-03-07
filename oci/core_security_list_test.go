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
	SecurityListRequiredOnlyResource = SecurityListResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation)

	securityListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"vcn_id":         Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         RepresentationGroup{Required, securityListDataSourceFilterRepresentation}}
	securityListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
	}

	securityListRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"vcn_id":                 Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":           Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{RepType: Optional, Create: `MyPrivateSubnetSecurityList`, Update: `displayName2`},
		"egress_security_rules":  []RepresentationGroup{{Required, securityListEgressSecurityRulesICMPRepresentation}, {Optional, securityListEgressSecurityRulesTCPRepresentation}, {Optional, securityListEgressSecurityRulesUDPRepresentation}},
		"freeform_tags":          Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ingress_security_rules": []RepresentationGroup{{Required, securityListIngressSecurityRulesICMPRepresentation}, {Optional, securityListIngressSecurityRulesTCPRepresentation}, {Optional, securityListIngressSecurityRulesUDPRepresentation}},
	}
	securityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `1`},
		"description":      Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"icmp_options":     RepresentationGroup{Optional, securityListEgressSecurityRulesIcmpOptionsRepresentation},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	securityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `6`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options":      RepresentationGroup{Optional, securityListEgressSecurityRulesTcpOptionsRepresentation},
	}
	securityListEgressSecurityRulesUDPRepresentation = map[string]interface{}{
		"destination":      Representation{RepType: Required, Create: `10.0.2.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"protocol":         Representation{RepType: Required, Create: `17`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"udp_options":      RepresentationGroup{Optional, securityListEgressSecurityRulesUdpOptionsRepresentation},
	}
	securityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":     Representation{RepType: Required, Create: `1`},
		"description":  Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"source":       Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"icmp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesIcmpOptionsRepresentation},
		"source_type":  Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":    Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	securityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `6`},
		"source":      Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesTcpOptionsRepresentation},
	}
	securityListIngressSecurityRulesUDPRepresentation = map[string]interface{}{
		"protocol":    Representation{RepType: Required, Create: `17`},
		"source":      Representation{RepType: Required, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"udp_options": RepresentationGroup{Optional, securityListIngressSecurityRulesUdpOptionsRepresentation},
	}
	securityListEgressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
		"code": Representation{RepType: Optional, Create: `4`, Update: `0`},
	}
	securityListEgressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
		"code": Representation{RepType: Optional, Create: `4`, Update: `0`},
	}
	securityListIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityListIngressSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"max":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"min":               Representation{RepType: Optional, Create: `1521`, Update: `1522`},
		"source_port_range": RepresentationGroup{Optional, securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation},
	}
	securityListEgressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListEgressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityListIngressSecurityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}

	SecurityListResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreSecurityListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreSecurityListResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_security_list.test_security_list"
	datasourceName := "data.oci_core_security_lists.test_security_lists"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+SecurityListResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Optional, Create, securityListRepresentation), "core", "securityList", t)

	ResourceTest(t, testAccCheckCoreSecurityListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Required, Create, securityListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination": "10.0.2.0/24",
					"protocol":    "1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol": "1",
					"source":   "10.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Optional, Create, securityListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":         "10.0.2.0/24",
					"description":         "description",
					"destination_type":    "CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "false",
				},
					[]string{}),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":                           "10.0.2.0/24",
					"destination_type":                      "CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "false",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1521",
					"tcp_options.0.min":                     "1521",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1521",
					"tcp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination":                           "10.0.2.0/24",
					"destination_type":                      "CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "false",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1521",
					"udp_options.0.min":                     "1521",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1521",
					"udp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"description":         "description",
					"protocol":            "1",
					"source":              "10.0.1.0/24",
					"source_type":         "CIDR_BLOCK",
					"stateless":           "false",
				},
					[]string{}),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source":                                "10.0.1.0/24",
					"source_type":                           "CIDR_BLOCK",
					"stateless":                             "false",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1521",
					"tcp_options.0.min":                     "1521",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1521",
					"tcp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source":                                "10.0.1.0/24",
					"source_type":                           "CIDR_BLOCK",
					"stateless":                             "false",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1521",
					"udp_options.0.min":                     "1521",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1521",
					"udp_options.0.source_port_range.0.min": "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SecurityListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Optional, Create,
					RepresentationCopyWithNewProperties(securityListRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"description":         "description",
					"destination":         "10.0.2.0/24",
					"destination_type":    "CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "false",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"description":         "description",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "4",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source":              "10.0.1.0/24",
					"source_type":         "CIDR_BLOCK",
					"stateless":           "false",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + SecurityListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Optional, Update, securityListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"description":         "description2",
					"destination_type":    "SERVICE_CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "true",
				},
					[]string{
						"destination",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"description":         "description2",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source_type":         "SERVICE_CIDR_BLOCK",
					"stateless":           "true",
				},
					[]string{
						"source",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_core_security_lists", "test_security_lists", Optional, Update, securityListDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityListResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", Optional, Update, securityListRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "security_lists.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"description":         "description2",
					"destination_type":    "SERVICE_CIDR_BLOCK",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"stateless":           "true",
				},
					[]string{
						"destination",
					}),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "6",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
					"destination_type":                      "SERVICE_CIDR_BLOCK",
					"protocol":                              "17",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"destination",
					}),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.#", "3"),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"description":         "description2",
					"icmp_options.#":      "1",
					"icmp_options.0.code": "0",
					"icmp_options.0.type": "3",
					"protocol":            "1",
					"source_type":         "SERVICE_CIDR_BLOCK",
					"stateless":           "true",
				},
					[]string{
						"source",
					}),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"protocol":                              "6",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"tcp_options.#":                         "1",
					"tcp_options.0.max":                     "1522",
					"tcp_options.0.min":                     "1522",
					"tcp_options.0.source_port_range.#":     "1",
					"tcp_options.0.source_port_range.0.max": "1522",
					"tcp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
					"protocol":                              "17",
					"source_type":                           "SERVICE_CIDR_BLOCK",
					"stateless":                             "true",
					"udp_options.#":                         "1",
					"udp_options.0.max":                     "1522",
					"udp_options.0.min":                     "1522",
					"udp_options.0.source_port_range.#":     "1",
					"udp_options.0.source_port_range.0.max": "1522",
					"udp_options.0.source_port_range.0.min": "1522",
				},
					[]string{
						"source",
					}),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.vcn_id"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreSecurityListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_security_list" {
			noResourceFound = false
			request := oci_core.GetSecurityListRequest{}

			tmp := rs.Primary.ID
			request.SecurityListId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetSecurityList(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.SecurityListLifecycleStateTerminated): true,
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
	if !InSweeperExcludeList("CoreSecurityList") {
		resource.AddTestSweepers("CoreSecurityList", &resource.Sweeper{
			Name:         "CoreSecurityList",
			Dependencies: DependencyGraph["securityList"],
			F:            sweepCoreSecurityListResource,
		})
	}
}

func sweepCoreSecurityListResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	securityListIds, err := getSecurityListIds(compartment)
	if err != nil {
		return err
	}
	for _, securityListId := range securityListIds {
		if ok := SweeperDefaultResourceId[securityListId]; !ok {
			deleteSecurityListRequest := oci_core.DeleteSecurityListRequest{}

			deleteSecurityListRequest.SecurityListId = &securityListId

			deleteSecurityListRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteSecurityList(context.Background(), deleteSecurityListRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityList %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityListId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &securityListId, securityListSweepWaitCondition, time.Duration(3*time.Minute),
				securityListSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getSecurityListIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "SecurityListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listSecurityListsRequest := oci_core.ListSecurityListsRequest{}
	listSecurityListsRequest.CompartmentId = &compartmentId
	listSecurityListsRequest.LifecycleState = oci_core.SecurityListLifecycleStateAvailable
	listSecurityListsResponse, err := virtualNetworkClient.ListSecurityLists(context.Background(), listSecurityListsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityList list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityList := range listSecurityListsResponse.Items {
		id := *securityList.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityListId", id)
	}
	return resourceIds, nil
}

func securityListSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityListResponse, ok := response.Response.(oci_core.GetSecurityListResponse); ok {
		return securityListResponse.LifecycleState != oci_core.SecurityListLifecycleStateTerminated
	}
	return false
}

func securityListSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetSecurityList(context.Background(), oci_core.GetSecurityListRequest{
		SecurityListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
