// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/oracle/oci-go-sdk/v52/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v52/datascience"
)

func init() {
	RegisterResource("oci_datascience_job", DatascienceJobResource())
}

func DatascienceJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatascienceJob,
		Read:     readDatascienceJob,
		Update:   updateDatascienceJob,
		Delete:   deleteDatascienceJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_configuration_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"job_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DEFAULT",
							}, true),
						},

						// Optional
						"command_line_arguments": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"environment_variables": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem:     schema.TypeString,
						},
						"maximum_runtime_in_minutes": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     ValidateInt64TypeString,
							DiffSuppressFunc: Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"job_infrastructure_configuration_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"job_infrastructure_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"STANDALONE",
							}, true),
						},
						"shape_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"job_artifact": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"artifact_content_length": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     ValidateInt64TypeString,
				DiffSuppressFunc: Int64StringDiffSuppressFunction,
			},
			"artifact_content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"delete_related_job_runs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
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
			"job_log_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"enable_auto_log_creation": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"enable_logging": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"log_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"artifact_content_md5": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"artifact_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"empty_artifact": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatascienceJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := d.GetOkExists("job_artifact"); ok {
		if e := sync.CreateArtifact(); e != nil {
			return e
		}
	}

	return ReadResource(sync)
}

func readDatascienceJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

func updateDatascienceJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return UpdateResource(d, sync)
}

func deleteDatascienceJob(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type HeadJobArtifact struct {
	ContentLength      *int64
	ContentDisposition *string
	ContentMd5         *string
	LastModified       *common.SDKTime
}

type DatascienceJobResourceCrud struct {
	BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.Job
	ArtifactHeadRes        *HeadJobArtifact
	DisableNotFoundRetries bool
}

func (s *DatascienceJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceJobResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatascienceJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.JobLifecycleStateCreating),
		string(oci_datascience.JobLifecycleStateActive),
	}
}

func (s *DatascienceJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.JobLifecycleStateDeleting),
	}
}

func (s *DatascienceJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.JobLifecycleStateDeleted),
	}
}

