// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v52/datalabelingservice"
)

func init() {
	RegisterDatasource("oci_data_labeling_service_dataset", DataLabelingServiceDatasetDataSource())
}

func DataLabelingServiceDatasetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dataset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DataLabelingServiceDatasetResource(), fieldMap, readSingularDataLabelingServiceDataset)
}

func readSingularDataLabelingServiceDataset(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataLabelingManagementClient()

	return ReadResource(sync)
}

type DataLabelingServiceDatasetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_labeling_service.DataLabelingManagementClient
	Res    *oci_data_labeling_service.GetDatasetResponse
}

func (s *DataLabelingServiceDatasetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataLabelingServiceDatasetDataSourceCrud) Get() error {
	request := oci_data_labeling_service.GetDatasetRequest{}

	if datasetId, ok := s.D.GetOkExists("dataset_id"); ok {
		tmp := datasetId.(string)
		request.DatasetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "data_labeling_service")

	response, err := s.Client.GetDataset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataLabelingServiceDatasetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AnnotationFormat != nil {
		s.D.Set("annotation_format", *s.Res.AnnotationFormat)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatasetFormatDetails != nil {
		datasetFormatDetailsArray := []interface{}{}
		if datasetFormatDetailsMap := DatasetFormatDetailsToMap(&s.Res.DatasetFormatDetails); datasetFormatDetailsMap != nil {
			datasetFormatDetailsArray = append(datasetFormatDetailsArray, datasetFormatDetailsMap)
		}
		s.D.Set("dataset_format_details", datasetFormatDetailsArray)
	} else {
		s.D.Set("dataset_format_details", nil)
	}

	if s.Res.DatasetSourceDetails != nil {
		datasetSourceDetailsArray := []interface{}{}
		if datasetSourceDetailsMap := DatasetSourceDetailsToMap(&s.Res.DatasetSourceDetails); datasetSourceDetailsMap != nil {
			datasetSourceDetailsArray = append(datasetSourceDetailsArray, datasetSourceDetailsMap)
		}
		s.D.Set("dataset_source_details", datasetSourceDetailsArray)
	} else {
		s.D.Set("dataset_source_details", nil)
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

	if s.Res.InitialRecordGenerationConfiguration != nil {
		s.D.Set("initial_record_generation_configuration", []interface{}{InitialRecordGenerationConfigurationToMap(s.Res.InitialRecordGenerationConfiguration)})
	} else {
		s.D.Set("initial_record_generation_configuration", nil)
	}

	if s.Res.LabelSet != nil {
		s.D.Set("label_set", []interface{}{LabelSetToMap(s.Res.LabelSet)})
	} else {
		s.D.Set("label_set", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
