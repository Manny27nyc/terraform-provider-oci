// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v52/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/v52/monitoring"
)

func init() {
	RegisterResource("oci_monitoring_alarm", MonitoringAlarmResource())
}

func MonitoringAlarmResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createMonitoringAlarm,
		Read:     readMonitoringAlarm,
		Update:   updateMonitoringAlarm,
		Delete:   deleteMonitoringAlarm,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destinations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"metric_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"body": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"message_format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"pending_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"repeat_notification_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolution": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"suppression": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"time_suppress_from": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: TimeDiffSuppressFunction,
						},
						"time_suppress_until": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: TimeDiffSuppressFunction,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient()

	return CreateResource(d, sync)
}

func readMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient()

	return ReadResource(sync)
}

func updateMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient()

	return UpdateResource(d, sync)
}

func deleteMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type MonitoringAlarmResourceCrud struct {
	BaseCrud
	Client                 *oci_monitoring.MonitoringClient
	Res                    *oci_monitoring.Alarm
	DisableNotFoundRetries bool
}

func (s *MonitoringAlarmResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MonitoringAlarmResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MonitoringAlarmResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_monitoring.AlarmLifecycleStateActive),
	}
}

func (s *MonitoringAlarmResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_monitoring.AlarmLifecycleStateDeleting),
	}
}

func (s *MonitoringAlarmResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_monitoring.AlarmLifecycleStateDeleted),
	}
}

