// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ScheduleKeyVersionDeletionDetails Details for scheduling key version deletion.
type ScheduleKeyVersionDeletionDetails struct {

	// An optional property to indicate when to delete the key version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format. The specified time must be between 7 and 30 days from the time
	// when the request is received. If this property is missing, it will be set to 30 days from the time of the request by default.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

func (m ScheduleKeyVersionDeletionDetails) String() string {
	return common.PointerString(m)
}
