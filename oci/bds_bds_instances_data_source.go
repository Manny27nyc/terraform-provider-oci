// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v52/bds"
)

func init() {
	RegisterDatasource("oci_bds_bds_instances", BdsBdsInstancesDataSource())
}

func BdsBdsInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstances,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(BdsBdsInstanceResource()),
			},
		},
	}
}

func readBdsBdsInstances(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).bdsClient()

	return ReadResource(sync)
}

type BdsBdsInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsInstancesResponse
}

func (s *BdsBdsInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstancesDataSourceCrud) Get() error {
	request := oci_bds.ListBdsInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	var displayName = ""
	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		displayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsInstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsInstances(context.Background(), request)
	if err != nil {
		return err
	}

	if displayName != "" {
		bdInstances := make([]oci_bds.BdsInstanceSummary, 0)
		for _, bdsInstance := range response.Items {
			if bdsInstance.DisplayName != nil && *bdsInstance.DisplayName == displayName {
				bdInstances = append(bdInstances, bdsInstance)
			}
		}
		response.Items = bdInstances
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsInstances(context.Background(), request)
		if err != nil {
			return err
		}
		if displayName != "" {
			bdInstances := make([]oci_bds.BdsInstanceSummary, 0)

			for _, bdsInstance := range listResponse.Items {
				if bdsInstance.DisplayName != nil && *bdsInstance.DisplayName == displayName {
					bdInstances = append(bdInstances, bdsInstance)
				}
			}
			listResponse.Items = bdInstances
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("BdsBdsInstancesDataSource-", BdsBdsInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		bdsInstance["cluster_version"] = r.ClusterVersion

		if r.DefinedTags != nil {
			bdsInstance["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			bdsInstance["display_name"] = *r.DisplayName
		}

		bdsInstance["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			bdsInstance["id"] = *r.Id
		}

		if r.IsCloudSqlConfigured != nil {
			bdsInstance["is_cloud_sql_configured"] = *r.IsCloudSqlConfigured
		}

		if r.IsHighAvailability != nil {
			bdsInstance["is_high_availability"] = *r.IsHighAvailability
		}

		if r.IsSecure != nil {
			bdsInstance["is_secure"] = *r.IsSecure
		}

		if r.NumberOfNodes != nil {
			bdsInstance["number_of_nodes"] = *r.NumberOfNodes
		}

		bdsInstance["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstance["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, bdsInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, BdsBdsInstancesDataSource().Schema["bds_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bds_instances", resources); err != nil {
		return err
	}

	return nil
}
