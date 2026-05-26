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

package cert

import (
	"crypto/tls"

	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	certDir          = "/tmp/k8s-webhook-server/serving-certs"
	caName           = "kubeflow-trainer-ca"
	caOrganization   = "kubeflow-trainer"
	defaultNamespace = "kubeflow-system"
)

func GetOperatorNamespace() string { _ = "STUB: not implemented"; return "" }

type Config struct {
	WebhookServiceName                 string
	WebhookSecretName                  string
	ValidatingWebhookConfigurationName string
	MutatingWebhookConfigurationName   string
}

//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update
//+kubebuilder:rbac:groups="admissionregistration.k8s.io",resources=validatingwebhookconfigurations,verbs=get;list;watch;update
//+kubebuilder:rbac:groups="admissionregistration.k8s.io",resources=mutatingwebhookconfigurations,verbs=get;list;watch;update

// ManageCerts creates all certs for webhooks.
func ManageCerts(mgr ctrl.Manager, cfg Config, setupFinished chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DNSName is <service name>.<namespace>.svc

// When Kubeflow Trainer is running in the leader election mode,
// we expect webhook server will run in primary and secondary instance

// SetupTLSConfig creates a TLS config with automatic certificate rotation support.
// It creates a cert watcher, adds it to the manager, and returns a TLS config
// that will automatically pick up rotated certificates.
func SetupTLSConfig(mgr ctrl.Manager, enableHTTP2 bool) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disable HTTP/2 unless explicitly enabled (CVE-2023-44487, CVE-2023-39325)
