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

package jax

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

type Jax struct{}

var _ framework.EnforceMLPolicyPlugin = (*Jax)(nil)
var _ framework.CustomValidationPlugin = (*Jax)(nil)

const Name = "JAX"

func New(context.Context, client.Client, client.FieldIndexer, *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (j *Jax) Name() string { _ = "STUB: not implemented"; return "" }

func (j *Jax) Validate(_ context.Context, runtimeInfo *runtime.Info, _, newObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

func (j *Jax) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	// Check if JAX MLPolicy is enabled
	return nil
}

// Find the trainer PodSet

// Set the number of nodes (JAX processes/hosts) from TrainJob

// Get the number of nodes for distributed setup

// Set JAX distributed environment variables

// Total number of JAX processes (one per node/host)

// Process ID - derived from job completion index

// Coordinator address - first pod in the headless service

// Add container port for the headless service (needed for pod-to-pod communication)
