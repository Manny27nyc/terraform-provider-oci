// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetBucketRequest wrapper for the GetBucket operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/GetBucket.go.html to see an example of how to use GetBucketRequest.
type GetBucketRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The entity tag (ETag) to match with the ETag of an existing resource. If the specified ETag matches the ETag of
	// the existing resource, GET and HEAD requests will return the resource and PUT and POST requests will upload
	// the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag (ETag) to avoid matching. Wildcards ('*') are not allowed. If the specified ETag does not
	// match the ETag of the existing resource, the request returns the expected response. If the ETag matches
	// the ETag of the existing resource, the request returns an HTTP 304 status without a response body.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Bucket summary includes the 'namespace', 'name', 'compartmentId', 'createdBy', 'timeCreated',
	// and 'etag' fields. This parameter can also include 'approximateCount' (approximate number of objects), 'approximateSize'
	// (total approximate size in bytes of all objects) and 'autoTiering' (state of auto tiering on the bucket).
	// For example 'approximateCount,approximateSize,autoTiering'.
	Fields []GetBucketFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"csv"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetBucketRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetBucketRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetBucketRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetBucketRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetBucketResponse wrapper for the GetBucket operation
type GetBucketResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Bucket instance
	Bucket `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current entity tag (ETag) for the bucket.
	ETag *string `presentIn:"header" name:"etag"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetBucketResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetBucketResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetBucketFieldsEnum Enum with underlying type: string
type GetBucketFieldsEnum string

// Set of constants representing the allowable values for GetBucketFieldsEnum
const (
	GetBucketFieldsApproximatecount GetBucketFieldsEnum = "approximateCount"
	GetBucketFieldsApproximatesize  GetBucketFieldsEnum = "approximateSize"
	GetBucketFieldsAutotiering      GetBucketFieldsEnum = "autoTiering"
)

var mappingGetBucketFields = map[string]GetBucketFieldsEnum{
	"approximateCount": GetBucketFieldsApproximatecount,
	"approximateSize":  GetBucketFieldsApproximatesize,
	"autoTiering":      GetBucketFieldsAutotiering,
}

// GetGetBucketFieldsEnumValues Enumerates the set of values for GetBucketFieldsEnum
func GetGetBucketFieldsEnumValues() []GetBucketFieldsEnum {
	values := make([]GetBucketFieldsEnum, 0)
	for _, v := range mappingGetBucketFields {
		values = append(values, v)
	}
	return values
}
