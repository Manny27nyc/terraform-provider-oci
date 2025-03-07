// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ScanDetails The Single Client Access Name (SCAN) details.
type ScanDetails struct {

	// The SCAN hostname.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The SCAN TCPIP port. Default is 1521.
	Port *int `mandatory:"true" json:"port"`

	// The list of SCAN IP addresses. Three addresses should be provided.
	Ips []string `mandatory:"true" json:"ips"`

	// The SCAN TCPIP port. Default is 1521.
	ScanListenerPortTcp *int `mandatory:"false" json:"scanListenerPortTcp"`

	// The SCAN TCPIP SSL port. Default is 2484.
	ScanListenerPortTcpSsl *int `mandatory:"false" json:"scanListenerPortTcpSsl"`
}

func (m ScanDetails) String() string {
	return common.PointerString(m)
}
