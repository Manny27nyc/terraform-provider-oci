---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_configurations"
sidebar_current: "docs-oci-datasource-core-instance_configurations"
description: |-
  Provides the list of Instance Configurations in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_configurations
This data source provides the list of Instance Configurations in Oracle Cloud Infrastructure Core service.

Lists the instance configurations in the specified compartment.


## Example Usage

```hcl
data "oci_core_instance_configurations" "test_instance_configurations" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `instance_configurations` - The list of instance_configurations.

### InstanceConfiguration Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance configuration. 
* `deferred_fields` - Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstanceConfiguration) operation. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the instance configuration. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration.
* `instance_details` - 
	* `block_volumes` - 
		* `attach_details` - Volume attachmentDetails. Please see [AttachVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AttachVolumeDetails/) 
			* `device` - The device name.
			* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
			* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `is_read_only` - Whether the attachment should be created in read-only mode.
			* `is_shareable` - Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
			* `type` - The type of volume. The only supported values are "iscsi" and "paravirtualized".
			* `use_chap` - Whether to use CHAP authentication for the volume attachment. Defaults to false. 
		* `create_details` - Creates a new block volume. Please see [CreateVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVolumeDetails/) 
			* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `compartment_id` - The OCID of the compartment that contains the volume.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `kms_key_id` - The OCID of the Key Management key to assign as the master encryption key for the volume. 
			* `size_in_gbs` - The size of the volume in GBs.
			* `source_details` - 
				* `id` - The OCID of the volume backup.
				* `type` - The type can be one of these values: `volume`, `volumeBackup`
			* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.

				Allowed values:
				* `0`: Represents Lower Cost option.
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option. 
		* `volume_id` - The OCID of the volume.
	* `instance_type` - The type of instance details. Supported instanceType is compute 
	* `launch_details` - Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.

		See [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) for more information. 
		* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
			* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

				To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
			* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

				These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

				The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
				* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
			* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

				These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

				The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
				* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
			* `plugins_config` - The configuration of plugins associated with this instance.
				* `desired_state` - Whether the plugin should be enabled or disabled. Accepted values are `ENABLED` and `DISABLED`

					To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
				* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `availability_config` - Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware. 
			* `recovery_action` - The lifecycle state for an instance when it is recovered after infrastructure maintenance.
				* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
				* `STOP_INSTANCE` - The instance is recovered in the stopped state. 
		* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under.
		* `compartment_id` - The OCID of the compartment.
		* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
            * `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `dedicated_vm_host_id` - The OCID of dedicated VM host.

			Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools. 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
		* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			If you do not specify the fault domain, the system selects one for you.

			 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `FAULT-DOMAIN-1` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `instance_options` - Optional mutable instance options. As a part of Instance Metadata Service Security Header, This allows user to disable the legacy imds endpoints.
			* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
		* `ipxe_script` - This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
		* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
			* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
			* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
			* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
			* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
		* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
			* `boot_volume_type` - Emulation type for the boot volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
			* `firmware` - Firmware used to boot VM. Select the option that matches your operating system.
				* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
				* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images. 
			* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
			* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationLaunchInstanceDetails). 
			* `network_type` - Emulation type for the physical network interface card (NIC).
				* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
				* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
			* `remote_data_volume_type` - Emulation type for volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
		* `metadata` - Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
			* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
			* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			**Providing Cloud-Init Metadata**

			You can use the following metadata key names to provide information to Cloud-Init:

			**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			**Metadata Example**

			"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `platform_config` - The platform configuration requested for the instance.

			If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `is_measured_boot_enabled` - Whether the Measured Boot is to be enabled on the instance 
			* `is_secure_boot_enabled` - Whether the Secure Boot is to be enabled on the instance 
			* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is to be enabled on the instance 
			* `numa_nodes_per_socket` - The number of NUMA nodes per socket. 
			* `type` - The type of platform being configured. (Supported types=[INTEL_VM, AMD_MILAN_BM, AMD_ROME_BM, INTEL_SKYLAKE_BM])
		* `preemptible_instance_config` - Configuration options for preemptible instances. 
			* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction. 
				* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - The type of action to run when the instance is interrupted for eviction.
		* `preferred_maintenance_action` - The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
			* `LIVE_MIGRATE` - Run maintenance using a live migration.
			* `REBOOT` - Run maintenance using a reboot. 
		* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
		* `shape_config` - The shape configuration requested for the instance.

			If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

				The following values are supported:
				* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
				* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
				* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
			* `ocpus` - The total number of OCPUs available to the instance. 
		* `source_details` - 
			* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
			* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
			* `image_id` - The OCID of the image used to boot the instance.
			* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `secondary_vnics` - 
		* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
            * `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `display_name` - A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
		* `nic_index` - Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `time_created` - The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

