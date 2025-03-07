// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListExternalContainerDatabasesRequest wrapper for the ListExternalContainerDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalContainerDatabases.go.html to see an example of how to use ListExternalContainerDatabasesRequest.
type ListExternalContainerDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListExternalContainerDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalContainerDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ExternalDatabaseBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalContainerDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalContainerDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalContainerDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalContainerDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExternalContainerDatabasesResponse wrapper for the ListExternalContainerDatabases operation
type ListExternalContainerDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExternalContainerDatabaseSummary instances
	Items []ExternalContainerDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalContainerDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalContainerDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalContainerDatabasesSortByEnum Enum with underlying type: string
type ListExternalContainerDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalContainerDatabasesSortByEnum
const (
	ListExternalContainerDatabasesSortByDisplayname ListExternalContainerDatabasesSortByEnum = "DISPLAYNAME"
	ListExternalContainerDatabasesSortByTimecreated ListExternalContainerDatabasesSortByEnum = "TIMECREATED"
)

var mappingListExternalContainerDatabasesSortBy = map[string]ListExternalContainerDatabasesSortByEnum{
	"DISPLAYNAME": ListExternalContainerDatabasesSortByDisplayname,
	"TIMECREATED": ListExternalContainerDatabasesSortByTimecreated,
}

// GetListExternalContainerDatabasesSortByEnumValues Enumerates the set of values for ListExternalContainerDatabasesSortByEnum
func GetListExternalContainerDatabasesSortByEnumValues() []ListExternalContainerDatabasesSortByEnum {
	values := make([]ListExternalContainerDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalContainerDatabasesSortBy {
		values = append(values, v)
	}
	return values
}

// ListExternalContainerDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalContainerDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalContainerDatabasesSortOrderEnum
const (
	ListExternalContainerDatabasesSortOrderAsc  ListExternalContainerDatabasesSortOrderEnum = "ASC"
	ListExternalContainerDatabasesSortOrderDesc ListExternalContainerDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalContainerDatabasesSortOrder = map[string]ListExternalContainerDatabasesSortOrderEnum{
	"ASC":  ListExternalContainerDatabasesSortOrderAsc,
	"DESC": ListExternalContainerDatabasesSortOrderDesc,
}

// GetListExternalContainerDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalContainerDatabasesSortOrderEnum
func GetListExternalContainerDatabasesSortOrderEnumValues() []ListExternalContainerDatabasesSortOrderEnum {
	values := make([]ListExternalContainerDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalContainerDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}
