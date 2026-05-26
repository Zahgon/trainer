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

package framework

import (
	"context"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

type Framework struct {
	testEnv *envtest.Environment
	cancel  context.CancelFunc
}

func (f *Framework) Init() *rest.Config { _ = "STUB: not implemented"; return nil }

func (f *Framework) RunManager(cfg *rest.Config, startControllers bool) (context.Context, client.Client) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(client.Client)
}

// disable metrics to avoid conflicts between packages.

// controller-runtime v0.19+ validates controller names are unique, to make sure
// exported Prometheus metrics for each controller do not conflict. The current check
// relies on static state that's not compatible with testing execution model.
// See the following resources for more context:
// https://github.com/kubernetes-sigs/controller-runtime/pull/2902#issuecomment-2284194683
// https://github.com/kubernetes-sigs/controller-runtime/issues/2994

func (f *Framework) Teardown() { _ = "STUB: not implemented"; return }
