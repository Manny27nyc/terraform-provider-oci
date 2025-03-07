// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_identity "github.com/oracle/oci-go-sdk/v52/identity"
)

func init() {
	RegisterResource("oci_identity_identity_provider", IdentityIdentityProviderResource())
}

func IdentityIdentityProviderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createIdentityIdentityProvider,
		Read:     readIdentityIdentityProvider,
		Update:   updateIdentityIdentityProvider,
		Delete:   deleteIdentityIdentityProvider,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"SAML2",
				}, true),
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"inactive_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redirect_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signing_certificate": {
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
		},
	}
}

func createIdentityIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return CreateResource(d, sync)
}

func readIdentityIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

func updateIdentityIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return UpdateResource(d, sync)
}

func deleteIdentityIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type IdentityIdentityProviderResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.IdentityProvider
	DisableNotFoundRetries bool
}

func (s *IdentityIdentityProviderResourceCrud) ID() string {
	identityProvider := *s.Res
	return *identityProvider.GetId()
}

func (s *IdentityIdentityProviderResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateCreating),
	}
}

func (s *IdentityIdentityProviderResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateActive),
	}
}

func (s *IdentityIdentityProviderResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateDeleting),
	}
}

func (s *IdentityIdentityProviderResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateDeleted),
	}
}

func (s *IdentityIdentityProviderResourceCrud) Create() error {
	request := oci_identity.CreateIdentityProviderRequest{}
	err := s.populateTopLevelPolymorphicCreateIdentityProviderRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityIdentityProviderResourceCrud) Get() error {
	request := oci_identity.GetIdentityProviderRequest{}

	tmp := s.D.Id()
	request.IdentityProviderId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityIdentityProviderResourceCrud) Update() error {
	request := oci_identity.UpdateIdentityProviderRequest{}
	err := s.populateTopLevelPolymorphicUpdateIdentityProviderRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityIdentityProviderResourceCrud) Delete() error {
	request := oci_identity.DeleteIdentityProviderRequest{}

	tmp := s.D.Id()
	request.IdentityProviderId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteIdentityProvider(context.Background(), request)
	return err
}

func (s *IdentityIdentityProviderResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_identity.Saml2IdentityProvider:
		s.D.Set("protocol", "SAML2")

		s.D.Set("freeform_attributes", v.FreeformAttributes)

		if v.Metadata != nil {
			s.D.Set("metadata", *v.Metadata)
		}

		if v.MetadataUrl != nil {
			s.D.Set("metadata_url", *v.MetadataUrl)
		}

		if v.RedirectUrl != nil {
			s.D.Set("redirect_url", *v.RedirectUrl)
		}

		if v.SigningCertificate != nil {
			s.D.Set("signing_certificate", *v.SigningCertificate)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", definedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.InactiveStatus != nil {
			s.D.Set("inactive_state", strconv.FormatInt(*v.InactiveStatus, 10))
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ProductType != nil {
			s.D.Set("product_type", *v.ProductType)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	default:
		log.Printf("[WARN] Received 'protocol' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *IdentityIdentityProviderResourceCrud) populateTopLevelPolymorphicCreateIdentityProviderRequest(request *oci_identity.CreateIdentityProviderRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}
	switch strings.ToLower(protocol) {
	case strings.ToLower("SAML2"):
		details := oci_identity.CreateSaml2IdentityProviderDetails{}
		if freeformAttributes, ok := s.D.GetOkExists("freeform_attributes"); ok {
			details.FreeformAttributes = ObjectMapToStringMap(freeformAttributes.(map[string]interface{}))
		}
		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			details.Metadata = &tmp
		}
		if metadataUrl, ok := s.D.GetOkExists("metadata_url"); ok {
			tmp := metadataUrl.(string)
			details.MetadataUrl = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if productType, ok := s.D.GetOkExists("product_type"); ok {
			details.ProductType = oci_identity.CreateIdentityProviderDetailsProductTypeEnum(productType.(string))
		}
		request.CreateIdentityProviderDetails = details
	default:
		return fmt.Errorf("unknown protocol '%v' was specified", protocol)
	}
	return nil
}

func (s *IdentityIdentityProviderResourceCrud) populateTopLevelPolymorphicUpdateIdentityProviderRequest(request *oci_identity.UpdateIdentityProviderRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}
	switch strings.ToLower(protocol) {
	case strings.ToLower("SAML2"):
		details := oci_identity.UpdateSaml2IdentityProviderDetails{}
		if freeformAttributes, ok := s.D.GetOkExists("freeform_attributes"); ok {
			details.FreeformAttributes = ObjectMapToStringMap(freeformAttributes.(map[string]interface{}))
		}
		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			details.Metadata = &tmp
		}
		if metadataUrl, ok := s.D.GetOkExists("metadata_url"); ok {
			tmp := metadataUrl.(string)
			details.MetadataUrl = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.IdentityProviderId = &tmp
		request.UpdateIdentityProviderDetails = details
	default:
		return fmt.Errorf("unknown protocol '%v' was specified", protocol)
	}
	return nil
}
