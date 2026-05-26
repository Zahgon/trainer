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
)

var (
	errorNotFoundSpecifiedClusterTrainingRuntime = errors.New("not found ClusterTrainingRuntime specified in TrainJob")
)

type ClusterTrainingRuntime struct {
	*TrainingRuntime
}

var _ runtime.Runtime = (*ClusterTrainingRuntime)(nil)

var ClusterTrainingRuntimeGroupKind = schema.GroupKind{
	Group: trainer.GroupVersion.Group,
	Kind:  trainer.ClusterTrainingRuntimeKind,
}.String()

func NewClusterTrainingRuntime(context.Context, client.Client, client.FieldIndexer, *configapi.Configuration) (runtime.Runtime, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Runtime), nil
}

func (r *ClusterTrainingRuntime) NewObjects(ctx context.Context, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClusterTrainingRuntime) RuntimeInfo(
	trainJob *trainer.TrainJob, runtimeTemplateSpec any, mlPolicy *trainer.MLPolicy, podGroupPolicy *trainer.PodGroupPolicy,
) (*runtime.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClusterTrainingRuntime) TrainJobStatus(ctx context.Context, trainJob *trainer.TrainJob) (*trainer.TrainJobStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClusterTrainingRuntime) EventHandlerRegistrars() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClusterTrainingRuntime) ValidateObjects(ctx context.Context, old, new *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}
