// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v52/dataflow"
)

func init() {
	RegisterDatasource("oci_dataflow_invoke_run", DataflowInvokeRunDataSource())
}

func DataflowInvokeRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DataflowInvokeRunResource(), fieldMap, readSingularDataflowInvokeRun)
}

func readSingularDataflowInvokeRun(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowInvokeRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient()

	return ReadResource(sync)
}

type DataflowInvokeRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetRunResponse
}

func (s *DataflowInvokeRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowInvokeRunDataSourceCrud) Get() error {
	request := oci_dataflow.GetRunRequest{}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowInvokeRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplicationId != nil {
		s.D.Set("application_id", *s.Res.ApplicationId)
	}

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

	if s.Res.DataReadInBytes != nil {
		s.D.Set("data_read_in_bytes", strconv.FormatInt(*s.Res.DataReadInBytes, 10))
	}

	if s.Res.DataWrittenInBytes != nil {
		s.D.Set("data_written_in_bytes", strconv.FormatInt(*s.Res.DataWrittenInBytes, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogsBucketUri != nil {
		s.D.Set("logs_bucket_uri", *s.Res.LogsBucketUri)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	if s.Res.NumExecutors != nil {
		s.D.Set("num_executors", *s.Res.NumExecutors)
	}

	if s.Res.OpcRequestId != nil {
		s.D.Set("opc_request_id", *s.Res.OpcRequestId)
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

	s.D.Set("private_endpoint_dns_zones", s.Res.PrivateEndpointDnsZones)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.PrivateEndpointMaxHostCount != nil {
		s.D.Set("private_endpoint_max_host_count", *s.Res.PrivateEndpointMaxHostCount)
	}

	s.D.Set("private_endpoint_nsg_ids", s.Res.PrivateEndpointNsgIds)

	if s.Res.PrivateEndpointSubnetId != nil {
		s.D.Set("private_endpoint_subnet_id", *s.Res.PrivateEndpointSubnetId)
	}

	if s.Res.RunDurationInMilliseconds != nil {
		s.D.Set("run_duration_in_milliseconds", strconv.FormatInt(*s.Res.RunDurationInMilliseconds, 10))
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

	if s.Res.TotalOCpu != nil {
		s.D.Set("total_ocpu", *s.Res.TotalOCpu)
	}

	if s.Res.WarehouseBucketUri != nil {
		s.D.Set("warehouse_bucket_uri", *s.Res.WarehouseBucketUri)
	}

	return nil
}
