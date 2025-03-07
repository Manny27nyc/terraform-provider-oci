// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PemEncodedPublicKey A PEM-encoded public key used for verifying the JWT signature.
type PemEncodedPublicKey struct {

	// A unique key ID. This key will be used to verify the signature of a
	// JWT with matching "kid".
	Kid *string `mandatory:"true" json:"kid"`

	// The content of the PEM-encoded public key.
	Key *string `mandatory:"true" json:"key"`
}

//GetKid returns Kid
func (m PemEncodedPublicKey) GetKid() *string {
	return m.Kid
}

func (m PemEncodedPublicKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PemEncodedPublicKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePemEncodedPublicKey PemEncodedPublicKey
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypePemEncodedPublicKey
	}{
		"PEM",
		(MarshalTypePemEncodedPublicKey)(m),
	}

	return json.Marshal(&s)
}
