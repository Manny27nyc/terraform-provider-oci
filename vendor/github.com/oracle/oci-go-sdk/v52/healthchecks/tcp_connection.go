// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// TcpConnection TCP connection results.  All durations are in milliseconds.
type TcpConnection struct {

	// The connection IP address.
	Address *string `mandatory:"false" json:"address"`

	// The port.
	Port *int `mandatory:"false" json:"port"`

	// Total connect duration, calculated using `connectEnd` minus `connectStart`.
	ConnectDuration *float64 `mandatory:"false" json:"connectDuration"`

	// The duration to secure the connection.  This value will be zero for
	// insecure connections.  Calculated using `connectEnd` minus `secureConnectionStart`.
	SecureConnectDuration *float64 `mandatory:"false" json:"secureConnectDuration"`
}

func (m TcpConnection) String() string {
	return common.PointerString(m)
}
