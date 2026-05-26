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

package indexer

import (
	"errors"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	TrainJobRuntimeRefKey        = ".spec.runtimeRef.kind=trainingRuntime"
	TrainJobClusterRuntimeRefKey = ".spec.runtimeRef.kind=clusterTrainingRuntime"
)

var (
	TrainingRuntimeContainerRuntimeClassKey                   = ".trainingRuntimeSpec.jobSetTemplateSpec.replicatedJobs.podTemplateSpec.runtimeClassName"
	ClusterTrainingRuntimeContainerRuntimeClassKey            = ".clusterTrainingRuntimeSpec.jobSetTemplateSpec.replicatedJobs.podTemplateSpec.runtimeClassName"
	ErrorCanNotSetupTrainingRuntimeRuntimeClassIndexer        = errors.New("setting index on runtimeClass for TrainingRuntime")
	ErrorCanNotSetupClusterTrainingRuntimeRuntimeClassIndexer = errors.New("setting index on runtimeClass for ClusterTrainingRuntime")
)

func IndexTrainingRuntimeContainerRuntimeClass(obj client.Object) []string {
	_ = "STUB: not implemented"
	return nil
}

func IndexClusterTrainingRuntimeContainerRuntimeClass(obj client.Object) []string {
	_ = "STUB: not implemented"
	return nil
}

func IndexTrainJobTrainingRuntime(obj client.Object) []string {
	_ = "STUB: not implemented"
	return nil
}

func IndexTrainJobClusterTrainingRuntime(obj client.Object) []string {
	_ = "STUB: not implemented"
	return nil
}
