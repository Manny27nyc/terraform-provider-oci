// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ManagementSavedSearchCollection List of saved searches.
type ManagementSavedSearchCollection struct {

	// Array of saved search summaries.
	Items []ManagementSavedSearchSummary `mandatory:"true" json:"items"`
}

func (m ManagementSavedSearchCollection) String() string {
	return common.PointerString(m)
}
