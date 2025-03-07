// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v52/core"
)

func init() {
	RegisterResource("oci_core_app_catalog_listing_resource_version_agreement", AppCatalogListingResourceVersionAgreementResource())
	RegisterResource("oci_core_listing_resource_version_agreement", AppCatalogListingResourceVersionAgreementResource())
}

func AppCatalogListingResourceVersionAgreementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createAppCatalogListingResourceVersionAgreement,
		Read:     readAppCatalogListingResourceVersionAgreement,
		Delete:   deleteAppCatalogListingResourceVersionAgreement,
		Schema: map[string]*schema.Schema{
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listing_resource_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Computed
			"eula_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_terms_of_use_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_retrieved": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAppCatalogListingResourceVersionAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &AppCatalogListingResourceVersionAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return CreateResource(d, sync)
}

func readAppCatalogListingResourceVersionAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &AppCatalogListingResourceVersionAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

func deleteAppCatalogListingResourceVersionAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &AppCatalogListingResourceVersionAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type AppCatalogListingResourceVersionAgreementResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.GetAppCatalogListingAgreementsResponse
	DisableNotFoundRetries bool
}

func (s *AppCatalogListingResourceVersionAgreementResourceCrud) ID() string {
	return s.Res.TimeRetrieved.Format(time.RFC3339Nano)
}

func (s *AppCatalogListingResourceVersionAgreementResourceCrud) Create() error {
	request := oci_core.GetAppCatalogListingAgreementsRequest{}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if resourceVersion, ok := s.D.GetOkExists("listing_resource_version"); ok {
		tmp := resourceVersion.(string)
		request.ResourceVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.GetAppCatalogListingAgreements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AppCatalogListingResourceVersionAgreementResourceCrud) Get() error {
	return nil
}

func (s *AppCatalogListingResourceVersionAgreementResourceCrud) Delete() error {
	return nil
}

func (s *AppCatalogListingResourceVersionAgreementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(s.Res.TimeRetrieved.String())

	if s.Res.EulaLink != nil {
		s.D.Set("eula_link", *s.Res.EulaLink)
	} else {
		s.D.Set("eula_link", "")
	}

	if s.Res.ListingResourceVersion != nil {
		s.D.Set("listing_resource_version", *s.Res.ListingResourceVersion)
	}

	if s.Res.OracleTermsOfUseLink != nil {
		s.D.Set("oracle_terms_of_use_link", *s.Res.OracleTermsOfUseLink)
	}

	if s.Res.Signature != nil {
		s.D.Set("signature", *s.Res.Signature)
	}

	if s.Res.TimeRetrieved != nil {
		s.D.Set("time_retrieved", s.Res.TimeRetrieved.Format(time.RFC3339Nano))
	}

	return nil
}
