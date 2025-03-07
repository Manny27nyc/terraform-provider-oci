// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// WorkRequest Description of the work request status.
type WorkRequest struct {

	// The ID of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request. Work requests should be scoped to the same compartment as the resource the work request affects. If the work request affects multiple resources, and those resources are not in the same compartment, it is up to the service team to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of this work request.
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// The status of this work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of this work request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, shown as an RFC 3339 formatted datetime string.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The time the request was started, shown as an RFC 3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the object was finished, shown as an RFC 3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}
