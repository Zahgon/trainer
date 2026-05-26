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

package runtime

import (
	"iter"

	corev1 "k8s.io/api/core/v1"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
)

type Info struct {
	// Labels and Annotations to add to the RuntimeJobTemplate.
	Labels      map[string]string
	Annotations map[string]string
	// Original policy values from the runtime.
	RuntimePolicy RuntimePolicy
	// Scheduler parameters to add to the RuntimeJobTemplate.
	Scheduler *Scheduler
	// TemplateSpec is TrainingRuntime Template object.
	// ObjApply podSpecs and this PodSets should be kept in sync by info.SyncPodSetsToTemplateSpec().
	TemplateSpec TemplateSpec
}

type RuntimePolicy struct {
	MLPolicySource *trainer.MLPolicySource
	PodGroupPolicy *trainer.PodGroupPolicy
	//FluxPolicySource *trainer.FluxMLPolicySource
}

type TemplateSpec struct {
	// ObjApply is ApplyConfiguration for the TrainingRuntimes Template field.
	ObjApply any
	// PodSets is a set of Pod extracted from ObjApply.
	// This is abstract concept to represent multiple PodSpec as a unit.
	PodSets []PodSet
}

type PodSet struct {
	// PodSet name is the name to identify PodSpec.
	// This typically has the name stored in each PodSpec.
	Name string
	// Ancestor is built by `trainer.kubeflow.org/trainjob-ancestor-step` label value
	// in Runtime CRDs.
	Ancestor       *string
	Count          *int32
	InitContainers []Container
	Containers     []Container
	Volumes        []corev1ac.VolumeApplyConfiguration
	Endpoints      iter.Seq[string]
	// The total PodSet requests can be calculated with
	// SinglePodRequests x Count.
	SinglePodRequests corev1.ResourceList
}

type Container struct {
	Name         string
	Image        string
	Command      []string
	Env          []corev1ac.EnvVarApplyConfiguration
	Ports        []corev1ac.ContainerPortApplyConfiguration
	VolumeMounts []corev1ac.VolumeMountApplyConfiguration
}

// TODO (andreyvelich): Potentially, we can add ScheduleTimeoutSeconds to the Scheduler for consistency.
type Scheduler struct {
	PodLabels      map[string]string
	PodAnnotations map[string]string
}

type InfoOptions struct {
	labels        map[string]string
	annotations   map[string]string
	runtimePolicy RuntimePolicy
	templateSpec  TemplateSpec
}

type InfoOption func(options *InfoOptions)

var defaultOptions = InfoOptions{}

func WithLabels(labels map[string]string) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

func WithAnnotations(annotations map[string]string) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

func WithMLPolicySource(mlPolicy *trainer.MLPolicy) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

func WithPodGroupPolicy(pgPolicy *trainer.PodGroupPolicy) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

func WithTemplateSpecObjApply(objApply any) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

// WithPodSet construct Info.TemplateSpec.PodSet from PodSpec.
// The forth argument, 'typedPodSpec' is used only to calculate requested resources.
func WithPodSet(
	psName string, ancestor *string, count int32, typedPodSpec corev1.PodSpec, podSpecApply *corev1ac.PodSpecApplyConfiguration,
) InfoOption {
	_ = "STUB: not implemented"
	return *new(InfoOption)
}

func toPodSetContainer(containerApply ...corev1ac.ContainerApplyConfiguration) iter.Seq[Container] {
	_ = "STUB: not implemented"
	return nil
}

func NewInfo(opts ...InfoOption) *Info { _ = "STUB: not implemented"; return nil }

func TemplateSpecApply[A any](info *Info) (*A, bool) { _ = "STUB: not implemented"; return nil, false }

// FindContainerByPodSetAncestorContainerName finds runtime.Container from Info.TemplateSpec.PodSet by PodSet Ancestor and Container name.
func (i *Info) FindContainerByPodSetAncestorContainerName(psAncestor, containerName string) *Container {
	_ = "STUB: not implemented"
	return nil
}

func (i *Info) FindPodSetByAncestor(ancestor string) *PodSet { _ = "STUB: not implemented"; return nil }

func (i *Info) FindPodSetByName(psName string) *PodSet { _ = "STUB: not implemented"; return nil }

func RuntimeRefToRuntimeRegistryKey(runtimeRef trainer.RuntimeRef) string {
	_ = "STUB: not implemented"
	return ""
}

// ExtractResourcePerNodeFromRuntime extracts the Trainer resource per node from the Info object.
func ExtractResourcePerNodeFromRuntime(info *Info) *corev1.ResourceRequirements {
	_ = "STUB: not implemented"
	return nil
}

// GetNumGPUPerNode returns the GPU count if found in container resources.
func GetNumGPUPerNode(res *corev1.ResourceRequirements) int { _ = "STUB: not implemented"; return 0 }

func numGPU(resourcePerNode corev1.ResourceList) int { _ = "STUB: not implemented"; return 0 }
