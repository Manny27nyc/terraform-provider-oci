// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreatePatchDetails Properties used in patch create operations.
type CreatePatchDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The type of the patch applied or being applied on the application.
	PatchType CreatePatchDetailsPatchTypeEnum `mandatory:"true" json:"patchType"`

	// The array of object keys to publish into application.
	ObjectKeys []string `mandatory:"true" json:"objectKeys"`

	// The object's key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreatePatchDetails) String() string {
	return common.PointerString(m)
}

// CreatePatchDetailsPatchTypeEnum Enum with underlying type: string
type CreatePatchDetailsPatchTypeEnum string

// Set of constants representing the allowable values for CreatePatchDetailsPatchTypeEnum
const (
	CreatePatchDetailsPatchTypePublish   CreatePatchDetailsPatchTypeEnum = "PUBLISH"
	CreatePatchDetailsPatchTypeRefresh   CreatePatchDetailsPatchTypeEnum = "REFRESH"
	CreatePatchDetailsPatchTypeUnpublish CreatePatchDetailsPatchTypeEnum = "UNPUBLISH"
)

var mappingCreatePatchDetailsPatchType = map[string]CreatePatchDetailsPatchTypeEnum{
	"PUBLISH":   CreatePatchDetailsPatchTypePublish,
	"REFRESH":   CreatePatchDetailsPatchTypeRefresh,
	"UNPUBLISH": CreatePatchDetailsPatchTypeUnpublish,
}

// GetCreatePatchDetailsPatchTypeEnumValues Enumerates the set of values for CreatePatchDetailsPatchTypeEnum
func GetCreatePatchDetailsPatchTypeEnumValues() []CreatePatchDetailsPatchTypeEnum {
	values := make([]CreatePatchDetailsPatchTypeEnum, 0)
	for _, v := range mappingCreatePatchDetailsPatchType {
		values = append(values, v)
	}
	return values
}
