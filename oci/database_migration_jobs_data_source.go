// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v52/databasemigration"
)

func init() {
	RegisterDatasource("oci_database_migration_jobs", DatabaseMigrationJobsDataSource())
}

func DatabaseMigrationJobsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationJobs,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(DatabaseMigrationJobResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationJobs(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return ReadResource(sync)
}

type DatabaseMigrationJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListJobsResponse
}

func (s *DatabaseMigrationJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationJobsDataSourceCrud) Get() error {
	request := oci_database_migration.ListJobsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_migration.ListJobsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListJobs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJobs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseMigrationJobsDataSource-", DatabaseMigrationJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	job := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JobSummaryToMap(item))
	}
	job["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationJobsDataSource().Schema["job_collection"].Elem.(*schema.Resource).Schema)
		job["items"] = items
	}

	resources = append(resources, job)
	if err := s.D.Set("job_collection", resources); err != nil {
		return err
	}

	return nil
}
