// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// UsageSummary The usage store result.
type UsageSummary struct {

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The tenancy OCID.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The tenancy name.
	TenantName *string `mandatory:"false" json:"tenantName"`

	// The compartment OCID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The compartment path, starting from root.
	CompartmentPath *string `mandatory:"false" json:"compartmentPath"`

	// The compartment name.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The service name that is incurring the cost.
	Service *string `mandatory:"false" json:"service"`

	// The resource name that is incurring the cost.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The resource OCID that is incurring the cost.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The region of the usage.
	Region *string `mandatory:"false" json:"region"`

	// The availability domain of the usage.
	Ad *string `mandatory:"false" json:"ad"`

	// The resource size being metered.
	Weight *float32 `mandatory:"false" json:"weight"`

	// The resource shape.
	Shape *string `mandatory:"false" json:"shape"`

	// The SKU part number.
	SkuPartNumber *string `mandatory:"false" json:"skuPartNumber"`

	// The SKU friendly name.
	SkuName *string `mandatory:"false" json:"skuName"`

	// The usage unit.
	Unit *string `mandatory:"false" json:"unit"`

	// The discretionary discount applied to the SKU.
	Discount *float32 `mandatory:"false" json:"discount"`

	// The SKU list rate (not discount).
	ListRate *float32 `mandatory:"false" json:"listRate"`

	// Platform for the cost.
	Platform *string `mandatory:"false" json:"platform"`

	// The computed cost.
	ComputedAmount *float32 `mandatory:"false" json:"computedAmount"`

	// The usage number.
	ComputedQuantity *float32 `mandatory:"false" json:"computedQuantity"`

	// The SPM OverageFlag.
	OveragesFlag *string `mandatory:"false" json:"overagesFlag"`

	// The price per unit.
	UnitPrice *float32 `mandatory:"false" json:"unitPrice"`

	// The price currency.
	Currency *string `mandatory:"false" json:"currency"`

	// The subscription ID.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The overage usage.
	Overage *string `mandatory:"false" json:"overage"`

	// The forecasted data.
	IsForecast *bool `mandatory:"false" json:"isForecast"`

	// For grouping, a tag definition. For filtering, a definition and key.
	Tags []Tag `mandatory:"false" json:"tags"`
}

func (m UsageSummary) String() string {
	return common.PointerString(m)
}
