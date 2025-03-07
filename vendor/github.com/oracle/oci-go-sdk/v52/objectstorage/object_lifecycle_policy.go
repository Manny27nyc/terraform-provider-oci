// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ObjectLifecyclePolicy The collection of lifecycle policy rules that together form the object lifecycle policy of a given bucket.
type ObjectLifecyclePolicy struct {

	// The date and time the object lifecycle policy was created, as described in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The live lifecycle policy on the bucket.
	// For an example of this value, see the
	// PutObjectLifecyclePolicy API documentation (https://docs.cloud.oracle.com/iaas/api/#/en/objectstorage/20160918/ObjectLifecyclePolicy/PutObjectLifecyclePolicy).
	Items []ObjectLifecycleRule `mandatory:"false" json:"items"`
}

func (m ObjectLifecyclePolicy) String() string {
	return common.PointerString(m)
}
