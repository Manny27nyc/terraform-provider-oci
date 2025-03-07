// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v52/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_load_balancers", LoadBalancerLoadBalancersDataSource())
}

func LoadBalancerLoadBalancersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerLoadBalancers,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(LoadBalancerLoadBalancerResource()),
			},
		},
	}
}

func readLoadBalancerLoadBalancers(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerLoadBalancersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListLoadBalancersResponse
}

func (s *LoadBalancerLoadBalancersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancersDataSourceCrud) Get() error {
	request := oci_load_balancer.ListLoadBalancersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if detail, ok := s.D.GetOkExists("detail"); ok {
		tmp := detail.(string)
		request.Detail = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_load_balancer.LoadBalancerLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListLoadBalancers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLoadBalancers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoadBalancerLoadBalancersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoadBalancerLoadBalancersDataSource-", LoadBalancerLoadBalancersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancer := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			loadBalancer["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			loadBalancer["display_name"] = *r.DisplayName
		}

		loadBalancer["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			loadBalancer["id"] = *r.Id
		}

		ipAddressDetails := []interface{}{}
		for _, item := range r.IpAddresses {
			ipAddressDetails = append(ipAddressDetails, IpAddressToMap(item))
		}
		loadBalancer["ip_address_details"] = ipAddressDetails

		ipAddresses := []string{}
		ipMode := "IPV4"
		for _, ad := range r.IpAddresses {
			if ad.IpAddress != nil {
				ipAddresses = append(ipAddresses, *ad.IpAddress)
			}
			tmp := *ad.IpAddress
			if !isIPV4(tmp) {
				ipMode = "IPV6"
			}
		}
		loadBalancer["ip_mode"] = ipMode
		loadBalancer["ip_addresses"] = ipAddresses

		if r.IsPrivate != nil {
			loadBalancer["is_private"] = *r.IsPrivate
		}

		loadBalancer["network_security_group_ids"] = r.NetworkSecurityGroupIds

		if r.ShapeName != nil {
			loadBalancer["shape"] = *r.ShapeName
		}

		if r.ShapeDetails != nil {
			loadBalancer["shape_details"] = []interface{}{ShapeDetailsToMap(r.ShapeDetails)}
		} else {
			loadBalancer["shape_details"] = nil
		}

		loadBalancer["state"] = r.LifecycleState

		loadBalancer["subnet_ids"] = r.SubnetIds

		if r.SystemTags != nil {
			loadBalancer["system_tags"] = systemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			loadBalancer["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, loadBalancer)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerLoadBalancersDataSource().Schema["load_balancers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("load_balancers", resources); err != nil {
		return err
	}

	return nil
}
