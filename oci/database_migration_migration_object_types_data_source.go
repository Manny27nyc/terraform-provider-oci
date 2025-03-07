// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v52/databasemigration"
)

func init() {
	RegisterDatasource("oci_database_migration_migration_object_types", DatabaseMigrationMigrationObjectTypesDataSource())
}

func DatabaseMigrationMigrationObjectTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationMigrationObjectTypes,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"migration_object_type_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationMigrationObjectTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationMigrationObjectTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return ReadResource(sync)
}

type DatabaseMigrationMigrationObjectTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListMigrationObjectTypesResponse
}

func (s *DatabaseMigrationMigrationObjectTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationObjectTypesDataSourceCrud) Get() error {
	request := oci_database_migration.ListMigrationObjectTypesRequest{}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListMigrationObjectTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMigrationObjectTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationMigrationObjectTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseMigrationMigrationObjectTypesDataSource-", DatabaseMigrationMigrationObjectTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	migrationObjectType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MigrationObjectTypeSummaryToMap(item))
	}
	migrationObjectType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationMigrationObjectTypesDataSource().Schema["migration_object_type_summary_collection"].Elem.(*schema.Resource).Schema)
		migrationObjectType["items"] = items
	}

	resources = append(resources, migrationObjectType)
	if err := s.D.Set("migration_object_type_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func MigrationObjectTypeSummaryToMap(obj oci_database_migration.MigrationObjectTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
