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

package jobset

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/constants"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

var (
	runtimeRefPath     = field.NewPath("spec").Child("runtimeRef")
	runtimePatchesPath = field.NewPath("spec").Child("runtimePatches")
)

type JobSet struct {
	client     client.Client
	restMapper meta.RESTMapper
	scheme     *apiruntime.Scheme
	logger     logr.Logger
}

var _ framework.WatchExtensionPlugin = (*JobSet)(nil)
var _ framework.PodNetworkPlugin = (*JobSet)(nil)
var _ framework.ComponentBuilderPlugin = (*JobSet)(nil)
var _ framework.TrainJobStatusPlugin = (*JobSet)(nil)
var _ framework.CustomValidationPlugin = (*JobSet)(nil)

const Name = constants.JobSetKind

// +kubebuilder:rbac:groups=jobset.x-k8s.io,resources=jobsets,verbs=create;delete;get;list;watch;update;patch

func New(ctx context.Context, client client.Client, _ client.FieldIndexer, _ *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (j *JobSet) Name() string { _ = "STUB: not implemented"; return "" }

func (j *JobSet) Validate(ctx context.Context, info *runtime.Info, oldObj, newObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// TODO (andreyvelich): Refactor this test to verify the ancestor label in PodTemplate.

// Names of initContainer and containers are unique.

// TODO (andreyvelich): Validate Volumes, VolumeMounts, and Tolerations.

func (j *JobSet) checkRuntimePatchesImmutability(ctx context.Context, oldObj, newObj *trainer.TrainJob) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// Checking immutability makes only sense on updates

// Allow modifying RuntimePatches if the TrainJob is suspended before or
// after the update (i.e. block only when it stays fully unsuspended).
// This lets external controllers (e.g. Kueue) update RuntimePatches and
// toggle spec.suspend in a single API request.

// If the JobSet exists, check whether it's inactive
// so changes won't have side effects on the JobSet's Pods
// that are still running.
// This can happen while the TrainJob is transitioning out
// from unsuspended state.

func (j *JobSet) ReconcilerBuilders() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// TODO (tenzen-y): After we provide the Configuration API, we should return errors based on the enabled plugins.

func (j *JobSet) IdentifyPodNetwork(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support multiple replicas for replicated Jobs.
// REF: https://github.com/kubeflow/trainer/issues/2318

func (j *JobSet) Build(ctx context.Context, info *runtime.Info, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not update the JobSet if it already exists and is not suspended

// Init the JobSet apply configuration from the runtime template spec

// TODO (andreyvelich): Refactor the builder with wrappers for PodSpec.
// TODO: Once we remove deprecated runtime.Info.Trainer, we should remove JobSet Builder with DeprecatedTrainer().

func (j *JobSet) Status(ctx context.Context, trainJob *trainer.TrainJob) (*trainer.TrainJobStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The JobSet may have been automatically deleted in case its TTL duration has been set
// and has expired.
