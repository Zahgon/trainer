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

package controller

import (
	"context"
	"iter"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/events"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	jobruntimes "github.com/kubeflow/trainer/v2/pkg/runtime"
)

type TrainJobWatcher interface {
	NotifyTrainJobUpdate(oldJob, newJob *trainer.TrainJob)
}

type TrainJobReconciler struct {
	log      logr.Logger
	client   client.Client
	recorder events.EventRecorder
	runtimes map[string]jobruntimes.Runtime
	watchers iter.Seq[TrainJobWatcher]
}

type TrainJobReconcilerOptions struct {
	Watchers iter.Seq[TrainJobWatcher]
}

type TrainJobReconcilerOption func(*TrainJobReconcilerOptions)

func WithWatchers(watchers ...TrainJobWatcher) TrainJobReconcilerOption {
	_ = "STUB: not implemented"
	return *new(TrainJobReconcilerOption)
}

var _ reconcile.Reconciler = (*TrainJobReconciler)(nil)
var _ predicate.TypedPredicate[*trainer.TrainJob] = (*TrainJobReconciler)(nil)

func NewTrainJobReconciler(client client.Client, recorder events.EventRecorder, runtimes map[string]jobruntimes.Runtime, opts ...TrainJobReconcilerOption) *TrainJobReconciler {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups="",resources=events,verbs=create;watch;update;patch
// +kubebuilder:rbac:groups=events.k8s.io,resources=events,verbs=create;watch;update;patch
// +kubebuilder:rbac:groups=trainer.kubeflow.org,resources=trainjobs,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=trainer.kubeflow.org,resources=trainjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=trainer.kubeflow.org,resources=trainjobs/finalizers,verbs=get;update;patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;update

func (r *TrainJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Keep track of the origin TrainJob status

// Let's clear the failed condition that could have been set previously.
// An external change to the TrainJob spec may transition it out of the Failed state.

// TODO (astefanutti): the error should be surfaced in the TrainJob status to indicate
//  the creation of the runtime resources failed and the TrainJob is backed off until
//  the next retry attempt.
// The event message is truncated to stay within the maximum length limit (1024 chars).

// TODO(astefanutti): Consider using SSA once controller-runtime client has SSA support
// for sub-resources. See: https://github.com/kubernetes-sigs/controller-runtime/issues/3183

func (r *TrainJobReconciler) reconcileObjects(ctx context.Context, runtime jobruntimes.Runtime, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainJobReconciler) reconcileDeadline(ctx context.Context, trainJob *trainer.TrainJob) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *TrainJobReconciler) Create(e event.TypedCreateEvent[*trainer.TrainJob]) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *TrainJobReconciler) Delete(e event.TypedDeleteEvent[*trainer.TrainJob]) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *TrainJobReconciler) Update(e event.TypedUpdateEvent[*trainer.TrainJob]) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *TrainJobReconciler) Generic(e event.TypedGenericEvent[*trainer.TrainJob]) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *TrainJobReconciler) notifyWatchers(oldJob, newJob *trainer.TrainJob) {
	_ = "STUB: not implemented"
	return
}

func setSuspendedCondition(trainJob *trainer.TrainJob) { _ = "STUB: not implemented"; return }

func setFailedCondition(trainJob *trainer.TrainJob, message, reason string) {
	_ = "STUB: not implemented"
	return
}

func removeFailedCondition(trainJob *trainer.TrainJob) { _ = "STUB: not implemented"; return }

func setTrainJobStatus(ctx context.Context, runtime jobruntimes.Runtime, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *TrainJobReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	_ = "STUB: not implemented"
	return nil
}
