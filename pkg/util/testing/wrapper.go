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

package testing

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	jobsetv1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"
	schedulerpluginsv1alpha1 "sigs.k8s.io/scheduler-plugins/apis/scheduling/v1alpha1"
	volcanov1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
)

type JobSetWrapper struct {
	jobsetv1alpha2.JobSet
}

func MakeJobSetWrapper(namespace, name string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Replicas(replicas int32, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) NumNodes(numNodes int32) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Parallelism(p int32, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Completions(c int32, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) LauncherReplica() *JobSetWrapper { _ = "STUB: not implemented"; return nil }

func (j *JobSetWrapper) InitContainer(rJobName, containerName, image string, envs ...corev1.EnvVar) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Container(rJobName, containerName, image string, command []string, args []string, res corev1.ResourceList) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ReplaceContainer(rJobName, containerName, newContainerName, image string, command, args []string, res corev1.ResourceList) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ContainerTrainerPorts(ports []corev1.ContainerPort) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) NodeSelector(rJobName string, selector map[string]string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

// NodeSelector field is atomic

func (j *JobSetWrapper) Affinity(rJobName string, affinity corev1.Affinity) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) SchedulingGates(rJobName string, schedulingGates ...corev1.PodSchedulingGate) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ImagePullSecrets(rJobName string, imagePullSecrets ...corev1.LocalObjectReference) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Tolerations(rJobName string, tolerations ...corev1.Toleration) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodSecurityContext(rJobName string, securityContext corev1.PodSecurityContext) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Volumes(rJobName string, v ...corev1.Volume) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) VolumeMounts(rJobName, containerName string, vms ...corev1.VolumeMount) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Env(rJobName, containerName string, envs ...corev1.EnvVar) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ContainerSecurityContext(rJobName, containerName string, securityContext corev1.SecurityContext) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) EnvFrom(rJobName, containerName string, envFrom ...corev1.EnvFromSource) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ServiceAccountName(rJobName string, serviceAccountName string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Suspend(suspend bool) *JobSetWrapper { _ = "STUB: not implemented"; return nil }

func (j *JobSetWrapper) ControllerReference(gvk schema.GroupVersionKind, name, uid string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ReplicatedJobLabel(key, value string, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodLabel(key, value string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodAnnotation(key, value string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodLabelForJobs(key, value string, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodAnnotationForJobs(key, value string, rJobNames ...string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) PodPriorityClassName(value string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Label(key, value string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Annotation(key, value string) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Conditions(conditions ...metav1.Condition) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) DependsOn(rJobName string, dependsOn ...jobsetv1alpha2.DependsOn) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) ReplicatedJobsStatuses(statuses []jobsetv1alpha2.ReplicatedJobStatus) *JobSetWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSetWrapper) Obj() *jobsetv1alpha2.JobSet { _ = "STUB: not implemented"; return nil }

type TrainJobWrapper struct {
	trainer.TrainJob
}

func MakeTrainJobWrapper(namespace, name string) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) Suspend(suspend bool) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) UID(uid string) *TrainJobWrapper { _ = "STUB: not implemented"; return nil }

func (t *TrainJobWrapper) ActiveDeadlineSeconds(deadline int64) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) RuntimeRef(gvk schema.GroupVersionKind, name string) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) Initializer(initializer *trainer.Initializer) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) Trainer(trainer *trainer.Trainer) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) RuntimePatches(patches []trainer.RuntimePatch) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) ManagedBy(m string) *TrainJobWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobWrapper) Obj() *trainer.TrainJob { _ = "STUB: not implemented"; return nil }

type TrainJobTrainerWrapper struct {
	trainer.Trainer
}

func MakeTrainJobTrainerWrapper() *TrainJobTrainerWrapper { _ = "STUB: not implemented"; return nil }

func (t *TrainJobTrainerWrapper) NumNodes(numNodes int32) *TrainJobTrainerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobTrainerWrapper) NumProcPerNode(numProcPerNode int32) *TrainJobTrainerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobTrainerWrapper) Container(image string, command []string, args []string, resRequests corev1.ResourceList) *TrainJobTrainerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobTrainerWrapper) Env(env ...corev1.EnvVar) *TrainJobTrainerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobTrainerWrapper) Obj() *trainer.Trainer { _ = "STUB: not implemented"; return nil }

