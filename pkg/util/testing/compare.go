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

package testing

import (
	"iter"
	"slices"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/types"
)

var (
	PodSetEndpointsCmpOpts = cmp.Transformer("Seq", func(a iter.Seq[string]) []string {
		if a == nil {
			return nil
		}
		return slices.Collect(a)
	})
	TrainJobUpdateReconcileRequestCmpOpts = cmp.Transformer("SeqTrainJobUpdateReconcileRequest",
		func(req iter.Seq[types.NamespacedName]) []types.NamespacedName {
			if req == nil {
				return nil
			}
			return slices.Collect(req)
		},
	)
)

func MPISecretDataComparer(a, b map[string][]byte) bool { _ = "STUB: not implemented"; return false }
