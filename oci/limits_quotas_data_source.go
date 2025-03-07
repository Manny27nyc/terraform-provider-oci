// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v52/limits"
)

func init() {
	RegisterDatasource("oci_limits_quotas", LimitsQuotasDataSource())
}

func LimitsQuotasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLimitsQuotas,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quotas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(LimitsQuotaResource()),
			},
		},
	}
}

func readLimitsQuotas(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).quotasClient()

	return ReadResource(sync)
}

type LimitsQuotasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.QuotasClient
	Res    *oci_limits.ListQuotasResponse
}

func (s *LimitsQuotasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsQuotasDataSourceCrud) Get() error {
	request := oci_limits.ListQuotasRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_limits.ListQuotasLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "limits")

	response, err := s.Client.ListQuotas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListQuotas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LimitsQuotasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LimitsQuotasDataSource-", LimitsQuotasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		quota := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			quota["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			quota["description"] = *r.Description
		}

		quota["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			quota["id"] = *r.Id
		}

		if r.Name != nil {
			quota["name"] = *r.Name
		}

		quota["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			quota["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, quota)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LimitsQuotasDataSource().Schema["quotas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("quotas", resources); err != nil {
		return err
	}

	return nil
}
