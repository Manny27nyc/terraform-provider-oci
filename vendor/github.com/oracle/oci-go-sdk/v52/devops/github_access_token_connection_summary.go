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

// GithubAccessTokenConnectionSummary Summary information for a connection of the type `GITHUB_ACCESS_TOKEN`.
// This type corresponds to a connection in GitHub that is authenticated with a personal access token.
type GithubAccessTokenConnectionSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Project Identifier
	ProjectId *string `mandatory:"true" json:"projectId"`

	// OCID of personal access token saved in secret store
	AccessToken *string `mandatory:"true" json:"accessToken"`

	// Connection identifier which can be renamed and is not necessarily unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the connection
	Description *string `mandatory:"false" json:"description"`

	// The time the the connection was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the connection was updated. An RFC3339 formatted daservertetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m GithubAccessTokenConnectionSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m GithubAccessTokenConnectionSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m GithubAccessTokenConnectionSummary) GetDescription() *string {
	return m.Description
}

//GetCompartmentId returns CompartmentId
func (m GithubAccessTokenConnectionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetProjectId returns ProjectId
func (m GithubAccessTokenConnectionSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetTimeCreated returns TimeCreated
func (m GithubAccessTokenConnectionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m GithubAccessTokenConnectionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m GithubAccessTokenConnectionSummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m GithubAccessTokenConnectionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GithubAccessTokenConnectionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m GithubAccessTokenConnectionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GithubAccessTokenConnectionSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubAccessTokenConnectionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubAccessTokenConnectionSummary GithubAccessTokenConnectionSummary
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeGithubAccessTokenConnectionSummary
	}{
		"GITHUB_ACCESS_TOKEN",
		(MarshalTypeGithubAccessTokenConnectionSummary)(m),
	}

	return json.Marshal(&s)
}
