// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// AutonomousDatabaseSummary An Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`

	// KMS key lifecycle details.
	KmsKeyLifecycleDetails *string `mandatory:"false" json:"kmsKeyLifecycleDetails"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.
	TimeReclamationOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeReclamationOfFreeAutonomousDatabase"`

	// The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.
	TimeDeletionOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeDeletionOfFreeAutonomousDatabase"`

	BackupConfig *AutonomousDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`

	// Key History Entry.
	KeyHistoryEntry []AutonomousDatabaseKeyHistoryEntry `mandatory:"false" json:"keyHistoryEntry"`

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	// - To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.
	// For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The quantity of data in the database, in gigabytes.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The infrastructure type this resource belongs to.
	InfrastructureType AutonomousDatabaseSummaryInfrastructureTypeEnum `mandatory:"false" json:"infrastructureType,omitempty"`

	// True if the database uses dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The date and time the Autonomous Database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The URL of the Service Console for the Autonomous Database.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	ConnectionStrings *AutonomousDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	ConnectionUrls *AutonomousDatabaseConnectionUrls `mandatory:"false" json:"connectionUrls"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel AutonomousDatabaseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The amount of storage that has been used, in terabytes.
	UsedDataStorageSizeInTBs *int `mandatory:"false" json:"usedDataStorageSizeInTBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private endpoint for the resource.
	PrivateEndpoint *string `mandatory:"false" json:"privateEndpoint"`

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// The private endpoint Ip address for the resource.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Indicates if the Autonomous Database version is a preview version.
	IsPreview *bool `mandatory:"false" json:"isPreview"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	IsAccessControlEnabled *bool `mandatory:"false" json:"isAccessControlEnabled"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby.
	// It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
	ArePrimaryWhitelistedIpsUsed *bool `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	StandbyWhitelistedIps []string `mandatory:"false" json:"standbyWhitelistedIps"`

	// Information about Oracle APEX Application Development.
	ApexDetails *AutonomousDatabaseApex `mandatory:"false" json:"apexDetails"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Status of the Data Safe registration for this Autonomous Database.
	DataSafeStatus AutonomousDatabaseSummaryDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`

	// Status of Operations Insights for this Autonomous Database.
	OperationsInsightsStatus AutonomousDatabaseSummaryOperationsInsightsStatusEnum `mandatory:"false" json:"operationsInsightsStatus,omitempty"`

	// The date and time when maintenance will begin.
	TimeMaintenanceBegin *common.SDKTime `mandatory:"false" json:"timeMaintenanceBegin"`

	// The date and time when maintenance will end.
	TimeMaintenanceEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceEnd"`

	// Indicates whether the Autonomous Database is a refreshable clone.
	IsRefreshableClone *bool `mandatory:"false" json:"isRefreshableClone"`

	// The date and time when last refresh happened.
	TimeOfLastRefresh *common.SDKTime `mandatory:"false" json:"timeOfLastRefresh"`

	// The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.
	TimeOfLastRefreshPoint *common.SDKTime `mandatory:"false" json:"timeOfLastRefreshPoint"`

	// The date and time of next refresh.
	TimeOfNextRefresh *common.SDKTime `mandatory:"false" json:"timeOfNextRefresh"`

	// The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
	OpenMode AutonomousDatabaseSummaryOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
	RefreshableStatus AutonomousDatabaseSummaryRefreshableStatusEnum `mandatory:"false" json:"refreshableStatus,omitempty"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode AutonomousDatabaseSummaryRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The Autonomous Database permission level. Restricted mode allows access only to admin users.
	PermissionLevel AutonomousDatabaseSummaryPermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

	// The timestamp of the last switchover operation for the Autonomous Database.
	TimeOfLastSwitchover *common.SDKTime `mandatory:"false" json:"timeOfLastSwitchover"`

	// The timestamp of the last failover operation.
	TimeOfLastFailover *common.SDKTime `mandatory:"false" json:"timeOfLastFailover"`

	// Indicates whether the Autonomous Database has Data Guard enabled.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// Indicates the number of seconds of data loss for a Data Guard failover.
	FailedDataRecoveryInSeconds *int `mandatory:"false" json:"failedDataRecoveryInSeconds"`

	StandbyDb *AutonomousDatabaseStandbySummary `mandatory:"false" json:"standbyDb"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousDatabaseSummaryRoleEnum `mandatory:"false" json:"role,omitempty"`

	// List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
	AvailableUpgradeVersions []string `mandatory:"false" json:"availableUpgradeVersions"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the key store.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// The wallet name for Oracle Key Vault.
	KeyStoreWalletName *string `mandatory:"false" json:"keyStoreWalletName"`

	// The list of regions that support the creation of Autonomous Data Guard standby database.
	SupportedRegionsToCloneTo []string `mandatory:"false" json:"supportedRegionsToCloneTo"`

	// Customer Contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// The date and time that Autonomous Data Guard was enabled for an Autonomous Database where the standby was provisioned in the same region as the primary database.
	TimeLocalDataGuardEnabled *common.SDKTime `mandatory:"false" json:"timeLocalDataGuardEnabled"`

	// The Autonomous Data Guard region type of the Autonomous Database. For Autonomous Databases on shared Exadata infrastructure, Data Guard associations have designated primary and standby regions, and these region types do not change when the database changes roles. The standby regions in Data Guard associations can be the same region designated as the primary region, or they can be remote regions. Certain database administrative operations may be available only in the primary region of the Data Guard association, and cannot be performed when the database using the "primary" role is operating in a remote Data Guard standby region.```
	DataguardRegionType AutonomousDatabaseSummaryDataguardRegionTypeEnum `mandatory:"false" json:"dataguardRegionType,omitempty"`

	// The date and time the Autonomous Data Guard role was switched for the Autonomous Database. For databases that have standbys in both the primary Data Guard region and a remote Data Guard standby region, this is the latest timestamp of either the database using the "primary" role in the primary Data Guard region, or database located in the remote Data Guard standby region.
	TimeDataGuardRoleChanged *common.SDKTime `mandatory:"false" json:"timeDataGuardRoleChanged"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of standby databases located in Autonomous Data Guard remote regions that are associated with the source database. Note that for shared Exadata infrastructure, standby databases located in the same region as the source primary database do not have OCIDs.
	PeerDbIds []string `mandatory:"false" json:"peerDbIds"`

	// Indicates whether the Autonomous Database requires mTLS connections.
	IsMtlsConnectionRequired *bool `mandatory:"false" json:"isMtlsConnectionRequired"`

	// The maintenance schedule type of the Autonomous Database on shared Exadata infrastructure. The EARLY maintenance schedule of this Autonomous Database
	// follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle.
	AutonomousMaintenanceScheduleType AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`
}

