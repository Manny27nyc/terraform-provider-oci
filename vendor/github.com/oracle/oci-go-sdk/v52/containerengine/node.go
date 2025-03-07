// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// Node The properties that define a node.
type Node struct {

	// The OCID of the compute instance backing this node.
	Id *string `mandatory:"false" json:"id"`

	// The name of the node.
	Name *string `mandatory:"false" json:"name"`

	// The version of Kubernetes this node is running.
	KubernetesVersion *string `mandatory:"false" json:"kubernetesVersion"`

	// The name of the availability domain in which this node is placed.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the subnet in which this node is placed.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID of the node pool to which this node belongs.
	NodePoolId *string `mandatory:"false" json:"nodePoolId"`

	// The fault domain of this node.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The private IP address of this node.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The public IP address of this node.
	PublicIp *string `mandatory:"false" json:"publicIp"`

	// An error that may be associated with the node.
	NodeError *NodeError `mandatory:"false" json:"nodeError"`

	// The state of the node.
	LifecycleState NodeLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the state of the node.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m Node) String() string {
	return common.PointerString(m)
}

// NodeLifecycleStateEnum Enum with underlying type: string
type NodeLifecycleStateEnum string

// Set of constants representing the allowable values for NodeLifecycleStateEnum
const (
	NodeLifecycleStateCreating NodeLifecycleStateEnum = "CREATING"
	NodeLifecycleStateActive   NodeLifecycleStateEnum = "ACTIVE"
	NodeLifecycleStateUpdating NodeLifecycleStateEnum = "UPDATING"
	NodeLifecycleStateDeleting NodeLifecycleStateEnum = "DELETING"
	NodeLifecycleStateDeleted  NodeLifecycleStateEnum = "DELETED"
	NodeLifecycleStateFailing  NodeLifecycleStateEnum = "FAILING"
	NodeLifecycleStateInactive NodeLifecycleStateEnum = "INACTIVE"
)

var mappingNodeLifecycleState = map[string]NodeLifecycleStateEnum{
	"CREATING": NodeLifecycleStateCreating,
	"ACTIVE":   NodeLifecycleStateActive,
	"UPDATING": NodeLifecycleStateUpdating,
	"DELETING": NodeLifecycleStateDeleting,
	"DELETED":  NodeLifecycleStateDeleted,
	"FAILING":  NodeLifecycleStateFailing,
	"INACTIVE": NodeLifecycleStateInactive,
}

// GetNodeLifecycleStateEnumValues Enumerates the set of values for NodeLifecycleStateEnum
func GetNodeLifecycleStateEnumValues() []NodeLifecycleStateEnum {
	values := make([]NodeLifecycleStateEnum, 0)
	for _, v := range mappingNodeLifecycleState {
		values = append(values, v)
	}
	return values
}
