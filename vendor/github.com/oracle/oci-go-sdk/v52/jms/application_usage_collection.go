// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ApplicationUsageCollection Results of an application search. Contains ApplicationUsage items.
type ApplicationUsageCollection struct {

	// A list of applications.
	Items []ApplicationUsage `mandatory:"true" json:"items"`
}

func (m ApplicationUsageCollection) String() string {
	return common.PointerString(m)
}
