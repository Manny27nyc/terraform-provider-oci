// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// PlanJobOperationDetailsSummary Job details that are specific to plan operations.
type PlanJobOperationDetailsSummary struct {
}

func (m PlanJobOperationDetailsSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PlanJobOperationDetailsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePlanJobOperationDetailsSummary PlanJobOperationDetailsSummary
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePlanJobOperationDetailsSummary
	}{
		"PLAN",
		(MarshalTypePlanJobOperationDetailsSummary)(m),
	}

	return json.Marshal(&s)
}
