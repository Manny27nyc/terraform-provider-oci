// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// LogAnalyticsParserSummary LoganParserDetails
type LogAnalyticsParserSummary struct {

	// The content.
	Content *string `mandatory:"false" json:"content"`

	// The parser description.
	Description *string `mandatory:"false" json:"description"`

	// The parser display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The parser edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The encoding.
	Encoding *string `mandatory:"false" json:"encoding"`

	// The example content.
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// The parser fields.
	FieldMaps []LogAnalyticsParserField `mandatory:"false" json:"fieldMaps"`

	// The footer regular expression.
	FooterContent *string `mandatory:"false" json:"footerContent"`

	// The header content.
	HeaderContent *string `mandatory:"false" json:"headerContent"`

	// The parser name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating if this is a default parser.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// A flag indicating if this is a single line content parser.
	IsSingleLineContent *bool `mandatory:"false" json:"isSingleLineContent"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The language.
	Language *string `mandatory:"false" json:"language"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The log type test request version.
	LogTypeTestRequestVersion *int `mandatory:"false" json:"logTypeTestRequestVersion"`

	// The mapped parser list.
	MappedParsers []LogAnalyticsParser `mandatory:"false" json:"mappedParsers"`

	// The line characters for the parser to ignore.
	ParserIgnorelineCharacters *string `mandatory:"false" json:"parserIgnorelineCharacters"`

	// A flag indicating if the parser is hidden or not.
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// The parser sequence.
	ParserSequence *int `mandatory:"false" json:"parserSequence"`

	// The time zone.
	ParserTimezone *string `mandatory:"false" json:"parserTimezone"`

	ParserFilter *LogAnalyticsParserFilter `mandatory:"false" json:"parserFilter"`

	// A flag indicating whther or not the parser is write once.
	IsParserWrittenOnce *bool `mandatory:"false" json:"isParserWrittenOnce"`

	// The parser function list.
	ParserFunctions []LogAnalyticsParserFunction `mandatory:"false" json:"parserFunctions"`

	// The number of sources using this parser.
	SourcesCount *int64 `mandatory:"false" json:"sourcesCount"`

	// The list of sources using this parser.
	Sources []LogAnalyticsSource `mandatory:"false" json:"sources"`

	// A flag indicating whether or not to tokenize the original text.
	ShouldTokenizeOriginalText *bool `mandatory:"false" json:"shouldTokenizeOriginalText"`

	// The parser field delimiter.
	FieldDelimiter *string `mandatory:"false" json:"fieldDelimiter"`

	// The parser field qualifier.
	FieldQualifier *string `mandatory:"false" json:"fieldQualifier"`

	// The parser type.  Default value is REGEX.
	Type LogAnalyticsParserSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A flag indicating whether or not the parser has been deleted.
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`

	// A flag indicating whether the XML parser should consider the namespace(s) while processing the log data.
	IsNamespaceAware *bool `mandatory:"false" json:"isNamespaceAware"`
}

func (m LogAnalyticsParserSummary) String() string {
	return common.PointerString(m)
}

// LogAnalyticsParserSummaryTypeEnum Enum with underlying type: string
type LogAnalyticsParserSummaryTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsParserSummaryTypeEnum
const (
	LogAnalyticsParserSummaryTypeXml       LogAnalyticsParserSummaryTypeEnum = "XML"
	LogAnalyticsParserSummaryTypeJson      LogAnalyticsParserSummaryTypeEnum = "JSON"
	LogAnalyticsParserSummaryTypeRegex     LogAnalyticsParserSummaryTypeEnum = "REGEX"
	LogAnalyticsParserSummaryTypeOdl       LogAnalyticsParserSummaryTypeEnum = "ODL"
	LogAnalyticsParserSummaryTypeDelimited LogAnalyticsParserSummaryTypeEnum = "DELIMITED"
)

var mappingLogAnalyticsParserSummaryType = map[string]LogAnalyticsParserSummaryTypeEnum{
	"XML":       LogAnalyticsParserSummaryTypeXml,
	"JSON":      LogAnalyticsParserSummaryTypeJson,
	"REGEX":     LogAnalyticsParserSummaryTypeRegex,
	"ODL":       LogAnalyticsParserSummaryTypeOdl,
	"DELIMITED": LogAnalyticsParserSummaryTypeDelimited,
}

// GetLogAnalyticsParserSummaryTypeEnumValues Enumerates the set of values for LogAnalyticsParserSummaryTypeEnum
func GetLogAnalyticsParserSummaryTypeEnumValues() []LogAnalyticsParserSummaryTypeEnum {
	values := make([]LogAnalyticsParserSummaryTypeEnum, 0)
	for _, v := range mappingLogAnalyticsParserSummaryType {
		values = append(values, v)
	}
	return values
}
