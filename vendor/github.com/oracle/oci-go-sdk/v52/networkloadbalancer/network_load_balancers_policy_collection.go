// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// NetworkLoadBalancersPolicyCollection Wrapper object for array of NetworkLoadBalancersPolicySummary objects.
type NetworkLoadBalancersPolicyCollection struct {

	// Array of NetworkLoadBalancersPolicySummary objects.
	Items []NetworkLoadBalancersPolicySummaryEnum `mandatory:"false" json:"items"`
}

func (m NetworkLoadBalancersPolicyCollection) String() string {
	return common.PointerString(m)
}
