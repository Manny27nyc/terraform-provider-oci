// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateDataGuardAssociationToExistingDbSystemDetails The configuration details for creating a Data Guard association for a bare metal or Exadata DB system database. For these types of DB system databases, the `creationType` should be `ExistingDbSystem`. A standby database will be created in the DB system you specify.
// To create a Data Guard association for a database in a virtual machine DB system, use the CreateDataGuardAssociationWithNewDbSystemDetails subtype instead.
type CreateDataGuardAssociationToExistingDbSystemDetails struct {

	// A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.
	// The password must contain no fewer than nine characters and include:
	// * At least two uppercase characters.
	// * At least two lowercase characters.
	// * At least two numeric characters.
	// * At least two special characters. Valid special characters include "_", "#", and "-" only.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// Specifies the `DB_UNIQUE_NAME` of the peer database to be created.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// Specifies a prefix for the `Oracle SID` of the database to be created.
	PeerSidPrefix *string `mandatory:"false" json:"peerSidPrefix"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system in which to create the standby database.
	// You must supply this value if creationType is `ExistingDbSystem`.
	PeerDbSystemId *string `mandatory:"false" json:"peerDbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB home in which to create the standby database.
	// You must supply this value to create standby database with an existing DB home
	PeerDbHomeId *string `mandatory:"false" json:"peerDbHomeId"`

	// The protection mode to set up between the primary and standby databases. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only protection mode currently supported by the Database service is MAXIMUM_PERFORMANCE.
	ProtectionMode CreateDataGuardAssociationDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	// * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	// * MAXIMUM_PERFORMANCE - ASYNC
	// * MAXIMUM_PROTECTION - SYNC
	// For more information, see
	// Redo Transport Services (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400)
	// in the Oracle Data Guard documentation.
	// **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC.
	TransportType CreateDataGuardAssociationDetailsTransportTypeEnum `mandatory:"true" json:"transportType"`
}

//GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

//GetDatabaseAdminPassword returns DatabaseAdminPassword
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetDatabaseAdminPassword() *string {
	return m.DatabaseAdminPassword
}

//GetProtectionMode returns ProtectionMode
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetProtectionMode() CreateDataGuardAssociationDetailsProtectionModeEnum {
	return m.ProtectionMode
}

//GetTransportType returns TransportType
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetTransportType() CreateDataGuardAssociationDetailsTransportTypeEnum {
	return m.TransportType
}

//GetPeerDbUniqueName returns PeerDbUniqueName
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetPeerDbUniqueName() *string {
	return m.PeerDbUniqueName
}

//GetPeerSidPrefix returns PeerSidPrefix
func (m CreateDataGuardAssociationToExistingDbSystemDetails) GetPeerSidPrefix() *string {
	return m.PeerSidPrefix
}

func (m CreateDataGuardAssociationToExistingDbSystemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDataGuardAssociationToExistingDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDataGuardAssociationToExistingDbSystemDetails CreateDataGuardAssociationToExistingDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"creationType"`
		MarshalTypeCreateDataGuardAssociationToExistingDbSystemDetails
	}{
		"ExistingDbSystem",
		(MarshalTypeCreateDataGuardAssociationToExistingDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}
