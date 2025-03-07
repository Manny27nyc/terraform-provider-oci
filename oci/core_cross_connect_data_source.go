// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterDatasource("oci_core_cross_connect", CoreCrossConnectDataSource())
}

func CoreCrossConnectDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cross_connect_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CoreCrossConnectResource(), fieldMap, readSingularCoreCrossConnect)
}

func readSingularCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreCrossConnectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectResponse
}

func (s *CoreCrossConnectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreCrossConnectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CrossConnectGroupId != nil {
		s.D.Set("cross_connect_group_id", *s.Res.CrossConnectGroupId)
	}

	if s.Res.CustomerReferenceName != nil {
		s.D.Set("customer_reference_name", *s.Res.CustomerReferenceName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LocationName != nil {
		s.D.Set("location_name", *s.Res.LocationName)
	}

	if s.Res.MacsecProperties != nil {
		s.D.Set("macsec_properties", []interface{}{MacsecPropertiesToMap(s.Res.MacsecProperties)})
	} else {
		s.D.Set("macsec_properties", nil)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.PortSpeedShapeName != nil {
		s.D.Set("port_speed_shape_name", *s.Res.PortSpeedShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
