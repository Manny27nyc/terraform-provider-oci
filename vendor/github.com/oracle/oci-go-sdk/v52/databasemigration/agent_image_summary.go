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

// AgentImageSummary Available ODMS Agent Images.
type AgentImageSummary struct {

	// ODMS Agent Image version.
	Version *string `mandatory:"true" json:"version"`

	// URL to download Agent Image of the ODMS Agent.
	DownloadUrl *string `mandatory:"true" json:"downloadUrl"`
}

func (m AgentImageSummary) String() string {
	return common.PointerString(m)
}
