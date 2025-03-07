// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ChangeProjectCompartmentDetails Details for changing the compartment of a project.
type ChangeProjectCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeProjectCompartmentDetails) String() string {
	return common.PointerString(m)
}
