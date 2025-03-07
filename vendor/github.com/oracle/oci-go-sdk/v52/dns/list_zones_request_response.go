// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListZonesRequest wrapper for the ListZones operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListZones.go.html to see an example of how to use ListZonesRequest.
type ListZonesRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A case-sensitive filter for zone names.
	// Will match any zone with a name that equals the provided value.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Search by zone name.
	// Will match any zone whose name (case-insensitive) contains the provided value.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// Search by zone type, `PRIMARY` or `SECONDARY`.
	// Will match any zone whose type equals the provided value.
	ZoneType ListZonesZoneTypeEnum `mandatory:"false" contributesTo:"query" name:"zoneType" omitEmpty:"true"`

	// An RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) timestamp that states
	// all returned resources were created on or after the indicated time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// An RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) timestamp that states
	// all returned resources were created before the indicated time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// The state of a resource.
	LifecycleState ListZonesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field by which to sort zones.
	SortBy ListZonesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder ListZonesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListZonesScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the resource is associated with.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// Search for zones that are associated with a TSIG key.
	TsigKeyId *string `mandatory:"false" contributesTo:"query" name:"tsigKeyId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZonesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZonesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZonesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZonesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListZonesResponse wrapper for the ListZones operation
type ListZonesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ZoneSummary instances
	Items []ZoneSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListZonesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZonesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZonesZoneTypeEnum Enum with underlying type: string
type ListZonesZoneTypeEnum string

// Set of constants representing the allowable values for ListZonesZoneTypeEnum
const (
	ListZonesZoneTypePrimary   ListZonesZoneTypeEnum = "PRIMARY"
	ListZonesZoneTypeSecondary ListZonesZoneTypeEnum = "SECONDARY"
)

var mappingListZonesZoneType = map[string]ListZonesZoneTypeEnum{
	"PRIMARY":   ListZonesZoneTypePrimary,
	"SECONDARY": ListZonesZoneTypeSecondary,
}

// GetListZonesZoneTypeEnumValues Enumerates the set of values for ListZonesZoneTypeEnum
func GetListZonesZoneTypeEnumValues() []ListZonesZoneTypeEnum {
	values := make([]ListZonesZoneTypeEnum, 0)
	for _, v := range mappingListZonesZoneType {
		values = append(values, v)
	}
	return values
}

// ListZonesLifecycleStateEnum Enum with underlying type: string
type ListZonesLifecycleStateEnum string

// Set of constants representing the allowable values for ListZonesLifecycleStateEnum
const (
	ListZonesLifecycleStateActive   ListZonesLifecycleStateEnum = "ACTIVE"
	ListZonesLifecycleStateCreating ListZonesLifecycleStateEnum = "CREATING"
	ListZonesLifecycleStateDeleted  ListZonesLifecycleStateEnum = "DELETED"
	ListZonesLifecycleStateDeleting ListZonesLifecycleStateEnum = "DELETING"
	ListZonesLifecycleStateFailed   ListZonesLifecycleStateEnum = "FAILED"
	ListZonesLifecycleStateUpdating ListZonesLifecycleStateEnum = "UPDATING"
)

var mappingListZonesLifecycleState = map[string]ListZonesLifecycleStateEnum{
	"ACTIVE":   ListZonesLifecycleStateActive,
	"CREATING": ListZonesLifecycleStateCreating,
	"DELETED":  ListZonesLifecycleStateDeleted,
	"DELETING": ListZonesLifecycleStateDeleting,
	"FAILED":   ListZonesLifecycleStateFailed,
	"UPDATING": ListZonesLifecycleStateUpdating,
}

// GetListZonesLifecycleStateEnumValues Enumerates the set of values for ListZonesLifecycleStateEnum
func GetListZonesLifecycleStateEnumValues() []ListZonesLifecycleStateEnum {
	values := make([]ListZonesLifecycleStateEnum, 0)
	for _, v := range mappingListZonesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListZonesSortByEnum Enum with underlying type: string
type ListZonesSortByEnum string

// Set of constants representing the allowable values for ListZonesSortByEnum
const (
	ListZonesSortByName        ListZonesSortByEnum = "name"
	ListZonesSortByZonetype    ListZonesSortByEnum = "zoneType"
	ListZonesSortByTimecreated ListZonesSortByEnum = "timeCreated"
)

var mappingListZonesSortBy = map[string]ListZonesSortByEnum{
	"name":        ListZonesSortByName,
	"zoneType":    ListZonesSortByZonetype,
	"timeCreated": ListZonesSortByTimecreated,
}

// GetListZonesSortByEnumValues Enumerates the set of values for ListZonesSortByEnum
func GetListZonesSortByEnumValues() []ListZonesSortByEnum {
	values := make([]ListZonesSortByEnum, 0)
	for _, v := range mappingListZonesSortBy {
		values = append(values, v)
	}
	return values
}

// ListZonesSortOrderEnum Enum with underlying type: string
type ListZonesSortOrderEnum string

// Set of constants representing the allowable values for ListZonesSortOrderEnum
const (
	ListZonesSortOrderAsc  ListZonesSortOrderEnum = "ASC"
	ListZonesSortOrderDesc ListZonesSortOrderEnum = "DESC"
)

var mappingListZonesSortOrder = map[string]ListZonesSortOrderEnum{
	"ASC":  ListZonesSortOrderAsc,
	"DESC": ListZonesSortOrderDesc,
}

// GetListZonesSortOrderEnumValues Enumerates the set of values for ListZonesSortOrderEnum
func GetListZonesSortOrderEnumValues() []ListZonesSortOrderEnum {
	values := make([]ListZonesSortOrderEnum, 0)
	for _, v := range mappingListZonesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListZonesScopeEnum Enum with underlying type: string
type ListZonesScopeEnum string

// Set of constants representing the allowable values for ListZonesScopeEnum
const (
	ListZonesScopeGlobal  ListZonesScopeEnum = "GLOBAL"
	ListZonesScopePrivate ListZonesScopeEnum = "PRIVATE"
)

var mappingListZonesScope = map[string]ListZonesScopeEnum{
	"GLOBAL":  ListZonesScopeGlobal,
	"PRIVATE": ListZonesScopePrivate,
}

// GetListZonesScopeEnumValues Enumerates the set of values for ListZonesScopeEnum
func GetListZonesScopeEnumValues() []ListZonesScopeEnum {
	values := make([]ListZonesScopeEnum, 0)
	for _, v := range mappingListZonesScope {
		values = append(values, v)
	}
	return values
}
