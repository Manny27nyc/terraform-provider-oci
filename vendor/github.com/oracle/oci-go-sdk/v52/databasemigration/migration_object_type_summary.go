// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// MigrationObjectTypeSummary Migration Object Type
type MigrationObjectTypeSummary struct {

	// Object type name
	Name *string `mandatory:"true" json:"name"`
}

func (m MigrationObjectTypeSummary) String() string {
	return common.PointerString(m)
}
