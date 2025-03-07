// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// MessageDetails The content of the message to be published.
type MessageDetails struct {

	// The body of the message to be published.
	// Avoid entering confidential information.
	Body *string `mandatory:"true" json:"body"`

	// The title of the message to be published. Avoid entering confidential information.
	Title *string `mandatory:"false" json:"title"`
}

func (m MessageDetails) String() string {
	return common.PointerString(m)
}
