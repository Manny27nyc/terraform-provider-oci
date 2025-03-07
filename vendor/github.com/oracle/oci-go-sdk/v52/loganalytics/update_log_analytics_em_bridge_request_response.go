// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// UpdateLogAnalyticsEmBridgeRequest wrapper for the UpdateLogAnalyticsEmBridge operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsEmBridge.go.html to see an example of how to use UpdateLogAnalyticsEmBridgeRequest.
type UpdateLogAnalyticsEmBridgeRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The log analytics enterprise manager bridge OCID.
	LogAnalyticsEmBridgeId *string `mandatory:"true" contributesTo:"path" name:"logAnalyticsEmBridgeId"`

	// Log analytics enterprise manager information to be updated.
	UpdateLogAnalyticsEmBridgeDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateLogAnalyticsEmBridgeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateLogAnalyticsEmBridgeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateLogAnalyticsEmBridgeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateLogAnalyticsEmBridgeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateLogAnalyticsEmBridgeResponse wrapper for the UpdateLogAnalyticsEmBridge operation
type UpdateLogAnalyticsEmBridgeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsEmBridge instance
	LogAnalyticsEmBridge `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateLogAnalyticsEmBridgeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateLogAnalyticsEmBridgeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
