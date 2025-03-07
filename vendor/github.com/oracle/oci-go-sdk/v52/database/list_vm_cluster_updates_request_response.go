// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListVmClusterUpdatesRequest wrapper for the ListVmClusterUpdates operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterUpdates.go.html to see an example of how to use ListVmClusterUpdatesRequest.
type ListVmClusterUpdatesRequest struct {

	// The VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	VmClusterId *string `mandatory:"true" contributesTo:"path" name:"vmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListVmClusterUpdatesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState VmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmClusterUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmClusterUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmClusterUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmClusterUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVmClusterUpdatesResponse wrapper for the ListVmClusterUpdates operation
type ListVmClusterUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VmClusterUpdateSummary instances
	Items []VmClusterUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmClusterUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmClusterUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmClusterUpdatesUpdateTypeEnum Enum with underlying type: string
type ListVmClusterUpdatesUpdateTypeEnum string

// Set of constants representing the allowable values for ListVmClusterUpdatesUpdateTypeEnum
const (
	ListVmClusterUpdatesUpdateTypeGiUpgrade ListVmClusterUpdatesUpdateTypeEnum = "GI_UPGRADE"
	ListVmClusterUpdatesUpdateTypeGiPatch   ListVmClusterUpdatesUpdateTypeEnum = "GI_PATCH"
	ListVmClusterUpdatesUpdateTypeOsUpdate  ListVmClusterUpdatesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListVmClusterUpdatesUpdateType = map[string]ListVmClusterUpdatesUpdateTypeEnum{
	"GI_UPGRADE": ListVmClusterUpdatesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListVmClusterUpdatesUpdateTypeGiPatch,
	"OS_UPDATE":  ListVmClusterUpdatesUpdateTypeOsUpdate,
}

// GetListVmClusterUpdatesUpdateTypeEnumValues Enumerates the set of values for ListVmClusterUpdatesUpdateTypeEnum
func GetListVmClusterUpdatesUpdateTypeEnumValues() []ListVmClusterUpdatesUpdateTypeEnum {
	values := make([]ListVmClusterUpdatesUpdateTypeEnum, 0)
	for _, v := range mappingListVmClusterUpdatesUpdateType {
		values = append(values, v)
	}
	return values
}
