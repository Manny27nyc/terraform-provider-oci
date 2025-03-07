// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListClustersRequest wrapper for the ListClusters operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/containerengine/ListClusters.go.html to see an example of how to use ListClustersRequest.
type ListClustersRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A cluster lifecycle state to filter on. Can have multiple parameters of this name.
	LifecycleState []ClusterLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// The name to filter on.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The optional order in which to sort the results.
	SortOrder ListClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The optional field to sort the results by.
	SortBy ListClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListClustersResponse wrapper for the ListClusters operation
type ListClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ClusterSummary instances
	Items []ClusterSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClustersSortOrderEnum Enum with underlying type: string
type ListClustersSortOrderEnum string

// Set of constants representing the allowable values for ListClustersSortOrderEnum
const (
	ListClustersSortOrderAsc  ListClustersSortOrderEnum = "ASC"
	ListClustersSortOrderDesc ListClustersSortOrderEnum = "DESC"
)

var mappingListClustersSortOrder = map[string]ListClustersSortOrderEnum{
	"ASC":  ListClustersSortOrderAsc,
	"DESC": ListClustersSortOrderDesc,
}

// GetListClustersSortOrderEnumValues Enumerates the set of values for ListClustersSortOrderEnum
func GetListClustersSortOrderEnumValues() []ListClustersSortOrderEnum {
	values := make([]ListClustersSortOrderEnum, 0)
	for _, v := range mappingListClustersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListClustersSortByEnum Enum with underlying type: string
type ListClustersSortByEnum string

// Set of constants representing the allowable values for ListClustersSortByEnum
const (
	ListClustersSortById          ListClustersSortByEnum = "ID"
	ListClustersSortByName        ListClustersSortByEnum = "NAME"
	ListClustersSortByTimeCreated ListClustersSortByEnum = "TIME_CREATED"
)

var mappingListClustersSortBy = map[string]ListClustersSortByEnum{
	"ID":           ListClustersSortById,
	"NAME":         ListClustersSortByName,
	"TIME_CREATED": ListClustersSortByTimeCreated,
}

// GetListClustersSortByEnumValues Enumerates the set of values for ListClustersSortByEnum
func GetListClustersSortByEnumValues() []ListClustersSortByEnum {
	values := make([]ListClustersSortByEnum, 0)
	for _, v := range mappingListClustersSortBy {
		values = append(values, v)
	}
	return values
}
