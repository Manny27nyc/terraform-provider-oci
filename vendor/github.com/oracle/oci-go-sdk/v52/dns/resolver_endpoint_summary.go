// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ResolverEndpointSummary An OCI DNS resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverEndpointSummary interface {

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	GetName() *string

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	GetIsForwarding() *bool

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	GetIsListening() *bool

	// The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under
	// and will be updated if the resolver's compartment is changed.
	GetCompartmentId() *string

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	GetTimeCreated() *common.SDKTime

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ"
	// format with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	GetTimeUpdated() *common.SDKTime

	// The current state of the resource.
	GetLifecycleState() ResolverEndpointSummaryLifecycleStateEnum

	// The canonical absolute URL of the resource.
	GetSelf() *string

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	GetForwardingAddress() *string

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	GetListeningAddress() *string
}

type resolverendpointsummary struct {
	JsonData          []byte
	Name              *string                                   `mandatory:"true" json:"name"`
	IsForwarding      *bool                                     `mandatory:"true" json:"isForwarding"`
	IsListening       *bool                                     `mandatory:"true" json:"isListening"`
	CompartmentId     *string                                   `mandatory:"true" json:"compartmentId"`
	TimeCreated       *common.SDKTime                           `mandatory:"true" json:"timeCreated"`
	TimeUpdated       *common.SDKTime                           `mandatory:"true" json:"timeUpdated"`
	LifecycleState    ResolverEndpointSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	Self              *string                                   `mandatory:"true" json:"self"`
	ForwardingAddress *string                                   `mandatory:"false" json:"forwardingAddress"`
	ListeningAddress  *string                                   `mandatory:"false" json:"listeningAddress"`
	EndpointType      string                                    `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *resolverendpointsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresolverendpointsummary resolverendpointsummary
	s := struct {
		Model Unmarshalerresolverendpointsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.IsForwarding = s.Model.IsForwarding
	m.IsListening = s.Model.IsListening
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.Self = s.Model.Self
	m.ForwardingAddress = s.Model.ForwardingAddress
	m.ListeningAddress = s.Model.ListeningAddress
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resolverendpointsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := ResolverVnicEndpointSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m resolverendpointsummary) GetName() *string {
	return m.Name
}

//GetIsForwarding returns IsForwarding
func (m resolverendpointsummary) GetIsForwarding() *bool {
	return m.IsForwarding
}

//GetIsListening returns IsListening
func (m resolverendpointsummary) GetIsListening() *bool {
	return m.IsListening
}

//GetCompartmentId returns CompartmentId
func (m resolverendpointsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m resolverendpointsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m resolverendpointsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m resolverendpointsummary) GetLifecycleState() ResolverEndpointSummaryLifecycleStateEnum {
	return m.LifecycleState
}

//GetSelf returns Self
func (m resolverendpointsummary) GetSelf() *string {
	return m.Self
}

//GetForwardingAddress returns ForwardingAddress
func (m resolverendpointsummary) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

//GetListeningAddress returns ListeningAddress
func (m resolverendpointsummary) GetListeningAddress() *string {
	return m.ListeningAddress
}

func (m resolverendpointsummary) String() string {
	return common.PointerString(m)
}

// ResolverEndpointSummaryLifecycleStateEnum Enum with underlying type: string
type ResolverEndpointSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ResolverEndpointSummaryLifecycleStateEnum
const (
	ResolverEndpointSummaryLifecycleStateActive   ResolverEndpointSummaryLifecycleStateEnum = "ACTIVE"
	ResolverEndpointSummaryLifecycleStateCreating ResolverEndpointSummaryLifecycleStateEnum = "CREATING"
	ResolverEndpointSummaryLifecycleStateDeleted  ResolverEndpointSummaryLifecycleStateEnum = "DELETED"
	ResolverEndpointSummaryLifecycleStateDeleting ResolverEndpointSummaryLifecycleStateEnum = "DELETING"
	ResolverEndpointSummaryLifecycleStateFailed   ResolverEndpointSummaryLifecycleStateEnum = "FAILED"
	ResolverEndpointSummaryLifecycleStateUpdating ResolverEndpointSummaryLifecycleStateEnum = "UPDATING"
)

var mappingResolverEndpointSummaryLifecycleState = map[string]ResolverEndpointSummaryLifecycleStateEnum{
	"ACTIVE":   ResolverEndpointSummaryLifecycleStateActive,
	"CREATING": ResolverEndpointSummaryLifecycleStateCreating,
	"DELETED":  ResolverEndpointSummaryLifecycleStateDeleted,
	"DELETING": ResolverEndpointSummaryLifecycleStateDeleting,
	"FAILED":   ResolverEndpointSummaryLifecycleStateFailed,
	"UPDATING": ResolverEndpointSummaryLifecycleStateUpdating,
}

// GetResolverEndpointSummaryLifecycleStateEnumValues Enumerates the set of values for ResolverEndpointSummaryLifecycleStateEnum
func GetResolverEndpointSummaryLifecycleStateEnumValues() []ResolverEndpointSummaryLifecycleStateEnum {
	values := make([]ResolverEndpointSummaryLifecycleStateEnum, 0)
	for _, v := range mappingResolverEndpointSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// ResolverEndpointSummaryEndpointTypeEnum Enum with underlying type: string
type ResolverEndpointSummaryEndpointTypeEnum string

// Set of constants representing the allowable values for ResolverEndpointSummaryEndpointTypeEnum
const (
	ResolverEndpointSummaryEndpointTypeVnic ResolverEndpointSummaryEndpointTypeEnum = "VNIC"
)

var mappingResolverEndpointSummaryEndpointType = map[string]ResolverEndpointSummaryEndpointTypeEnum{
	"VNIC": ResolverEndpointSummaryEndpointTypeVnic,
}

// GetResolverEndpointSummaryEndpointTypeEnumValues Enumerates the set of values for ResolverEndpointSummaryEndpointTypeEnum
func GetResolverEndpointSummaryEndpointTypeEnumValues() []ResolverEndpointSummaryEndpointTypeEnum {
	values := make([]ResolverEndpointSummaryEndpointTypeEnum, 0)
	for _, v := range mappingResolverEndpointSummaryEndpointType {
		values = append(values, v)
	}
	return values
}
