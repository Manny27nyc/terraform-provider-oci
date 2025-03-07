// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateSpanFilterDetails A span filter is a named setting that specifies filter criteria to match a subset of the spans.
type CreateSpanFilterDetails struct {

	// The name by which this filter can be displayed in the UI.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The string that defines the Span Filter expression.
	FilterText *string `mandatory:"true" json:"filterText"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An optional string that describes what the filter is intended or used for.
	Description *string `mandatory:"false" json:"description"`
}

//GetFreeformTags returns FreeformTags
func (m CreateSpanFilterDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateSpanFilterDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSpanFilterDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateSpanFilterDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSpanFilterDetails CreateSpanFilterDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateSpanFilterDetails
	}{
		"SPAN_FILTER",
		(MarshalTypeCreateSpanFilterDetails)(m),
	}

	return json.Marshal(&s)
}