func (s *DatascienceJobResourceCrud) Create() error {
	request := oci_datascience.CreateJobRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobConfigurationDetails, ok := s.D.GetOkExists("job_configuration_details"); ok {
		if tmpList := jobConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_configuration_details", 0)
			tmp, err := s.mapToJobConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobConfigurationDetails = tmp
		}
	}

	if jobInfrastructureConfigurationDetails, ok := s.D.GetOkExists("job_infrastructure_configuration_details"); ok {
		if tmpList := jobInfrastructureConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_infrastructure_configuration_details", 0)
			tmp, err := s.mapToJobInfrastructureConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobInfrastructureConfigurationDetails = tmp
		}
	}

	if jobLogConfigurationDetails, ok := s.D.GetOkExists("job_log_configuration_details"); ok {
		if tmpList := jobLogConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_log_configuration_details", 0)
			tmp, err := s.mapToJobLogConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobLogConfigurationDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *DatascienceJobResourceCrud) Get() error {
	request := oci_datascience.GetJobRequest{}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	if emptyArtifact, ok := s.D.GetOkExists("empty_artifact"); ok {
		tmp := emptyArtifact.(bool)
		if !tmp {
			err := s.GetArtifactHead()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *DatascienceJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateJobRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.JobId = &tmp

	if jobInfrastructureConfigurationDetails, ok := s.D.GetOkExists("job_infrastructure_configuration_details"); ok {
		if tmpList := jobInfrastructureConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_infrastructure_configuration_details", 0)
			tmp, err := s.mapToJobInfrastructureConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobInfrastructureConfigurationDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *DatascienceJobResourceCrud) Delete() error {
	request := oci_datascience.DeleteJobRequest{}

	if deleteRelatedJobRuns, ok := s.D.GetOkExists("delete_related_job_runs"); ok {
		tmp := deleteRelatedJobRuns.(bool)
		request.DeleteRelatedJobRuns = &tmp
	}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteJob(context.Background(), request)
	return err
}

func (s *DatascienceJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.JobConfigurationDetails != nil {
		jobConfigurationDetailsArray := []interface{}{}
		if jobConfigurationDetailsMap := JobConfigurationDetailsToMap(&s.Res.JobConfigurationDetails); jobConfigurationDetailsMap != nil {
			jobConfigurationDetailsArray = append(jobConfigurationDetailsArray, jobConfigurationDetailsMap)
		}
		s.D.Set("job_configuration_details", jobConfigurationDetailsArray)
	} else {
		s.D.Set("job_configuration_details", nil)
	}

	if s.Res.JobInfrastructureConfigurationDetails != nil {
		jobInfrastructureConfigurationDetailsArray := []interface{}{}
		if jobInfrastructureConfigurationDetailsMap := JobInfrastructureConfigurationDetailsToMap(&s.Res.JobInfrastructureConfigurationDetails); jobInfrastructureConfigurationDetailsMap != nil {
			jobInfrastructureConfigurationDetailsArray = append(jobInfrastructureConfigurationDetailsArray, jobInfrastructureConfigurationDetailsMap)
		}
		s.D.Set("job_infrastructure_configuration_details", jobInfrastructureConfigurationDetailsArray)
	} else {
		s.D.Set("job_infrastructure_configuration_details", nil)
	}

	if s.Res.JobLogConfigurationDetails != nil {
		s.D.Set("job_log_configuration_details", []interface{}{JobLogConfigurationDetailsToMap(s.Res.JobLogConfigurationDetails)})
	} else {
		s.D.Set("job_log_configuration_details", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return s.SetArtifactData()
}

func (s *DatascienceJobResourceCrud) mapToJobConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobConfigurationDetails, error) {
	var baseObject oci_datascience.JobConfigurationDetails
	//discriminator
	jobTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_type"))
	var jobType string
	if ok {
		jobType = jobTypeRaw.(string)
	} else {
		jobType = "" // default value
	}
	switch strings.ToLower(jobType) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.DefaultJobConfigurationDetails{}
		if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
			tmp := commandLineArguments.(string)
			details.CommandLineArguments = &tmp
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
			tmp := maximumRuntimeInMinutes.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaximumRuntimeInMinutes = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown job_type '%v' was specified", jobType)
	}
	return baseObject, nil
}

func JobConfigurationDetailsToMap(obj *oci_datascience.JobConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.DefaultJobConfigurationDetails:
		result["job_type"] = "DEFAULT"

		if v.CommandLineArguments != nil {
			result["command_line_arguments"] = string(*v.CommandLineArguments)
		}

		result["environment_variables"] = v.EnvironmentVariables

		if v.MaximumRuntimeInMinutes != nil {
			result["maximum_runtime_in_minutes"] = strconv.FormatInt(*v.MaximumRuntimeInMinutes, 10)
		}
	default:
		log.Printf("[WARN] Received 'job_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceJobResourceCrud) mapToJobInfrastructureConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobInfrastructureConfigurationDetails, error) {
	var baseObject oci_datascience.JobInfrastructureConfigurationDetails
	//discriminator
	jobInfrastructureTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_infrastructure_type"))
	var jobInfrastructureType string
	if ok {
		jobInfrastructureType = jobInfrastructureTypeRaw.(string)
	} else {
		jobInfrastructureType = "" // default value
	}
	switch strings.ToLower(jobInfrastructureType) {
	case strings.ToLower("STANDALONE"):
		details := oci_datascience.StandaloneJobInfrastructureConfigurationDetails{}
		if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
			tmp := blockStorageSizeInGBs.(int)
			details.BlockStorageSizeInGBs = &tmp
		}
		if shapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_name")); ok {
			tmp := shapeName.(string)
			details.ShapeName = &tmp
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown job_infrastructure_type '%v' was specified", jobInfrastructureType)
	}
	return baseObject, nil
}

func JobInfrastructureConfigurationDetailsToMap(obj *oci_datascience.JobInfrastructureConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.StandaloneJobInfrastructureConfigurationDetails:
		result["job_infrastructure_type"] = "STANDALONE"

		if v.BlockStorageSizeInGBs != nil {
			result["block_storage_size_in_gbs"] = int(*v.BlockStorageSizeInGBs)
		}

		if v.ShapeName != nil {
			result["shape_name"] = string(*v.ShapeName)
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	default:
		log.Printf("[WARN] Received 'job_infrastructure_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceJobResourceCrud) mapToJobLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobLogConfigurationDetails, error) {
	result := oci_datascience.JobLogConfigurationDetails{}

	if enableAutoLogCreation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_auto_log_creation")); ok {
		tmp := enableAutoLogCreation.(bool)
		result.EnableAutoLogCreation = &tmp
	}

	if enableLogging, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_logging")); ok {
		tmp := enableLogging.(bool)
		result.EnableLogging = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func JobLogConfigurationDetailsToMap(obj *oci_datascience.JobLogConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnableAutoLogCreation != nil {
		result["enable_auto_log_creation"] = bool(*obj.EnableAutoLogCreation)
	}

	if obj.EnableLogging != nil {
		result["enable_logging"] = bool(*obj.EnableLogging)
	}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatascienceJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.JobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceJobResourceCrud) CreateArtifact() error {
	request := oci_datascience.CreateJobArtifactRequest{}

	if contentDisposition, ok := s.D.GetOkExists("artifact_content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("artifact_content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert Content-Length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if jobArtifact, ok := s.D.GetOkExists("job_artifact"); ok {
		tmp := jobArtifact.(string)
		var artifactReader io.Reader
		artifactReader, err := os.Open(tmp)
		if err != nil {
			return fmt.Errorf("the specified job_artifact is not available: %q", err)
		}
		request.JobArtifact = ioutil.NopCloser(artifactReader)
	}

	request.JobId = s.Res.Id

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateJobArtifact(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceJobResourceCrud) GetArtifactHead() error {
	request := oci_datascience.HeadJobArtifactRequest{}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datascience")

	response, err := s.Client.HeadJobArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.ArtifactHeadRes = &HeadJobArtifact{
		ContentLength:      response.ContentLength,
		ContentDisposition: response.ContentDisposition,
		ContentMd5:         response.ContentMd5,
		LastModified:       response.LastModified,
	}
	return nil
}

func (s *DatascienceJobResourceCrud) SetArtifactData() error {
	if s.ArtifactHeadRes == nil {
		s.D.Set("empty_artifact", true)
		return nil
	}

	if s.ArtifactHeadRes.ContentDisposition != nil {
		s.D.Set("artifact_content_disposition", *s.ArtifactHeadRes.ContentDisposition)
	}

	if s.ArtifactHeadRes.ContentLength != nil {
		s.D.Set("artifact_content_length", *s.ArtifactHeadRes.ContentLength)
	}

	if s.ArtifactHeadRes.ContentMd5 != nil {
		s.D.Set("artifact_content_md5", *s.ArtifactHeadRes.ContentMd5)
	}

	if s.ArtifactHeadRes.LastModified != nil {
		s.D.Set("artifact_last_modified", s.ArtifactHeadRes.LastModified.String())
	}

	s.D.Set("empty_artifact", false)

	return nil
}
