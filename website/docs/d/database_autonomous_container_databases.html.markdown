---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_databases"
sidebar_current: "docs-oci-datasource-database-autonomous_container_databases"
description: |-
  Provides the list of Autonomous Container Databases in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_databases
This data source provides the list of Autonomous Container Databases in Oracle Cloud Infrastructure Database service.

Gets a list of the Autonomous Container Databases in the specified compartment.


## Example Usage

```hcl
data "oci_database_autonomous_container_databases" "test_autonomous_container_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	availability_domain = var.autonomous_container_database_availability_domain
	display_name = var.autonomous_container_database_display_name
	infrastructure_type = var.autonomous_container_database_infrastructure_type
	service_level_agreement_type = var.autonomous_container_database_service_level_agreement_type
	state = var.autonomous_container_database_state
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Optional) The Autonomous Exadata Infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_vm_cluster_id` - (Optional) The Autonomous VM Cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `infrastructure_type` - (Optional) A filter to return only resources that match the given Infrastructure Type.
* `service_level_agreement_type` - (Optional) A filter to return only resources that match the given service level agreement type exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_container_databases` - The list of autonomous_container_databases.

### AutonomousContainerDatabase Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_id` - The OCID of the Autonomous Exadata Infrastructure.
* `autonomous_vm_cluster_id` - The OCID of the Autonomous VM Cluster.
* `availability_domain` - The availability domain of the Autonomous Container Database.
* `backup_config` - Backup options for the Autonomous Container Database. 
	* `backup_destination_details` - Backup destination details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - The OCID of the compartment.
* `db_unique_name` - **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - Oracle Database version of the Autonomous Container Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
* `patch_model` - Database patch model preference.
* `role` - The role of the dataguard enabled Autonomous Container Database.
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `state` - The current state of the Autonomous Container Database.
* `time_created` - The date and time the Autonomous Container Database was created.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).

