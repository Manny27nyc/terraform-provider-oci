// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_waas "github.com/oracle/oci-go-sdk/v52/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	WaasCertificateRequiredOnlyResource = WaasCertificateResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Required, Create, waasCertificateRepresentation)

	WaasCertificateResourceConfig = WaasCertificateResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Update, waasCertificateRepresentation)

	waasCertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id": Representation{RepType: Required, Create: `${oci_waas_certificate.test_certificate.id}`},
	}

	waasCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_names":                         Representation{RepType: Optional, Create: []string{`displayName2`}},
		"ids":                                   Representation{RepType: Optional, Create: []string{`${oci_waas_certificate.test_certificate.id}`}},
		"states":                                Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{RepType: Optional, Create: `${timestamp()}`},
		"filter":                                RepresentationGroup{Required, waasCertificateDataSourceFilterRepresentation}}
	waasCertificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waas_certificate.test_certificate.id}`}},
	}

	waasCertificateRepresentation = map[string]interface{}{
		"certificate_data":               Representation{RepType: Required, Create: "${var.ca_certificate_value}"},
		"compartment_id":                 Representation{RepType: Required, Create: `${var.compartment_id}`},
		"private_key_data":               Representation{RepType: Required, Create: "${var.private_key_value}"},
		"defined_tags":                   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_trust_verification_disabled": Representation{RepType: Required, Create: `true`},
		"timeouts":                       RepresentationGroup{Required, waasCertificateTimeoutsRepresentation},
	}
	// Add timeout for delete upto the same time as the dependency to ensure clean delete
	waasCertificateTimeoutsRepresentation = map[string]interface{}{
		"delete": Representation{RepType: Required, Create: `60m`},
	}

	WaasCertificateResourceDependencies = DefinedTagsDependencies + caCertificateVariableStr + privateKeyVariableStr
)

// issue-routing-tag: waas/default
func TestWaasCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_certificate.test_certificate"
	datasourceName := "data.oci_waas_certificates.test_certificates"
	singularDatasourceName := "data.oci_waas_certificate.test_certificate"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+CertificateResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Create, certificateRepresentation), "waas", "certificate", t)

	ResourceTest(t, testAccCheckWaasCertificateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WaasCertificateResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Required, Create, waasCertificateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "certificate_data", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "private_key_data", regexp.MustCompile("-----BEGIN RSA.*")),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaasCertificateResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WaasCertificateResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Create, waasCertificateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "certificate_data", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_trust_verification_disabled", "true"),
				resource.TestMatchResourceAttr(resourceName, "private_key_data", regexp.MustCompile("-----BEGIN RSA.*")),
				resource.TestCheckResourceAttr(resourceName, "public_key_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
				//resource.TestCheckResourceAttrSet(resourceName, "signature_algorithm"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_before"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WaasCertificateResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Create,
					RepresentationCopyWithNewProperties(waasCertificateRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "certificate_data", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_trust_verification_disabled", "true"),
				resource.TestMatchResourceAttr(resourceName, "private_key_data", regexp.MustCompile("-----BEGIN RSA.*")),
				resource.TestCheckResourceAttr(resourceName, "public_key_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
				//resource.TestCheckResourceAttrSet(resourceName, "signature_algorithm"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_before"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),

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
			Config: config + compartmentIdVariableStr + WaasCertificateResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Update, waasCertificateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "certificate_data", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_trust_verification_disabled", "true"),
				resource.TestMatchResourceAttr(resourceName, "private_key_data", regexp.MustCompile("-----BEGIN RSA.*")),
				resource.TestCheckResourceAttr(resourceName, "public_key_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
				//resource.TestCheckResourceAttrSet(resourceName, "signature_algorithm"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),
				resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_before"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),

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
				GenerateDataSourceFromRepresentationMap("oci_waas_certificates", "test_certificates", Optional, Update, waasCertificateDataSourceRepresentation) +
				compartmentIdVariableStr + WaasCertificateResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Optional, Update, waasCertificateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_names.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.time_not_valid_after"),
				resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.version"),
			),
			// Non empty plan expected because the data source input relies on interpolation syntax
			ExpectNonEmptyPlan: true,
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Required, Create, waasCertificateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + WaasCertificateResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_data", "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_trust_verification_disabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "issued_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "issuer_name.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "public_key_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subject_name.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_not_valid_after"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_not_valid_before"),
			),
		},
	})
}

func testAccCheckWaasCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).waasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_certificate" {
			noResourceFound = false
			request := oci_waas.GetCertificateRequest{}

			tmp := rs.Primary.ID
			request.CertificateId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")

			response, err := client.GetCertificate(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waas.CertificateLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("WaasCertificate") {
		resource.AddTestSweepers("WaasCertificate", &resource.Sweeper{
			Name:         "WaasCertificate",
			Dependencies: DependencyGraph["certificate"],
			F:            sweepWaasCertificateResource,
		})
	}
}

func sweepWaasCertificateResource(compartment string) error {
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()
	certificateIds, err := getCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, certificateId := range certificateIds {
		if ok := SweeperDefaultResourceId[certificateId]; !ok {
			deleteCertificateRequest := oci_waas.DeleteCertificateRequest{}

			deleteCertificateRequest.CertificateId = &certificateId

			deleteCertificateRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")
			_, error := waasClient.DeleteCertificate(context.Background(), deleteCertificateRequest)
			if error != nil {
				fmt.Printf("Error deleting Certificate %s %s, It is possible that the resource is already deleted. Please verify manually \n", certificateId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &certificateId, certificateSweepWaitCondition, time.Duration(3*time.Minute),
				certificateSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getCertificateIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "CertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()

	listCertificatesRequest := oci_waas.ListCertificatesRequest{}
	listCertificatesRequest.CompartmentId = &compartmentId
	listCertificatesResponse, err := waasClient.ListCertificates(context.Background(), listCertificatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Certificate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, certificate := range listCertificatesResponse.Items {
		id := *certificate.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "CertificateId", id)
	}
	return resourceIds, nil
}

func certificateSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if certificateResponse, ok := response.Response.(oci_waas.GetCertificateResponse); ok {
		return certificateResponse.LifecycleState != oci_waas.CertificateLifecycleStateDeleted
	}
	return false
}

func certificateSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.waasClient().GetCertificate(context.Background(), oci_waas.GetCertificateRequest{
		CertificateId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
