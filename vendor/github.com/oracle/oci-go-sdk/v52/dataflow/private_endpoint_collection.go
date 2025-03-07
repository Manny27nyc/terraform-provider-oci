// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PrivateEndpointCollection The results of a query for a list of private endpoints. It contains PrivateEndpointSummary items.
type PrivateEndpointCollection struct {

	// A list of private endpoints.
	Items []PrivateEndpointSummary `mandatory:"true" json:"items"`
}

func (m PrivateEndpointCollection) String() string {
	return common.PointerString(m)
}
