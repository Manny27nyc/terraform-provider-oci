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

// AutomatedMountDetails Used for creating NFS Auto Mount backup destinations for autonomous on ExaCC.
type AutomatedMountDetails struct {

	// IP addresses for NFS Auto mount.
	NfsServer []string `mandatory:"true" json:"nfsServer"`

	// Specifies the directory on which to mount the file system
	NfsServerExport *string `mandatory:"true" json:"nfsServerExport"`
}

func (m AutomatedMountDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AutomatedMountDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutomatedMountDetails AutomatedMountDetails
	s := struct {
		DiscriminatorParam string `json:"mountType"`
		MarshalTypeAutomatedMountDetails
	}{
		"AUTOMATED_MOUNT",
		(MarshalTypeAutomatedMountDetails)(m),
	}

	return json.Marshal(&s)
}
