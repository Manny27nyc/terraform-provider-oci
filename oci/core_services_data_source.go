// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterDatasource("oci_core_services", CoreServicesDataSource())
}

func CoreServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreServices,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreServices(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListServicesResponse
}

func (s *CoreServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreServicesDataSourceCrud) Get() error {
	request := oci_core.ListServicesRequest{}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.ListServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreServicesDataSource-", CoreServicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		service := map[string]interface{}{}

		if r.CidrBlock != nil {
			service["cidr_block"] = *r.CidrBlock
		}

		if r.Description != nil {
			service["description"] = *r.Description
		}

		if r.Id != nil {
			service["id"] = *r.Id
		}

		if r.Name != nil {
			service["name"] = *r.Name
		}

		resources = append(resources, service)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreServicesDataSource().Schema["services"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("services", resources); err != nil {
		return err
	}

	return nil
}