type TrainJobInitializerWrapper struct {
	trainer.Initializer
}

func MakeTrainJobInitializerWrapper() *TrainJobInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobInitializerWrapper) DatasetInitializer(datasetInitializer *trainer.DatasetInitializer) *TrainJobInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobInitializerWrapper) ModelInitializer(modelInitializer *trainer.ModelInitializer) *TrainJobInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobInitializerWrapper) Obj() *trainer.Initializer {
	_ = "STUB: not implemented"
	return nil
}

type TrainJobDatasetInitializerWrapper struct {
	trainer.DatasetInitializer
}

func MakeTrainJobDatasetInitializerWrapper() *TrainJobDatasetInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobDatasetInitializerWrapper) StorageUri(storageUri string) *TrainJobDatasetInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobDatasetInitializerWrapper) Env(env ...corev1.EnvVar) *TrainJobDatasetInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobDatasetInitializerWrapper) SecretRef(secretRef corev1.LocalObjectReference) *TrainJobDatasetInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobDatasetInitializerWrapper) Obj() *trainer.DatasetInitializer {
	_ = "STUB: not implemented"
	return nil
}

type TrainJobModelInitializerWrapper struct {
	trainer.ModelInitializer
}

func MakeTrainJobModelInitializerWrapper() *TrainJobModelInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobModelInitializerWrapper) StorageUri(storageUri string) *TrainJobModelInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobModelInitializerWrapper) Env(env ...corev1.EnvVar) *TrainJobModelInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobModelInitializerWrapper) SecretRef(secretRef corev1.LocalObjectReference) *TrainJobModelInitializerWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrainJobModelInitializerWrapper) Obj() *trainer.ModelInitializer {
	_ = "STUB: not implemented"
	return nil
}

type TrainingRuntimeWrapper struct {
	trainer.TrainingRuntime
}

func MakeTrainingRuntimeWrapper(namespace, name string) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) Label(key, value string) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) Annotation(key, value string) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) Finalizers(f ...string) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) DeletionTimestamp(t metav1.Time) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) RuntimeSpec(spec trainer.TrainingRuntimeSpec) *TrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainingRuntimeWrapper) Obj() *trainer.TrainingRuntime {
	_ = "STUB: not implemented"
	return nil
}

type ClusterTrainingRuntimeWrapper struct {
	trainer.ClusterTrainingRuntime
}

func MakeClusterTrainingRuntimeWrapper(name string) *ClusterTrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClusterTrainingRuntimeWrapper) Finalizers(f ...string) *ClusterTrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClusterTrainingRuntimeWrapper) DeletionTimestamp(t metav1.Time) *ClusterTrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClusterTrainingRuntimeWrapper) RuntimeSpec(spec trainer.TrainingRuntimeSpec) *ClusterTrainingRuntimeWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClusterTrainingRuntimeWrapper) Obj() *trainer.ClusterTrainingRuntime {
	_ = "STUB: not implemented"
	return nil
}

type TrainingRuntimeSpecWrapper struct {
	trainer.TrainingRuntimeSpec
}

func MakeTrainingRuntimeSpecWrapper(spec trainer.TrainingRuntimeSpec) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) JobSetSpec(spec jobsetv1alpha2.JobSetSpec) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) WithMLPolicy(mlPolicy *trainer.MLPolicy) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) LauncherReplica() *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) Replicas(replicas int32, rJobNames ...string) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) InitContainer(rJobName, containerName, image string, envs ...corev1.EnvVar) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) Container(rJobName, containerName, image string, command []string, args []string, res corev1.ResourceList) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) Env(rJobName, containerName string, envs ...corev1.EnvVar) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) PodGroupPolicyCoscheduling(src *trainer.CoschedulingPodGroupPolicySource) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) PodGroupPolicyCoschedulingSchedulingTimeout(timeout int32) *TrainingRuntimeSpecWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrainingRuntimeSpecWrapper) Obj() trainer.TrainingRuntimeSpec {
	_ = "STUB: not implemented"
	return *new(trainer.TrainingRuntimeSpec)
}

