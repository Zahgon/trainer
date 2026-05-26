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

package controller

import (
	"context"
	"iter"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
)

const (
	updateChBuffer = 10
)

type TrainingRuntimeReconciler struct {
	log                      logr.Logger
	client                   client.Client
	recorder                 events.EventRecorder
	nonRuntimeObjectUpdateCh chan event.TypedGenericEvent[iter.Seq[types.NamespacedName]]
}

var _ reconcile.Reconciler = (*TrainingRuntimeReconciler)(nil)
var _ TrainJobWatcher = (*TrainingRuntimeReconciler)(nil)

func NewTrainingRuntimeReconciler(cli client.Client, recorder events.EventRecorder) *TrainingRuntimeReconciler {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups=trainer.kubeflow.org,resources=trainingruntimes,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=trainer.kubeflow.org,resources=trainingruntimes/finalizers,verbs=get;update;patch

func (r *TrainingRuntimeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *TrainingRuntimeReconciler) NotifyTrainJobUpdate(oldJob, newJob *trainer.TrainJob) {
	_ = "STUB: not implemented"
	return
}

// UPDATE Event.

// CREATE Event.

// DELETE Event.

type nonRuntimeObjectHandler struct{}

var _ handler.TypedEventHandler[iter.Seq[types.NamespacedName], reconcile.Request] = (*nonRuntimeObjectHandler)(nil)

func (h *nonRuntimeObjectHandler) Create(context.Context, event.TypedCreateEvent[iter.Seq[types.NamespacedName]], workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nonRuntimeObjectHandler) Update(context.Context, event.TypedUpdateEvent[iter.Seq[types.NamespacedName]], workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nonRuntimeObjectHandler) Delete(context.Context, event.TypedDeleteEvent[iter.Seq[types.NamespacedName]], workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nonRuntimeObjectHandler) Generic(_ context.Context, e event.TypedGenericEvent[iter.Seq[types.NamespacedName]], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (r *TrainingRuntimeReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	_ = "STUB: not implemented"
	return nil
}
