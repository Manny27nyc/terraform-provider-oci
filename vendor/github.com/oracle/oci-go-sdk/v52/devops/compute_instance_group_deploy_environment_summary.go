// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ComputeInstanceGroupDeployEnvironmentSummary Specifies the Compute instance group environment.
type ComputeInstanceGroupDeployEnvironmentSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ComputeInstanceGroupSelectors *ComputeInstanceGroupSelectorCollection `mandatory:"true" json:"computeInstanceGroupSelectors"`

	// Optional description about the deployment environment.
	Description *string `mandatory:"false" json:"description"`

	// Deployment environment display name, which can be renamed and is not necessarily unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment environment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment environment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the deployment environment.
	LifecycleState DeployEnvironmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetId() *string {
	return m.Id
}

//GetDescription returns Description
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetLifecycleState() DeployEnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m ComputeInstanceGroupDeployEnvironmentSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ComputeInstanceGroupDeployEnvironmentSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupDeployEnvironmentSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupDeployEnvironmentSummary ComputeInstanceGroupDeployEnvironmentSummary
	s := struct {
		DiscriminatorParam string `json:"deployEnvironmentType"`
		MarshalTypeComputeInstanceGroupDeployEnvironmentSummary
	}{
		"COMPUTE_INSTANCE_GROUP",
		(MarshalTypeComputeInstanceGroupDeployEnvironmentSummary)(m),
	}

	return json.Marshal(&s)
}
