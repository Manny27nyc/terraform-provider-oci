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

// CreateNetworkAddressListVcnAddressesDetails The information about new NetworkAddressListVcnAddresses.
type CreateNetworkAddressListVcnAddressesDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A list of private address prefixes, each associated with a particular VCN.
	// To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6.
	VcnAddresses []PrivateAddresses `mandatory:"true" json:"vcnAddresses"`

	// NetworkAddressList display name, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

//GetDisplayName returns DisplayName
func (m CreateNetworkAddressListVcnAddressesDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateNetworkAddressListVcnAddressesDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateNetworkAddressListVcnAddressesDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateNetworkAddressListVcnAddressesDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m CreateNetworkAddressListVcnAddressesDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CreateNetworkAddressListVcnAddressesDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNetworkAddressListVcnAddressesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNetworkAddressListVcnAddressesDetails CreateNetworkAddressListVcnAddressesDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateNetworkAddressListVcnAddressesDetails
	}{
		"VCN_ADDRESSES",
		(MarshalTypeCreateNetworkAddressListVcnAddressesDetails)(m),
	}

	return json.Marshal(&s)
}
