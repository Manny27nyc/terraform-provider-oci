// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetPdbConversionHistoryEntryRequest wrapper for the GetPdbConversionHistoryEntry operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetPdbConversionHistoryEntry.go.html to see an example of how to use GetPdbConversionHistoryEntryRequest.
type GetPdbConversionHistoryEntryRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// The database conversion history OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	PdbConversionHistoryEntryId *string `mandatory:"true" contributesTo:"path" name:"pdbConversionHistoryEntryId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPdbConversionHistoryEntryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPdbConversionHistoryEntryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPdbConversionHistoryEntryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPdbConversionHistoryEntryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPdbConversionHistoryEntryResponse wrapper for the GetPdbConversionHistoryEntry operation
type GetPdbConversionHistoryEntryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PdbConversionHistoryEntry instance
	PdbConversionHistoryEntry `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPdbConversionHistoryEntryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPdbConversionHistoryEntryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
