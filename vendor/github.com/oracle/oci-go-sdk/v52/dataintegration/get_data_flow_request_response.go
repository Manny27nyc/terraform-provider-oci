// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetDataFlowRequest wrapper for the GetDataFlow operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetDataFlow.go.html to see an example of how to use GetDataFlowRequest.
type GetDataFlowRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The data flow key.
	DataFlowKey *string `mandatory:"true" contributesTo:"path" name:"dataFlowKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to expand references of the object. If value is true, then all referenced objects are expanded. If value is false, then shallow objects are returned in place of references. Default is false. <br><br><B>Example:</B><br> <ul> <li><B>?expandReferences=true</B> returns all objects of type data loader task</li> </ul>
	ExpandReferences *string `mandatory:"false" contributesTo:"query" name:"expandReferences"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataFlowRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataFlowRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDataFlowRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataFlowRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDataFlowResponse wrapper for the GetDataFlow operation
type GetDataFlowResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataFlow instance
	DataFlow `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataFlowResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataFlowResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