func (m AutonomousDatabaseSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLifecycleStateEnum
const (
	AutonomousDatabaseSummaryLifecycleStateProvisioning            AutonomousDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseSummaryLifecycleStateAvailable               AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateStopping                AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseSummaryLifecycleStateStopped                 AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseSummaryLifecycleStateStarting                AutonomousDatabaseSummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseSummaryLifecycleStateTerminating             AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseSummaryLifecycleStateTerminated              AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseSummaryLifecycleStateUnavailable             AutonomousDatabaseSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateRestoreInProgress       AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateRestoreFailed           AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseSummaryLifecycleStateBackupInProgress        AutonomousDatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateScaleInProgress         AutonomousDatabaseSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseSummaryLifecycleStateUpdating                AutonomousDatabaseSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateRestarting              AutonomousDatabaseSummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseSummaryLifecycleStateRecreating              AutonomousDatabaseSummaryLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseSummaryLifecycleStateRoleChangeInProgress    AutonomousDatabaseSummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateUpgrading               AutonomousDatabaseSummaryLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseSummaryLifecycleStateInaccessible            AutonomousDatabaseSummaryLifecycleStateEnum = "INACCESSIBLE"
)

var mappingAutonomousDatabaseSummaryLifecycleState = map[string]AutonomousDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseSummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseSummaryLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseSummaryLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseSummaryLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseSummaryLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseSummaryLifecycleStateInaccessible,
}

// GetAutonomousDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLifecycleStateEnum
func GetAutonomousDatabaseSummaryLifecycleStateEnumValues() []AutonomousDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryInfrastructureTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryInfrastructureTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryInfrastructureTypeEnum
const (
	AutonomousDatabaseSummaryInfrastructureTypeCloud           AutonomousDatabaseSummaryInfrastructureTypeEnum = "CLOUD"
	AutonomousDatabaseSummaryInfrastructureTypeCloudAtCustomer AutonomousDatabaseSummaryInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
)

var mappingAutonomousDatabaseSummaryInfrastructureType = map[string]AutonomousDatabaseSummaryInfrastructureTypeEnum{
	"CLOUD":             AutonomousDatabaseSummaryInfrastructureTypeCloud,
	"CLOUD_AT_CUSTOMER": AutonomousDatabaseSummaryInfrastructureTypeCloudAtCustomer,
}

// GetAutonomousDatabaseSummaryInfrastructureTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryInfrastructureTypeEnum
func GetAutonomousDatabaseSummaryInfrastructureTypeEnumValues() []AutonomousDatabaseSummaryInfrastructureTypeEnum {
	values := make([]AutonomousDatabaseSummaryInfrastructureTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryInfrastructureType {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLicenseModelEnum
const (
	AutonomousDatabaseSummaryLicenseModelLicenseIncluded     AutonomousDatabaseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense AutonomousDatabaseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseSummaryLicenseModel = map[string]AutonomousDatabaseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLicenseModelEnum
func GetAutonomousDatabaseSummaryLicenseModelEnumValues() []AutonomousDatabaseSummaryLicenseModelEnum {
	values := make([]AutonomousDatabaseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDbWorkloadEnum
const (
	AutonomousDatabaseSummaryDbWorkloadOltp AutonomousDatabaseSummaryDbWorkloadEnum = "OLTP"
	AutonomousDatabaseSummaryDbWorkloadDw   AutonomousDatabaseSummaryDbWorkloadEnum = "DW"
	AutonomousDatabaseSummaryDbWorkloadAjd  AutonomousDatabaseSummaryDbWorkloadEnum = "AJD"
	AutonomousDatabaseSummaryDbWorkloadApex AutonomousDatabaseSummaryDbWorkloadEnum = "APEX"
)

var mappingAutonomousDatabaseSummaryDbWorkload = map[string]AutonomousDatabaseSummaryDbWorkloadEnum{
	"OLTP": AutonomousDatabaseSummaryDbWorkloadOltp,
	"DW":   AutonomousDatabaseSummaryDbWorkloadDw,
	"AJD":  AutonomousDatabaseSummaryDbWorkloadAjd,
	"APEX": AutonomousDatabaseSummaryDbWorkloadApex,
}

// GetAutonomousDatabaseSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDbWorkloadEnum
func GetAutonomousDatabaseSummaryDbWorkloadEnumValues() []AutonomousDatabaseSummaryDbWorkloadEnum {
	values := make([]AutonomousDatabaseSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDataSafeStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDataSafeStatusEnum
const (
	AutonomousDatabaseSummaryDataSafeStatusRegistering   AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusRegistered    AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusDeregistering AutonomousDatabaseSummaryDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusNotRegistered AutonomousDatabaseSummaryDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusFailed        AutonomousDatabaseSummaryDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDatabaseSummaryDataSafeStatus = map[string]AutonomousDatabaseSummaryDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseSummaryDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseSummaryDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseSummaryDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseSummaryDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseSummaryDataSafeStatusFailed,
}

// GetAutonomousDatabaseSummaryDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataSafeStatusEnum
func GetAutonomousDatabaseSummaryDataSafeStatusEnumValues() []AutonomousDatabaseSummaryDataSafeStatusEnum {
	values := make([]AutonomousDatabaseSummaryDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataSafeStatus {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryOperationsInsightsStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryOperationsInsightsStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryOperationsInsightsStatusEnum
const (
	AutonomousDatabaseSummaryOperationsInsightsStatusEnabling        AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "ENABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusEnabled         AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "ENABLED"
	AutonomousDatabaseSummaryOperationsInsightsStatusDisabling       AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "DISABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusNotEnabled      AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "NOT_ENABLED"
	AutonomousDatabaseSummaryOperationsInsightsStatusFailedEnabling  AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "FAILED_ENABLING"
	AutonomousDatabaseSummaryOperationsInsightsStatusFailedDisabling AutonomousDatabaseSummaryOperationsInsightsStatusEnum = "FAILED_DISABLING"
)

var mappingAutonomousDatabaseSummaryOperationsInsightsStatus = map[string]AutonomousDatabaseSummaryOperationsInsightsStatusEnum{
	"ENABLING":         AutonomousDatabaseSummaryOperationsInsightsStatusEnabling,
	"ENABLED":          AutonomousDatabaseSummaryOperationsInsightsStatusEnabled,
	"DISABLING":        AutonomousDatabaseSummaryOperationsInsightsStatusDisabling,
	"NOT_ENABLED":      AutonomousDatabaseSummaryOperationsInsightsStatusNotEnabled,
	"FAILED_ENABLING":  AutonomousDatabaseSummaryOperationsInsightsStatusFailedEnabling,
	"FAILED_DISABLING": AutonomousDatabaseSummaryOperationsInsightsStatusFailedDisabling,
}

// GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryOperationsInsightsStatusEnum
func GetAutonomousDatabaseSummaryOperationsInsightsStatusEnumValues() []AutonomousDatabaseSummaryOperationsInsightsStatusEnum {
	values := make([]AutonomousDatabaseSummaryOperationsInsightsStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryOperationsInsightsStatus {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryOpenModeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryOpenModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryOpenModeEnum
const (
	AutonomousDatabaseSummaryOpenModeOnly  AutonomousDatabaseSummaryOpenModeEnum = "READ_ONLY"
	AutonomousDatabaseSummaryOpenModeWrite AutonomousDatabaseSummaryOpenModeEnum = "READ_WRITE"
)

var mappingAutonomousDatabaseSummaryOpenMode = map[string]AutonomousDatabaseSummaryOpenModeEnum{
	"READ_ONLY":  AutonomousDatabaseSummaryOpenModeOnly,
	"READ_WRITE": AutonomousDatabaseSummaryOpenModeWrite,
}

// GetAutonomousDatabaseSummaryOpenModeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryOpenModeEnum
func GetAutonomousDatabaseSummaryOpenModeEnumValues() []AutonomousDatabaseSummaryOpenModeEnum {
	values := make([]AutonomousDatabaseSummaryOpenModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryOpenMode {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryRefreshableStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRefreshableStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRefreshableStatusEnum
const (
	AutonomousDatabaseSummaryRefreshableStatusRefreshing    AutonomousDatabaseSummaryRefreshableStatusEnum = "REFRESHING"
	AutonomousDatabaseSummaryRefreshableStatusNotRefreshing AutonomousDatabaseSummaryRefreshableStatusEnum = "NOT_REFRESHING"
)

var mappingAutonomousDatabaseSummaryRefreshableStatus = map[string]AutonomousDatabaseSummaryRefreshableStatusEnum{
	"REFRESHING":     AutonomousDatabaseSummaryRefreshableStatusRefreshing,
	"NOT_REFRESHING": AutonomousDatabaseSummaryRefreshableStatusNotRefreshing,
}

// GetAutonomousDatabaseSummaryRefreshableStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRefreshableStatusEnum
func GetAutonomousDatabaseSummaryRefreshableStatusEnumValues() []AutonomousDatabaseSummaryRefreshableStatusEnum {
	values := make([]AutonomousDatabaseSummaryRefreshableStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRefreshableStatus {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryRefreshableModeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRefreshableModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRefreshableModeEnum
const (
	AutonomousDatabaseSummaryRefreshableModeAutomatic AutonomousDatabaseSummaryRefreshableModeEnum = "AUTOMATIC"
	AutonomousDatabaseSummaryRefreshableModeManual    AutonomousDatabaseSummaryRefreshableModeEnum = "MANUAL"
)

var mappingAutonomousDatabaseSummaryRefreshableMode = map[string]AutonomousDatabaseSummaryRefreshableModeEnum{
	"AUTOMATIC": AutonomousDatabaseSummaryRefreshableModeAutomatic,
	"MANUAL":    AutonomousDatabaseSummaryRefreshableModeManual,
}

// GetAutonomousDatabaseSummaryRefreshableModeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRefreshableModeEnum
func GetAutonomousDatabaseSummaryRefreshableModeEnumValues() []AutonomousDatabaseSummaryRefreshableModeEnum {
	values := make([]AutonomousDatabaseSummaryRefreshableModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRefreshableMode {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryPermissionLevelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryPermissionLevelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryPermissionLevelEnum
const (
	AutonomousDatabaseSummaryPermissionLevelRestricted   AutonomousDatabaseSummaryPermissionLevelEnum = "RESTRICTED"
	AutonomousDatabaseSummaryPermissionLevelUnrestricted AutonomousDatabaseSummaryPermissionLevelEnum = "UNRESTRICTED"
)

var mappingAutonomousDatabaseSummaryPermissionLevel = map[string]AutonomousDatabaseSummaryPermissionLevelEnum{
	"RESTRICTED":   AutonomousDatabaseSummaryPermissionLevelRestricted,
	"UNRESTRICTED": AutonomousDatabaseSummaryPermissionLevelUnrestricted,
}

// GetAutonomousDatabaseSummaryPermissionLevelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryPermissionLevelEnum
func GetAutonomousDatabaseSummaryPermissionLevelEnumValues() []AutonomousDatabaseSummaryPermissionLevelEnum {
	values := make([]AutonomousDatabaseSummaryPermissionLevelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryPermissionLevel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryRoleEnum Enum with underlying type: string
type AutonomousDatabaseSummaryRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryRoleEnum
const (
	AutonomousDatabaseSummaryRolePrimary         AutonomousDatabaseSummaryRoleEnum = "PRIMARY"
	AutonomousDatabaseSummaryRoleStandby         AutonomousDatabaseSummaryRoleEnum = "STANDBY"
	AutonomousDatabaseSummaryRoleDisabledStandby AutonomousDatabaseSummaryRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousDatabaseSummaryRole = map[string]AutonomousDatabaseSummaryRoleEnum{
	"PRIMARY":          AutonomousDatabaseSummaryRolePrimary,
	"STANDBY":          AutonomousDatabaseSummaryRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseSummaryRoleDisabledStandby,
}

// GetAutonomousDatabaseSummaryRoleEnumValues Enumerates the set of values for AutonomousDatabaseSummaryRoleEnum
func GetAutonomousDatabaseSummaryRoleEnumValues() []AutonomousDatabaseSummaryRoleEnum {
	values := make([]AutonomousDatabaseSummaryRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryRole {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDataguardRegionTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDataguardRegionTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDataguardRegionTypeEnum
const (
	AutonomousDatabaseSummaryDataguardRegionTypePrimaryDgRegion       AutonomousDatabaseSummaryDataguardRegionTypeEnum = "PRIMARY_DG_REGION"
	AutonomousDatabaseSummaryDataguardRegionTypeRemoteStandbyDgRegion AutonomousDatabaseSummaryDataguardRegionTypeEnum = "REMOTE_STANDBY_DG_REGION"
)

var mappingAutonomousDatabaseSummaryDataguardRegionType = map[string]AutonomousDatabaseSummaryDataguardRegionTypeEnum{
	"PRIMARY_DG_REGION":        AutonomousDatabaseSummaryDataguardRegionTypePrimaryDgRegion,
	"REMOTE_STANDBY_DG_REGION": AutonomousDatabaseSummaryDataguardRegionTypeRemoteStandbyDgRegion,
}

// GetAutonomousDatabaseSummaryDataguardRegionTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataguardRegionTypeEnum
func GetAutonomousDatabaseSummaryDataguardRegionTypeEnumValues() []AutonomousDatabaseSummaryDataguardRegionTypeEnum {
	values := make([]AutonomousDatabaseSummaryDataguardRegionTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataguardRegionType {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum Enum with underlying type: string
type AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum
const (
	AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEarly   AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum = "EARLY"
	AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeRegular AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleType = map[string]AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum{
	"EARLY":   AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEarly,
	"REGULAR": AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeRegular,
}

// GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumValues Enumerates the set of values for AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum
func GetAutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnumValues() []AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum {
	values := make([]AutonomousDatabaseSummaryAutonomousMaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryAutonomousMaintenanceScheduleType {
		values = append(values, v)
	}
	return values
}
