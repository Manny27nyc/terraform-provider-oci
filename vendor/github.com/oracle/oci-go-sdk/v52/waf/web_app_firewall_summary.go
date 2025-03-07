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

// WebAppFirewallSummary Summary of the WebAppFirewall.
type WebAppFirewallSummary interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WebAppFirewall.
	GetId() *string

	// WebAppFirewall display name, can be renamed.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of WebAppFirewallPolicy, which is attached to the resource.
	GetWebAppFirewallPolicyId() *string

	// The time the WebAppFirewall was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the WebAppFirewall.
	GetLifecycleState() WebAppFirewallLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the WebAppFirewall was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in FAILED state.
	GetLifecycleDetails() *string
}

type webappfirewallsummary struct {
	JsonData               []byte
	Id                     *string                           `mandatory:"true" json:"id"`
	DisplayName            *string                           `mandatory:"true" json:"displayName"`
	CompartmentId          *string                           `mandatory:"true" json:"compartmentId"`
	WebAppFirewallPolicyId *string                           `mandatory:"true" json:"webAppFirewallPolicyId"`
	TimeCreated            *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState         WebAppFirewallLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	FreeformTags           map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags            map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	SystemTags             map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
	TimeUpdated            *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails       *string                           `mandatory:"false" json:"lifecycleDetails"`
	BackendType            string                            `json:"backendType"`
}

// UnmarshalJSON unmarshals json
func (m *webappfirewallsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerwebappfirewallsummary webappfirewallsummary
	s := struct {
		Model Unmarshalerwebappfirewallsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.WebAppFirewallPolicyId = s.Model.WebAppFirewallPolicyId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.BackendType = s.Model.BackendType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *webappfirewallsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BackendType {
	case "LOAD_BALANCER":
		mm := WebAppFirewallLoadBalancerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m webappfirewallsummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m webappfirewallsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m webappfirewallsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetWebAppFirewallPolicyId returns WebAppFirewallPolicyId
func (m webappfirewallsummary) GetWebAppFirewallPolicyId() *string {
	return m.WebAppFirewallPolicyId
}

//GetTimeCreated returns TimeCreated
func (m webappfirewallsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m webappfirewallsummary) GetLifecycleState() WebAppFirewallLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m webappfirewallsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m webappfirewallsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m webappfirewallsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetTimeUpdated returns TimeUpdated
func (m webappfirewallsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m webappfirewallsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m webappfirewallsummary) String() string {
	return common.PointerString(m)
}
