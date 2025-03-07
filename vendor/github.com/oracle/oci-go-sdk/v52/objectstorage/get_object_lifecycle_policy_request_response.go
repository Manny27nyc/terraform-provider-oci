// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetObjectLifecyclePolicyRequest wrapper for the GetObjectLifecyclePolicy operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/GetObjectLifecyclePolicy.go.html to see an example of how to use GetObjectLifecyclePolicyRequest.
type GetObjectLifecyclePolicyRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetObjectLifecyclePolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetObjectLifecyclePolicyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetObjectLifecyclePolicyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetObjectLifecyclePolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetObjectLifecyclePolicyResponse wrapper for the GetObjectLifecyclePolicy operation
type GetObjectLifecyclePolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ObjectLifecyclePolicy instance
	ObjectLifecyclePolicy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// The entity tag (ETag) for the object lifecycle policy.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response GetObjectLifecyclePolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetObjectLifecyclePolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
