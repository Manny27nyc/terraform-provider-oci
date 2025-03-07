// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListLogAnalyticsObjectCollectionRulesRequest wrapper for the ListLogAnalyticsObjectCollectionRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsObjectCollectionRules.go.html to see an example of how to use ListLogAnalyticsObjectCollectionRulesRequest.
type ListLogAnalyticsObjectCollectionRulesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return rules only matching with this name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Lifecycle state filter.
	LifecycleState ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsObjectCollectionRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeUpdated is descending.
	// Default order for name is ascending. If no value is specified timeUpdated is default.
	SortBy ListLogAnalyticsObjectCollectionRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsObjectCollectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsObjectCollectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsObjectCollectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsObjectCollectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsObjectCollectionRulesResponse wrapper for the ListLogAnalyticsObjectCollectionRules operation
type ListLogAnalyticsObjectCollectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsObjectCollectionRuleCollection instances
	LogAnalyticsObjectCollectionRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsObjectCollectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsObjectCollectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum
const (
	ListLogAnalyticsObjectCollectionRulesLifecycleStateActive  ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsObjectCollectionRulesLifecycleStateDeleted ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsObjectCollectionRulesLifecycleState = map[string]ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsObjectCollectionRulesLifecycleStateActive,
	"DELETED": ListLogAnalyticsObjectCollectionRulesLifecycleStateDeleted,
}

// GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum
func GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumValues() []ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsObjectCollectionRulesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesSortOrderEnum
const (
	ListLogAnalyticsObjectCollectionRulesSortOrderAsc  ListLogAnalyticsObjectCollectionRulesSortOrderEnum = "ASC"
	ListLogAnalyticsObjectCollectionRulesSortOrderDesc ListLogAnalyticsObjectCollectionRulesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsObjectCollectionRulesSortOrder = map[string]ListLogAnalyticsObjectCollectionRulesSortOrderEnum{
	"ASC":  ListLogAnalyticsObjectCollectionRulesSortOrderAsc,
	"DESC": ListLogAnalyticsObjectCollectionRulesSortOrderDesc,
}

// GetListLogAnalyticsObjectCollectionRulesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesSortOrderEnum
func GetListLogAnalyticsObjectCollectionRulesSortOrderEnumValues() []ListLogAnalyticsObjectCollectionRulesSortOrderEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsObjectCollectionRulesSortByEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesSortByEnum
const (
	ListLogAnalyticsObjectCollectionRulesSortByTimeupdated ListLogAnalyticsObjectCollectionRulesSortByEnum = "timeUpdated"
	ListLogAnalyticsObjectCollectionRulesSortByTimecreated ListLogAnalyticsObjectCollectionRulesSortByEnum = "timeCreated"
	ListLogAnalyticsObjectCollectionRulesSortByName        ListLogAnalyticsObjectCollectionRulesSortByEnum = "name"
)

var mappingListLogAnalyticsObjectCollectionRulesSortBy = map[string]ListLogAnalyticsObjectCollectionRulesSortByEnum{
	"timeUpdated": ListLogAnalyticsObjectCollectionRulesSortByTimeupdated,
	"timeCreated": ListLogAnalyticsObjectCollectionRulesSortByTimecreated,
	"name":        ListLogAnalyticsObjectCollectionRulesSortByName,
}

// GetListLogAnalyticsObjectCollectionRulesSortByEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesSortByEnum
func GetListLogAnalyticsObjectCollectionRulesSortByEnumValues() []ListLogAnalyticsObjectCollectionRulesSortByEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesSortBy {
		values = append(values, v)
	}
	return values
}
