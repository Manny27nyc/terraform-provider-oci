// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PreparedStatement The result of query preparation.
type PreparedStatement struct {

	// A base64-encoded, compiled and parameterized version of
	// a SQL statement.
	Statement *string `mandatory:"false" json:"statement"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m PreparedStatement) String() string {
	return common.PointerString(m)
}
