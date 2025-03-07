// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v52/devops"
)

func init() {
	RegisterDatasource("oci_devops_repository_object_content", DevopsRepositoryObjectContentDataSource())
}

func DevopsRepositoryObjectContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryObjectContent,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sha": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDevopsRepositoryObjectContent(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryObjectContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsRepositoryObjectContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetObjectContentResponse
}

func (s *DevopsRepositoryObjectContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryObjectContentDataSourceCrud) Get() error {
	request := oci_devops.GetObjectContentRequest{}

	if filePath, ok := s.D.GetOkExists("file_path"); ok {
		tmp := filePath.(string)
		request.FilePath = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if sha, ok := s.D.GetOkExists("sha"); ok {
		tmp := sha.(string)
		request.Sha = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")

	response, err := s.Client.GetObjectContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryObjectContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DevopsRepositoryObjectContentDataSource-", DevopsRepositoryObjectContentDataSource(), s.D))

	return nil
}
