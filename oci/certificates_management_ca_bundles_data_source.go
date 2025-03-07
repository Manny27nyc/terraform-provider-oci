// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v52/certificatesmanagement"
)

func init() {
	RegisterDatasource("oci_certificates_management_ca_bundles", CertificatesManagementCaBundlesDataSource())
}

func CertificatesManagementCaBundlesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementCaBundles,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"ca_bundle_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "compartment_id"},
			},
			"compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ca_bundle_id"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ca_bundle_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca_bundle_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(CertificatesManagementCaBundleResource()),
						},
					},
				},
			},
		},
	}
}

func readCertificatesManagementCaBundles(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundlesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).certificatesManagementClient()

	return ReadResource(sync)
}

type CertificatesManagementCaBundlesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListCaBundlesResponse
}

func (s *CertificatesManagementCaBundlesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCaBundlesDataSourceCrud) Get() error {
	request := oci_certificates_management.ListCaBundlesRequest{}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_certificates_management.ListCaBundlesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.ListCaBundles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCaBundles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementCaBundlesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CertificatesManagementCaBundlesDataSource-", CertificatesManagementCaBundlesDataSource(), s.D))
	resources := []map[string]interface{}{}
	caBundle := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CaBundleSummaryToMap(item))
	}
	caBundle["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementCaBundlesDataSource().Schema["ca_bundle_collection"].Elem.(*schema.Resource).Schema)
		caBundle["items"] = items
	}

	resources = append(resources, caBundle)
	if err := s.D.Set("ca_bundle_collection", resources); err != nil {
		return err
	}

	return nil
}
