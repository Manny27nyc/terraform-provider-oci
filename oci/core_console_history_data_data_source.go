// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterDatasource("oci_core_console_history_data", CoreConsoleHistoryContentDataSource())
}

func CoreConsoleHistoryContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreConsoleHistoryContent,
		Schema: map[string]*schema.Schema{
			"console_history_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offset": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreConsoleHistoryContent(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoryContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreConsoleHistoryContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetConsoleHistoryContentResponse
}

func (s *CoreConsoleHistoryContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreConsoleHistoryContentDataSourceCrud) Get() error {
	request := oci_core.GetConsoleHistoryContentRequest{}

	if consoleHistoryId, ok := s.D.GetOkExists("console_history_id"); ok {
		tmp := consoleHistoryId.(string)
		request.InstanceConsoleHistoryId = &tmp
	}

	if length, ok := s.D.GetOkExists("length"); ok {
		tmp := length.(int)
		request.Length = &tmp
	}

	if offset, ok := s.D.GetOkExists("offset"); ok {
		tmp := offset.(int)
		request.Offset = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.GetConsoleHistoryContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreConsoleHistoryContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreConsoleHistoryContentDataSource-", CoreConsoleHistoryContentDataSource(), s.D))

	if s.Res.Value != nil {
		s.D.Set("data", *s.Res.Value)
	}

	return nil
}
