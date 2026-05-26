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

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
)

// +kubebuilder:webhook:path=/validate-trainer-kubeflow-org-v1alpha1-clustertrainingruntime,mutating=false,failurePolicy=fail,sideEffects=None,groups=trainer.kubeflow.org,resources=clustertrainingruntimes,verbs=create;update,versions=v1alpha1,name=validator.clustertrainingruntime.trainer.kubeflow.org,admissionReviewVersions=v1

// ClusterTrainingRuntimeValidator validates ClusterTrainingRuntimes
type ClusterTrainingRuntimeValidator struct{}

var _ admission.Validator[*trainer.ClusterTrainingRuntime] = (*ClusterTrainingRuntimeValidator)(nil)

func setupWebhookForClusterTrainingRuntime(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *ClusterTrainingRuntimeValidator) ValidateCreate(ctx context.Context, obj *trainer.ClusterTrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (w *ClusterTrainingRuntimeValidator) ValidateUpdate(ctx context.Context, oldObj, newObj *trainer.ClusterTrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (w *ClusterTrainingRuntimeValidator) ValidateDelete(ctx context.Context, obj *trainer.ClusterTrainingRuntime) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
