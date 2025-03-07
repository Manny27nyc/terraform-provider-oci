// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// HostResourceStatistics Contains host resource base statistics.
type HostResourceStatistics interface {

	// Total amount used of the resource metric type (CPU, STORAGE).
	GetUsage() *float64

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	GetCapacity() *float64

	// Resource utilization in percentage.
	GetUtilizationPercent() *float64

	// Change in resource utilization in percentage
	GetUsageChangePercent() *float64
}

type hostresourcestatistics struct {
	JsonData           []byte
	Usage              *float64 `mandatory:"true" json:"usage"`
	Capacity           *float64 `mandatory:"true" json:"capacity"`
	UtilizationPercent *float64 `mandatory:"true" json:"utilizationPercent"`
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`
	ResourceName       string   `json:"resourceName"`
}

// UnmarshalJSON unmarshals json
func (m *hostresourcestatistics) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostresourcestatistics hostresourcestatistics
	s := struct {
		Model Unmarshalerhostresourcestatistics
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Usage = s.Model.Usage
	m.Capacity = s.Model.Capacity
	m.UtilizationPercent = s.Model.UtilizationPercent
	m.UsageChangePercent = s.Model.UsageChangePercent
	m.ResourceName = s.Model.ResourceName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostresourcestatistics) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResourceName {
	case "HOST_MEMORY_STATISTICS":
		mm := HostMemoryStatistics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_CPU_STATISTICS":
		mm := HostCpuStatistics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetUsage returns Usage
func (m hostresourcestatistics) GetUsage() *float64 {
	return m.Usage
}

//GetCapacity returns Capacity
func (m hostresourcestatistics) GetCapacity() *float64 {
	return m.Capacity
}

//GetUtilizationPercent returns UtilizationPercent
func (m hostresourcestatistics) GetUtilizationPercent() *float64 {
	return m.UtilizationPercent
}

//GetUsageChangePercent returns UsageChangePercent
func (m hostresourcestatistics) GetUsageChangePercent() *float64 {
	return m.UsageChangePercent
}

func (m hostresourcestatistics) String() string {
	return common.PointerString(m)
}

// HostResourceStatisticsResourceNameEnum Enum with underlying type: string
type HostResourceStatisticsResourceNameEnum string

// Set of constants representing the allowable values for HostResourceStatisticsResourceNameEnum
const (
	HostResourceStatisticsResourceNameCpuStatistics    HostResourceStatisticsResourceNameEnum = "HOST_CPU_STATISTICS"
	HostResourceStatisticsResourceNameMemoryStatistics HostResourceStatisticsResourceNameEnum = "HOST_MEMORY_STATISTICS"
)

var mappingHostResourceStatisticsResourceName = map[string]HostResourceStatisticsResourceNameEnum{
	"HOST_CPU_STATISTICS":    HostResourceStatisticsResourceNameCpuStatistics,
	"HOST_MEMORY_STATISTICS": HostResourceStatisticsResourceNameMemoryStatistics,
}

// GetHostResourceStatisticsResourceNameEnumValues Enumerates the set of values for HostResourceStatisticsResourceNameEnum
func GetHostResourceStatisticsResourceNameEnumValues() []HostResourceStatisticsResourceNameEnum {
	values := make([]HostResourceStatisticsResourceNameEnum, 0)
	for _, v := range mappingHostResourceStatisticsResourceName {
		values = append(values, v)
	}
	return values
}
