// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	tenancySingularDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	TenancyResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityTenancyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTenancyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	singularDatasourceName := "data.oci_identity_tenancy.test_tenancy"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_tenancy", "test_tenancy", Required, Create, tenancySingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
			),
		},
	})
}
