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

// NetworkLoadBalancersProtocolCollection Wrapper object for array of ProtocolSummary objects.
type NetworkLoadBalancersProtocolCollection struct {

	// Array of NetworkLoadBalancersProtocolSummary objects.
	Items []NetworkLoadBalancersProtocolSummaryEnum `mandatory:"false" json:"items"`
}

func (m NetworkLoadBalancersProtocolCollection) String() string {
	return common.PointerString(m)
}
