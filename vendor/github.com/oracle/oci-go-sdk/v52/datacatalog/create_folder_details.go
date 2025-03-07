// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateFolderDetails Properties used in folder create operations.
type CreateFolderDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"true" json:"timeExternal"`

	// Optional user friendly business name of the folder. If set, this supplements the harvested display name of the object.
	BusinessName *string `mandatory:"false" json:"businessName"`

	// Detailed description of a folder.
	Description *string `mandatory:"false" json:"description"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertySetUsage `mandatory:"false" json:"customPropertyMembers"`

	// A map of maps that contains the properties which are specific to the folder type. Each folder type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// folders have required properties within the "default" category. To determine the set of optional and
	// required properties for a folder type, a query can be done on '/types?type=folder' that returns a
	// collection of all folder types. The appropriate folder type, which includes definitions of all of
	// it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// The key of the containing folder or null if there isn't a parent folder.
	ParentFolderKey *string `mandatory:"false" json:"parentFolderKey"`

	// The job key of the harvest process that updated the folder definition from the source system.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`

	// Folder harvesting status.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`
}

func (m CreateFolderDetails) String() string {
	return common.PointerString(m)
}
