/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindAppscodeOtelStack = "AppscodeOtelStack"
	ResourceAppscodeOtelStack     = "appscodeOtelStack"
	ResourceAppscodeOtelStacks    = "appscodeOtelStacks"
)

// AppscodeOtelStack defines the schema for AppscodeOtelStack installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=appscodeOtelStacks,singular=appscodeOtelStack,categories={kubeops,appscode}
type AppscodeOtelStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppscodeOtelStackSpec `json:"spec,omitempty"`
}

// AppscodeOtelStackSpec is the schema for AppscodeOtelStack values file
type AppscodeOtelStackSpec struct {
	OpentelemetryKubeStack OpentelemetryKubeStackValues `json:"opentelemetry-kube-stack"`
}

type OpentelemetryKubeStackValues struct {
	ClusterName string `json:"clusterName"`
	// +optional
	OpentelemetryOperator     *apiextensionsv1.JSON            `json:"opentelemetry-operator"`
	AdmissionWebhooks         OtelAdmissionWebhooks            `json:"admissionWebhooks"`
	Collectors                OpentelemetryKubeStackCollectors `json:"collectors"`
	Instrumentation           OpentelemetryFeatureFlag         `json:"instrumentation"`
	OpAMPBridge               OpentelemetryFeatureFlag         `json:"opAMPBridge"`
	KubernetesServiceMonitors OpentelemetryFeatureFlag         `json:"kubernetesServiceMonitors"`
	KubeApiServer             OpentelemetryFeatureFlag         `json:"kubeApiServer"`
	Kubelet                   OpentelemetryFeatureFlag         `json:"kubelet"`
	KubeControllerManager     OpentelemetryFeatureFlag         `json:"kubeControllerManager"`
	KubeDns                   OpentelemetryFeatureFlag         `json:"kubeDns"`
	KubeEtcd                  OpentelemetryFeatureFlag         `json:"kubeEtcd"`
	KubeScheduler             OpentelemetryFeatureFlag         `json:"kubeScheduler"`
	KubeProxy                 OpentelemetryFeatureFlag         `json:"kubeProxy"`
	KubeStateMetrics          OpentelemetryFeatureFlag         `json:"kubeStateMetrics"`
	NodeExporter              OpentelemetryFeatureFlag         `json:"nodeExporter"`
}

type OtelAdmissionWebhooks struct {
	AutoGenerateCert OtelAutoGenerateCert `json:"autoGenerateCert"`
}

type OtelAutoGenerateCert struct {
	Enabled bool `json:"enabled"`
}

type OpentelemetryFeatureFlag struct {
	Enabled bool `json:"enabled"`
}

type OpentelemetryKubeStackCollectors struct {
	Daemon  OtelCollector `json:"daemon"`
	Cluster OtelCollector `json:"cluster"`
}

type OtelCollector struct {
	//+optional
	Suffix  string `json:"suffix,omitempty"`
	Enabled bool   `json:"enabled"`
	//+optional
	Env []core.EnvVar `json:"env,omitempty"`
	//+optional
	ScrapeConfigsFile string `json:"scrape_configs_file,omitempty"`
	//+optional
	Labels map[string]string `json:"labels,omitempty"`
	//+optional
	Image *OtelCollectorImage `json:"image,omitempty"`
	// +optional
	TargetAllocator *OtelTargetAllocator `json:"targetAllocator,omitempty"`
	// Config holds the arbitrary OpenTelemetry collector pipeline configuration.
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Config *apiextensionsv1.JSON `json:"config,omitempty"`
	// +optional
	Presets *OtelCollectorPresets `json:"presets,omitempty"`
	//+optional
	VolumeMounts []core.VolumeMount `json:"volumeMounts,omitempty"`
	//+optional
	Volumes []core.Volume `json:"volumes,omitempty"`
	// +optional
	Resources core.ResourceRequirements `json:"resources,omitempty"`
	// +optional
	HostAliases []core.HostAlias `json:"hostAliases,omitempty"`
}

type OtelCollectorImage struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type OtelTargetAllocator struct {
	Enabled            bool   `json:"enabled"`
	ServiceAccount     string `json:"serviceAccount"`
	Image              string `json:"image"`
	AllocationStrategy string `json:"allocationStrategy"`
	// +optional
	PrometheusCR *OtelTargetAllocatorPromCR `json:"prometheusCR"`
}

type OtelTargetAllocatorPromCR struct {
	Enabled        bool   `json:"enabled"`
	ScrapeInterval string `json:"scrapeInterval"`
	//+optional
	PodMonitorSelector map[string]string `json:"podMonitorSelector"`
	//+optional
	ServiceMonitorSelector map[string]string `json:"serviceMonitorSelector"`
}

type OtelCollectorPresets struct {
	// +optional
	LogsCollection *OtelPresetFlag `json:"logsCollection,omitempty"`
	// +optional
	HostMetrics *OtelPresetFlag `json:"hostMetrics,omitempty"`
	// +optional
	KubeletMetrics *OtelPresetFlag `json:"kubeletMetrics,omitempty"`
	// +optional
	KubernetesAttributes *OtelPresetFlag `json:"kubernetesAttributes,omitempty"`
	// +optional
	ClusterMetrics *OtelPresetFlag `json:"clusterMetrics,omitempty"`
}

type OtelPresetFlag struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppscodeOtelStackList is a list of AppscodeOtelStacks
type AppscodeOtelStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppscodeOtelStack CRD objects
	Items []AppscodeOtelStack `json:"items,omitempty"`
}
