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
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"k8s.io/client-go/rest"
)

type TokenAuthorizer interface {
	Init(ctx context.Context) error
	Authorize(ctx context.Context, rawIDToken, namespace, trainJobName string) (bool, error)
}

type projectedServiceAccountTokenAuthorizer struct {
	oidcProvider *oidc.Provider
	config       *rest.Config
}

var _ TokenAuthorizer = &projectedServiceAccountTokenAuthorizer{}

// projectedToken is the decoded JWT claims of a k8s projected service account token.
// Note: we only decode the subset of claims we actually need.
type projectedToken struct {
	Issuer     string `json:"iss"`
	Kubernetes struct {
		Namespace string `json:"namespace"`
	} `json:"kubernetes.io"`
}

// NewProjectedServiceAccountTokenAuthorizer creates a validator for checking a bearer token has permission to
// update the requested train job.
func NewProjectedServiceAccountTokenAuthorizer(config *rest.Config) TokenAuthorizer {
	_ = "STUB: not implemented"
	return *new(TokenAuthorizer)
}

func (p *projectedServiceAccountTokenAuthorizer) Init(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Create an authenticated HTTP client using the provided rest config

// Create context with the authenticated HTTP client

func (p *projectedServiceAccountTokenAuthorizer) Authorize(ctx context.Context, authHeader, namespace, trainJobName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Create authorizer with TrainJob-specific audience

// Check token signature, expiry, and audience

// Check token is bound to a pod in the same namespace as the train job

func extractRawToken(authHeader string) string { _ = "STUB: not implemented"; return "" }

// getClusterOIDCIssuerURL tries to look up the cluster token issuer from the in-cluster service account token
// Different clusters may use different issuers. This is a reliable way of discovering the issuer.
func getClusterOIDCIssuerURL() (string, error) { _ = "STUB: not implemented"; return "", nil }
