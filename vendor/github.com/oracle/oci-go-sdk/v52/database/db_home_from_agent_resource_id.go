// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// DbHomeFromAgentResourceId The representation of DbHomeFromAgentResourceId
type DbHomeFromAgentResourceId struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Database Home. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Database Home.
	LifecycleState DbHomeFromAgentResourceIdLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Oracle Database version.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The location of the Oracle Database Home.
	DbHomeLocation *string `mandatory:"true" json:"dbHomeLocation"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Database Home was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// List of one-off patches for Database Homes.
	OneOffPatches []string `mandatory:"false" json:"oneOffPatches"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`
}

func (m DbHomeFromAgentResourceId) String() string {
	return common.PointerString(m)
}

// DbHomeFromAgentResourceIdLifecycleStateEnum Enum with underlying type: string
type DbHomeFromAgentResourceIdLifecycleStateEnum string

// Set of constants representing the allowable values for DbHomeFromAgentResourceIdLifecycleStateEnum
const (
	DbHomeFromAgentResourceIdLifecycleStateProvisioning DbHomeFromAgentResourceIdLifecycleStateEnum = "PROVISIONING"
	DbHomeFromAgentResourceIdLifecycleStateAvailable    DbHomeFromAgentResourceIdLifecycleStateEnum = "AVAILABLE"
	DbHomeFromAgentResourceIdLifecycleStateUpdating     DbHomeFromAgentResourceIdLifecycleStateEnum = "UPDATING"
	DbHomeFromAgentResourceIdLifecycleStateTerminating  DbHomeFromAgentResourceIdLifecycleStateEnum = "TERMINATING"
	DbHomeFromAgentResourceIdLifecycleStateTerminated   DbHomeFromAgentResourceIdLifecycleStateEnum = "TERMINATED"
	DbHomeFromAgentResourceIdLifecycleStateFailed       DbHomeFromAgentResourceIdLifecycleStateEnum = "FAILED"
)

var mappingDbHomeFromAgentResourceIdLifecycleState = map[string]DbHomeFromAgentResourceIdLifecycleStateEnum{
	"PROVISIONING": DbHomeFromAgentResourceIdLifecycleStateProvisioning,
	"AVAILABLE":    DbHomeFromAgentResourceIdLifecycleStateAvailable,
	"UPDATING":     DbHomeFromAgentResourceIdLifecycleStateUpdating,
	"TERMINATING":  DbHomeFromAgentResourceIdLifecycleStateTerminating,
	"TERMINATED":   DbHomeFromAgentResourceIdLifecycleStateTerminated,
	"FAILED":       DbHomeFromAgentResourceIdLifecycleStateFailed,
}

// GetDbHomeFromAgentResourceIdLifecycleStateEnumValues Enumerates the set of values for DbHomeFromAgentResourceIdLifecycleStateEnum
func GetDbHomeFromAgentResourceIdLifecycleStateEnumValues() []DbHomeFromAgentResourceIdLifecycleStateEnum {
	values := make([]DbHomeFromAgentResourceIdLifecycleStateEnum, 0)
	for _, v := range mappingDbHomeFromAgentResourceIdLifecycleState {
		values = append(values, v)
	}
	return values
}
