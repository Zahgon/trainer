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

package statusserver

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
)

func SetupServer(mgr ctrl.Manager, cfg *configapi.StatusServer, enableHTTP2 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a separate client with its own QPS/Burst limits
// to avoid impacting the main reconciler's rate limits

// Initialize OIDC provider for token authentication
// The provider will be used to create verifiers with TrainJob-specific audiences

func createClient(mgr ctrl.Manager, cfg *configapi.StatusServer) (client.Client, error) {
	_ = "STUB: not implemented"
	// Copy the manager's rest config and override rate limits
	return *new(client.Client), nil
}
