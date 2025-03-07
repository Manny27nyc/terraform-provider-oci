// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterDatasource("oci_core_compute_image_capability_schemas", CoreComputeImageCapabilitySchemasDataSource())
}

func CoreComputeImageCapabilitySchemasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeImageCapabilitySchemas,
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
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_image_capability_schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreComputeImageCapabilitySchemaResource()),
			},
		},
	}
}

func readCoreComputeImageCapabilitySchemas(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreComputeImageCapabilitySchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeImageCapabilitySchemasResponse
}

func (s *CoreComputeImageCapabilitySchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeImageCapabilitySchemasDataSourceCrud) Get() error {
	request := oci_core.ListComputeImageCapabilitySchemasRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeImageCapabilitySchemas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeImageCapabilitySchemas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeImageCapabilitySchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreComputeImageCapabilitySchemasDataSource-", CoreComputeImageCapabilitySchemasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeImageCapabilitySchema := map[string]interface{}{}

		if r.CompartmentId != nil {
			computeImageCapabilitySchema["compartment_id"] = *r.CompartmentId
		}

		if r.ComputeGlobalImageCapabilitySchemaVersionName != nil {
			computeImageCapabilitySchema["compute_global_image_capability_schema_version_name"] = *r.ComputeGlobalImageCapabilitySchemaVersionName
		}

		if r.DefinedTags != nil {
			computeImageCapabilitySchema["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			computeImageCapabilitySchema["display_name"] = *r.DisplayName
		}

		computeImageCapabilitySchema["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			computeImageCapabilitySchema["id"] = *r.Id
		}

		if r.ImageId != nil {
			computeImageCapabilitySchema["image_id"] = *r.ImageId
		}

		computeImageCapabilitySchema["schema_data"] = r.SchemaData

		if r.TimeCreated != nil {
			computeImageCapabilitySchema["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, computeImageCapabilitySchema)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreComputeImageCapabilitySchemasDataSource().Schema["compute_image_capability_schemas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_image_capability_schemas", resources); err != nil {
		return err
	}

	return nil
}
