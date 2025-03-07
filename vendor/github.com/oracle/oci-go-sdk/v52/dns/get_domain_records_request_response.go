// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// GetDomainRecordsRequest wrapper for the GetDomainRecords operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetDomainRecords.go.html to see an example of how to use GetDomainRecordsRequest.
type GetDomainRecordsRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

	// The target fully-qualified domain name (FQDN) within the target zone.
	Domain *string `mandatory:"true" contributesTo:"path" name:"domain"`

	// The `If-None-Match` header field makes the request method conditional on
	// the absence of any current representation of the target resource, when
	// the field-value is `*`, or having a selected representation with an
	// entity-tag that does not match any of those listed in the field-value.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"If-None-Match"`

	// The `If-Modified-Since` header field makes a GET or HEAD request method
	// conditional on the selected representation's modification date being more
	// recent than the date provided in the field-value.  Transfer of the
	// selected representation's data is avoided if that data has not changed.
	IfModifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Modified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The version of the zone for which data is requested.
	ZoneVersion *string `mandatory:"false" contributesTo:"query" name:"zoneVersion"`

	// Search by record type.
	// Will match any record whose type (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4) (case-insensitive) equals the provided value.
	Rtype *string `mandatory:"false" contributesTo:"query" name:"rtype"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope GetDomainRecordsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the resource is associated with.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// The field by which to sort records.
	SortBy GetDomainRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder GetDomainRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDomainRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDomainRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDomainRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDomainRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDomainRecordsResponse wrapper for the GetDomainRecords operation
type GetDomainRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecordCollection instances
	RecordCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response GetDomainRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDomainRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDomainRecordsScopeEnum Enum with underlying type: string
type GetDomainRecordsScopeEnum string

// Set of constants representing the allowable values for GetDomainRecordsScopeEnum
const (
	GetDomainRecordsScopeGlobal  GetDomainRecordsScopeEnum = "GLOBAL"
	GetDomainRecordsScopePrivate GetDomainRecordsScopeEnum = "PRIVATE"
)

var mappingGetDomainRecordsScope = map[string]GetDomainRecordsScopeEnum{
	"GLOBAL":  GetDomainRecordsScopeGlobal,
	"PRIVATE": GetDomainRecordsScopePrivate,
}

// GetGetDomainRecordsScopeEnumValues Enumerates the set of values for GetDomainRecordsScopeEnum
func GetGetDomainRecordsScopeEnumValues() []GetDomainRecordsScopeEnum {
	values := make([]GetDomainRecordsScopeEnum, 0)
	for _, v := range mappingGetDomainRecordsScope {
		values = append(values, v)
	}
	return values
}

// GetDomainRecordsSortByEnum Enum with underlying type: string
type GetDomainRecordsSortByEnum string

// Set of constants representing the allowable values for GetDomainRecordsSortByEnum
const (
	GetDomainRecordsSortByRtype GetDomainRecordsSortByEnum = "rtype"
	GetDomainRecordsSortByTtl   GetDomainRecordsSortByEnum = "ttl"
)

var mappingGetDomainRecordsSortBy = map[string]GetDomainRecordsSortByEnum{
	"rtype": GetDomainRecordsSortByRtype,
	"ttl":   GetDomainRecordsSortByTtl,
}

// GetGetDomainRecordsSortByEnumValues Enumerates the set of values for GetDomainRecordsSortByEnum
func GetGetDomainRecordsSortByEnumValues() []GetDomainRecordsSortByEnum {
	values := make([]GetDomainRecordsSortByEnum, 0)
	for _, v := range mappingGetDomainRecordsSortBy {
		values = append(values, v)
	}
	return values
}

// GetDomainRecordsSortOrderEnum Enum with underlying type: string
type GetDomainRecordsSortOrderEnum string

// Set of constants representing the allowable values for GetDomainRecordsSortOrderEnum
const (
	GetDomainRecordsSortOrderAsc  GetDomainRecordsSortOrderEnum = "ASC"
	GetDomainRecordsSortOrderDesc GetDomainRecordsSortOrderEnum = "DESC"
)

var mappingGetDomainRecordsSortOrder = map[string]GetDomainRecordsSortOrderEnum{
	"ASC":  GetDomainRecordsSortOrderAsc,
	"DESC": GetDomainRecordsSortOrderDesc,
}

// GetGetDomainRecordsSortOrderEnumValues Enumerates the set of values for GetDomainRecordsSortOrderEnum
func GetGetDomainRecordsSortOrderEnumValues() []GetDomainRecordsSortOrderEnum {
	values := make([]GetDomainRecordsSortOrderEnum, 0)
	for _, v := range mappingGetDomainRecordsSortOrder {
		values = append(values, v)
	}
	return values
}
