// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v52/artifacts"
)

func init() {
	RegisterDatasource("oci_artifacts_container_repository", ArtifactsContainerRepositoryDataSource())
}

func ArtifactsContainerRepositoryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["repository_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ArtifactsContainerRepositoryResource(), fieldMap, readSingularArtifactsContainerRepository)
}

func readSingularArtifactsContainerRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return ReadResource(sync)
}

type ArtifactsContainerRepositoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetContainerRepositoryResponse
}

func (s *ArtifactsContainerRepositoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerRepositoryDataSourceCrud) Get() error {
	request := oci_artifacts.GetContainerRepositoryRequest{}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "artifacts")

	response, err := s.Client.GetContainerRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsContainerRepositoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BillableSizeInGBs != nil {
		s.D.Set("billable_size_in_gbs", strconv.FormatInt(*s.Res.BillableSizeInGBs, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ImageCount != nil {
		s.D.Set("image_count", *s.Res.ImageCount)
	}

	if s.Res.IsImmutable != nil {
		s.D.Set("is_immutable", *s.Res.IsImmutable)
	}

	if s.Res.IsPublic != nil {
		s.D.Set("is_public", *s.Res.IsPublic)
	}

	if s.Res.LayerCount != nil {
		s.D.Set("layer_count", *s.Res.LayerCount)
	}

	if s.Res.LayersSizeInBytes != nil {
		s.D.Set("layers_size_in_bytes", strconv.FormatInt(*s.Res.LayersSizeInBytes, 10))
	}

	if s.Res.Readme != nil {
		s.D.Set("readme", []interface{}{ContainerRepositoryReadmeToMap(s.Res.Readme)})
	} else {
		s.D.Set("readme", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastPushed != nil {
		s.D.Set("time_last_pushed", s.Res.TimeLastPushed.String())
	}

	return nil
}
