// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// OperatorControlAssignmentSummary Details of the operator control assignment.
type OperatorControlAssignmentSummary struct {

	// The OCID of the operator control assignment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the operator control.
	OperatorControlId *string `mandatory:"true" json:"operatorControlId"`

	// The OCID of the target resource being governed by the operator control.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The OCID of the compartment that contains the operator control assignment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// resourceType for which the OperatorControlAssignment is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// The time at which the target resource will be brought under the governance of the operator control in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentFrom *common.SDKTime `mandatory:"false" json:"timeAssignmentFrom"`

	// The time at which the target resource will leave the governance of the operator control in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeAssignmentTo *common.SDKTime `mandatory:"false" json:"timeAssignmentTo"`

	// If true, then the target resource is always governed by the operator control. Otherwise governance is time-based as specified by timeAssignmentTo and timeAssignmentFrom.
	IsEnforcedAlways *bool `mandatory:"false" json:"isEnforcedAlways"`

	// Time when the operator control assignment is created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfAssignment *common.SDKTime `mandatory:"false" json:"timeOfAssignment"`

	// The code identifying the error occurred during Assignment operation.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// The message describing the error occurred during Assignment operation.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// If set, then the audit logs are being forwarded to the relevant remote logging server
	IsLogForwarded *bool `mandatory:"false" json:"isLogForwarded"`

	// The address of the remote syslog server where the audit logs are being forwarded to. Address in host or IP format.
	RemoteSyslogServerAddress *string `mandatory:"false" json:"remoteSyslogServerAddress"`

	// The listening port of the remote syslog server. The port range is 0 - 65535.
	RemoteSyslogServerPort *int `mandatory:"false" json:"remoteSyslogServerPort"`

	// The current lifcycle state of the OperatorControl.
	LifecycleState OperatorControlAssignmentLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OperatorControlAssignmentSummary) String() string {
	return common.PointerString(m)
}
