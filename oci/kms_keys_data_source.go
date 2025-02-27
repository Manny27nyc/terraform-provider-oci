// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v52/keymanagement"
)

func init() {
	RegisterDatasource("oci_kms_keys", KmsKeysDataSource())
}

func KmsKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsKeys,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"curve_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(KmsKeyResource()),
			},
		},
	}
}

func readKmsKeys(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeysDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

type KmsKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.ListKeysResponse
}

func (s *KmsKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsKeysDataSourceCrud) Get() error {
	request := oci_kms.ListKeysRequest{}

	if algorithm, ok := s.D.GetOkExists("algorithm"); ok {
		request.Algorithm = oci_kms.ListKeysAlgorithmEnum(algorithm.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if curveId, ok := s.D.GetOkExists("curve_id"); ok {
		request.CurveId = oci_kms.ListKeysCurveIdEnum(curveId.(string))
	}

	if length, ok := s.D.GetOkExists("length"); ok {
		tmp := length.(int)
		request.Length = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_kms.ListKeysProtectionModeEnum(protectionMode.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "kms")

	response, err := s.Client.ListKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *KmsKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("KmsKeysDataSource-", KmsKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		key := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			key["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			key["display_name"] = *r.DisplayName
		}

		key["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			key["id"] = *r.Id
		}

		key["protection_mode"] = r.ProtectionMode

		key["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			key["time_created"] = r.TimeCreated.String()
		}

		if r.VaultId != nil {
			key["vault_id"] = *r.VaultId
		}

		resources = append(resources, key)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, KmsKeysDataSource().Schema["keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("keys", resources); err != nil {
		return err
	}

	return nil
}
