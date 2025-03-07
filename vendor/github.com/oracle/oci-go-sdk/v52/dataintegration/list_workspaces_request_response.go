// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListWorkspacesRequest wrapper for the ListWorkspaces operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkspaces.go.html to see an example of how to use ListWorkspacesRequest.
type ListWorkspacesRequest struct {

	// The OCID of the compartment containing the resources you want to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The lifecycle state of a resource. When specified, the operation only returns resources that match the given lifecycle state. When not specified, all lifecycle states are processed as a match.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkspacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkspacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkspacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkspacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkspacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkspacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkspacesResponse wrapper for the ListWorkspaces operation
type ListWorkspacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkspaceSummary instances
	Items []WorkspaceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkspacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkspacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkspacesSortOrderEnum Enum with underlying type: string
type ListWorkspacesSortOrderEnum string

// Set of constants representing the allowable values for ListWorkspacesSortOrderEnum
const (
	ListWorkspacesSortOrderAsc  ListWorkspacesSortOrderEnum = "ASC"
	ListWorkspacesSortOrderDesc ListWorkspacesSortOrderEnum = "DESC"
)

var mappingListWorkspacesSortOrder = map[string]ListWorkspacesSortOrderEnum{
	"ASC":  ListWorkspacesSortOrderAsc,
	"DESC": ListWorkspacesSortOrderDesc,
}

// GetListWorkspacesSortOrderEnumValues Enumerates the set of values for ListWorkspacesSortOrderEnum
func GetListWorkspacesSortOrderEnumValues() []ListWorkspacesSortOrderEnum {
	values := make([]ListWorkspacesSortOrderEnum, 0)
	for _, v := range mappingListWorkspacesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkspacesSortByEnum Enum with underlying type: string
type ListWorkspacesSortByEnum string

// Set of constants representing the allowable values for ListWorkspacesSortByEnum
const (
	ListWorkspacesSortByTimeCreated ListWorkspacesSortByEnum = "TIME_CREATED"
	ListWorkspacesSortByDisplayName ListWorkspacesSortByEnum = "DISPLAY_NAME"
)

var mappingListWorkspacesSortBy = map[string]ListWorkspacesSortByEnum{
	"TIME_CREATED": ListWorkspacesSortByTimeCreated,
	"DISPLAY_NAME": ListWorkspacesSortByDisplayName,
}

// GetListWorkspacesSortByEnumValues Enumerates the set of values for ListWorkspacesSortByEnum
func GetListWorkspacesSortByEnumValues() []ListWorkspacesSortByEnum {
	values := make([]ListWorkspacesSortByEnum, 0)
	for _, v := range mappingListWorkspacesSortBy {
		values = append(values, v)
	}
	return values
}
