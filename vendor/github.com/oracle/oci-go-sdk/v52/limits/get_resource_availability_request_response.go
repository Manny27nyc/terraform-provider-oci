// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetResourceAvailabilityRequest wrapper for the GetResourceAvailability operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limits/GetResourceAvailability.go.html to see an example of how to use GetResourceAvailabilityRequest.
type GetResourceAvailabilityRequest struct {

	// The service name of the target quota.
	ServiceName *string `mandatory:"true" contributesTo:"path" name:"serviceName"`

	// The limit name for which to fetch the data.
	LimitName *string `mandatory:"true" contributesTo:"path" name:"limitName"`

	// The OCID of the compartment for which data is being fetched.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// This field is mandatory if the scopeType of the target resource limit is AD.
	// Otherwise, this field should be omitted.
	// If the above requirements are not met, the API returns a 400 - InvalidParameter response.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetResourceAvailabilityRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetResourceAvailabilityRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetResourceAvailabilityRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetResourceAvailabilityRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetResourceAvailabilityResponse wrapper for the GetResourceAvailability operation
type GetResourceAvailabilityResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ResourceAvailability instance
	ResourceAvailability `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetResourceAvailabilityResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetResourceAvailabilityResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
