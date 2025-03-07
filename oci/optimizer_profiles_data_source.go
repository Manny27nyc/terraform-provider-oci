// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v52/optimizer"
)

func init() {
	RegisterDatasource("oci_optimizer_profiles", OptimizerProfilesDataSource())
}

func OptimizerProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerProfiles,
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
			"profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(OptimizerProfileResource()),
						},
					},
				},
			},
		},
	}
}

func readOptimizerProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).optimizerClient()

	return ReadResource(sync)
}

type OptimizerProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListProfilesResponse
}

func (s *OptimizerProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerProfilesDataSourceCrud) Get() error {
	request := oci_optimizer.ListProfilesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_optimizer.ListProfilesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "optimizer")

	response, err := s.Client.ListProfiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("OptimizerProfilesDataSource-", OptimizerProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	profile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProfileSummaryToMap(item))
	}
	profile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerProfilesDataSource().Schema["profile_collection"].Elem.(*schema.Resource).Schema)
		profile["items"] = items
	}

	resources = append(resources, profile)
	if err := s.D.Set("profile_collection", resources); err != nil {
		return err
	}

	return nil
}
