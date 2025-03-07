// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v52/cloudguard"

	oci_common "github.com/oracle/oci-go-sdk/v52/common"
)

func init() {
	RegisterOracleClient("oci_cloud_guard.CloudGuardClient", &OracleClient{InitClientFn: initCloudguardCloudGuardClient})
}

func initCloudguardCloudGuardClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_guard.NewCloudGuardClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) cloudGuardClient() *oci_cloud_guard.CloudGuardClient {
	return m.GetClient("oci_cloud_guard.CloudGuardClient").(*oci_cloud_guard.CloudGuardClient)
}
