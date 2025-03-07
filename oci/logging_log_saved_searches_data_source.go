// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v52/logging"
)

func init() {
	RegisterDatasource("oci_logging_log_saved_searches", LoggingLogSavedSearchesDataSource())
}

func LoggingLogSavedSearchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoggingLogSavedSearches,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_saved_search_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_saved_search_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(LoggingLogSavedSearchResource()),
						},
					},
				},
			},
		},
	}
}

func readLoggingLogSavedSearches(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return ReadResource(sync)
}

type LoggingLogSavedSearchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.ListLogSavedSearchesResponse
}

func (s *LoggingLogSavedSearchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogSavedSearchesDataSourceCrud) Get() error {
	request := oci_logging.ListLogSavedSearchesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if logSavedSearchId, ok := s.D.GetOkExists("id"); ok {
		tmp := logSavedSearchId.(string)
		request.LogSavedSearchId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "logging")

	response, err := s.Client.ListLogSavedSearches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogSavedSearches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoggingLogSavedSearchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoggingLogSavedSearchesDataSource-", LoggingLogSavedSearchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	logSavedSearch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogSavedSearchSummaryToMap(item))
	}
	logSavedSearch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, LoggingLogSavedSearchesDataSource().Schema["log_saved_search_summary_collection"].Elem.(*schema.Resource).Schema)
		logSavedSearch["items"] = items
	}

	resources = append(resources, logSavedSearch)
	if err := s.D.Set("log_saved_search_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
