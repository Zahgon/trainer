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

package apply

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	errorRequestedFieldPathNotFound = errors.New("requested field path not found")
)

func UpsertEnvVars(envVars *[]corev1ac.EnvVarApplyConfiguration, upEnvVars ...corev1ac.EnvVarApplyConfiguration) {
	_ = "STUB: not implemented"
	return
}

func UpsertPort(ports *[]corev1ac.ContainerPortApplyConfiguration, port ...corev1ac.ContainerPortApplyConfiguration) {
	_ = "STUB: not implemented"
	return
}

func UpsertVolumes(volumes *[]corev1ac.VolumeApplyConfiguration, upVolumes ...corev1ac.VolumeApplyConfiguration) {
	_ = "STUB: not implemented"
	return
}

func UpsertVolumeMounts(mounts *[]corev1ac.VolumeMountApplyConfiguration, upMounts ...corev1ac.VolumeMountApplyConfiguration) {
	_ = "STUB: not implemented"
	return
}

func byEnvVarName(a, b corev1ac.EnvVarApplyConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

func byContainerPortOrName(a, b corev1ac.ContainerPortApplyConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

func byVolumeName(a, b corev1ac.VolumeApplyConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

func byVolumeMountPath(a, b corev1ac.VolumeMountApplyConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

type compare[T any] func(T, T) bool

func upsert[T any](items *[]T, item T, predicate compare[T]) { _ = "STUB: not implemented"; return }

func EnvVar(e corev1.EnvVar) *corev1ac.EnvVarApplyConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func EnvVars(e ...corev1.EnvVar) []corev1ac.EnvVarApplyConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func FromTypedObjWithFields[A any](typed client.Object, fields ...string) (*A, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
