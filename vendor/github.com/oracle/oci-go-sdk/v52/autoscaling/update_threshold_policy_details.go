// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// UpdateThresholdPolicyDetails The representation of UpdateThresholdPolicyDetails
type UpdateThresholdPolicyDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The capacity requirements of the autoscaling policy.
	Capacity *Capacity `mandatory:"false" json:"capacity"`

	// Whether the autoscaling policy is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Rules []UpdateConditionDetails `mandatory:"false" json:"rules"`
}

//GetDisplayName returns DisplayName
func (m UpdateThresholdPolicyDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCapacity returns Capacity
func (m UpdateThresholdPolicyDetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetIsEnabled returns IsEnabled
func (m UpdateThresholdPolicyDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m UpdateThresholdPolicyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateThresholdPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateThresholdPolicyDetails UpdateThresholdPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeUpdateThresholdPolicyDetails
	}{
		"threshold",
		(MarshalTypeUpdateThresholdPolicyDetails)(m),
	}

	return json.Marshal(&s)
}
