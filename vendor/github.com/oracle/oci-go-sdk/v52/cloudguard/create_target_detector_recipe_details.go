// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateTargetDetectorRecipeDetails The information required to create TargetDetectorRecipe
type CreateTargetDetectorRecipeDetails struct {

	// Identifier for DetectorRecipe.
	DetectorRecipeId *string `mandatory:"true" json:"detectorRecipeId"`

	// Overrides to be applied to Detector Rule associated with the target
	DetectorRules []UpdateTargetRecipeDetectorRuleDetails `mandatory:"false" json:"detectorRules"`
}

func (m CreateTargetDetectorRecipeDetails) String() string {
	return common.PointerString(m)
}
