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

package coscheduling

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	nodev1 "k8s.io/api/node/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

type CoScheduling struct {
	client     client.Client
	restMapper meta.RESTMapper
	scheme     *apiruntime.Scheme
	logger     logr.Logger
}

var _ framework.EnforcePodGroupPolicyPlugin = (*CoScheduling)(nil)
var _ framework.WatchExtensionPlugin = (*CoScheduling)(nil)
var _ framework.ComponentBuilderPlugin = (*CoScheduling)(nil)

const Name = "CoScheduling"

// +kubebuilder:rbac:groups=scheduling.x-k8s.io,resources=podgroups,verbs=create;get;list;watch;update;patch
// +kubebuilder:rbac:groups=node.k8s.io,resources=runtimeclasses,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=limitranges,verbs=get;list;watch

func New(_ context.Context, client client.Client, _ client.FieldIndexer, _ *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (c *CoScheduling) Name() string { _ = "STUB: not implemented"; return "" }

func (c *CoScheduling) EnforcePodGroupPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoScheduling) Build(ctx context.Context, info *runtime.Info, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not update the PodGroup if it already exists and the TrainJob is not suspended

type PodGroupRuntimeClassHandler struct {
	client client.Client
}

var _ handler.TypedEventHandler[*nodev1.RuntimeClass, reconcile.Request] = (*PodGroupRuntimeClassHandler)(nil)

func (h *PodGroupRuntimeClassHandler) Create(ctx context.Context, e event.TypedCreateEvent[*nodev1.RuntimeClass], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupRuntimeClassHandler) Update(ctx context.Context, e event.TypedUpdateEvent[*nodev1.RuntimeClass], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupRuntimeClassHandler) Delete(ctx context.Context, e event.TypedDeleteEvent[*nodev1.RuntimeClass], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupRuntimeClassHandler) Generic(context.Context, event.TypedGenericEvent[*nodev1.RuntimeClass], workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupRuntimeClassHandler) queueSuspendedTrainJobs(ctx context.Context, runtimeClass *nodev1.RuntimeClass, q workqueue.TypedRateLimitingInterface[reconcile.Request]) error {
	_ = "STUB: not implemented"
	return nil
}

type PodGroupLimitRangeHandler struct {
	client client.Client
}

var _ handler.TypedEventHandler[*corev1.LimitRange, reconcile.Request] = (*PodGroupLimitRangeHandler)(nil)

func (h *PodGroupLimitRangeHandler) Create(ctx context.Context, e event.TypedCreateEvent[*corev1.LimitRange], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupLimitRangeHandler) Update(ctx context.Context, e event.TypedUpdateEvent[*corev1.LimitRange], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupLimitRangeHandler) Delete(ctx context.Context, e event.TypedDeleteEvent[*corev1.LimitRange], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupLimitRangeHandler) Generic(context.Context, event.TypedGenericEvent[*corev1.LimitRange], workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *PodGroupLimitRangeHandler) queueSuspendedTrainJob(ctx context.Context, ns string, q workqueue.TypedRateLimitingInterface[reconcile.Request]) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoScheduling) ReconcilerBuilders() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}
