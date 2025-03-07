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

// LocalClonePluggableDatabaseDetails Parameters for cloning a pluggable database (PDB) within the same database (CDB).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type LocalClonePluggableDatabaseDetails struct {

	// The name for the pluggable database (PDB). The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName *string `mandatory:"true" json:"clonedPdbName"`

	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword *string `mandatory:"false" json:"pdbAdminPassword"`

	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword *string `mandatory:"false" json:"targetTdeWalletPassword"`

	// The locked mode of the pluggable database admin account. If false, the user needs to provide the PDB Admin Password to connect to it.
	// If true, the pluggable database will be locked and user cannot login to it.
	ShouldPdbAdminAccountBeLocked *bool `mandatory:"false" json:"shouldPdbAdminAccountBeLocked"`
}

func (m LocalClonePluggableDatabaseDetails) String() string {
	return common.PointerString(m)
}
