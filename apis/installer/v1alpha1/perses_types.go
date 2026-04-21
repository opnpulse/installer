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
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindPerses = "Perses"
	ResourcePerses     = "perses"
	ResourcePerseses   = "perseses"
)

// Perses defines the schema for Perses installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=perseses,singular=perses,categories={kubeops,appscode}
type Perses struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PersesSpec `json:"spec,omitempty"`
}

// PersesSpec is the schema for Perses values file
type PersesSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string      `json:"fullnameOverride"`
	Image            PersesImage `json:"image"`
	//+optional
	AdditionalLabels map[string]string    `json:"additionalLabels"`
	ServiceAccount   ServiceAccountSpec   `json:"serviceAccount"`
	Service          PersesService        `json:"service"`
	ServiceMonitor   PersesServiceMonitor `json:"serviceMonitor"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	Replicas    int32             `json:"replicas"`
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	//+optional
	PodAntiAffinity string `json:"podAntiAffinity"`
	//+optional
	PodAntiAffinityTopologyKey string `json:"podAntiAffinityTopologyKey"`
	//+optional
	LogLevel string `json:"logLevel"`
	// +optional
	LivenessProbe PersesProbe `json:"livenessProbe"`
	// +optional
	ReadinessProbe PersesProbe `json:"readinessProbe"`
	// +optional
	Resources               core.ResourceRequirements     `json:"resources"`
	Config                  PersesConfig                  `json:"config"`
	ProvisioningPersistence PersesProvisioningPersistence `json:"provisioningPersistence"`
	//+optional
	EnvVars []core.EnvVar `json:"envVars"`
	//+optional
	EnvVarsExternalSecretName string            `json:"envVarsExternalSecretName"`
	Sidecar                   PersesSidecar     `json:"sidecar"`
	Persistence               PersesPersistence `json:"persistence"`
	Ingress                   PersesIngress     `json:"ingress"`
	TLS                       PersesTLS         `json:"tls"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type PersesImage struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	PullPolicy string `json:"pullPolicy"`
}

type PersesService struct {
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	Labels     map[string]string `json:"labels"`
	Type       string            `json:"type"`
	PortName   string            `json:"portName"`
	Port       int32             `json:"port"`
	TargetPort int32             `json:"targetPort"`
}

type PersesServiceMonitor struct {
	SelfMonitor bool `json:"selfMonitor"`
	//+optional
	Labels map[string]string `json:"labels"`
	//+optional
	Interval string `json:"interval"`
}

type PersesProbe struct {
	Enabled             bool  `json:"enabled"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	PeriodSeconds       int32 `json:"periodSeconds"`
	TimeoutSeconds      int32 `json:"timeoutSeconds"`
	SuccessThreshold    int32 `json:"successThreshold"`
	FailureThreshold    int32 `json:"failureThreshold"`
}

type PersesConfig struct {
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	ApiPrefix    string                   `json:"api_prefix"`
	Security     PersesConfigSecurity     `json:"security"`
	Frontend     PersesConfigFrontend     `json:"frontend"`
	Database     PersesConfigDatabase     `json:"database"`
	Provisioning PersesConfigProvisioning `json:"provisioning"`
}

type PersesConfigSecurity struct {
	Readonly       bool                       `json:"readonly"`
	EnableAuth     bool                       `json:"enable_auth"`
	Authentication PersesConfigAuthentication `json:"authentication"`
	Cookie         PersesConfigCookie         `json:"cookie"`
}

type PersesConfigAuthentication struct {
	AccessTokenTTL  string `json:"access_token_ttl"`
	RefreshTokenTTL string `json:"refresh_token_ttl"`
}

type PersesConfigCookie struct {
	SameSite string `json:"same_site"`
	Secure   bool   `json:"secure"`
}

type PersesConfigFrontend struct {
	Explorer PersesConfigFrontendExplorer `json:"explorer"`
	//+optional
	Information string `json:"information"`
}

type PersesConfigFrontendExplorer struct {
	Enable bool `json:"enable"`
}

type PersesConfigDatabase struct {
	// +optional
	SQL PersesConfigDatabaseSQL `json:"sql"`
}

type PersesConfigDatabaseSQL struct {
	User                 string `json:"user"`
	Password             string `json:"password"`
	Addr                 string `json:"addr"`
	DbName               string `json:"db_name"`
	AllowNativePasswords bool   `json:"allow_native_passwords"`
}

type PersesConfigProvisioning struct {
	//+optional
	Folders  []string `json:"folders"`
	Interval string   `json:"interval"`
}

type PersesProvisioningPersistence struct {
	Enabled bool `json:"enabled"`
	//+optional
	StorageClass string `json:"storageClass"`
	//+optional
	AccessModes []string `json:"accessModes"`
	Size        string   `json:"size"`
	//+optional
	Labels map[string]string `json:"labels"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type PersesSidecar struct {
	Enabled       bool               `json:"enabled"`
	Image         PersesSidecarImage `json:"image"`
	Label         string             `json:"label"`
	LabelValue    string             `json:"labelValue"`
	AllNamespaces bool               `json:"allNamespaces"`
	//+optional
	ExtraEnvVars       []core.EnvVar `json:"extraEnvVars"`
	EnableSecretAccess bool          `json:"enableSecretAccess"`
}

type PersesSidecarImage struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type PersesPersistence struct {
	Enabled bool `json:"enabled"`
	//+optional
	StorageClass string `json:"storageClass"`
	//+optional
	AccessModes []string `json:"accessModes"`
	Size        string   `json:"size"`
	// +optional
	SecurityContext *core.PodSecurityContext `json:"securityContext"`
	//+optional
	Labels map[string]string `json:"labels"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type PersesIngress struct {
	Enabled bool `json:"enabled"`
	//+optional
	IngressClassName string `json:"ingressClassName"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	Hosts []PersesIngressHost `json:"hosts"`
	//+optional
	TLS []PersesIngressTLS `json:"tls"`
}

type PersesIngressHost struct {
	Host  string              `json:"host"`
	Paths []PersesIngressPath `json:"paths"`
}

type PersesIngressPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type PersesIngressTLS struct {
	SecretName string `json:"secretName"`
	//+optional
	Hosts []string `json:"hosts"`
}

type PersesTLS struct {
	Enabled    bool          `json:"enabled"`
	CaCert     PersesTLSCert `json:"caCert"`
	ClientCert PersesTLSCert `json:"clientCert"`
}

type PersesTLSCert struct {
	Enabled bool `json:"enabled"`
	//+optional
	SecretName string `json:"secretName"`
	MountPath  string `json:"mountPath"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PersesList is a list of Perses
type PersesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Perses CRD objects
	Items []Perses `json:"items,omitempty"`
}
