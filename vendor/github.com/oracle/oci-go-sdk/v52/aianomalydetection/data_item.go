// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// DataItem Simple object representing signal values at a certain point in time.
type DataItem struct {

	// Array of double precision values.
	Values []float64 `mandatory:"true" json:"values"`

	// Nullable string representing timestamp.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m DataItem) String() string {
	return common.PointerString(m)
}
