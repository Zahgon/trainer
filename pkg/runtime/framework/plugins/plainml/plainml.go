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

package plainml

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

var _ framework.EnforceMLPolicyPlugin = (*PlainML)(nil)

type PlainML struct{}

const Name = "PlainML"

func New(context.Context, client.Client, client.FieldIndexer, *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (p *PlainML) Name() string { _ = "STUB: not implemented"; return "" }

func (p *PlainML) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// TrainJob contains the actual information for the number of nodes.

// Add envs from the TrainJob.
