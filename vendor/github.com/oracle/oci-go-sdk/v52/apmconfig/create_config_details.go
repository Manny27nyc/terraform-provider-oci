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

// CreateConfigDetails The request body used to create new Configuration entities. It must specify the configuration type of the item to
// create, as well as the actual data to populate the item with.
type CreateConfigDetails interface {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createconfigdetails struct {
	JsonData     []byte
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigType   string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *createconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconfigdetails createconfigdetails
	s := struct {
		Model Unmarshalercreateconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "SPAN_FILTER":
		mm := CreateSpanFilterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "METRIC_GROUP":
		mm := CreateMetricGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APDEX":
		mm := CreateApdexRulesDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFreeformTags returns FreeformTags
func (m createconfigdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createconfigdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createconfigdetails) String() string {
	return common.PointerString(m)
}
