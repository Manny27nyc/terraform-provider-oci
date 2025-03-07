// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// AnalyticsClusterSchemaMemoryEstimate DEPRECATED -- please use HeatWave API instead.
// Schema with estimated memory footprints for each MySQL user table
// of the schema when loaded to Analytics Cluster memory.
type AnalyticsClusterSchemaMemoryEstimate struct {

	// The name of the schema.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// Estimated memory footprints for MySQL user tables of the schema
	// when loaded to Analytics Cluster memory.
	PerTableEstimates []AnalyticsClusterTableMemoryEstimate `mandatory:"true" json:"perTableEstimates"`
}

func (m AnalyticsClusterSchemaMemoryEstimate) String() string {
	return common.PointerString(m)
}
