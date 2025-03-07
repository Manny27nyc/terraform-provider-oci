// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v52/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_listing_package_agreements", MarketplaceListingPackageAgreementsDataSource())
}

func MarketplaceListingPackageAgreementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceListingPackageAgreements,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"agreements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"author": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"content_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prompt": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplaceListingPackageAgreements(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageAgreementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplaceListingPackageAgreementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListAgreementsResponse
}

func (s *MarketplaceListingPackageAgreementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingPackageAgreementsDataSourceCrud) Get() error {
	request := oci_marketplace.ListAgreementsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListAgreements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgreements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceListingPackageAgreementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MarketplaceListingPackageAgreementsDataSource-", MarketplaceListingPackageAgreementsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listingPackageAgreement := map[string]interface{}{}

		listingPackageAgreement["author"] = r.Author

		if r.ContentUrl != nil {
			listingPackageAgreement["content_url"] = *r.ContentUrl
		}

		if r.Id != nil {
			listingPackageAgreement["id"] = *r.Id
		}

		if r.Prompt != nil {
			listingPackageAgreement["prompt"] = *r.Prompt
		}

		resources = append(resources, listingPackageAgreement)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MarketplaceListingPackageAgreementsDataSource().Schema["agreements"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("agreements", resources); err != nil {
		return err
	}

	return nil
}
