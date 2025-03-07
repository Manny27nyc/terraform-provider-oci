// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PublicVantagePointCollection The results of a public vantage point search, which contains PublicVantagePointSummary items and other data in an APM domain.
type PublicVantagePointCollection struct {

	// List of PublicVantagePointSummary items.
	Items []PublicVantagePointSummary `mandatory:"true" json:"items"`
}

func (m PublicVantagePointCollection) String() string {
	return common.PointerString(m)
}
