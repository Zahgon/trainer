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

package flux

import (
	"context"

	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
	trainer "github.com/kubeflow/trainer/v2/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/v2/pkg/constants"
	"github.com/kubeflow/trainer/v2/pkg/runtime"
	"github.com/kubeflow/trainer/v2/pkg/runtime/framework"

	_ "embed"
)

//go:embed templates/broker.toml
var brokerTemplate string

//go:embed templates/entrypoint.sh
var entrypointTemplate string

//go:embed templates/init.sh
var initTemplate string

// We can customize not easily exposed MiniCluster attributes with envars
var (
	brokerDefaults = map[string]string{

		// the flux view image is the base OS / version for the view to install flux
		// ghcr.io/converged-computing/flux-view-rocky:arm-9
		// ghcr.io/converged-computing/flux-view-rocky:arm-8
		// ghcr.io/converged-computing/flux-view-rocky:tag-9
		// ghcr.io/converged-computing/flux-view-rocky:tag-8
		// ghcr.io/converged-computing/flux-view-ubuntu:tag-noble
		// ghcr.io/converged-computing/flux-view-ubuntu:tag-jammy
		// ghcr.io/converged-computing/flux-view-ubuntu:tag-focal
		// ghcr.io/converged-computing/flux-view-ubuntu:arm-jammy
		// ghcr.io/converged-computing/flux-view-ubuntu:arm-focal
		// We use an ubuntu (more recent) default since it is common
		"FLUX_VIEW_IMAGE":     constants.FluxInstallerImage,
		"FLUX_NETWORK_DEVICE": constants.FluxNewtowkDevice,
		"FLUX_QUEUE_POLICY":   constants.FluxQueuePolicy,
		// Extra flux or broker options can be added as needed.
	}
)

var _ framework.CustomValidationPlugin = (*Flux)(nil)
var _ framework.ComponentBuilderPlugin = (*Flux)(nil)
var _ framework.EnforceMLPolicyPlugin = (*Flux)(nil)
var _ framework.WatchExtensionPlugin = (*Flux)(nil)

const Name = "Flux"

type Flux struct {
	client client.Client
	scheme *apiruntime.Scheme
}

func New(_ context.Context, client client.Client, _ client.FieldIndexer, _ *configapi.Configuration) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (f *Flux) Name() string { _ = "STUB: not implemented"; return "" }

func (f *Flux) Validate(_ context.Context, runtimeInfo *runtime.Info, _, newJobObj *trainer.TrainJob) (admission.Warnings, field.ErrorList) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), *new(field.ErrorList)
}

// We require at least 1 proc per node. Zero or fewer does not make sense.

// Iterate through Trainer's internal PodSet abstraction

// EnforceMLPolicy updates the JobSet
func (f *Flux) EnforceMLPolicy(info *runtime.Info, trainJob *trainer.TrainJob) error {
	_ = "STUB: not implemented"
	return nil
}

// Capture and set the values as annotations
// This effectively "saves" the state onto the Kubernetes resource itself
// We provide the original command to the entrypoint

// Update the command here so we wrap the original command saved earlier
// Also clear existing args so only the Flux entrypoint controls execution

// Define the Init Container. This has a spack view with flux pre-built, and we add to an emptyDir
// with configuration that is then accessible to the application. The OS/version should match.
// For VolumeMounts, you can still use corev1ac because runtime.Container
// methods accept the corev1ac types for nested fields

// Making changes directly to the PodSet allows them to persist

// Update the PodSets (Abstractions for the ReplicatedJobs)

// Add Volumes to the PodSet

// Important! We have to add this to the JobSet to actually take

// Update Containers in the PodSet

// Build creates the extra config map (configuration) and curve secret for Flux.
func (f *Flux) Build(ctx context.Context, info *runtime.Info, trainJob *trainer.TrainJob) ([]apiruntime.ApplyConfiguration, error) {
	_ = "STUB: not implemented"

	// If the user's chosen runtime does not have the flux policy enabled, skip this plugin
	return nil, nil
}

// Note that for Flux, we currently support a design that allows for
// derivation of options from envars that are associated with the job.
// We get these from the designated node container.

// We need a custom entrypoint to prepare the view and configure flux

// Generate/Apply the Curve Secret deterministically based on trainjob id

// Return both. SSA will ensure they are created/merged correctly.

