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

// RestoreAutonomousDatabaseDetails Details to restore an Oracle Autonomous Database.
type RestoreAutonomousDatabaseDetails struct {

	// The time to restore the database to.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Restores using the backup with the System Change Number (SCN) specified.
	DatabaseSCN *string `mandatory:"false" json:"databaseSCN"`

	// Restores to the last known good state with the least possible data loss.
	Latest *bool `mandatory:"false" json:"latest"`
}

func (m RestoreAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}
