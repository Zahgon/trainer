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

package core

import (
	"context"
	"errors"

	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	fwkcore "github.com/kubeflow/trainer/v2/pkg/runtime/framework/core"
)

var (
	errorNotFoundSpecifiedTrainingRuntime = errors.New("TrainingRuntime specified in TrainJob is not found")
)

type TrainingRuntime struct {
	framework *fwkcore.Framework
	client    client.Client
}

var TrainingRuntimeGroupKind = schema.GroupKind{
	Group: trainer.GroupVersion.Group,
	Kind:  trainer.TrainingRuntimeKind,
}.String()

var _ runtime.Runtime = (*TrainingRuntime)(nil)

var trainingRuntimeFactory *TrainingRuntime

func NewTrainingRuntime(ctx context.Context, c client.Client, indexer client.FieldIndexer, cfg *configapi.Configuration) (runtime.Runtime, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Runtime), nil
}

func (r *TrainingRuntime) NewObjects(ctx context.Context, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TrainingRuntime) RuntimeInfo(
	trainJob *trainer.TrainJob, runtimeTemplateSpec any, mlPolicy *trainer.MLPolicy, podGroupPolicy *trainer.PodGroupPolicy,
) (*runtime.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TrainingRuntime) newRuntimeInfo(
	trainJob *trainer.TrainJob, jobSetTemplateSpec trainer.JobSetTemplateSpec, mlPolicy *trainer.MLPolicy, podGroupPolicy *trainer.PodGroupPolicy,
) (*runtime.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Support multiple replicas ('.template.spec.replicatedJobs[*].replicas') for replicated Jobs.
// REF: https://github.com/kubeflow/trainer/issues/2318

// Apply resourcesPerNode from TrainJob to the template spec

func (r *TrainingRuntime) mergeRuntimePatches(trainJob *trainer.TrainJob, jobSetTemplateSpec *trainer.JobSetTemplateSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntime) TrainJobStatus(ctx context.Context, trainJob *trainer.TrainJob) (*trainer.TrainJobStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TrainingRuntime) EventHandlerRegistrars() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntime) ValidateObjects(ctx context.Context, old, new *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// ignoring the error here as the runtime configured should be valid
