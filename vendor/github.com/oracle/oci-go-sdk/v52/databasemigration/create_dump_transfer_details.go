// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateDumpTransferDetails Optional additional properties for dump transfer.
type CreateDumpTransferDetails struct {
	Source CreateHostDumpTransferDetails `mandatory:"false" json:"source"`

	Target CreateHostDumpTransferDetails `mandatory:"false" json:"target"`
}

func (m CreateDumpTransferDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDumpTransferDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Source createhostdumptransferdetails `json:"source"`
		Target createhostdumptransferdetails `json:"target"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(CreateHostDumpTransferDetails)
	} else {
		m.Source = nil
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(CreateHostDumpTransferDetails)
	} else {
		m.Target = nil
	}

	return
}
