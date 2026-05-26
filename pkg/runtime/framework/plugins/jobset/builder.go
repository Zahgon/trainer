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
	jobsetv1alpha2ac "sigs.k8s.io/jobset/client-go/applyconfiguration/jobset/v1alpha2"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
)

type Builder struct {
	*jobsetv1alpha2ac.JobSetApplyConfiguration
}

func NewBuilder(jobSet *jobsetv1alpha2ac.JobSetApplyConfiguration) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Initializer updates JobSet values for the initializer Job.
func (b *Builder) Initializer(trainJob *trainer.TrainJob) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Update values for the Dataset Initializer Job.

// TODO: Support multiple replicas ('.template.spec.replicatedJobs[*].replicas') for replicated Jobs.
// REF: https://github.com/kubeflow/trainer/issues/2318

// Update values for the dataset initializer container.

// Update the dataset initializer envs.

// Update the dataset initializer secret reference.

// Update values for the Model Initializer Job.

// TODO: Support multiple replicas ('.template.spec.replicatedJobs[*].replicas') for replicated Jobs.
// REF: https://github.com/kubeflow/trainer/issues/2318

// Update values for the model initializer container.

// Update the model initializer envs.

// Update the model initializer secret reference.

// isRunLauncherAsNode returns true if runLauncherAsNode is set to true in the MPI policy.
func (b *Builder) isRunLauncherAsNode(info *runtime.Info) bool {
	_ = "STUB: not implemented"
	return false
}

// Trainer updates JobSet values for the trainer Job.
func (b *Builder) Trainer(info *runtime.Info, trainJob *trainer.TrainJob) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support multiple replicas ('.template.spec.replicatedJobs[*].replicas') for replicated Jobs.
// REF: https://github.com/kubeflow/trainer/issues/2318

// Update values for the Trainer container.

// Update values from the TrainJob trainer.

// TODO (andreyvelich): For MPI we should apply container resources to the Node ReplicatedJob also.
// Eventually, we should find better way to propagate resources from TrainJob to JobSet.

// TODO: Supporting merge labels would be great.

func (b *Builder) PodLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) PodAnnotations(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) Suspend(suspend *bool) *Builder { _ = "STUB: not implemented"; return nil }

func (b *Builder) Build() *jobsetv1alpha2ac.JobSetApplyConfiguration {
	_ = "STUB: not implemented"
	return nil
}
