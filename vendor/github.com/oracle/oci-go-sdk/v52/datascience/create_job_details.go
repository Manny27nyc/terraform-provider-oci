// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateJobDetails Parameters needed to create a new job.
type CreateJobDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	JobConfigurationDetails JobConfigurationDetails `mandatory:"true" json:"jobConfigurationDetails"`

	JobInfrastructureConfigurationDetails JobInfrastructureConfigurationDetails `mandatory:"true" json:"jobInfrastructureConfigurationDetails"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the job.
	Description *string `mandatory:"false" json:"description"`

	JobLogConfigurationDetails *JobLogConfigurationDetails `mandatory:"false" json:"jobLogConfigurationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                           *string                               `json:"displayName"`
		Description                           *string                               `json:"description"`
		JobLogConfigurationDetails            *JobLogConfigurationDetails           `json:"jobLogConfigurationDetails"`
		FreeformTags                          map[string]string                     `json:"freeformTags"`
		DefinedTags                           map[string]map[string]interface{}     `json:"definedTags"`
		ProjectId                             *string                               `json:"projectId"`
		CompartmentId                         *string                               `json:"compartmentId"`
		JobConfigurationDetails               jobconfigurationdetails               `json:"jobConfigurationDetails"`
		JobInfrastructureConfigurationDetails jobinfrastructureconfigurationdetails `json:"jobInfrastructureConfigurationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.JobLogConfigurationDetails = model.JobLogConfigurationDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	nn, e = model.JobConfigurationDetails.UnmarshalPolymorphicJSON(model.JobConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobConfigurationDetails = nn.(JobConfigurationDetails)
	} else {
		m.JobConfigurationDetails = nil
	}

	nn, e = model.JobInfrastructureConfigurationDetails.UnmarshalPolymorphicJSON(model.JobInfrastructureConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobInfrastructureConfigurationDetails = nn.(JobInfrastructureConfigurationDetails)
	} else {
		m.JobInfrastructureConfigurationDetails = nil
	}

	return
}
