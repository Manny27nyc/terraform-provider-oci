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
	apiContentSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": Representation{RepType: Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApiContentResourceConfig = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create, apiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiContentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_content.test_api_content"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_apigateway_api_content", "test_api_content", Required, Create, apiContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiContentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),
			),
		},
	})
}
