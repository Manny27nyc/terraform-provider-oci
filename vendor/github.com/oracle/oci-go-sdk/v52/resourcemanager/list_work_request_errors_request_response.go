// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListWorkRequestErrorsRequest wrapper for the ListWorkRequestErrors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrorsRequest.
type ListWorkRequestErrorsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// A filter to return only resources that exist in the compartment, identified by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder ListWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestErrorsResponse wrapper for the ListWorkRequestErrors operation
type ListWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestError instances
	Items []WorkRequestError `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortOrderEnum
const (
	ListWorkRequestErrorsSortOrderAsc  ListWorkRequestErrorsSortOrderEnum = "ASC"
	ListWorkRequestErrorsSortOrderDesc ListWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListWorkRequestErrorsSortOrder = map[string]ListWorkRequestErrorsSortOrderEnum{
	"ASC":  ListWorkRequestErrorsSortOrderAsc,
	"DESC": ListWorkRequestErrorsSortOrderDesc,
}

// GetListWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListWorkRequestErrorsSortOrderEnum
func GetListWorkRequestErrorsSortOrderEnumValues() []ListWorkRequestErrorsSortOrderEnum {
	values := make([]ListWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortOrder {
		values = append(values, v)
	}
	return values
}