func (s *MonitoringAlarmResourceCrud) Create() error {
	request := oci_monitoring.CreateAlarmRequest{}

	if body, ok := s.D.GetOkExists("body"); ok {
		tmp := body.(string)
		request.Body = &tmp
	}

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

	if destinations, ok := s.D.GetOkExists("destinations"); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("destinations") {
			request.Destinations = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if messageFormat, ok := s.D.GetOkExists("message_format"); ok {
		request.MessageFormat = oci_monitoring.CreateAlarmDetailsMessageFormatEnum(messageFormat.(string))
	}

	if metricCompartmentId, ok := s.D.GetOkExists("metric_compartment_id"); ok {
		tmp := metricCompartmentId.(string)
		request.MetricCompartmentId = &tmp
	}

	if metricCompartmentIdInSubtree, ok := s.D.GetOkExists("metric_compartment_id_in_subtree"); ok {
		tmp := metricCompartmentIdInSubtree.(bool)
		request.MetricCompartmentIdInSubtree = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if pendingDuration, ok := s.D.GetOkExists("pending_duration"); ok {
		tmp := pendingDuration.(string)
		request.PendingDuration = &tmp
	}

	if query, ok := s.D.GetOkExists("query"); ok {
		tmp := query.(string)
		request.Query = &tmp
	}

	if repeatNotificationDuration, ok := s.D.GetOkExists("repeat_notification_duration"); ok {
		tmp := repeatNotificationDuration.(string)
		request.RepeatNotificationDuration = &tmp
	}

	if resolution, ok := s.D.GetOkExists("resolution"); ok {
		tmp := resolution.(string)
		request.Resolution = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_monitoring.AlarmSeverityEnum(severity.(string))
	}

	if suppression, ok := s.D.GetOkExists("suppression"); ok {
		if tmpList := suppression.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "suppression", 0)
			tmp, err := s.mapToSuppression(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Suppression = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "monitoring")

	response, err := s.Client.CreateAlarm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alarm
	return nil
}

func (s *MonitoringAlarmResourceCrud) Get() error {
	request := oci_monitoring.GetAlarmRequest{}

	tmp := s.D.Id()
	request.AlarmId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "monitoring")

	response, err := s.Client.GetAlarm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alarm
	return nil
}

func (s *MonitoringAlarmResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_monitoring.UpdateAlarmRequest{}

	tmp := s.D.Id()
	request.AlarmId = &tmp

	if body, ok := s.D.GetOkExists("body"); ok {
		tmp := body.(string)
		request.Body = &tmp
	}

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

	if destinations, ok := s.D.GetOkExists("destinations"); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("destinations") {
			request.Destinations = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if messageFormat, ok := s.D.GetOkExists("message_format"); ok {
		request.MessageFormat = oci_monitoring.UpdateAlarmDetailsMessageFormatEnum(messageFormat.(string))
	}

	if metricCompartmentId, ok := s.D.GetOkExists("metric_compartment_id"); ok {
		tmp := metricCompartmentId.(string)
		request.MetricCompartmentId = &tmp
	}

	if metricCompartmentIdInSubtree, ok := s.D.GetOkExists("metric_compartment_id_in_subtree"); ok {
		tmp := metricCompartmentIdInSubtree.(bool)
		request.MetricCompartmentIdInSubtree = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if pendingDuration, ok := s.D.GetOkExists("pending_duration"); ok {
		tmp := pendingDuration.(string)
		request.PendingDuration = &tmp
	}

	if query, ok := s.D.GetOkExists("query"); ok {
		tmp := query.(string)
		request.Query = &tmp
	}

	if repeatNotificationDuration, ok := s.D.GetOkExists("repeat_notification_duration"); ok {
		tmp := repeatNotificationDuration.(string)
		request.RepeatNotificationDuration = &tmp
	}

	if resolution, ok := s.D.GetOkExists("resolution"); ok {
		tmp := resolution.(string)
		request.Resolution = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_monitoring.AlarmSeverityEnum(severity.(string))
	}

	if suppression, ok := s.D.GetOkExists("suppression"); ok {
		if tmpList := suppression.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "suppression", 0)
			tmp, err := s.mapToSuppression(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Suppression = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "monitoring")

	response, err := s.Client.UpdateAlarm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alarm
	return nil
}

func (s *MonitoringAlarmResourceCrud) Delete() error {
	request := oci_monitoring.DeleteAlarmRequest{}

	tmp := s.D.Id()
	request.AlarmId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "monitoring")

	_, err := s.Client.DeleteAlarm(context.Background(), request)
	return err
}

func (s *MonitoringAlarmResourceCrud) SetData() error {
	if s.Res.Body != nil {
		s.D.Set("body", *s.Res.Body)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("destinations", s.Res.Destinations)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("message_format", s.Res.MessageFormat)

	if s.Res.MetricCompartmentId != nil {
		s.D.Set("metric_compartment_id", *s.Res.MetricCompartmentId)
	}

	if s.Res.MetricCompartmentIdInSubtree != nil {
		s.D.Set("metric_compartment_id_in_subtree", *s.Res.MetricCompartmentIdInSubtree)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.PendingDuration != nil {
		s.D.Set("pending_duration", *s.Res.PendingDuration)
	}

	if s.Res.Query != nil {
		s.D.Set("query", *s.Res.Query)
	}

	if s.Res.RepeatNotificationDuration != nil {
		s.D.Set("repeat_notification_duration", *s.Res.RepeatNotificationDuration)
	}

	if s.Res.Resolution != nil {
		s.D.Set("resolution", *s.Res.Resolution)
	}

	if s.Res.ResourceGroup != nil {
		s.D.Set("resource_group", *s.Res.ResourceGroup)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Suppression != nil {
		s.D.Set("suppression", []interface{}{SuppressionToMap(s.Res.Suppression)})
	} else {
		s.D.Set("suppression", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MonitoringAlarmResourceCrud) mapToSuppression(fieldKeyFormat string) (oci_monitoring.Suppression, error) {
	result := oci_monitoring.Suppression{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if timeSuppressFrom, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_suppress_from")); ok {
		tmp, err := time.Parse(time.RFC3339, timeSuppressFrom.(string))
		if err != nil {
			return result, err
		}
		result.TimeSuppressFrom = &oci_common.SDKTime{Time: tmp}
	}

	if timeSuppressUntil, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_suppress_until")); ok {
		tmp, err := time.Parse(time.RFC3339, timeSuppressUntil.(string))
		if err != nil {
			return result, err
		}
		result.TimeSuppressUntil = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func SuppressionToMap(obj *oci_monitoring.Suppression) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.TimeSuppressFrom != nil {
		result["time_suppress_from"] = obj.TimeSuppressFrom.Format(time.RFC3339Nano)
	}

	if obj.TimeSuppressUntil != nil {
		result["time_suppress_until"] = obj.TimeSuppressUntil.Format(time.RFC3339Nano)
	}

	return result
}

func (s *MonitoringAlarmResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_monitoring.ChangeAlarmCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AlarmId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "monitoring")

	_, err := s.Client.ChangeAlarmCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
