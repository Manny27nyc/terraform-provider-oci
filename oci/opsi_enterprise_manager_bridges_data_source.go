// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v52/opsi"
)

func init() {
	RegisterDatasource("oci_opsi_enterprise_manager_bridges", OpsiEnterpriseManagerBridgesDataSource())
}

func OpsiEnterpriseManagerBridgesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiEnterpriseManagerBridges,
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
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enterprise_manager_bridge_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(OpsiEnterpriseManagerBridgeResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiEnterpriseManagerBridges(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).operationsInsightsClient()

	return ReadResource(sync)
}

type OpsiEnterpriseManagerBridgesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListEnterpriseManagerBridgesResponse
}

func (s *OpsiEnterpriseManagerBridgesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiEnterpriseManagerBridgesDataSourceCrud) Get() error {
	request := oci_opsi.ListEnterpriseManagerBridgesRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.LifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.LifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListEnterpriseManagerBridges(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEnterpriseManagerBridges(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiEnterpriseManagerBridgesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("OpsiEnterpriseManagerBridgesDataSource-", OpsiEnterpriseManagerBridgesDataSource(), s.D))
	resources := []map[string]interface{}{}
	enterpriseManagerBridge := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EnterpriseManagerBridgeSummaryToMap(item))
	}
	enterpriseManagerBridge["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, OpsiEnterpriseManagerBridgesDataSource().Schema["enterprise_manager_bridge_collection"].Elem.(*schema.Resource).Schema)
		enterpriseManagerBridge["items"] = items
	}

	resources = append(resources, enterpriseManagerBridge)
	if err := s.D.Set("enterprise_manager_bridge_collection", resources); err != nil {
		return err
	}

	return nil
}
