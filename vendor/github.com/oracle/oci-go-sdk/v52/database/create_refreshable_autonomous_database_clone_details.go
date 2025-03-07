// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateRefreshableAutonomousDatabaseCloneDetails Details to create an Oracle Autonomous Database refreshable clone.
type CreateRefreshableAutonomousDatabaseCloneDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"true" json:"dbName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	// - To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.
	// For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Notes**
	// - This parameter is only supported for dedicated Exadata infrastructure.
	// - This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
	IsPreviewVersionWithServiceTermsAccepted *bool `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`

	// Indicates if auto scaling is enabled for the Autonomous Database OCPU core count. The default value is `FALSE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// True if the database is on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

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

	// Indicates whether the Autonomous Database has Data Guard enabled.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

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

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Customer Contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// Indicates whether the Autonomous Database requires mTLS connections.
	IsMtlsConnectionRequired *bool `mandatory:"false" json:"isMtlsConnectionRequired"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	DbWorkload CreateAutonomousDatabaseBaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel CreateAutonomousDatabaseBaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The maintenance schedule type of the Autonomous Database on shared Exadata infrastructure. The EARLY maintenance schedule of this Autonomous Database
	// follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle.
	AutonomousMaintenanceScheduleType CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum `mandatory:"false" json:"autonomousMaintenanceScheduleType,omitempty"`
}

//GetCompartmentId returns CompartmentId
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDbName returns DbName
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDbName() *string {
	return m.DbName
}

//GetCpuCoreCount returns CpuCoreCount
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

//GetOcpuCount returns OcpuCount
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetOcpuCount() *float32 {
	return m.OcpuCount
}

//GetDbWorkload returns DbWorkload
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

//GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

//GetDataStorageSizeInGBs returns DataStorageSizeInGBs
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDataStorageSizeInGBs() *int {
	return m.DataStorageSizeInGBs
}

//GetIsFreeTier returns IsFreeTier
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

//GetKmsKeyId returns KmsKeyId
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

//GetVaultId returns VaultId
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetVaultId() *string {
	return m.VaultId
}

//GetAdminPassword returns AdminPassword
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetAdminPassword() *string {
	return m.AdminPassword
}

//GetDisplayName returns DisplayName
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetLicenseModel returns LicenseModel
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

//GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

//GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

//GetIsDedicated returns IsDedicated
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsDedicated() *bool {
	return m.IsDedicated
}

//GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

//GetIsAccessControlEnabled returns IsAccessControlEnabled
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsAccessControlEnabled() *bool {
	return m.IsAccessControlEnabled
}

//GetWhitelistedIps returns WhitelistedIps
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

//GetArePrimaryWhitelistedIpsUsed returns ArePrimaryWhitelistedIpsUsed
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetArePrimaryWhitelistedIpsUsed() *bool {
	return m.ArePrimaryWhitelistedIpsUsed
}

//GetStandbyWhitelistedIps returns StandbyWhitelistedIps
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetStandbyWhitelistedIps() []string {
	return m.StandbyWhitelistedIps
}

//GetIsDataGuardEnabled returns IsDataGuardEnabled
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsDataGuardEnabled() *bool {
	return m.IsDataGuardEnabled
}

//GetSubnetId returns SubnetId
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetNsgIds() []string {
	return m.NsgIds
}

//GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

//GetFreeformTags returns FreeformTags
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDbVersion returns DbVersion
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetDbVersion() *string {
	return m.DbVersion
}

//GetCustomerContacts returns CustomerContacts
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

//GetIsMtlsConnectionRequired returns IsMtlsConnectionRequired
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetIsMtlsConnectionRequired() *bool {
	return m.IsMtlsConnectionRequired
}

//GetAutonomousMaintenanceScheduleType returns AutonomousMaintenanceScheduleType
func (m CreateRefreshableAutonomousDatabaseCloneDetails) GetAutonomousMaintenanceScheduleType() CreateAutonomousDatabaseBaseAutonomousMaintenanceScheduleTypeEnum {
	return m.AutonomousMaintenanceScheduleType
}

func (m CreateRefreshableAutonomousDatabaseCloneDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateRefreshableAutonomousDatabaseCloneDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRefreshableAutonomousDatabaseCloneDetails CreateRefreshableAutonomousDatabaseCloneDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateRefreshableAutonomousDatabaseCloneDetails
	}{
		"CLONE_TO_REFRESHABLE",
		(MarshalTypeCreateRefreshableAutonomousDatabaseCloneDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum Enum with underlying type: string
type CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum string

// Set of constants representing the allowable values for CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum
const (
	CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeAutomatic CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum = "AUTOMATIC"
	CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeManual    CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum = "MANUAL"
)

var mappingCreateRefreshableAutonomousDatabaseCloneDetailsRefreshableMode = map[string]CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum{
	"AUTOMATIC": CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeAutomatic,
	"MANUAL":    CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeManual,
}

// GetCreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnumValues Enumerates the set of values for CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum
func GetCreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnumValues() []CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum {
	values := make([]CreateRefreshableAutonomousDatabaseCloneDetailsRefreshableModeEnum, 0)
	for _, v := range mappingCreateRefreshableAutonomousDatabaseCloneDetailsRefreshableMode {
		values = append(values, v)
	}
	return values
}
