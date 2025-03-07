// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ExtendHttpRequestHeaderValueRule An object that represents the action of modifying a request header value. This rule applies only to HTTP listeners.
// This rule adds a prefix, a suffix, or both to the header value.
// **NOTES:**
// *  This rule requires a value for a prefix, suffix, or both.
// *  The system does not support this rule for headers with multiple values.
// *  The system does not distinquish between underscore and dash characters in headers. That is, it treats
//    `example_header_name` and `example-header-name` as identical.  If two such headers appear in a request, the system
//    applies the action to the first header it finds. The affected header cannot be determined in advance. Oracle
//    recommends that you do not rely on underscore or dash characters to uniquely distinguish header names.
type ExtendHttpRequestHeaderValueRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A string to prepend to the header value. The resulting header value must conform to RFC 7230.
	// With the following exceptions:
	// *  value cannot contain `$`
	// *  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.
	// Example: `example_prefix_value`
	Prefix *string `mandatory:"false" json:"prefix"`

	// A string to append to the header value. The resulting header value must conform to RFC 7230.
	// With the following exceptions:
	// *  value cannot contain `$`
	// *  value cannot contain patterns like `{variable_name}`. They are reserved for future extensions. Currently, such values are invalid.
	// Example: `example_suffix_value`
	Suffix *string `mandatory:"false" json:"suffix"`
}

func (m ExtendHttpRequestHeaderValueRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExtendHttpRequestHeaderValueRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExtendHttpRequestHeaderValueRule ExtendHttpRequestHeaderValueRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeExtendHttpRequestHeaderValueRule
	}{
		"EXTEND_HTTP_REQUEST_HEADER_VALUE",
		(MarshalTypeExtendHttpRequestHeaderValueRule)(m),
	}

	return json.Marshal(&s)
}
