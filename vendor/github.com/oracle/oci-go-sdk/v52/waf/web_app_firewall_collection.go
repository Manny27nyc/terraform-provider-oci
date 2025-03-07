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

// WebAppFirewallCollection Result of a WebAppFirewall list operation.
type WebAppFirewallCollection struct {

	// List of WebAppFirewalls.
	Items []WebAppFirewallSummary `mandatory:"true" json:"items"`
}

func (m WebAppFirewallCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *WebAppFirewallCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []webappfirewallsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]WebAppFirewallSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(WebAppFirewallSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
