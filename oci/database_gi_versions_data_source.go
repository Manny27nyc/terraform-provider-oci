// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v52/database"
)

func init() {
	RegisterDatasource("oci_database_gi_versions", DatabaseGiVersionsDataSource())
}

func DatabaseGiVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseGiVersions,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gi_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseGiVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseGiVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseGiVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListGiVersionsResponse
}

func (s *DatabaseGiVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseGiVersionsDataSourceCrud) Get() error {
	request := oci_database.ListGiVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "database")

	response, err := s.Client.ListGiVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGiVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseGiVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseGiVersionsDataSource-", DatabaseGiVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		giVersion := map[string]interface{}{}

		if r.Version != nil {
			giVersion["version"] = *r.Version
		}

		resources = append(resources, giVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseGiVersionsDataSource().Schema["gi_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("gi_versions", resources); err != nil {
		return err
	}

	return nil
}
