---
subcategory: "Data Labeling Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_labeling_service_dataset"
sidebar_current: "docs-oci-datasource-data_labeling_service-dataset"
description: |-
  Provides details about a specific Dataset in Oracle Cloud Infrastructure Data Labeling Service service
---

# Data Source: oci_data_labeling_service_dataset
This data source provides details about a specific Dataset resource in Oracle Cloud Infrastructure Data Labeling Service service.

Gets a Dataset by identifier

## Example Usage

```hcl
data "oci_data_labeling_service_dataset" "test_dataset" {
	#Required
	dataset_id = oci_data_labeling_service_dataset.test_dataset.id
}
```

## Argument Reference

The following arguments are supported:

* `dataset_id` - (Required) Unique Dataset OCID


## Attributes Reference

The following attributes are exported:

* `annotation_format` - The annotation format name required for labeling records.
* `compartment_id` - The OCID of the compartment of the resource.
* `dataset_format_details` - Specifies how to process the data. Supported formats include DOCUMENT, IMAGE and TEXT.
	* `format_type` - Format type. DOCUMENT format is for record contents that are PDFs or TIFFs. IMAGE format is for record contents that are JPEGs or PNGs. TEXT format is for record contents that are txt files.
* `dataset_source_details` - This allows the customer to specify the source of the dataset.
	* `bucket` - The object storage bucket that contains the dataset data source
	* `namespace` - Namespace of the bucket that contains the dataset data source
	* `prefix` - A common path prefix shared by the objects that make up the dataset.
	* `source_type` - Source type.  OBJECT_STORAGE allows the customer to describe where the dataset is in object storage.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user provided description of the dataset
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Dataset.
* `initial_record_generation_configuration` - Initial Generate Records configuration, generates records from the Dataset's source.
* `label_set` - An ordered collection of Labels that are unique by name. 
	* `items` - An ordered collection of Labels that are unique by name.
		* `name` - An unique name for a label within its dataset.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in FAILED or NEEDS_ATTENTION state.
* `state` - The state of a dataset. CREATING - The dataset is being created.  It will transition to ACTIVE when it is ready for labeling. ACTIVE   - The dataset is ready for labeling. UPDATING - The dataset is being updated.  It and its related resources may be unavailable for other updates until it returns to ACTIVE. NEEDS_ATTENTION - A dataset updation operation has failed due to validation or other errors and needs attention. DELETING - The dataset and its related resources are being deleted. DELETED  - The dataset has been deleted and is no longer available. FAILED   - The dataset has failed due to validation or other errors. 
* `time_created` - The date and time the resource was created, in the timestamp format defined by RFC3339.
* `time_updated` - The date and time the resource was last updated, in the timestamp format defined by RFC3339.

