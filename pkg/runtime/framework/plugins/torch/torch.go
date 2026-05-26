/*
Copyright 2024 The Kubeflow Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package torch

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

type Torch struct{}

var _ framework.EnforceMLPolicyPlugin = (*Torch)(nil)
var _ framework.CustomValidationPlugin = (*Torch)(nil)

const Name = "Torch"

func New(context.Context, client.Client, client.FieldIndexer, *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (t *Torch) Name() string { _ = "STUB: not implemented"; return "" }

func (t *Torch) Validate(_ context.Context, runtimeInfo *runtime.Info, _, newObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// Check reserved envs.

// Check supported pretrained models for torchtune.
// TODO(Electronic-Waste): Add more validation for torchtune when we support more arguments.

// TODO (andreyvelich): Add support for PyTorch elastic when JobSet supports Elastic Jobs.
func (t *Torch) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// TrainJob contains the actual information for the Trainer.

// Determine numProcPerNode based on the resourcesPerNode.

// If no GPU is set in resource, calculate numProcPerNode based on CPU.

// Update envs for Info object.

// Add PyTorch distributed "PET_" values for torchrun and torchtune.
// TODO (andreyvelich): We should validate that envs from different plugins don't conflict with each other.
// Ref: https://github.com/kubeflow/trainer/pull/2308#discussion_r1823229940

// Add PET_MASTER_ADDR and PET_MASTER_PORT envs for torchrun.

// Mutate trainer command for torchtune.
// Ref: https://github.com/kubeflow/trainer/tree/master/docs/proposals/2401-llm-trainer-v2#complement-torch-plugin
// 1. Add rendezvous backend arg for torchtune.
// Rendezvous backend is only enabled for multi-nodes or multi-devices training.

// 2. Get the recipe and config from old args and append them to newCommand.

// 3. Extract output directory, tokenizer path and model mount path from (Cluster)TrainingRuntime.

// Add container port for the headless service.

// getNumCPUPerNode calculates the number of CPU processes per node based on the provided resources.
func getNumCPUPerNode(res *corev1.ResourceRequirements) int { _ = "STUB: not implemented"; return 0 }
