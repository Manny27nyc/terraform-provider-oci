// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// UpdateConnectionFromAdwc The details to update an Autonomous Data Warehouse data asset connection.
type UpdateConnectionFromAdwc struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The Autonomous Data Warehouse instance service name.
	TnsAlias *string `mandatory:"false" json:"tnsAlias"`

	// Array of service names that are available for selection in the tnsAlias property.
	TnsNames []string `mandatory:"false" json:"tnsNames"`

	// The user name for the connection.
	Username *string `mandatory:"false" json:"username"`

	// The password for the connection.
	Password *string `mandatory:"false" json:"password"`

	PasswordSecret *SensitiveAttribute `mandatory:"false" json:"passwordSecret"`
}

//GetKey returns Key
func (m UpdateConnectionFromAdwc) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateConnectionFromAdwc) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateConnectionFromAdwc) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateConnectionFromAdwc) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateConnectionFromAdwc) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateConnectionFromAdwc) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateConnectionFromAdwc) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateConnectionFromAdwc) GetIdentifier() *string {
	return m.Identifier
}

//GetConnectionProperties returns ConnectionProperties
func (m UpdateConnectionFromAdwc) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateConnectionFromAdwc) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateConnectionFromAdwc) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateConnectionFromAdwc) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateConnectionFromAdwc UpdateConnectionFromAdwc
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateConnectionFromAdwc
	}{
		"ORACLE_ADWC_CONNECTION",
		(MarshalTypeUpdateConnectionFromAdwc)(m),
	}

	return json.Marshal(&s)
}
