/*
Copyright 2026 The Kubeflow Authors.

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

package xgboost

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

type XGBoost struct{}

var _ framework.EnforceMLPolicyPlugin = (*XGBoost)(nil)
var _ framework.CustomValidationPlugin = (*XGBoost)(nil)

const Name = "XGBoost"

func New(context.Context, client.Client, client.FieldIndexer, *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (x *XGBoost) Name() string { _ = "STUB: not implemented"; return "" }

func (x *XGBoost) Validate(_ context.Context, runtimeInfo *runtime.Info, _, newObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

func (x *XGBoost) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	// Guard: Return early if XGBoost policy not configured.
	return nil
}

// Find the trainer PodSet.

// Set the number of nodes from TrainJob if specified.

// Find the trainer container and inject environment variables.

// Auto-derive numWorkersPerNode from GPU resources.
// GPU training: 1 worker per GPU | CPU training: 1 worker per node.

// Step 1: Get resources from Runtime (ClusterTrainingRuntime template).

// Step 2: Override with TrainJob resources if specified.

// Step 3: Derive GPU count from the final resolved resources.

// Build tracker URI: <trainjob-name>-node-0-0.<trainjob-name>

// Inject DMLC_* environment variables.

// DMLC_TRACKER_URI - DNS name for rank-0 worker running tracker.

// DMLC_TRACKER_PORT - Default tracker port.

// DMLC_TASK_ID - Worker rank from Job completion index.

// DMLC_NUM_WORKER - Total number of workers.

// Add container port for tracker communication.
