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

// LogAnalyticsSourceEntityType LogAnalyticsSourceEntityType
type LogAnalyticsSourceEntityType struct {

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The entity type.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The type category.
	EntityTypeCategory *string `mandatory:"false" json:"entityTypeCategory"`

	// The entity type display name.
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`
}

func (m LogAnalyticsSourceEntityType) String() string {
	return common.PointerString(m)
}
