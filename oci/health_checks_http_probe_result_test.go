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
)

var (
	startTime                               = time.Now()
	httpProbeResultDataSourceRepresentation = map[string]interface{}{
		"probe_configuration_id":              Representation{RepType: Required, Create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"start_time_greater_than_or_equal_to": Representation{RepType: Optional, Create: strconv.FormatInt(startTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"start_time_less_than_or_equal_to":    Representation{RepType: Optional, Create: strconv.FormatInt(startTime.Add(30*time.Minute).UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"target":                              Representation{RepType: Optional, Create: `www.oracle.com`},
	}

	HttpProbeResultResourceConfig = GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Required, Create, httpMonitorRepresentation)
)

// issue-routing-tag: health_checks/default
func TestHealthChecksHttpProbeResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpProbeResultResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_http_probe_results.test_http_probe_results"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + HttpProbeResultResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					if httpreplay.ShouldRetryImmediately() {
						time.Sleep(10 * time.Millisecond)
					} else {
						time.Sleep(5 * time.Minute)
					}
					return nil
				},
			),
		},
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_health_checks_http_probe_results", "test_http_probe_results", Optional, Create, httpProbeResultDataSourceRepresentation) +
				compartmentIdVariableStr + HttpProbeResultResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "probe_configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_time_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_time_less_than_or_equal_to"),
				resource.TestCheckResourceAttr(datasourceName, "target", "www.oracle.com"),

				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.connect_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.connect_start"),
				resource.TestCheckResourceAttr(datasourceName, "http_probe_results.0.connection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "http_probe_results.0.dns.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.domain_lookup_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.domain_lookup_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.duration"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.encoded_body_size"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.error_category"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.fetch_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.is_healthy"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.is_timed_out"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.probe_configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.protocol"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.request_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.response_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.response_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.secure_connection_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.start_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.status_code"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.target"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.vantage_point_name"),
			),
		},
	})
}
