// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v52/dns"
)

func init() {
	RegisterResource("oci_dns_steering_policy_attachment", DnsSteeringPolicyAttachmentResource())
}

func DnsSteeringPolicyAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDnsSteeringPolicyAttachment,
		Read:     readDnsSteeringPolicyAttachment,
		Update:   updateDnsSteeringPolicyAttachment,
		Delete:   deleteDnsSteeringPolicyAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"domain_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"steering_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rtypes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"self": {
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

func createDnsSteeringPolicyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return CreateResource(d, sync)
}

func readDnsSteeringPolicyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return ReadResource(sync)
}

func updateDnsSteeringPolicyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return UpdateResource(d, sync)
}

func deleteDnsSteeringPolicyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DnsSteeringPolicyAttachmentResourceCrud struct {
	BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.SteeringPolicyAttachment
	DisableNotFoundRetries bool
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.SteeringPolicyAttachmentLifecycleStateCreating),
	}
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.SteeringPolicyAttachmentLifecycleStateActive),
	}
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.SteeringPolicyAttachmentLifecycleStateDeleting),
	}
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) Create() error {
	request := oci_dns.CreateSteeringPolicyAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domainName, ok := s.D.GetOkExists("domain_name"); ok {
		tmp := domainName.(string)
		request.DomainName = &tmp
	}

	if steeringPolicyId, ok := s.D.GetOkExists("steering_policy_id"); ok {
		tmp := steeringPolicyId.(string)
		request.SteeringPolicyId = &tmp
	}

	if zoneId, ok := s.D.GetOkExists("zone_id"); ok {
		tmp := zoneId.(string)
		request.ZoneId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateSteeringPolicyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicyAttachment
	return nil
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) Get() error {
	request := oci_dns.GetSteeringPolicyAttachmentRequest{}

	tmp := s.D.Id()
	request.SteeringPolicyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetSteeringPolicyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicyAttachment
	return nil
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) Update() error {
	request := oci_dns.UpdateSteeringPolicyAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.SteeringPolicyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateSteeringPolicyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicyAttachment
	return nil
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) Delete() error {
	request := oci_dns.DeleteSteeringPolicyAttachmentRequest{}

	tmp := s.D.Id()
	request.SteeringPolicyAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteSteeringPolicyAttachment(context.Background(), request)
	return err
}

func (s *DnsSteeringPolicyAttachmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainName != nil {
		s.D.Set("domain_name", *s.Res.DomainName)
	}

	s.D.Set("rtypes", s.Res.Rtypes)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SteeringPolicyId != nil {
		s.D.Set("steering_policy_id", *s.Res.SteeringPolicyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.ZoneId != nil {
		s.D.Set("zone_id", *s.Res.ZoneId)
	}

	return nil
}
