// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// UpdateCloudExadataInfrastructureRequest wrapper for the UpdateCloudExadataInfrastructure operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/UpdateCloudExadataInfrastructure.go.html to see an example of how to use UpdateCloudExadataInfrastructureRequest.
type UpdateCloudExadataInfrastructureRequest struct {

	// The cloud Exadata infrastructure OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudExadataInfrastructureId *string `mandatory:"true" contributesTo:"path" name:"cloudExadataInfrastructureId"`

	// Request to update the properties of an cloud Exadata infrastructure resource.
	UpdateCloudExadataInfrastructureDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateCloudExadataInfrastructureRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateCloudExadataInfrastructureRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateCloudExadataInfrastructureRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateCloudExadataInfrastructureRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateCloudExadataInfrastructureResponse wrapper for the UpdateCloudExadataInfrastructure operation
type UpdateCloudExadataInfrastructureResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CloudExadataInfrastructure instance
	CloudExadataInfrastructure `presentIn:"body"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateCloudExadataInfrastructureResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateCloudExadataInfrastructureResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
