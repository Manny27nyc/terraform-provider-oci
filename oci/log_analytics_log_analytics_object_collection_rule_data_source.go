// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v52/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_log_analytics_object_collection_rule", LogAnalyticsLogAnalyticsObjectCollectionRuleDataSource())
}

func LogAnalyticsLogAnalyticsObjectCollectionRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["log_analytics_object_collection_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LogAnalyticsLogAnalyticsObjectCollectionRuleResource(), fieldMap, readSingularLogAnalyticsLogAnalyticsObjectCollectionRule)
}

func readSingularLogAnalyticsLogAnalyticsObjectCollectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLogAnalyticsObjectCollectionRuleResponse
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{}

	if logAnalyticsObjectCollectionRuleId, ok := s.D.GetOkExists("log_analytics_object_collection_rule_id"); ok {
		tmp := logAnalyticsObjectCollectionRuleId.(string)
		request.LogAnalyticsObjectCollectionRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLogAnalyticsObjectCollectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CharEncoding != nil {
		s.D.Set("char_encoding", *s.Res.CharEncoding)
	}

	s.D.Set("collection_type", s.Res.CollectionType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EntityId != nil {
		s.D.Set("entity_id", *s.Res.EntityId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogGroupId != nil {
		s.D.Set("log_group_id", *s.Res.LogGroupId)
	}

	if s.Res.LogSourceName != nil {
		s.D.Set("log_source_name", *s.Res.LogSourceName)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("object_name_filters", s.Res.ObjectNameFilters)

	if s.Res.OsBucketName != nil {
		s.D.Set("os_bucket_name", *s.Res.OsBucketName)
	}

	if s.Res.OsNamespace != nil {
		s.D.Set("os_namespace", *s.Res.OsNamespace)
	}

	if s.Res.Overrides != nil {
		s.D.Set("overrides", propertyOverridesToMap(s.Res.Overrides))
	} else {
		s.D.Set("overrides", nil)
	}

	if s.Res.PollSince != nil {
		s.D.Set("poll_since", *s.Res.PollSince)
	}

	if s.Res.PollTill != nil {
		s.D.Set("poll_till", *s.Res.PollTill)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