func (f *Flux) ReconcilerBuilders() []runtime.ReconcilerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// brokerSettingsFromTrainJob derives Flux broker config settings from the jobspet node container environment.
func (f *Flux) brokerSettingsFromEnvironment(trainJob *trainer.TrainJob, info *runtime.Info) map[string]string {
	_ = "STUB: not implemented"

	// All settings defaults that we support are already defined here
	return nil
}

// Look through the envars in the runtime spec.
// We only care about the environment defined for the main workers/nodes

// TrainJob (user) gets first preference
// If the variable name matches one of our Flux settings, override it

// getViewVolumes returns the volume apply configurations for the flux view setup
// We need everything here except the curve certificate
func getViewVolumes(configMapName string) []corev1ac.VolumeApplyConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// buildInitScriptConfigMap creates a ConfigMapApplyConfiguration to support server-side Apply
func (f *Flux) buildInitScriptConfigMap(
	trainJob *trainer.TrainJob,
	info *runtime.Info,
	settings map[string]string,
) (*corev1ac.ConfigMapApplyConfiguration, error) {
	_ = "STUB: not implemented"

	// The entrypoint script finishes Flux setup and executes the wrapped application
	return nil, nil
}

// Build the ConfigMap using the Apply Configuration pattern

// entrypoint for application container

// entrypoint for init container (configuration)

// generateBrokerConfig writes the entrypoint file, which prepares the install and configures Flux
func generateBrokerConfig(
	trainJob *trainer.TrainJob,
	hosts string,
	settings map[string]string,
) string {
	_ = "STUB: not implemented"

	// Get the network device for Flux to use (or fall back to default)
	return ""
}

// TODO: we can eventually derive network device from init container
// These shouldn't be formatted in block

// getOriginalCommand derives the original Kubeflow command we need to wrap / handoff to Flux
func getOriginalCommand(trainJob *trainer.TrainJob, info *runtime.Info) string {
	_ = "STUB: not implemented"
	return ""
}

// check PodSets first

// Override if user defined them in the top-level Trainer spec

// Combine into a single string for the shell script

// generateFluxEntrypoint generates the flux entrypoint to prepare the view and run the job
func (f *Flux) generateFluxEntrypoint(trainJob *trainer.TrainJob, info *runtime.Info) string {
	_ = "STUB: not implemented"
	return ""
}

// Derive number of tasks
// This may not technically be the number of processes per node,
// but that is all the TrainJob can currently represent.

// Derive number of GPUs from resources. In Flux, -g is --gpus-per-task

// Resource file for cluster includes GPUs or not
// flux R encode --hosts=${hosts} --cores=0-1 --gpu=0

// generateInitEntrypoint generates the flux entrypoint to prepare flux
func generateInitEntrypoint(
	trainJob *trainer.TrainJob,
	settings map[string]string,
) string {
	_ = "STUB: not implemented"

	// fluxRoot for the view is in /opt/view/lib
	// This must be consistent between the flux-view containers
	// github.com:converged-computing/flux-views.git
	return ""
}

// Generate hostlists. The hostname (prefix) is the trainJob Name
// We need the initial jobset size, and container command	size := *trainJob.Spec.Trainer.NumNodes

// generateHostlist for a specific size given a host prefix and a size
// This is a replicated job so format is different
// lammps-flux-interactive-node-0-0
func generateHostlist(prefix string, size int32) string {
	_ = "STUB: not implemented"

	// Assume a setup without bursting / changing size.
	// We can extend this in the future to allow adding hosts
	return ""
}

// generateRange is a shared function to generate a range string
func generateRange(size int32, start int32) string { _ = "STUB: not implemented"; return "" }

func encodeZ85(data []byte) string { _ = "STUB: not implemented"; return "" }

// Encode into 5 characters (Base 85)

// buildCurveSecret generates a cluster wide curve certificate for flux
func (f *Flux) buildCurveSecret(trainJob *trainer.TrainJob) (*corev1ac.SecretApplyConfiguration, error) {
	_ = "STUB: not implemented"

	// Generate a deterministic Secret Key from the UID
	return nil, nil
}

// Derive the Public Key using standard X25519 (CURVE25519)
// ZeroMQ/Flux uses X25519.

// Encode both to Z85 (40 characters each)

// Follow template from flux keygen curve.cert
