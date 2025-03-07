// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v52/certificatesmanagement"
)

func init() {
	RegisterDatasource("oci_certificates_management_ca_bundle", CertificatesManagementCaBundleDataSource())
}

func CertificatesManagementCaBundleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ca_bundle_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CertificatesManagementCaBundleResource(), fieldMap, readSingularCertificatesManagementCaBundle)
}

func readSingularCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).certificatesManagementClient()

	return ReadResource(sync)
}

type CertificatesManagementCaBundleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetCaBundleResponse
}

func (s *CertificatesManagementCaBundleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCaBundleDataSourceCrud) Get() error {
	request := oci_certificates_management.GetCaBundleRequest{}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetCaBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementCaBundleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
