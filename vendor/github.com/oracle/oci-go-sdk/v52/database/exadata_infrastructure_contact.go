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

// ExadataInfrastructureContact Contact details for Exadata Infrastructure.
type ExadataInfrastructureContact struct {

	// The name of the Exadata Infrastructure contact.
	Name *string `mandatory:"true" json:"name"`

	// The email for the Exadata Infrastructure contact.
	Email *string `mandatory:"true" json:"email"`

	// If `true`, this Exadata Infrastructure contact is a primary contact. If `false`, this Exadata Infrastructure is a secondary contact.
	IsPrimary *bool `mandatory:"true" json:"isPrimary"`

	// The phone number for the Exadata Infrastructure contact.
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// If `true`, this Exadata Infrastructure contact is a valid My Oracle Support (MOS) contact. If `false`, this Exadata Infrastructure contact is not a valid MOS contact.
	IsContactMosValidated *bool `mandatory:"false" json:"isContactMosValidated"`
}

func (m ExadataInfrastructureContact) String() string {
	return common.PointerString(m)
}
