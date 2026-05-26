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

package config

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	configapi "github.com/kubeflow/trainer/v2/pkg/apis/config/v1alpha1"
)

// fromFile loads configuration from a file.
func fromFile(path string, scheme *runtime.Scheme, cfg *configapi.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// Decode the configuration file into the Configuration object

// addTo applies the configuration to controller runtime Options.
func addTo(o *ctrl.Options, cfg *configapi.Configuration, enableHTTP2 bool) {
	_ = "STUB: not implemented"
	// Set metrics server options
	return
}

// Disable http/2 for security reasons (CVE-2023-44487, CVE-2023-39325)

// Set webhook server options

// Set health probe bind address

// Set leader election

// Set controller concurrency if specified

// Load loads configuration from file and returns controller Options and Configuration.
// If configFile is empty, default configuration is used.
func Load(scheme *runtime.Scheme, configFile string, enableHTTP2 bool) (ctrl.Options, configapi.Configuration, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Options), *new(configapi.Configuration), nil
}

// Apply defaults

// Load from file

// Validate configuration

// Apply configuration to options

// IsCertManagementEnabled returns true if certificate management is enabled.
// Returns true by default if not explicitly disabled.
func IsCertManagementEnabled(cfg *configapi.Configuration) bool {
	_ = "STUB: not implemented"
	return false
}

// Enabled by default
