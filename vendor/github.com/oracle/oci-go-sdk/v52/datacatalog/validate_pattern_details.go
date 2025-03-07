// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ValidatePatternDetails Validate pattern using the expression and file list.
type ValidatePatternDetails struct {

	// The expression used in the pattern that may include qualifiers. Refer to the user documentation for details of the format and examples.
	Expression *string `mandatory:"false" json:"expression"`

	// List of file paths against which the expression can be tried, as a check. This documents, for reference
	// purposes, some example objects a pattern is meant to work with.
	// If provided with the request,this overrides the list which already exists as part of the pattern, if any.
	CheckFilePathList []string `mandatory:"false" json:"checkFilePathList"`

	// The maximum number of UNMATCHED files, in checkFilePathList, above which the check fails.
	// Optional, if checkFilePathList is provided.
	// If provided with the request, this overrides the value which already exists as part of the pattern, if any.
	CheckFailureLimit *int `mandatory:"false" json:"checkFailureLimit"`
}

func (m ValidatePatternDetails) String() string {
	return common.PointerString(m)
}
