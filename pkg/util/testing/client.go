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
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func NewClientBuilder(addToSchemes ...func(s *runtime.Scheme) error) *fake.ClientBuilder {
	_ = "STUB: not implemented"
	return nil
}

type builderIndexer struct {
	*fake.ClientBuilder
}

var _ client.FieldIndexer = (*builderIndexer)(nil)

func (b *builderIndexer) IndexField(_ context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func AsIndex(builder *fake.ClientBuilder) client.FieldIndexer {
	_ = "STUB: not implemented"
	return *new(client.FieldIndexer)
}
