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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindThanosOperator = "ThanosOperator"
	ResourceThanosOperator     = "thanosoperator"
	ResourceThanosOperators    = "thanosoperators"
)

// ThanosOperator defines the schema for ThanosOperator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=thanosoperators,singular=thanosoperator,categories={kubeops,appscode}
type ThanosOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThanosOperatorSpec `json:"spec,omitempty"`
}

// ThanosOperatorSpec is the schema for ThanosOperator values file
type ThanosOperatorSpec struct {
	Manager     ThanosOperatorManager `json:"manager"`
	RbacHelpers ThanosRbacHelpers     `json:"rbacHelpers"`
	Crd         ThanosOperatorCrd     `json:"crd"`
	Metrics     ThanosOperatorMetrics `json:"metrics"`
	CertManager ThanosFeatureFlag     `json:"certManager"`
	Prometheus  ThanosFeatureFlag     `json:"prometheus"`
	// +optional
	ObjectStorage ThanosObjectStorage `json:"objectStorage"`
	// +optional
	ThanosRuler ThanosRulerSpec `json:"thanosRuler"`
}

type ThanosRulerSpec struct {
	Enabled bool               `json:"enabled"`
	Storage ThanosRulerStorage `json:"storage"`
	//+optional
	AdditionalArgs []string `json:"additionalArgs"`
	//+optional
	MonitoringClusterBaseURL string `json:"monitoringClusterBaseURL"`
	// +optional
	AdditionalVolumeMounts []core.VolumeMount `json:"additionalVolumeMounts"`
	// +optional
	AdditionalVolumes []core.Volume `json:"additionalVolumes"`
}

type ThanosRulerStorage struct {
	Size string `json:"size"`
}

type ThanosObjectStorage struct {
	//+optional
	Bucket string `json:"bucket"`
	//+optional
	Endpoint string `json:"endpoint"`
	//+optional
	AccessKey string `json:"accessKey"`
	//+optional
	SecretKey string `json:"secretKey"`
	//+optional
	Prefix string `json:"prefix"`
	//+optional
	Region string `json:"region"`
}

type ThanosOperatorManager struct {
	Replicas int                 `json:"replicas"`
	Image    ThanosOperatorImage `json:"image"`
	//+optional
	Args []string `json:"args"`
	//+optional
	Env []core.EnvVar `json:"env"`
	//+optional
	EnvOverrides map[string]string `json:"envOverrides"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
}

type ThanosOperatorImage struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type ThanosRbacHelpers struct {
	Enable bool `json:"enable"`
}

type ThanosOperatorCrd struct {
	Enable bool `json:"enable"`
	Keep   bool `json:"keep"`
}

type ThanosOperatorMetrics struct {
	Enable bool  `json:"enable"`
	Port   int32 `json:"port"`
}

type ThanosFeatureFlag struct {
	Enable bool `json:"enable"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ThanosOperatorList is a list of ThanosOperators
type ThanosOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ThanosOperator CRD objects
	Items []ThanosOperator `json:"items,omitempty"`
}
