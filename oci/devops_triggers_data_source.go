// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v52/devops"
)

func init() {
	RegisterDatasource("oci_devops_triggers", DevopsTriggersDataSource())
}

func DevopsTriggersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsTriggers,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trigger_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(DevopsTriggerResource()),
						},
					},
				},
			},
		},
	}
}

func readDevopsTriggers(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsTriggersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsTriggersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListTriggersResponse
}

func (s *DevopsTriggersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsTriggersDataSourceCrud) Get() error {
	request := oci_devops.ListTriggersRequest{}

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

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.TriggerLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")

	response, err := s.Client.ListTriggers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTriggers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsTriggersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DevopsTriggersDataSource-", DevopsTriggersDataSource(), s.D))
	resources := []map[string]interface{}{}
	trigger := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TriggerSummaryToMap(item))
	}
	trigger["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DevopsTriggersDataSource().Schema["trigger_collection"].Elem.(*schema.Resource).Schema)
		trigger["items"] = items
	}

	resources = append(resources, trigger)
	if err := s.D.Set("trigger_collection", resources); err != nil {
		return err
	}

	return nil
}
