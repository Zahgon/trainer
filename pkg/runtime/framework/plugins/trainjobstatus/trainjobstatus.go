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

package trainjobstatus

import (
	"context"

	apiruntime "k8s.io/apimachinery/pkg/runtime"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"
)

const (
	Name = "TrainJobStatus"

	// Environment variable names
	envNameStatusURL = "KUBEFLOW_TRAINER_SERVER_URL"
	envNameCACert    = "KUBEFLOW_TRAINER_SERVER_CA_CERT"
	envNameToken     = "KUBEFLOW_TRAINER_SERVER_TOKEN"

	// Volume and mount configuration
	configMountPath = "/var/run/secrets/kubeflow/trainer"
	caCertFileName  = "ca.crt"
	tokenFileName   = "token"
	tokenVolumeName = "kubeflow-trainer-token"

	// Service account token configuration
	tokenExpirySeconds = 3600

	// Server tls config
	caCertKey = "ca.crt"
)

type Status struct {
	client client.Client
	cfg    *configapi.Configuration
}

var _ framework.ComponentBuilderPlugin = (*Status)(nil)
var _ framework.EnforceMLPolicyPlugin = (*Status)(nil)

func New(_ context.Context, c client.Client, _ client.FieldIndexer, cfg *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (p *Status) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Status) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// Inject into all trainer containers

func (p *Status) Build(ctx context.Context, info *runtime.Info, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Status) createEnvVars(trainJob *trainer.TrainJob) ([]corev1ac.EnvVarApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: consider renaming the CertManagement.WebhookServiceName name?

func createTokenVolumeMount() corev1ac.VolumeMountApplyConfiguration {
	_ = "STUB: not implemented"
	return *new(corev1ac.VolumeMountApplyConfiguration)
}

func createTokenVolume(trainJob *trainer.TrainJob) corev1ac.VolumeApplyConfiguration {
	_ = "STUB: not implemented"
	return *new(corev1ac.VolumeApplyConfiguration)
}

// buildStatusServerCaCrtConfigMap creates a ConfigMap that will copy the ca.crt from the webhook secret
func (p *Status) buildStatusServerCaCrtConfigMap(ctx context.Context, trainJob *trainer.TrainJob) (*corev1ac.ConfigMapApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the CA cert from the webhook secret
