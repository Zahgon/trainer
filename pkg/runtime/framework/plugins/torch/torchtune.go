/*
Copyright 2025 The Kubeflow Authors.

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
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
)

func validateTorchTune(runtimeInfo *runtime.Info, newObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// LoRA fine-tuning is not supported for multi-node training in TorchTune.
// getRecipeAndConfig silently falls through to full_finetune_distributed when
// numNodes > 1, discarding LoRA args without any user-visible error.

// Immutable runtime configs must not be set in spec.trainer.args.
// output_dir, tokenizer.path, checkpointer.checkpoint_dir, and tokenizer.merges_file
// are injected by the runtime via extractOverridesFromRuntime and must not be
// overridden by the user.

// getRecipeAndConfig returns the recipe and config file name based on the number of nodes,
// number of processes per node, gpu count, resource per node, model name, and command line arguments.
func getRecipeAndConfig(numNodes int32, numProcPerNode intstr.IntOrString, gpuQ int, trainJob *trainer.TrainJob) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// isLoraConfigEnabled checks if we enables LoraConfig.
func isLoraConfigEnabled(args []string) bool { _ = "STUB: not implemented"; return false }

// isUseQLoraFinetune checks if QLoRA fine-tuning should be used.
func isUseQLoraFinetune(args []string) bool { _ = "STUB: not implemented"; return false }

// If Dora is enabled, no need to continue

// extractOverridesFromRuntime extracts overrides from the TorchTune Trainer Node.
func extractOverridesFromRuntime(info *runtime.Info) []string {
	_ = "STUB: not implemented"
	return nil
}

func getModelFromRuntimeRef(runtimeRefName string) string { _ = "STUB: not implemented"; return "" }
