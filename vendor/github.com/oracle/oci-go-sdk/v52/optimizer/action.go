// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// Action Details about the recommended action.
type Action struct {

	// The status of the resource action.
	Type ActionTypeEnum `mandatory:"true" json:"type"`

	// Text describing the recommended action.
	Description *string `mandatory:"true" json:"description"`

	// The URL path to documentation that explains how to perform the action.
	Url *string `mandatory:"true" json:"url"`
}

func (m Action) String() string {
	return common.PointerString(m)
}
