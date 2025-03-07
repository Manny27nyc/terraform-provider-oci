// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	oci_work_requests "github.com/oracle/oci-go-sdk/v52/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v52/database"
)

func init() {
	RegisterResource("oci_database_exadata_infrastructure", DatabaseExadataInfrastructureResource())
}

func DatabaseExadataInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExadataInfrastructure,
		Read:     readDatabaseExadataInfrastructure,
		Update:   updateDatabaseExadataInfrastructure,
		Delete:   deleteDatabaseExadataInfrastructure,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_network_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_control_plane_server1": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_control_plane_server2": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dns_server": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"infini_band_network_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ntp_server": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"create_async": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"activation_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_primary": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_contact_mos_validated": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"corporate_proxy": {
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
			"maintenance_window": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"preference": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"days_of_week": {
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

									// Optional

									// Computed
								},
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 20,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"months": {
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

									// Optional

									// Computed
								},
							},
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Computed
					},
				},
			},
			"storage_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"activated_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"additional_storage_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"csi_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_slo_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_data_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_db_node_storage_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_memory_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
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

func createDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient

	return CreateResource(d, sync)
}

func readDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient

	return ReadResource(sync)
}

func updateDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient

	return UpdateResource(d, sync)
}

func deleteDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseExadataInfrastructureResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.ExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseExadataInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExadataInfrastructureResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateCreating),
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) CreatedTarget() []string {
	if createAsyn, ok := s.D.GetOk("create_async"); ok {
		tmp := createAsyn.(bool)
		if tmp {
			return []string{
				string(oci_database.ExadataInfrastructureLifecycleStateCreating),
				string(oci_database.ExadataInfrastructureLifecycleStateActivating),
				string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
				string(oci_database.ExadataInfrastructureLifecycleStateActive),
			}
		}
	}
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleting),
		"TERMINATING",
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleted),
		"TERMINATED",
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
		string(oci_database.ExadataInfrastructureLifecycleStateUpdating),
		string(oci_database.ExadataInfrastructureLifecycleStateMaintenanceInProgress),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
		string(oci_database.ExadataInfrastructureLifecycleStateActivationFailed),
		string(oci_database.ExadataInfrastructureLifecycleStateDisconnected),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) Create() error {
	request := oci_database.CreateExadataInfrastructureRequest{}

	if adminNetworkCIDR, ok := s.D.GetOkExists("admin_network_cidr"); ok {
		tmp := adminNetworkCIDR.(string)
		request.AdminNetworkCIDR = &tmp
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists("compute_count"); ok {
		tmp := computeCount.(int)
		request.ComputeCount = &tmp
	}

	if contacts, ok := s.D.GetOkExists("contacts"); ok {
		interfaces := contacts.([]interface{})
		tmp := make([]oci_database.ExadataInfrastructureContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "contacts", stateDataIndex)
			converted, err := s.mapToExadataInfrastructureContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("contacts") {
			request.Contacts = tmp
		}
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsServer, ok := s.D.GetOkExists("dns_server"); ok {
		request.DnsServer = []string{}
		interfaces := dnsServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_server") {
			request.DnsServer = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if infiniBandNetworkCIDR, ok := s.D.GetOkExists("infini_band_network_cidr"); ok {
		tmp := infiniBandNetworkCIDR.(string)
		request.InfiniBandNetworkCIDR = &tmp
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if netmask, ok := s.D.GetOkExists("netmask"); ok {
		tmp := netmask.(string)
		request.Netmask = &tmp
	}

	if ntpServer, ok := s.D.GetOkExists("ntp_server"); ok {
		request.NtpServer = []string{}
		interfaces := ntpServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp_server") {
			request.NtpServer = tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if storageCount, ok := s.D.GetOkExists("storage_count"); ok {
		tmp := storageCount.(int)
		request.StorageCount = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure

	if waitErr := waitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if activationFile, ok := s.D.GetOkExists("activation_file"); ok {
		response, err := s.activateExadataInfrastructure(activationFile.(string), s.D.Id())
		if err != nil {
			s.D.Set("activation_file", "")
			return err
		}
		s.Res = &response.ExadataInfrastructure
	}

	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Get() error {
	request := oci_database.GetExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure
	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	request := oci_database.UpdateExadataInfrastructureRequest{}

	if additionalStorageCount, ok := s.D.GetOkExists("additional_storage_count"); ok {
		tmp := additionalStorageCount.(int)
		request.AdditionalStorageCount = &tmp
	}

	if adminNetworkCIDR, ok := s.D.GetOkExists("admin_network_cidr"); ok && s.D.HasChange("admin_network_cidr") {
		tmp := adminNetworkCIDR.(string)
		request.AdminNetworkCIDR = &tmp
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok && s.D.HasChange("cloud_control_plane_server1") {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok && s.D.HasChange("cloud_control_plane_server2") {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if contacts, ok := s.D.GetOkExists("contacts"); ok {
		interfaces := contacts.([]interface{})
		tmp := make([]oci_database.ExadataInfrastructureContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "contacts", stateDataIndex)
			converted, err := s.mapToExadataInfrastructureContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("contacts") {
			request.Contacts = tmp
		}
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok && s.D.HasChange("corporate_proxy") {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dnsServer, ok := s.D.GetOkExists("dns_server"); ok && s.D.HasChange("dns_server") {
		request.DnsServer = []string{}
		interfaces := dnsServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_server") {
			request.DnsServer = tmp
		}
	}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok && s.D.HasChange("gateway") {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if infiniBandNetworkCIDR, ok := s.D.GetOkExists("infini_band_network_cidr"); ok && s.D.HasChange("infini_band_network_cidr") {
		tmp := infiniBandNetworkCIDR.(string)
		request.InfiniBandNetworkCIDR = &tmp
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if netmask, ok := s.D.GetOkExists("netmask"); ok && s.D.HasChange("netmask") {
		tmp := netmask.(string)
		request.Netmask = &tmp
	}

	if ntpServer, ok := s.D.GetOkExists("ntp_server"); ok && s.D.HasChange("ntp_server") {
		request.NtpServer = []string{}
		interfaces := ntpServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp_server") {
			request.NtpServer = tmp
		}
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok && s.D.HasChange("time_zone") {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if s.D.Get("state").(string) == string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation) ||
		s.D.Get("additional_storage_count").(int) > 0 {
		if activationFile, ok := s.D.GetOkExists("activation_file"); ok &&
			s.D.Get("activation_file").(string) != "" {
			response, err := s.activateExadataInfrastructure(activationFile.(string), s.D.Id())
			if err != nil {
				s.D.Set("activation_file", "")
				return err
			}
			s.Res = &response.ExadataInfrastructure
			return nil
		}
	}

	response, err := s.Client.UpdateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Delete() error {
	request := oci_database.DeleteExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteExadataInfrastructure(context.Background(), request)
	return err
}

func (s *DatabaseExadataInfrastructureResourceCrud) SetData() error {
	if s.Res.ActivatedStorageCount != nil {
		s.D.Set("activated_storage_count", *s.Res.ActivatedStorageCount)
	}

	if s.Res.AdditionalStorageCount != nil {
		s.D.Set("additional_storage_count", *s.Res.AdditionalStorageCount)
	}

	if s.Res.AdminNetworkCIDR != nil {
		s.D.Set("admin_network_cidr", *s.Res.AdminNetworkCIDR)
	}

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	contacts := []interface{}{}
	for _, item := range s.Res.Contacts {
		contacts = append(contacts, ExadataInfrastructureContactToMap(item))
	}
	s.D.Set("contacts", contacts)

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.CsiNumber != nil {
		s.D.Set("csi_number", *s.Res.CsiNumber)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_server", s.Res.DnsServer)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.InfiniBandNetworkCIDR != nil {
		s.D.Set("infini_band_network_cidr", *s.Res.InfiniBandNetworkCIDR)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_slo_status", s.Res.MaintenanceSLOStatus)

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{ExadataInfrastructureMaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MaxCpuCount != nil {
		s.D.Set("max_cpu_count", *s.Res.MaxCpuCount)
	}

	if s.Res.MaxDataStorageInTBs != nil {
		s.D.Set("max_data_storage_in_tbs", *s.Res.MaxDataStorageInTBs)
	}

	if s.Res.MaxDbNodeStorageInGBs != nil {
		s.D.Set("max_db_node_storage_in_gbs", *s.Res.MaxDbNodeStorageInGBs)
	}

	if s.Res.MaxMemoryInGBs != nil {
		s.D.Set("max_memory_in_gbs", *s.Res.MaxMemoryInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	s.D.Set("ntp_server", s.Res.NtpServer)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageCount != nil {
		s.D.Set("storage_count", *s.Res.StorageCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func ExadataInfrastructureDayOfWeekToMap(obj oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseExadataInfrastructureResourceCrud) mapToExadataInfrastructureContact(fieldKeyFormat string) (oci_database.ExadataInfrastructureContact, error) {
	result := oci_database.ExadataInfrastructureContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	if isContactMosValidated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_contact_mos_validated")); ok {
		tmp := isContactMosValidated.(bool)
		result.IsContactMosValidated = &tmp
	}

	if isPrimary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_primary")); ok {
		tmp := isPrimary.(bool)
		result.IsPrimary = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if phoneNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "phone_number")); ok {
		tmp := phoneNumber.(string)
		result.PhoneNumber = &tmp
	}

	return result, nil
}

func ExadataInfrastructureContactToMap(obj oci_database.ExadataInfrastructureContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.IsContactMosValidated != nil {
		result["is_contact_mos_validated"] = bool(*obj.IsContactMosValidated)
	}

	if obj.IsPrimary != nil {
		result["is_primary"] = bool(*obj.IsPrimary)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PhoneNumber != nil {
		result["phone_number"] = string(*obj.PhoneNumber)
	}

	return result
}

func (s *DatabaseExadataInfrastructureResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))
		if result.Preference == oci_database.MaintenanceWindowPreferenceNoPreference {
			return result, nil
		}
	}

	if daysOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_week")); ok {
		interfaces := daysOfWeek.([]interface{})
		tmp := make([]oci_database.DayOfWeek, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "days_of_week"), stateDataIndex)
			converted, err := s.mapToDayOfWeek(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_week")) {
			result.DaysOfWeek = tmp
		}
	}

	if hoursOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hours_of_day")); ok {
		interfaces := hoursOfDay.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hours_of_day")) {
			result.HoursOfDay = tmp
		}
	}

	if leadTimeInWeeks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lead_time_in_weeks")); ok {
		tmp := leadTimeInWeeks.(int)
		result.LeadTimeInWeeks = &tmp
	}

	if months, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "months")); ok {
		interfaces := months.([]interface{})
		tmp := make([]oci_database.Month, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "months"), stateDataIndex)
			converted, err := s.mapToMonth(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "months")) {
			result.Months = tmp
		}
	}

	if weeksOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")); ok {
		interfaces := weeksOfMonth.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")) {
			result.WeeksOfMonth = tmp
		}
	}

	return result, nil
}

func ExadataInfrastructureMaintenanceWindowToMap(obj *oci_database.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	daysOfWeek := []interface{}{}
	for _, item := range obj.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, ExadataInfrastructureDayOfWeekToMap(item))
	}
	result["days_of_week"] = daysOfWeek

	result["hours_of_day"] = obj.HoursOfDay

	if obj.LeadTimeInWeeks != nil {
		result["lead_time_in_weeks"] = int(*obj.LeadTimeInWeeks)
	}

	months := []interface{}{}
	for _, item := range obj.Months {
		months = append(months, ExadataInfrastructureMonthToMap(item))
	}
	result["months"] = months

	result["preference"] = string(obj.Preference)

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}

func (s *DatabaseExadataInfrastructureResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func ExadataInfrastructureMonthToMap(obj oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseExadataInfrastructureResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExadataInfrastructureCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExadataInfrastructureId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExadataInfrastructureCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) activateExadataInfrastructure(activationFile string, exadataInfrastructureId string) (*oci_database.ActivateExadataInfrastructureResponse, error) {
	request := oci_database.ActivateExadataInfrastructureRequest{}

	activationKeyFile, err := ioutil.ReadFile(activationFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open activation key file: %s", err)
	}

	actionKeyFileBase64Encoded := []byte(base64.StdEncoding.EncodeToString(activationKeyFile))
	request.ActivationFile = actionKeyFileBase64Encoded

	request.ExadataInfrastructureId = &exadataInfrastructureId

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ActivateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return nil, err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadataInfrastructure", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}
