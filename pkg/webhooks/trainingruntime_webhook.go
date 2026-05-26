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

package webhooks

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	jobsetv1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/constants"
)

const (
	rJobReplicasErrorMsg       = "always must be 1"
	rJobContainerNamesErrorMsg = "must contain the required container for the ancestor: %s"
)

var (
	expectedContainerNames = map[string]string{
		constants.AncestorTrainer:    constants.Node,
		constants.ModelInitializer:   constants.ModelInitializer,
		constants.DatasetInitializer: constants.DatasetInitializer,
	}
)

// +kubebuilder:webhook:path=/validate-trainer-kubeflow-org-v1alpha1-trainingruntime,mutating=false,failurePolicy=fail,sideEffects=None,groups=trainer.kubeflow.org,resources=trainingruntimes,verbs=create;update,versions=v1alpha1,name=validator.trainingruntime.trainer.kubeflow.org,admissionReviewVersions=v1

// TrainingRuntimeValidator validates TrainingRuntimes
type TrainingRuntimeValidator struct{}

var _ admission.Validator[*trainer.TrainingRuntime] = (*TrainingRuntimeValidator)(nil)

func setupWebhookForTrainingRuntime(mgr ctrl.Manager) error { _ = "STUB: not implemented"; return nil }

func (w *TrainingRuntimeValidator) ValidateCreate(ctx context.Context, obj *trainer.TrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func validateReplicatedJobs(rJobs []jobsetv1alpha2.ReplicatedJob) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// Validate replicated job contains the required containers.
// Mapping of the ancestor labels to the containers:
// 1. dataset-initializer - dataset-initializer
// 2. model-initializer - model-initializer
// 3. trainer - node

func (w *TrainingRuntimeValidator) ValidateUpdate(ctx context.Context, oldObj, newObj *trainer.TrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (w *TrainingRuntimeValidator) ValidateDelete(ctx context.Context, obj *trainer.TrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
