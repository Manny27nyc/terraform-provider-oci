// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// DeleteBackendRequest wrapper for the DeleteBackend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loadbalancer/DeleteBackend.go.html to see an example of how to use DeleteBackendRequest.
type DeleteBackendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and server.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the backend set associated with the backend server.
	// Example: `example_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The IP address and port of the backend server to remove.
	// Example: `10.0.0.3:8080`
	BackendName *string `mandatory:"true" contributesTo:"path" name:"backendName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteBackendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteBackendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteBackendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteBackendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteBackendResponse wrapper for the DeleteBackend operation
type DeleteBackendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteBackendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteBackendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
