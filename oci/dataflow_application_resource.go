// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_dataflow "github.com/oracle/oci-go-sdk/v52/dataflow"
)

func init() {
	RegisterResource("oci_dataflow_application", DataflowApplicationResource())
}

func DataflowApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDataflowApplication,
		Read:     readDataflowApplication,
		Update:   updateDataflowApplication,
		Delete:   deleteDataflowApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"driver_shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"executor_shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"num_executors": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"spark_version": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"archive_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"arguments": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"class_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
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
			"execute": {
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
			"logs_bucket_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metastore_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warehouse_bucket_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"owner_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_user_name": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataflowApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient()

	return CreateResource(d, sync)
}

func readDataflowApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient()

	return ReadResource(sync)
}

func updateDataflowApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient()

	return UpdateResource(d, sync)
}

func deleteDataflowApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DataflowApplicationResourceCrud struct {
	BaseCrud
	Client                 *oci_dataflow.DataFlowClient
	Res                    *oci_dataflow.Application
	DisableNotFoundRetries bool
}

func (s *DataflowApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataflowApplicationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DataflowApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataflow.ApplicationLifecycleStateActive),
	}
}

func (s *DataflowApplicationResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DataflowApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataflow.ApplicationLifecycleStateDeleted),
	}
}

func (s *DataflowApplicationResourceCrud) Create() error {
	request := oci_dataflow.CreateApplicationRequest{}

	if archiveUri, ok := s.D.GetOkExists("archive_uri"); ok {
		tmp := archiveUri.(string)
		request.ArchiveUri = &tmp
	}

	if arguments, ok := s.D.GetOkExists("arguments"); ok {
		interfaces := arguments.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("arguments") {
			request.Arguments = tmp
		}
	}

	if className, ok := s.D.GetOkExists("class_name"); ok {
		tmp := className.(string)
		request.ClassName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		request.Configuration = ObjectMapToStringMap(configuration.(map[string]interface{}))
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if driverShape, ok := s.D.GetOkExists("driver_shape"); ok {
		tmp := driverShape.(string)
		request.DriverShape = &tmp
	}

	if execute, ok := s.D.GetOkExists("execute"); ok {
		tmp := execute.(string)
		request.Execute = &tmp
	}

	if executorShape, ok := s.D.GetOkExists("executor_shape"); ok {
		tmp := executorShape.(string)
		request.ExecutorShape = &tmp
	}

	if fileUri, ok := s.D.GetOkExists("file_uri"); ok {
		tmp := fileUri.(string)
		request.FileUri = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if language, ok := s.D.GetOkExists("language"); ok {
		request.Language = oci_dataflow.ApplicationLanguageEnum(language.(string))
	}

	if logsBucketUri, ok := s.D.GetOkExists("logs_bucket_uri"); ok {
		tmp := logsBucketUri.(string)
		request.LogsBucketUri = &tmp
	}

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	if numExecutors, ok := s.D.GetOkExists("num_executors"); ok {
		tmp := numExecutors.(int)
		request.NumExecutors = &tmp
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_dataflow.ApplicationParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToApplicationParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if sparkVersion, ok := s.D.GetOkExists("spark_version"); ok {
		tmp := sparkVersion.(string)
		request.SparkVersion = &tmp
	}

	if warehouseBucketUri, ok := s.D.GetOkExists("warehouse_bucket_uri"); ok {
		tmp := warehouseBucketUri.(string)
		request.WarehouseBucketUri = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataflowApplicationResourceCrud) Get() error {
	request := oci_dataflow.GetApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataflowApplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dataflow.UpdateApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	if archiveUri, ok := s.D.GetOkExists("archive_uri"); ok {
		tmp := archiveUri.(string)
		request.ArchiveUri = &tmp
	}

	if arguments, ok := s.D.GetOkExists("arguments"); ok {
		interfaces := arguments.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("arguments") {
			request.Arguments = tmp
		}
	}

	if className, ok := s.D.GetOkExists("class_name"); ok {
		tmp := className.(string)
		request.ClassName = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		request.Configuration = ObjectMapToStringMap(configuration.(map[string]interface{}))
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if driverShape, ok := s.D.GetOkExists("driver_shape"); ok {
		tmp := driverShape.(string)
		request.DriverShape = &tmp
	}

	if execute, ok := s.D.GetOkExists("execute"); ok {
		tmp := execute.(string)
		request.Execute = &tmp
	}

	if executorShape, ok := s.D.GetOkExists("executor_shape"); ok {
		tmp := executorShape.(string)
		request.ExecutorShape = &tmp
	}

	if fileUri, ok := s.D.GetOkExists("file_uri"); ok {
		tmp := fileUri.(string)
		request.FileUri = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if language, ok := s.D.GetOkExists("language"); ok {
		request.Language = oci_dataflow.ApplicationLanguageEnum(language.(string))
	}

	if logsBucketUri, ok := s.D.GetOkExists("logs_bucket_uri"); ok {
		tmp := logsBucketUri.(string)
		request.LogsBucketUri = &tmp
	}

	if metastoreId, ok := s.D.GetOkExists("metastore_id"); ok {
		tmp := metastoreId.(string)
		request.MetastoreId = &tmp
	}

	if numExecutors, ok := s.D.GetOkExists("num_executors"); ok {
		tmp := numExecutors.(int)
		request.NumExecutors = &tmp
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_dataflow.ApplicationParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToApplicationParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if sparkVersion, ok := s.D.GetOkExists("spark_version"); ok {
		tmp := sparkVersion.(string)
		request.SparkVersion = &tmp
	}

	if warehouseBucketUri, ok := s.D.GetOkExists("warehouse_bucket_uri"); ok {
		tmp := warehouseBucketUri.(string)
		request.WarehouseBucketUri = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataflowApplicationResourceCrud) Delete() error {
	request := oci_dataflow.DeleteApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.DeleteApplication(context.Background(), request)
	return err
}

func (s *DataflowApplicationResourceCrud) SetData() error {
	if s.Res.ArchiveUri != nil {
		s.D.Set("archive_uri", *s.Res.ArchiveUri)
	}

	s.D.Set("arguments", s.Res.Arguments)

	if s.Res.ClassName != nil {
		s.D.Set("class_name", *s.Res.ClassName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration", s.Res.Configuration)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DriverShape != nil {
		s.D.Set("driver_shape", *s.Res.DriverShape)
	}

	if s.Res.Execute != nil {
		s.D.Set("execute", *s.Res.Execute)
	}

	if s.Res.ExecutorShape != nil {
		s.D.Set("executor_shape", *s.Res.ExecutorShape)
	}

	if s.Res.FileUri != nil {
		s.D.Set("file_uri", *s.Res.FileUri)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("language", s.Res.Language)

	if s.Res.LogsBucketUri != nil {
		s.D.Set("logs_bucket_uri", *s.Res.LogsBucketUri)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	if s.Res.NumExecutors != nil {
		s.D.Set("num_executors", *s.Res.NumExecutors)
	}

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, ApplicationParameterToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.SparkVersion != nil {
		s.D.Set("spark_version", *s.Res.SparkVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WarehouseBucketUri != nil {
		s.D.Set("warehouse_bucket_uri", *s.Res.WarehouseBucketUri)
	}

	return nil
}

func (s *DataflowApplicationResourceCrud) mapToApplicationParameter(fieldKeyFormat string) (oci_dataflow.ApplicationParameter, error) {
	result := oci_dataflow.ApplicationParameter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func ApplicationParameterToMap(obj oci_dataflow.ApplicationParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DataflowApplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dataflow.ChangeApplicationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApplicationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.ChangeApplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
