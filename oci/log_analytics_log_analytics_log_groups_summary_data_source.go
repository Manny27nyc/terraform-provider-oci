// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v52/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_log_analytics_log_groups_summary", LogAnalyticsLogAnalyticsLogGroupsSummaryDataSource())
}

func LogAnalyticsLogAnalyticsLogGroupsSummaryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsLogGroupsSummary,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"log_group_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsLogGroupsSummary(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupsSummaryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsLogAnalyticsLogGroupsSummaryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLogAnalyticsLogGroupsSummaryResponse
}

func (s *LogAnalyticsLogAnalyticsLogGroupsSummaryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsLogGroupsSummaryDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsLogGroupsSummaryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLogAnalyticsLogGroupsSummary(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsLogGroupsSummaryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LogAnalyticsLogAnalyticsLogGroupsSummaryDataSource-", LogAnalyticsLogAnalyticsLogGroupsSummaryDataSource(), s.D))

	if s.Res.Count != nil {
		s.D.Set("log_group_count", *s.Res.Count)
	}

	return nil
}
