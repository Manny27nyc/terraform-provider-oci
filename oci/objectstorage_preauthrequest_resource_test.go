// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"regexp"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstoragePARTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceObjectstoragePARTestSuite) SetupTest() {
	s.Token, s.TokenFn = TokenizeWithHttpReplay("object_storage_resource")
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	data "oci_objectstorage_namespace" "t" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "{{.token}}"
		access_type="ObjectRead"
	}

	resource "oci_objectstorage_object" "t" {
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		bucket = "${oci_objectstorage_bucket.t.name}"
		object = "-tf-object"
		content = "123"
	}`, nil)

	s.ResourceName = "oci_objectstorage_preauthrequest.t"
}

func (s *ResourceObjectstoragePARTestSuite) TestAccResourceObjectstoragePAR_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config + `
				resource "oci_objectstorage_preauthrequest" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					name = "-tf-par"
					access_type = "ObjectRead"
					time_expires = "` + expirationTimeForPar.Format(time.RFC3339Nano) + `"
					object = "-tf-object"
				}`,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(s.ResourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "access_uri"),
					// regex match example: /p/QJ1Geyhs3WKZvJr8jhw0TeqqqKd4OE1i9ZsGcJ5bzi8/n/internalbriangustafson/b/2018-02-05-130953-145201650/o/
					resource.TestMatchResourceAttr(s.ResourceName, "access_uri", regexp.MustCompile("/p/.*/n/.*/b/"+s.Token+"/o/")),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
			// verify access_uri is still available after subsequent refreshes (api only returns this value on Create)
			{
				Config: s.Config + `
				resource "oci_objectstorage_preauthrequest" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					name = "-tf-par"
					access_type = "ObjectRead"
					time_expires = "` + expirationTimeForPar.Format(time.RFC3339Nano) + `"
					object = "-tf-object"
				}`,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestMatchResourceAttr(s.ResourceName, "access_uri", regexp.MustCompile("/p/.*/n/.*/b/"+s.Token+"/o/")),
				),
			},
		},
	})
}

// issue-routing-tag: terraform/default
func TestUnitResourceObjectstoragePAR_parseIds(t *testing.T) {
	t.Run("Parse Composite Ids", func(t *testing.T) {
		tests := []struct {
			parId       string
			expectError bool
			error       string
			parsedParId string
		}{
			{`n/dxterraformdev/b/bucket/p/dJoeW0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJoeW0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W/0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W/0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W/0iJzmjVX4x6rAKn/UUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W/0iJzmjVX4x6rAKn/UUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0iJzmj/n/VX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W0iJzmj/n/VX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAY%2FzwI5YcACNtzyY=:object`, false, "", "8Wx4XAY/zwI5YcACNtzyY=:object"},
			{`dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`n/dxterraformdev/p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId n/dxterraformdev/p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`n/dxterraformdev/b/bucket/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId n/dxterraformdev/b/bucket/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`/b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId /b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
		}

		for _, test := range tests {
			if _, _, parId, err := parsePreauthenticatedRequestCompositeId(test.parId); err != nil {

				if test.expectError && err == nil {
					t.Fatalf("expected an error but got none")
				}
				if !test.expectError && err != nil {
					t.Fatalf("did not expect an error but got one %s ", err.Error())
				}
				if test.expectError && err != nil && err.Error() != test.error {
					t.Fatalf("unexpected error %s, expected: %s ", err.Error(), test.error)
				}

				if !test.expectError && err == nil && parId != test.parsedParId {
					t.Fatalf("parId parsed incorrectly, got: %s, expected: %s ", parId, test.parsedParId)
				}
			}
		}
	})
}

// Tests the usage of newly created parameter "object_name" in place of the deprecated "object" parameter in object_storage_PAR
// issue-routing-tag: object_storage/default
func TestObjectStoragePreauthenticatedRequestResource_newObjectNameParam(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePreauthenticatedRequestResource_newObjectNameParam")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_preauthrequest.test_preauthenticated_request"
	datasourceName := "data.oci_objectstorage_preauthrequests.test_preauthenticated_requests"
	singularDatasourceName := "data.oci_objectstorage_preauthrequest.test_preauthenticated_request"

	updatedRepresentation := RepresentationCopyWithNewProperties(
		RepresentationCopyWithRemovedProperties(
			preauthenticatedRequestRepresentation,
			[]string{"object"}), map[string]interface{}{
			"object_name": Representation{
				RepType: Optional,
				Create:  `my-test-object-1`,
			},
		})

	var resId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+PreauthenticatedRequestResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Create, updatedRepresentation), "objectstorage", "preauthenticatedRequest", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStoragePreauthenticatedRequestDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, updatedRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "AnyObjectWrite"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, updatedRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(resourceName, "bucket_listing_action", ""),
					resource.TestCheckResourceAttrSet(resourceName, "access_uri"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_name", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),

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

			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequests", "test_preauthenticated_requests", Optional, Update, preauthenticatedRequestDataSourceRepresentation) +
					compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, updatedRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "bucket", testPreAuthBucketName),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttr(datasourceName, "object_name_prefix", "my-test-object"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.bucket_listing_action", ""),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.name", "-tf-par"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.object_name", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.time_expires", expirationTimeForPar.String()),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, updatedRepresentation),

				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testPreAuthBucketName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "par_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket_listing_action", ""),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object_name", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_expires", expirationTimeForPar.String()),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
			},
			//verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"access_uri",
					"time_expires",
				},
				ResourceName: resourceName,
			},
		},
	})
}

// issue-routing-tag: object_storage/default
func TestResourceObjectstoragePARTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceObjectstoragePARTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceObjectstoragePARTestSuite))
}
