// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v52/databasemigration"
)

func init() {
	RegisterDatasource("oci_database_migration_agents", DatabaseMigrationAgentsDataSource())
}

func DatabaseMigrationAgentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationAgents,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(DatabaseMigrationAgentResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationAgents(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAgentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseMigrationClient()

	return ReadResource(sync)
}

type DatabaseMigrationAgentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAgentsResponse
}

func (s *DatabaseMigrationAgentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAgentsDataSourceCrud) Get() error {
	request := oci_database_migration.ListAgentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_migration.ListAgentsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAgents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationAgentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseMigrationAgentsDataSource-", DatabaseMigrationAgentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	agent := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AgentSummaryToMap(item))
	}
	agent["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAgentsDataSource().Schema["agent_collection"].Elem.(*schema.Resource).Schema)
		agent["items"] = items
	}

	resources = append(resources, agent)
	if err := s.D.Set("agent_collection", resources); err != nil {
		return err
	}

	return nil
}
