// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// AllowAction An object that represents an action which upon matching rule skips all remaining rules in the current module.
type AllowAction struct {

	// Action name. Can be used to reference the action.
	Name *string `mandatory:"true" json:"name"`
}

//GetName returns Name
func (m AllowAction) GetName() *string {
	return m.Name
}

func (m AllowAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AllowAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllowAction AllowAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAllowAction
	}{
		"ALLOW",
		(MarshalTypeAllowAction)(m),
	}

	return json.Marshal(&s)
}
