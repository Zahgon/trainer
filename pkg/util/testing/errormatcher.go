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
	"github.com/onsi/gomega/types"
)

func BeNotFoundError() types.GomegaMatcher {
	_ = "STUB: not implemented"
	return *new(types.GomegaMatcher)
}

func BeForbiddenError() types.GomegaMatcher {
	_ = "STUB: not implemented"
	return *new(types.GomegaMatcher)
}

func BeInvalidError() types.GomegaMatcher {
	_ = "STUB: not implemented"
	return *new(types.GomegaMatcher)
}

type errorMatcher int

const (
	NotFoundError errorMatcher = iota
	ForbiddenError
	InvalidError
)

func (em errorMatcher) String() string { _ = "STUB: not implemented"; return "" }

type apiError func(error) bool

func (em errorMatcher) isAPIError(err error) bool { _ = "STUB: not implemented"; return false }

type isErrorMatch struct {
	name errorMatcher
}

func BeAPIError(name errorMatcher) types.GomegaMatcher {
	_ = "STUB: not implemented"
	return *new(types.GomegaMatcher)
}

func (matcher *isErrorMatch) Match(actual interface{}) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (matcher *isErrorMatch) FailureMessage(actual interface{}) (message string) {
	_ = "STUB: not implemented"
	return ""
}

func (matcher *isErrorMatch) NegatedFailureMessage(actual interface{}) (message string) {
	_ = "STUB: not implemented"
	return ""
}
