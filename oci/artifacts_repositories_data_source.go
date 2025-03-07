// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v52/artifacts"
)

func init() {
	RegisterDatasource("oci_artifacts_repositories", ArtifactsRepositoriesDataSource())
}

func ArtifactsRepositoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readArtifactsRepositories,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_immutable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(ArtifactsRepositoryResource()),
						},
					},
				},
			},
		},
	}
}

func readArtifactsRepositories(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return ReadResource(sync)
}

type ArtifactsRepositoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.ListRepositoriesResponse
}

func (s *ArtifactsRepositoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsRepositoriesDataSourceCrud) Get() error {
	request := oci_artifacts.ListRepositoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isImmutable, ok := s.D.GetOkExists("is_immutable"); ok {
		tmp := isImmutable.(bool)
		request.IsImmutable = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "artifacts")

	response, err := s.Client.ListRepositories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRepositories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ArtifactsRepositoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ArtifactsRepositoriesDataSource-", ArtifactsRepositoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	repository := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositorySummaryToMap(item))
	}
	repository["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, ArtifactsRepositoriesDataSource().Schema["repository_collection"].Elem.(*schema.Resource).Schema)
		repository["items"] = items
	}

	resources = append(resources, repository)
	if err := s.D.Set("repository_collection", resources); err != nil {
		return err
	}

	return nil
}
