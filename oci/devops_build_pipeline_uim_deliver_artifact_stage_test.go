// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BuildPipelineStageUIMDeliverArtifactRequiredOnlyResource = BuildPipelineStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageUIMDeliverArtifactRepresentation)

	BuildPipelineStageUIMDeliverArtifactResourceConfig = BuildPipelineStageUIMDeliverArtifactResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageUIMDeliverArtifactRepresentation)

	buildPipelineStageUIMDeliverArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_stage_id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
	}

	buildPipelineStageUIMDeliverArtifactDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": Representation{RepType: Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":      Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                Representation{RepType: Optional, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
		"state":             Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":            RepresentationGroup{Required, buildPipelineStageUIMDeliverArtifactDataSourceFilterRepresentation}}
	buildPipelineStageUIMDeliverArtifactDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`}},
	}

	// TODO: stage timeout is missing in alll the steps
	buildPipelineStageUIMDeliverArtifactRepresentation = map[string]interface{}{
		"build_pipeline_id":                           Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": RepresentationGroup{Required, buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentation},
		"build_pipeline_stage_type":                   Representation{RepType: Required, Create: `DELIVER_ARTIFACT`},
		"defined_tags":                                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deliver_artifact_collection":                 RepresentationGroup{Required, buildPipelineStageDeliverArtifactCollectionRepresentation},
		"description":                                 Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":                                Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	buildPipelineWaitStageRepresentationForDeliverArtifact = map[string]interface{}{
		"build_pipeline_id":                           Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": RepresentationGroup{Required, buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentationWait},
		"build_pipeline_stage_type":                   Representation{RepType: Required, Create: `WAIT`},
		"defined_tags":                                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                                 Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":                                Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"wait_criteria":                               RepresentationGroup{Required, buildPipelineStageWaitCriteriaRepresentation},
	}
	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentation},
	}

	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionRepresentationWait = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentationWait},
	}

	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentationWait = map[string]interface{}{
		"id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}
	buildPipelineStageUIMDeliverArtifactBuildPipelineStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage_wait.id}`},
	}

	buildPipelineStageDeliverArtifactCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineStageDeliverArtifactCollectionItemsRepresentation},
	}

	buildPipelineStageDeliverArtifactCollectionItemsRepresentation = map[string]interface{}{
		"artifact_id":   Representation{RepType: Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
		"artifact_name": Representation{RepType: Required, Create: `artifactName`, Update: `artifactName2`},
	}

	BuildPipelineStageUIMDeliverArtifactResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployGenericArtifactRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage_wait", Required, Create, buildPipelineWaitStageRepresentationForDeliverArtifact)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineStageUIMDeliverArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineStageUIMDeliverArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline_stage.test_build_pipeline_stage"
	datasourceName := "data.oci_devops_build_pipeline_stages.test_build_pipeline_stages"
	singularDatasourceName := "data.oci_devops_build_pipeline_stage.test_build_pipeline_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineStageUIMDeliverArtifactResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineStageUIMDeliverArtifactRepresentation), "devops", "buildPipelineStage", t)

	ResourceTest(t, testAccCheckDevopsBuildPipelineStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deliver_artifact_collection.0.items.0.artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName2"),

				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stages", "test_build_pipeline_stages", Optional, Update, buildPipelineStageUIMDeliverArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageUIMDeliverArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.#", "1"),
			),
		},
		//verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageUIMDeliverArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_stage_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_type", "DELIVER_ARTIFACT"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deliver_artifact_collection.0.items.0.artifact_name", "artifactName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageUIMDeliverArtifactResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
