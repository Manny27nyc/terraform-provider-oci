// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// TypeListRule The type list rule that defines how fields are projected.
type TypeListRule struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Specifies whether the rule uses a java regex syntax.
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether to skip remaining rules when a match is found.
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// Reference to a typed object. This can be either a key value to an object within the document, a shall referenced to a `TypedObject`, or a full `TypedObject` definition.
	Scope *interface{} `mandatory:"false" json:"scope"`

	// Specifies whether to cascade or not.
	IsCascade *bool `mandatory:"false" json:"isCascade"`

	// Specifies if the rule is case sensitive.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// An arry of types.
	Types []interface{} `mandatory:"false" json:"types"`

	// The pattern matching strategy.
	MatchingStrategy TypeListRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// The rule type.
	RuleType TypeListRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

//GetKey returns Key
func (m TypeListRule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TypeListRule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TypeListRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m TypeListRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m TypeListRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m TypeListRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m TypeListRule) GetDescription() *string {
	return m.Description
}

func (m TypeListRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TypeListRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTypeListRule TypeListRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTypeListRule
	}{
		"TYPE_LIST_RULE",
		(MarshalTypeTypeListRule)(m),
	}

	return json.Marshal(&s)
}

// TypeListRuleMatchingStrategyEnum Enum with underlying type: string
type TypeListRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for TypeListRuleMatchingStrategyEnum
const (
	TypeListRuleMatchingStrategyNameOrTags TypeListRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	TypeListRuleMatchingStrategyTagsOnly   TypeListRuleMatchingStrategyEnum = "TAGS_ONLY"
	TypeListRuleMatchingStrategyNameOnly   TypeListRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingTypeListRuleMatchingStrategy = map[string]TypeListRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": TypeListRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    TypeListRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    TypeListRuleMatchingStrategyNameOnly,
}

// GetTypeListRuleMatchingStrategyEnumValues Enumerates the set of values for TypeListRuleMatchingStrategyEnum
func GetTypeListRuleMatchingStrategyEnumValues() []TypeListRuleMatchingStrategyEnum {
	values := make([]TypeListRuleMatchingStrategyEnum, 0)
	for _, v := range mappingTypeListRuleMatchingStrategy {
		values = append(values, v)
	}
	return values
}

// TypeListRuleRuleTypeEnum Enum with underlying type: string
type TypeListRuleRuleTypeEnum string

// Set of constants representing the allowable values for TypeListRuleRuleTypeEnum
const (
	TypeListRuleRuleTypeInclude TypeListRuleRuleTypeEnum = "INCLUDE"
	TypeListRuleRuleTypeExclude TypeListRuleRuleTypeEnum = "EXCLUDE"
)

var mappingTypeListRuleRuleType = map[string]TypeListRuleRuleTypeEnum{
	"INCLUDE": TypeListRuleRuleTypeInclude,
	"EXCLUDE": TypeListRuleRuleTypeExclude,
}

// GetTypeListRuleRuleTypeEnumValues Enumerates the set of values for TypeListRuleRuleTypeEnum
func GetTypeListRuleRuleTypeEnumValues() []TypeListRuleRuleTypeEnum {
	values := make([]TypeListRuleRuleTypeEnum, 0)
	for _, v := range mappingTypeListRuleRuleType {
		values = append(values, v)
	}
	return values
}
