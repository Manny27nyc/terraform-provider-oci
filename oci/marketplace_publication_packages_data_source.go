// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v52/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_publication_packages", MarketplacePublicationPackagesDataSource())
}

func MarketplacePublicationPackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplacePublicationPackages,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"package_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publication_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"publication_packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"listing_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplacePublicationPackages(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationPackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplacePublicationPackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListPublicationPackagesResponse
}

func (s *MarketplacePublicationPackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplacePublicationPackagesDataSourceCrud) Get() error {
	request := oci_marketplace.ListPublicationPackagesRequest{}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		tmp := packageType.(string)
		request.PackageType = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	if publicationId, ok := s.D.GetOkExists("publication_id"); ok {
		tmp := publicationId.(string)
		request.PublicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListPublicationPackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublicationPackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplacePublicationPackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MarketplacePublicationPackagesDataSource-", MarketplacePublicationPackagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		publicationPackage := map[string]interface{}{}

		if r.ListingId != nil {
			publicationPackage["listing_id"] = *r.ListingId
		}

		publicationPackage["package_type"] = r.PackageType

		if r.PackageVersion != nil {
			publicationPackage["package_version"] = *r.PackageVersion
		}

		if r.ResourceId != nil {
			publicationPackage["resource_id"] = *r.ResourceId
		}

		if r.TimeCreated != nil {
			publicationPackage["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, publicationPackage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MarketplacePublicationPackagesDataSource().Schema["publication_packages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("publication_packages", resources); err != nil {
		return err
	}

	return nil
}
