// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// BuildSourceCollection Collection of Build Sources.
type BuildSourceCollection struct {

	// Collection of Build sources. In case of UPDATE operation, replaces existing Build sources list. Merging with existing Build Sources is not supported.
	Items []BuildSource `mandatory:"true" json:"items"`
}

func (m BuildSourceCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *BuildSourceCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []buildsource `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]BuildSource, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(BuildSource)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
