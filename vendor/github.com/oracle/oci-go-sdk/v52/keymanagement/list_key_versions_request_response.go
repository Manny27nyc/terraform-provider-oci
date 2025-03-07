// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListKeyVersionsRequest wrapper for the ListKeyVersions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ListKeyVersions.go.html to see an example of how to use ListKeyVersionsRequest.
type ListKeyVersionsRequest struct {

	// The OCID of the key.
	KeyId *string `mandatory:"true" contributesTo:"path" name:"keyId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `TIMECREATED` is descending. The default order for `DISPLAYNAME`
	// is ascending.
	SortBy ListKeyVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListKeyVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKeyVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKeyVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKeyVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKeyVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListKeyVersionsResponse wrapper for the ListKeyVersions operation
type ListKeyVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []KeyVersionSummary instances
	Items []KeyVersionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListKeyVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKeyVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKeyVersionsSortByEnum Enum with underlying type: string
type ListKeyVersionsSortByEnum string

// Set of constants representing the allowable values for ListKeyVersionsSortByEnum
const (
	ListKeyVersionsSortByTimecreated ListKeyVersionsSortByEnum = "TIMECREATED"
	ListKeyVersionsSortByDisplayname ListKeyVersionsSortByEnum = "DISPLAYNAME"
)

var mappingListKeyVersionsSortBy = map[string]ListKeyVersionsSortByEnum{
	"TIMECREATED": ListKeyVersionsSortByTimecreated,
	"DISPLAYNAME": ListKeyVersionsSortByDisplayname,
}

// GetListKeyVersionsSortByEnumValues Enumerates the set of values for ListKeyVersionsSortByEnum
func GetListKeyVersionsSortByEnumValues() []ListKeyVersionsSortByEnum {
	values := make([]ListKeyVersionsSortByEnum, 0)
	for _, v := range mappingListKeyVersionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListKeyVersionsSortOrderEnum Enum with underlying type: string
type ListKeyVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListKeyVersionsSortOrderEnum
const (
	ListKeyVersionsSortOrderAsc  ListKeyVersionsSortOrderEnum = "ASC"
	ListKeyVersionsSortOrderDesc ListKeyVersionsSortOrderEnum = "DESC"
)

var mappingListKeyVersionsSortOrder = map[string]ListKeyVersionsSortOrderEnum{
	"ASC":  ListKeyVersionsSortOrderAsc,
	"DESC": ListKeyVersionsSortOrderDesc,
}

// GetListKeyVersionsSortOrderEnumValues Enumerates the set of values for ListKeyVersionsSortOrderEnum
func GetListKeyVersionsSortOrderEnumValues() []ListKeyVersionsSortOrderEnum {
	values := make([]ListKeyVersionsSortOrderEnum, 0)
	for _, v := range mappingListKeyVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
