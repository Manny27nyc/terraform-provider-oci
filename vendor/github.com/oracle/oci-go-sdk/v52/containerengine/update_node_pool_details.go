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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// UpdateNodePoolDetails The properties that define a request to update a node pool.
type UpdateNodePoolDetails struct {

	// The new name for the cluster. Avoid entering confidential information.
	Name *string `mandatory:"false" json:"name"`

	// The version of Kubernetes to which the nodes in the node pool should be upgraded.
	KubernetesVersion *string `mandatory:"false" json:"kubernetesVersion"`

	// A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	InitialNodeLabels []KeyValue `mandatory:"false" json:"initialNodeLabels"`

	// The number of nodes to have in each subnet specified in the subnetIds property. This property is deprecated,
	// use nodeConfigDetails instead. If the current value of quantityPerSubnet is greater than 0, you can only
	// use quantityPerSubnet to scale the node pool. If the current value of quantityPerSubnet is equal to 0 and
	// the current value of size in nodeConfigDetails is greater than 0, before you can use quantityPerSubnet,
	// you must first scale the node pool to 0 nodes using nodeConfigDetails.
	QuantityPerSubnet *int `mandatory:"false" json:"quantityPerSubnet"`

	// The OCIDs of the subnets in which to place nodes for this node pool. This property is deprecated,
	// use nodeConfigDetails instead. Only one of the subnetIds or nodeConfigDetails
	// properties can be specified.
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// The configuration of nodes in the node pool. Only one of the subnetIds or nodeConfigDetails
	// properties should be specified. If the current value of quantityPerSubnet is greater than 0, the node
	// pool may still be scaled using quantityPerSubnet. Before you can use nodeConfigDetails,
	// you must first scale the node pool to 0 nodes using quantityPerSubnet.
	NodeConfigDetails *UpdateNodePoolNodeConfigDetails `mandatory:"false" json:"nodeConfigDetails"`

	// A list of key/value pairs to add to each underlying OCI instance in the node pool on launch.
	NodeMetadata map[string]string `mandatory:"false" json:"nodeMetadata"`

	// Specify the source to use to launch nodes in the node pool. Currently, image is the only supported source.
	NodeSourceDetails NodeSourceDetails `mandatory:"false" json:"nodeSourceDetails"`

	// The SSH public key to add to each node in the node pool on launch.
	SshPublicKey *string `mandatory:"false" json:"sshPublicKey"`

	// The name of the node shape of the nodes in the node pool used on launch.
	NodeShape *string `mandatory:"false" json:"nodeShape"`

	// Specify the configuration of the shape to launch nodes in the node pool.
	NodeShapeConfig *UpdateNodeShapeConfigDetails `mandatory:"false" json:"nodeShapeConfig"`
}

func (m UpdateNodePoolDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateNodePoolDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name              *string                          `json:"name"`
		KubernetesVersion *string                          `json:"kubernetesVersion"`
		InitialNodeLabels []KeyValue                       `json:"initialNodeLabels"`
		QuantityPerSubnet *int                             `json:"quantityPerSubnet"`
		SubnetIds         []string                         `json:"subnetIds"`
		NodeConfigDetails *UpdateNodePoolNodeConfigDetails `json:"nodeConfigDetails"`
		NodeMetadata      map[string]string                `json:"nodeMetadata"`
		NodeSourceDetails nodesourcedetails                `json:"nodeSourceDetails"`
		SshPublicKey      *string                          `json:"sshPublicKey"`
		NodeShape         *string                          `json:"nodeShape"`
		NodeShapeConfig   *UpdateNodeShapeConfigDetails    `json:"nodeShapeConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.KubernetesVersion = model.KubernetesVersion

	m.InitialNodeLabels = make([]KeyValue, len(model.InitialNodeLabels))
	for i, n := range model.InitialNodeLabels {
		m.InitialNodeLabels[i] = n
	}

	m.QuantityPerSubnet = model.QuantityPerSubnet

	m.SubnetIds = make([]string, len(model.SubnetIds))
	for i, n := range model.SubnetIds {
		m.SubnetIds[i] = n
	}

	m.NodeConfigDetails = model.NodeConfigDetails

	m.NodeMetadata = model.NodeMetadata

	nn, e = model.NodeSourceDetails.UnmarshalPolymorphicJSON(model.NodeSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NodeSourceDetails = nn.(NodeSourceDetails)
	} else {
		m.NodeSourceDetails = nil
	}

	m.SshPublicKey = model.SshPublicKey

	m.NodeShape = model.NodeShape

	m.NodeShapeConfig = model.NodeShapeConfig

	return
}
