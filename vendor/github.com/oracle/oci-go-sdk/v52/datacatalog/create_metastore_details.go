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

// CreateMetastoreDetails Information about a new metastore.
type CreateMetastoreDetails struct {

	// OCID of the compartment which holds the metastore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Location under which managed tables will be created by default. This references Object Storage
	// using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/
	DefaultManagedTableLocation *string `mandatory:"true" json:"defaultManagedTableLocation"`

	// Location under which external tables will be created by default. This references Object Storage
	// using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/
	DefaultExternalTableLocation *string `mandatory:"true" json:"defaultExternalTableLocation"`

	// Mutable name of the metastore.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMetastoreDetails) String() string {
	return common.PointerString(m)
}
