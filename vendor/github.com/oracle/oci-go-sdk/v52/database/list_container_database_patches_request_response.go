// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListContainerDatabasePatchesRequest wrapper for the ListContainerDatabasePatches operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListContainerDatabasePatches.go.html to see an example of how to use ListContainerDatabasePatchesRequest.
type ListContainerDatabasePatchesRequest struct {

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousContainerDatabaseId"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainerDatabasePatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainerDatabasePatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainerDatabasePatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainerDatabasePatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListContainerDatabasePatchesResponse wrapper for the ListContainerDatabasePatches operation
type ListContainerDatabasePatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousPatchSummary instances
	Items []AutonomousPatchSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListContainerDatabasePatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainerDatabasePatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
