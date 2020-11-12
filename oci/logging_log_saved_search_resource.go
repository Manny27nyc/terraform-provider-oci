// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_logging "github.com/oracle/oci-go-sdk/v28/logging"
)

func init() {
	RegisterResource("oci_logging_log_saved_search", LoggingLogSavedSearchResource())
}

func LoggingLogSavedSearchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createLoggingLogSavedSearch,
		Read:     readLoggingLogSavedSearch,
		Update:   updateLoggingLogSavedSearch,
		Delete:   deleteLoggingLogSavedSearch,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoggingLogSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return CreateResource(d, sync)
}

func readLoggingLogSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return ReadResource(sync)
}

func updateLoggingLogSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return UpdateResource(d, sync)
}

func deleteLoggingLogSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type LoggingLogSavedSearchResourceCrud struct {
	BaseCrud
	Client                 *oci_logging.LoggingManagementClient
	Res                    *oci_logging.LogSavedSearch
	DisableNotFoundRetries bool
}

func (s *LoggingLogSavedSearchResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LoggingLogSavedSearchResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_logging.LogSavedSearchLifecycleStateCreating),
	}
}

func (s *LoggingLogSavedSearchResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_logging.LogSavedSearchLifecycleStateActive),
	}
}

func (s *LoggingLogSavedSearchResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_logging.LogSavedSearchLifecycleStateDeleting),
	}
}

func (s *LoggingLogSavedSearchResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LoggingLogSavedSearchResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_logging.LogSavedSearchLifecycleStateUpdating),
	}
}

func (s *LoggingLogSavedSearchResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_logging.LogSavedSearchLifecycleStateActive),
	}
}

func (s *LoggingLogSavedSearchResourceCrud) Create() error {
	request := oci_logging.CreateLogSavedSearchRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if query, ok := s.D.GetOkExists("query"); ok {
		tmp := query.(string)
		request.Query = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.CreateLogSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogSavedSearch
	return nil
}

func (s *LoggingLogSavedSearchResourceCrud) Get() error {
	request := oci_logging.GetLogSavedSearchRequest{}

	tmp := s.D.Id()
	request.LogSavedSearchId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.GetLogSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogSavedSearch
	return nil
}

func (s *LoggingLogSavedSearchResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_logging.UpdateLogSavedSearchRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LogSavedSearchId = &tmp

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if query, ok := s.D.GetOkExists("query"); ok {
		tmp := query.(string)
		request.Query = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.UpdateLogSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogSavedSearch
	return nil
}

func (s *LoggingLogSavedSearchResourceCrud) Delete() error {
	request := oci_logging.DeleteLogSavedSearchRequest{}

	tmp := s.D.Id()
	request.LogSavedSearchId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "logging")

	_, err := s.Client.DeleteLogSavedSearch(context.Background(), request)
	return err
}

func (s *LoggingLogSavedSearchResourceCrud) SetData() error {
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Query != nil {
		s.D.Set("query", *s.Res.Query)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func LogSavedSearchSummaryToMap(obj oci_logging.LogSavedSearchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}

func (s *LoggingLogSavedSearchResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_logging.ChangeLogSavedSearchCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LogSavedSearchId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "logging")

	_, err := s.Client.ChangeLogSavedSearchCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
