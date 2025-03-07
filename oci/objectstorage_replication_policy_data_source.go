// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v52/objectstorage"
)

func init() {
	RegisterDatasource("oci_objectstorage_replication_policy", ObjectStorageReplicationPolicyDataSource())
}

func ObjectStorageReplicationPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bucket"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["replication_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ObjectStorageReplicationPolicyResource(), fieldMap, readSingularObjectStorageReplicationPolicy)
}

func readSingularObjectStorageReplicationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient()

	return ReadResource(sync)
}

type ObjectStorageReplicationPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetReplicationPolicyResponse
}

func (s *ObjectStorageReplicationPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageReplicationPolicyDataSourceCrud) Get() error {
	request := oci_object_storage.GetReplicationPolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if replicationId, ok := s.D.GetOkExists("replication_id"); ok {
		tmp := replicationId.(string)
		request.ReplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "object_storage")

	response, err := s.Client.GetReplicationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStorageReplicationPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DestinationBucketName != nil {
		s.D.Set("destination_bucket_name", *s.Res.DestinationBucketName)
	}

	if s.Res.DestinationRegionName != nil {
		s.D.Set("destination_region_name", *s.Res.DestinationRegionName)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusMessage != nil {
		s.D.Set("status_message", *s.Res.StatusMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSync != nil {
		s.D.Set("time_last_sync", s.Res.TimeLastSync.String())
	}

	return nil
}
