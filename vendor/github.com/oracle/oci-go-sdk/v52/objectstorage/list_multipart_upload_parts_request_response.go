// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListMultipartUploadPartsRequest wrapper for the ListMultipartUploadParts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/ListMultipartUploadParts.go.html to see an example of how to use ListMultipartUploadPartsRequest.
type ListMultipartUploadPartsRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object. Avoid entering confidential information.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// The upload ID for a multipart upload.
	UploadId *string `mandatory:"true" contributesTo:"query" name:"uploadId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMultipartUploadPartsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMultipartUploadPartsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMultipartUploadPartsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMultipartUploadPartsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMultipartUploadPartsResponse wrapper for the ListMultipartUploadParts operation
type ListMultipartUploadPartsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MultipartUploadPartSummary instances
	Items []MultipartUploadPartSummary `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For paginating a list of multipart upload parts.
	// In the GET request, set the limit to the number of multipart upload parts that you want returned in the
	// response. If the `opc-next-page` header appears in the response, then this is a partial list and there are
	// additional multipart upload parts to get. Include the header's value as the `page` parameter in the subsequent
	// GET request to get the next batch of multipart upload parts. Repeat this process to retrieve the entire list
	// of multipart upload parts.
	// For more details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMultipartUploadPartsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMultipartUploadPartsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
