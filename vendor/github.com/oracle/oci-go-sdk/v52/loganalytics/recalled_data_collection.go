// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// RecalledDataCollection This is the list of recalled data
type RecalledDataCollection struct {

	// This is the array of recalled data
	Items []RecalledData `mandatory:"true" json:"items"`
}

func (m RecalledDataCollection) String() string {
	return common.PointerString(m)
}
