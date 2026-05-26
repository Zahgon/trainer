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

	"k8s.io/utils/clock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
)

// +kubebuilder:webhook:path=/mutate-trainer-kubeflow-org-v1alpha1-trainjob,mutating=true,failurePolicy=fail,sideEffects=None,groups=trainer.kubeflow.org,resources=trainjobs,verbs=create;update,versions=v1alpha1,name=defaulter.trainjob.trainer.kubeflow.org,admissionReviewVersions=v1

// TrainJobDefaulter defaults TrainJobs.
type TrainJobDefaulter struct {
	clock clock.PassiveClock
}

var _ admission.Defaulter[*trainer.TrainJob] = (*TrainJobDefaulter)(nil)

func (d *TrainJobDefaulter) Default(ctx context.Context, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:webhook:path=/validate-trainer-kubeflow-org-v1alpha1-trainjob,mutating=false,failurePolicy=fail,sideEffects=None,groups=trainer.kubeflow.org,resources=trainjobs,verbs=create;update,versions=v1alpha1,name=validator.trainjob.trainer.kubeflow.org,admissionReviewVersions=v1

// TrainJobValidator validates TrainJobs
type TrainJobValidator struct {
	runtimes map[string]runtime.Runtime
}

var _ admission.Validator[*trainer.TrainJob] = (*TrainJobValidator)(nil)

func setupWebhookForTrainJob(mgr ctrl.Manager, run map[string]runtime.Runtime) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *TrainJobValidator) ValidateCreate(ctx context.Context, obj *trainer.TrainJob) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (w *TrainJobValidator) ValidateUpdate(ctx context.Context, oldObj, newObj *trainer.TrainJob) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (w *TrainJobValidator) ValidateDelete(ctx context.Context, obj *trainer.TrainJob) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
