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

// LogAnalyticsLookupFields LogAnalyticsLookupFields
type LogAnalyticsLookupFields struct {

	// The common field name.
	CommonFieldName *string `mandatory:"false" json:"commonFieldName"`

	// The default match value.
	DefaultMatchValue *string `mandatory:"false" json:"defaultMatchValue"`

	// The display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A flag indicating whether or not the field is a common field.
	IsCommonField *bool `mandatory:"false" json:"isCommonField"`

	// The match operator.
	MatchOperator *string `mandatory:"false" json:"matchOperator"`

	// The field name.
	Name *string `mandatory:"false" json:"name"`

	// The position.
	Position *int64 `mandatory:"false" json:"position"`
}

func (m LogAnalyticsLookupFields) String() string {
	return common.PointerString(m)
}
