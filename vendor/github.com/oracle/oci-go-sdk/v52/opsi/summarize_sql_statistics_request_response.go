// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// SummarizeSqlStatisticsRequest wrapper for the SummarizeSqlStatistics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatistics.go.html to see an example of how to use SummarizeSqlStatisticsRequest.
type SummarizeSqlStatisticsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []SummarizeSqlStatisticsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Filter sqls by percentage of db time.
	DatabaseTimePctGreaterThan *float64 `mandatory:"false" contributesTo:"query" name:"databaseTimePctGreaterThan"`

	// One or more unique SQL_IDs for a SQL Statement.
	// Example: `6rgjh9bjmy2s7`
	SqlIdentifier []string `contributesTo:"query" name:"sqlIdentifier" collectionFormat:"multi"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeSqlStatisticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to use when sorting SQL statistics.
	// Example: databaseTimeInSec
	SortBy SummarizeSqlStatisticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter sqls by one or more performance categories.
	Category []SummarizeSqlStatisticsCategoryEnum `contributesTo:"query" name:"category" omitEmpty:"true" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeSqlStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeSqlStatisticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeSqlStatisticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeSqlStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeSqlStatisticsResponse wrapper for the SummarizeSqlStatistics operation
type SummarizeSqlStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlStatisticAggregationCollection instances
	SqlStatisticAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeSqlStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeSqlStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeSqlStatisticsDatabaseTypeEnum Enum with underlying type: string
type SummarizeSqlStatisticsDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsDatabaseTypeEnum
const (
	SummarizeSqlStatisticsDatabaseTypeAdwS           SummarizeSqlStatisticsDatabaseTypeEnum = "ADW-S"
	SummarizeSqlStatisticsDatabaseTypeAtpS           SummarizeSqlStatisticsDatabaseTypeEnum = "ATP-S"
	SummarizeSqlStatisticsDatabaseTypeAdwD           SummarizeSqlStatisticsDatabaseTypeEnum = "ADW-D"
	SummarizeSqlStatisticsDatabaseTypeAtpD           SummarizeSqlStatisticsDatabaseTypeEnum = "ATP-D"
	SummarizeSqlStatisticsDatabaseTypeExternalPdb    SummarizeSqlStatisticsDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeSqlStatisticsDatabaseTypeExternalNoncdb SummarizeSqlStatisticsDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingSummarizeSqlStatisticsDatabaseType = map[string]SummarizeSqlStatisticsDatabaseTypeEnum{
	"ADW-S":           SummarizeSqlStatisticsDatabaseTypeAdwS,
	"ATP-S":           SummarizeSqlStatisticsDatabaseTypeAtpS,
	"ADW-D":           SummarizeSqlStatisticsDatabaseTypeAdwD,
	"ATP-D":           SummarizeSqlStatisticsDatabaseTypeAtpD,
	"EXTERNAL-PDB":    SummarizeSqlStatisticsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": SummarizeSqlStatisticsDatabaseTypeExternalNoncdb,
}

// GetSummarizeSqlStatisticsDatabaseTypeEnumValues Enumerates the set of values for SummarizeSqlStatisticsDatabaseTypeEnum
func GetSummarizeSqlStatisticsDatabaseTypeEnumValues() []SummarizeSqlStatisticsDatabaseTypeEnum {
	values := make([]SummarizeSqlStatisticsDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsDatabaseType {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsSortOrderEnum Enum with underlying type: string
type SummarizeSqlStatisticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsSortOrderEnum
const (
	SummarizeSqlStatisticsSortOrderAsc  SummarizeSqlStatisticsSortOrderEnum = "ASC"
	SummarizeSqlStatisticsSortOrderDesc SummarizeSqlStatisticsSortOrderEnum = "DESC"
)

var mappingSummarizeSqlStatisticsSortOrder = map[string]SummarizeSqlStatisticsSortOrderEnum{
	"ASC":  SummarizeSqlStatisticsSortOrderAsc,
	"DESC": SummarizeSqlStatisticsSortOrderDesc,
}

// GetSummarizeSqlStatisticsSortOrderEnumValues Enumerates the set of values for SummarizeSqlStatisticsSortOrderEnum
func GetSummarizeSqlStatisticsSortOrderEnumValues() []SummarizeSqlStatisticsSortOrderEnum {
	values := make([]SummarizeSqlStatisticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsSortByEnum Enum with underlying type: string
type SummarizeSqlStatisticsSortByEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsSortByEnum
const (
	SummarizeSqlStatisticsSortByDatabasetimeinsec                  SummarizeSqlStatisticsSortByEnum = "databaseTimeInSec"
	SummarizeSqlStatisticsSortByExecutionsperhour                  SummarizeSqlStatisticsSortByEnum = "executionsPerHour"
	SummarizeSqlStatisticsSortByExecutionscount                    SummarizeSqlStatisticsSortByEnum = "executionsCount"
	SummarizeSqlStatisticsSortByCputimeinsec                       SummarizeSqlStatisticsSortByEnum = "cpuTimeInSec"
	SummarizeSqlStatisticsSortByIotimeinsec                        SummarizeSqlStatisticsSortByEnum = "ioTimeInSec"
	SummarizeSqlStatisticsSortByInefficientwaittimeinsec           SummarizeSqlStatisticsSortByEnum = "inefficientWaitTimeInSec"
	SummarizeSqlStatisticsSortByResponsetimeinsec                  SummarizeSqlStatisticsSortByEnum = "responseTimeInSec"
	SummarizeSqlStatisticsSortByPlancount                          SummarizeSqlStatisticsSortByEnum = "planCount"
	SummarizeSqlStatisticsSortByVariability                        SummarizeSqlStatisticsSortByEnum = "variability"
	SummarizeSqlStatisticsSortByAverageactivesessions              SummarizeSqlStatisticsSortByEnum = "averageActiveSessions"
	SummarizeSqlStatisticsSortByDatabasetimepct                    SummarizeSqlStatisticsSortByEnum = "databaseTimePct"
	SummarizeSqlStatisticsSortByInefficiencyinpct                  SummarizeSqlStatisticsSortByEnum = "inefficiencyInPct"
	SummarizeSqlStatisticsSortByChangeincputimeinpct               SummarizeSqlStatisticsSortByEnum = "changeInCpuTimeInPct"
	SummarizeSqlStatisticsSortByChangeiniotimeinpct                SummarizeSqlStatisticsSortByEnum = "changeInIoTimeInPct"
	SummarizeSqlStatisticsSortByChangeininefficientwaittimeinpct   SummarizeSqlStatisticsSortByEnum = "changeInInefficientWaitTimeInPct"
	SummarizeSqlStatisticsSortByChangeinresponsetimeinpct          SummarizeSqlStatisticsSortByEnum = "changeInResponseTimeInPct"
	SummarizeSqlStatisticsSortByChangeinaverageactivesessionsinpct SummarizeSqlStatisticsSortByEnum = "changeInAverageActiveSessionsInPct"
	SummarizeSqlStatisticsSortByChangeinexecutionsperhourinpct     SummarizeSqlStatisticsSortByEnum = "changeInExecutionsPerHourInPct"
	SummarizeSqlStatisticsSortByChangeininefficiencyinpct          SummarizeSqlStatisticsSortByEnum = "changeInInefficiencyInPct"
)

var mappingSummarizeSqlStatisticsSortBy = map[string]SummarizeSqlStatisticsSortByEnum{
	"databaseTimeInSec":                  SummarizeSqlStatisticsSortByDatabasetimeinsec,
	"executionsPerHour":                  SummarizeSqlStatisticsSortByExecutionsperhour,
	"executionsCount":                    SummarizeSqlStatisticsSortByExecutionscount,
	"cpuTimeInSec":                       SummarizeSqlStatisticsSortByCputimeinsec,
	"ioTimeInSec":                        SummarizeSqlStatisticsSortByIotimeinsec,
	"inefficientWaitTimeInSec":           SummarizeSqlStatisticsSortByInefficientwaittimeinsec,
	"responseTimeInSec":                  SummarizeSqlStatisticsSortByResponsetimeinsec,
	"planCount":                          SummarizeSqlStatisticsSortByPlancount,
	"variability":                        SummarizeSqlStatisticsSortByVariability,
	"averageActiveSessions":              SummarizeSqlStatisticsSortByAverageactivesessions,
	"databaseTimePct":                    SummarizeSqlStatisticsSortByDatabasetimepct,
	"inefficiencyInPct":                  SummarizeSqlStatisticsSortByInefficiencyinpct,
	"changeInCpuTimeInPct":               SummarizeSqlStatisticsSortByChangeincputimeinpct,
	"changeInIoTimeInPct":                SummarizeSqlStatisticsSortByChangeiniotimeinpct,
	"changeInInefficientWaitTimeInPct":   SummarizeSqlStatisticsSortByChangeininefficientwaittimeinpct,
	"changeInResponseTimeInPct":          SummarizeSqlStatisticsSortByChangeinresponsetimeinpct,
	"changeInAverageActiveSessionsInPct": SummarizeSqlStatisticsSortByChangeinaverageactivesessionsinpct,
	"changeInExecutionsPerHourInPct":     SummarizeSqlStatisticsSortByChangeinexecutionsperhourinpct,
	"changeInInefficiencyInPct":          SummarizeSqlStatisticsSortByChangeininefficiencyinpct,
}

// GetSummarizeSqlStatisticsSortByEnumValues Enumerates the set of values for SummarizeSqlStatisticsSortByEnum
func GetSummarizeSqlStatisticsSortByEnumValues() []SummarizeSqlStatisticsSortByEnum {
	values := make([]SummarizeSqlStatisticsSortByEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeSqlStatisticsCategoryEnum Enum with underlying type: string
type SummarizeSqlStatisticsCategoryEnum string

// Set of constants representing the allowable values for SummarizeSqlStatisticsCategoryEnum
const (
	SummarizeSqlStatisticsCategoryDegrading                                            SummarizeSqlStatisticsCategoryEnum = "DEGRADING"
	SummarizeSqlStatisticsCategoryVariant                                              SummarizeSqlStatisticsCategoryEnum = "VARIANT"
	SummarizeSqlStatisticsCategoryInefficient                                          SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT"
	SummarizeSqlStatisticsCategoryChangingPlans                                        SummarizeSqlStatisticsCategoryEnum = "CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryImproving                                            SummarizeSqlStatisticsCategoryEnum = "IMPROVING"
	SummarizeSqlStatisticsCategoryDegradingVariant                                     SummarizeSqlStatisticsCategoryEnum = "DEGRADING_VARIANT"
	SummarizeSqlStatisticsCategoryDegradingInefficient                                 SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INEFFICIENT"
	SummarizeSqlStatisticsCategoryDegradingChangingPlans                               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryDegradingIncreasingIo                                SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_IO"
	SummarizeSqlStatisticsCategoryDegradingIncreasingCpu                               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryDegradingIncreasingInefficientWait                   SummarizeSqlStatisticsCategoryEnum = "DEGRADING_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingIo                SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_IO"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingCpu               SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingInefficientWait   SummarizeSqlStatisticsCategoryEnum = "DEGRADING_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryVariantInefficient                                   SummarizeSqlStatisticsCategoryEnum = "VARIANT_INEFFICIENT"
	SummarizeSqlStatisticsCategoryVariantChangingPlans                                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryVariantIncreasingIo                                  SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_IO"
	SummarizeSqlStatisticsCategoryVariantIncreasingCpu                                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryVariantIncreasingInefficientWait                     SummarizeSqlStatisticsCategoryEnum = "VARIANT_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingIo                  SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_IO"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingCpu                 SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_CPU"
	SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingInefficientWait     SummarizeSqlStatisticsCategoryEnum = "VARIANT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryInefficientChangingPlans                             SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_CHANGING_PLANS"
	SummarizeSqlStatisticsCategoryInefficientIncreasingInefficientWait                 SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_INCREASING_INEFFICIENT_WAIT"
	SummarizeSqlStatisticsCategoryInefficientChangingPlansAndIncreasingInefficientWait SummarizeSqlStatisticsCategoryEnum = "INEFFICIENT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT"
)

var mappingSummarizeSqlStatisticsCategory = map[string]SummarizeSqlStatisticsCategoryEnum{
	"DEGRADING":                                   SummarizeSqlStatisticsCategoryDegrading,
	"VARIANT":                                     SummarizeSqlStatisticsCategoryVariant,
	"INEFFICIENT":                                 SummarizeSqlStatisticsCategoryInefficient,
	"CHANGING_PLANS":                              SummarizeSqlStatisticsCategoryChangingPlans,
	"IMPROVING":                                   SummarizeSqlStatisticsCategoryImproving,
	"DEGRADING_VARIANT":                           SummarizeSqlStatisticsCategoryDegradingVariant,
	"DEGRADING_INEFFICIENT":                       SummarizeSqlStatisticsCategoryDegradingInefficient,
	"DEGRADING_CHANGING_PLANS":                    SummarizeSqlStatisticsCategoryDegradingChangingPlans,
	"DEGRADING_INCREASING_IO":                     SummarizeSqlStatisticsCategoryDegradingIncreasingIo,
	"DEGRADING_INCREASING_CPU":                    SummarizeSqlStatisticsCategoryDegradingIncreasingCpu,
	"DEGRADING_INCREASING_INEFFICIENT_WAIT":       SummarizeSqlStatisticsCategoryDegradingIncreasingInefficientWait,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_IO":  SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingIo,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_CPU": SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingCpu,
	"DEGRADING_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT": SummarizeSqlStatisticsCategoryDegradingChangingPlansAndIncreasingInefficientWait,
	"VARIANT_INEFFICIENT":                                        SummarizeSqlStatisticsCategoryVariantInefficient,
	"VARIANT_CHANGING_PLANS":                                     SummarizeSqlStatisticsCategoryVariantChangingPlans,
	"VARIANT_INCREASING_IO":                                      SummarizeSqlStatisticsCategoryVariantIncreasingIo,
	"VARIANT_INCREASING_CPU":                                     SummarizeSqlStatisticsCategoryVariantIncreasingCpu,
	"VARIANT_INCREASING_INEFFICIENT_WAIT":                        SummarizeSqlStatisticsCategoryVariantIncreasingInefficientWait,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_IO":                   SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingIo,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_CPU":                  SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingCpu,
	"VARIANT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT":     SummarizeSqlStatisticsCategoryVariantChangingPlansAndIncreasingInefficientWait,
	"INEFFICIENT_CHANGING_PLANS":                                 SummarizeSqlStatisticsCategoryInefficientChangingPlans,
	"INEFFICIENT_INCREASING_INEFFICIENT_WAIT":                    SummarizeSqlStatisticsCategoryInefficientIncreasingInefficientWait,
	"INEFFICIENT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT": SummarizeSqlStatisticsCategoryInefficientChangingPlansAndIncreasingInefficientWait,
}

// GetSummarizeSqlStatisticsCategoryEnumValues Enumerates the set of values for SummarizeSqlStatisticsCategoryEnum
func GetSummarizeSqlStatisticsCategoryEnumValues() []SummarizeSqlStatisticsCategoryEnum {
	values := make([]SummarizeSqlStatisticsCategoryEnum, 0)
	for _, v := range mappingSummarizeSqlStatisticsCategory {
		values = append(values, v)
	}
	return values
}
