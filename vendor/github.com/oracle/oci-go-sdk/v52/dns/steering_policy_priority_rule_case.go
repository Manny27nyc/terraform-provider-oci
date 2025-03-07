// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// SteeringPolicyPriorityRuleCase The representation of SteeringPolicyPriorityRuleCase
type SteeringPolicyPriorityRuleCase struct {

	// An expression that uses conditions at the time of a DNS query to indicate
	// whether a case matches. Conditions may include the geographical location, IP
	// subnet, or ASN the DNS query originated. **Example:** If you have an
	// office that uses the subnet `192.0.2.0/24` you could use a `caseCondition`
	// expression `query.client.subnet in ('192.0.2.0/24')` to define a case that
	// matches queries from that office.
	CaseCondition *string `mandatory:"false" json:"caseCondition"`

	// An array of `SteeringPolicyPriorityAnswerData` objects.
	AnswerData []SteeringPolicyPriorityAnswerData `mandatory:"false" json:"answerData"`
}

func (m SteeringPolicyPriorityRuleCase) String() string {
	return common.PointerString(m)
}
