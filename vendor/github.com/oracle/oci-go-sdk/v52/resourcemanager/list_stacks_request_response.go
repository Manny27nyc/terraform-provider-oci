// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListStacksRequest wrapper for the ListStacks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListStacks.go.html to see an example of how to use ListStacksRequest.
type ListStacksRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that exist in the compartment, identified by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a stack.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter that returns only those resources that match the specified
	// lifecycle state. The state value is case-insensitive.
	// For more information about stack lifecycle states, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates).
	// Allowable values:
	// - CREATING
	// - ACTIVE
	// - DELETING
	// - DELETED
	// - FAILED
	LifecycleState StackLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list a resource by name.
	// Requires `sortBy` set to `DISPLAYNAME`.
	// Alternatively, when you know the resource OCID, use the related Get operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use when sorting returned resources.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListStacksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder ListStacksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStacksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStacksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStacksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStacksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStacksResponse wrapper for the ListStacks operation
type ListStacksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StackSummary instances
	Items []StackSummary `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStacksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStacksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStacksSortByEnum Enum with underlying type: string
type ListStacksSortByEnum string

// Set of constants representing the allowable values for ListStacksSortByEnum
const (
	ListStacksSortByTimecreated ListStacksSortByEnum = "TIMECREATED"
	ListStacksSortByDisplayname ListStacksSortByEnum = "DISPLAYNAME"
)

var mappingListStacksSortBy = map[string]ListStacksSortByEnum{
	"TIMECREATED": ListStacksSortByTimecreated,
	"DISPLAYNAME": ListStacksSortByDisplayname,
}

// GetListStacksSortByEnumValues Enumerates the set of values for ListStacksSortByEnum
func GetListStacksSortByEnumValues() []ListStacksSortByEnum {
	values := make([]ListStacksSortByEnum, 0)
	for _, v := range mappingListStacksSortBy {
		values = append(values, v)
	}
	return values
}

// ListStacksSortOrderEnum Enum with underlying type: string
type ListStacksSortOrderEnum string

// Set of constants representing the allowable values for ListStacksSortOrderEnum
const (
	ListStacksSortOrderAsc  ListStacksSortOrderEnum = "ASC"
	ListStacksSortOrderDesc ListStacksSortOrderEnum = "DESC"
)

var mappingListStacksSortOrder = map[string]ListStacksSortOrderEnum{
	"ASC":  ListStacksSortOrderAsc,
	"DESC": ListStacksSortOrderDesc,
}

// GetListStacksSortOrderEnumValues Enumerates the set of values for ListStacksSortOrderEnum
func GetListStacksSortOrderEnumValues() []ListStacksSortOrderEnum {
	values := make([]ListStacksSortOrderEnum, 0)
	for _, v := range mappingListStacksSortOrder {
		values = append(values, v)
	}
	return values
}
