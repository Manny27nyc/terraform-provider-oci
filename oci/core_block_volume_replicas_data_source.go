// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterDatasource("oci_core_block_volume_replicas", CoreBlockVolumeReplicasDataSource())
}

func CoreBlockVolumeReplicasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreBlockVolumeReplicas,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"block_volume_replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"block_volume_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_synced": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreBlockVolumeReplicas(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBlockVolumeReplicasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

type CoreBlockVolumeReplicasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListBlockVolumeReplicasResponse
}

func (s *CoreBlockVolumeReplicasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBlockVolumeReplicasDataSourceCrud) Get() error {
	request := oci_core.ListBlockVolumeReplicasRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.BlockVolumeReplicaLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.ListBlockVolumeReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBlockVolumeReplicas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreBlockVolumeReplicasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreBlockVolumeReplicasDataSource-", CoreBlockVolumeReplicasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		blockVolumeReplica := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.BlockVolumeId != nil {
			blockVolumeReplica["block_volume_id"] = *r.BlockVolumeId
		}

		if r.DefinedTags != nil {
			blockVolumeReplica["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			blockVolumeReplica["display_name"] = *r.DisplayName
		}

		blockVolumeReplica["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			blockVolumeReplica["id"] = *r.Id
		}

		if r.SizeInGBs != nil {
			blockVolumeReplica["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		blockVolumeReplica["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			blockVolumeReplica["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastSynced != nil {
			blockVolumeReplica["time_last_synced"] = r.TimeLastSynced.String()
		}

		resources = append(resources, blockVolumeReplica)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreBlockVolumeReplicasDataSource().Schema["block_volume_replicas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("block_volume_replicas", resources); err != nil {
		return err
	}

	return nil
}
