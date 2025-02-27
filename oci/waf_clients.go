// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_waf "github.com/oracle/oci-go-sdk/v52/waf"

	oci_common "github.com/oracle/oci-go-sdk/v52/common"
)

func init() {
	RegisterOracleClient("oci_waf.WafClient", &OracleClient{InitClientFn: initWafWafClient})
}

func initWafWafClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_waf.NewWafClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) wafClient() *oci_waf.WafClient {
	return m.GetClient("oci_waf.WafClient").(*oci_waf.WafClient)
}