type MLPolicyWrapper struct {
	trainer.MLPolicy
}

func MakeMLPolicyWrapper() *MLPolicyWrapper { _ = "STUB: not implemented"; return nil }

func (m *MLPolicyWrapper) WithNumNodes(numNodes int32) *MLPolicyWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (m *MLPolicyWrapper) WithMLPolicySource(source trainer.MLPolicySource) *MLPolicyWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (m *MLPolicyWrapper) Obj() *trainer.MLPolicy { _ = "STUB: not implemented"; return nil }

type MLPolicySourceWrapper struct {
	trainer.MLPolicySource
}

func MakeMLPolicySourceWrapper() *MLPolicySourceWrapper { _ = "STUB: not implemented"; return nil }

func (m *MLPolicySourceWrapper) TorchPolicy() *MLPolicySourceWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (w *MLPolicySourceWrapper) JAXPolicy() *MLPolicySourceWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (w *MLPolicySourceWrapper) XGBoostPolicy() *MLPolicySourceWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (m *MLPolicySourceWrapper) MPIPolicy(numProcPerNode *int32, MPImplementation trainer.MPIImplementation, sshAuthMountPath *string, runLauncherAsNode *bool) *MLPolicySourceWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (m *MLPolicySourceWrapper) Obj() *trainer.MLPolicySource {
	_ = "STUB: not implemented"
	return nil
}

type SchedulerPluginsPodGroupWrapper struct {
	schedulerpluginsv1alpha1.PodGroup
}

func MakeSchedulerPluginsPodGroup(namespace, name string) *SchedulerPluginsPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *SchedulerPluginsPodGroupWrapper) MinMember(members int32) *SchedulerPluginsPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *SchedulerPluginsPodGroupWrapper) MinResources(resources corev1.ResourceList) *SchedulerPluginsPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *SchedulerPluginsPodGroupWrapper) SchedulingTimeout(timeout int32) *SchedulerPluginsPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *SchedulerPluginsPodGroupWrapper) ControllerReference(gvk schema.GroupVersionKind, name, uid string) *SchedulerPluginsPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *SchedulerPluginsPodGroupWrapper) Obj() *schedulerpluginsv1alpha1.PodGroup {
	_ = "STUB: not implemented"
	return nil
}

type VolcanoPodGroupWrapper struct {
	volcanov1beta1.PodGroup
}

func MakeVolcanoPodGroup(namespace, name string) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) MinMember(members int32) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) MinResources(resources *corev1.ResourceList) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) Queue(queue string) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) PriorityClassName(pc string) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) NetworkTopology(mode volcanov1beta1.NetworkTopologyMode, highestTier int) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) ControllerReference(gvk schema.GroupVersionKind, name, uid string) *VolcanoPodGroupWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *VolcanoPodGroupWrapper) Obj() *volcanov1beta1.PodGroup {
	_ = "STUB: not implemented"
	return nil
}

type ConfigMapWrapper struct {
	corev1.ConfigMap
}

func MakeConfigMapWrapper(name, ns string) *ConfigMapWrapper { _ = "STUB: not implemented"; return nil }

func (c *ConfigMapWrapper) WithData(data map[string]string) *ConfigMapWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigMapWrapper) ControllerReference(gvk schema.GroupVersionKind, name, uid string) *ConfigMapWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigMapWrapper) Obj() *corev1.ConfigMap { _ = "STUB: not implemented"; return nil }

type SecretWrapper struct {
	corev1.Secret
}

func MakeSecretWrapper(name, ns string) *SecretWrapper { _ = "STUB: not implemented"; return nil }

func (s *SecretWrapper) WithType(t corev1.SecretType) *SecretWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretWrapper) WithData(data map[string][]byte) *SecretWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretWrapper) WithImmutable(immutable bool) *SecretWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretWrapper) ControllerReference(gvk schema.GroupVersionKind, name, uid string) *SecretWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretWrapper) Obj() *corev1.Secret { _ = "STUB: not implemented"; return nil }
