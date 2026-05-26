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

package mpi

import (
	"context"

	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

// TODO : Support MPICH and IntelMPI implementations.

type MPI struct {
	client client.Client
	scheme *apiruntime.Scheme
}

var _ framework.CustomValidationPlugin = (*MPI)(nil)
var _ framework.EnforceMLPolicyPlugin = (*MPI)(nil)
var _ framework.WatchExtensionPlugin = (*MPI)(nil)
var _ framework.ComponentBuilderPlugin = (*MPI)(nil)

const Name = "MPI"

// +kubebuilder:rbac:groups="",resources=secrets,verbs=create;get;list;watch;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=create;get;list;watch;update;patch

func New(_ context.Context, client client.Client, _ client.FieldIndexer, _ *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (m *MPI) Name() string {
	_ = "STUB: not implemented"

	// TODO (andreyvelich): We should validate that envs from different plugins don't conflict with each other.
	// Ref: https://github.com/kubeflow/trainer/pull/2308#discussion_r1823229940
	return ""
}

func (m *MPI) Validate(_ context.Context, runtimeInfo *runtime.Info, _, newJobObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// validate PodSet configurations based on NumNodes and RunLauncherAsNode.

// Check reserved MPI envs.

func (m *MPI) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// TrainJob contains the actual information for the Trainer.

// When runLauncherAsNode is enabled, 1 nodes should be allocated to launcher.

// If numProcPerNode is set to 1 in runtime, we make it equal to number of GPUs.

// Add Secret and ConfigMap volumes to the Info object

func (m *MPI) ReconcilerBuilders() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MPI) Build(ctx context.Context, info *runtime.Info, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SSHAuthSecret is immutable.

func (m *MPI) buildSSHAuthSecret(trainJob *trainer.TrainJob) (*corev1ac.SecretApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sshAuthSecretName(trainJobName string) string { _ = "STUB: not implemented"; return "" }

func (m *MPI) buildHostFileConfigMap(info *runtime.Info, trainJob *trainer.TrainJob) *corev1ac.ConfigMapApplyConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func isNode(runLauncherAsNode bool, ps runtime.PodSet) bool {
	_ = "STUB: not implemented"
	return false
}
