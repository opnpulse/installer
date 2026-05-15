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
	catgwapi "go.bytebuilders.dev/catalog/api/gateway/v1alpha1"
	core "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindPromLabelProxy = "PromLabelProxy"
	ResourcePromLabelProxy     = "promlabelproxy"
	ResourcePromLabelProxies   = "promlabelproxies"
)

// PromLabelProxy defines the schema for PromLabelProxy installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=promlabelproxies,singular=promlabelproxy,categories={kubeops,appscode}
type PromLabelProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PromLabelProxySpec `json:"spec,omitempty"`
}

// PromLabelProxySpec is the schema for PromLabelProxy values file
type PromLabelProxySpec struct {
	ReplicaCount int                 `json:"replicaCount"`
	Image        PromLabelProxyImage `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	NamespaceOverride string `json:"namespaceOverride"`
	//+optional
	FullnameOverride string             `json:"fullnameOverride"`
	ServiceAccount   ServiceAccountSpec `json:"serviceAccount"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	Service         PromLabelProxyService `json:"service"`
	// +optional
	Resources   core.ResourceRequirements `json:"resources"`
	Autoscaling PromLabelProxyAutoscaling `json:"autoscaling"`
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// +optional
	LivenessProbe PromLabelProxyProbe `json:"livenessProbe"`
	// +optional
	ReadinessProbe PromLabelProxyProbe           `json:"readinessProbe"`
	Ingress        PromLabelProxyIngress         `json:"ingress"`
	Config         PromLabelProxyConfig          `json:"config"`
	Metrics        PromLabelProxyMetrics         `json:"metrics"`
	KubeRBACProxy  PromLabelProxyKubeRBACProxy   `json:"kubeRBACProxy"`
	Clickhouse     PromLabelProxyClickhouse      `json:"clickhouse"`
	Infra          catgwapi.ServiceProviderInfra `json:"infra"`
	// +optional
	Platform PromLabelProxyPlatform `json:"platform"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
	// +optional
	ExtraManifests *apiextensionsv1.JSON `json:"extraManifests"`
}

type PromLabelProxyProbe struct {
	// +optional
	HttpGet PromLabelProxyProbeHttpGet `json:"httpGet"`
}

type PromLabelProxyProbeHttpGet struct {
	Path   string `json:"path"`
	Port   string `json:"port"`
	Scheme string `json:"scheme"`
}

type PromLabelProxyPlatform struct {
	//+optional
	BaseURL string `json:"baseURL"`
}

type PromLabelProxyImage struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type PromLabelProxyService struct {
	Port int32  `json:"port"`
	Type string `json:"type"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type PromLabelProxyAutoscaling struct {
	Enabled                        bool  `json:"enabled"`
	MinReplicas                    int32 `json:"minReplicas"`
	MaxReplicas                    int32 `json:"maxReplicas"`
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage"`
}

type PromLabelProxyIngress struct {
	Enabled   bool   `json:"enabled"`
	ClassName string `json:"className"`
	//+optional
	Labels map[string]string `json:"labels"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	Hosts []PromLabelProxyIngressHost `json:"hosts"`
	//+optional
	TLS []PromLabelProxyIngressTLS `json:"tls"`
}

type PromLabelProxyIngressHost struct {
	Host  string                      `json:"host"`
	Paths []PromLabelProxyIngressPath `json:"paths"`
}

type PromLabelProxyIngressPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type PromLabelProxyIngressTLS struct {
	SecretName string `json:"secretName"`
	//+optional
	Hosts []string `json:"hosts"`
}

type PromLabelProxyConfig struct {
	ListenAddress string `json:"listenAddress"`
	Upstream      string `json:"upstream"`
	Label         string `json:"label"`
	//+optional
	ExtraArgs []string `json:"extraArgs"`
}

type PromLabelProxyMetrics struct {
	Enabled        bool                         `json:"enabled"`
	ListenAddress  string                       `json:"listenAddress"`
	ServiceMonitor PromLabelProxyServiceMonitor `json:"serviceMonitor"`
}

type PromLabelProxyServiceMonitor struct {
	Enabled bool  `json:"enabled"`
	Port    int32 `json:"port"`
	//+optional
	AdditionalLabels map[string]string `json:"additionalLabels"`
	//+optional
	JobLabel string `json:"jobLabel"`
	//+optional
	TargetLabels []string `json:"targetLabels"`
	//+optional
	PodTargetLabels []string `json:"podTargetLabels"`
	//+optional
	SampleLimit int32 `json:"sampleLimit"`
	//+optional
	TargetLimit int32 `json:"targetLimit"`
	//+optional
	LabelLimit int32 `json:"labelLimit"`
	//+optional
	LabelNameLengthLimit int32 `json:"labelNameLengthLimit"`
	//+optional
	LabelValueLengthLimit int32 `json:"labelValueLengthLimit"`
	HonorLabels           bool  `json:"honorLabels"`
	// +optional
	HonorTimestamps *bool `json:"honorTimestamps"`
	//+optional
	Interval string `json:"interval"`
	//+optional
	ScrapeTimeout string `json:"scrapeTimeout"`
	// +optional
	AttachMetadata *apiextensionsv1.JSON `json:"attachMetadata"`
	// +optional
	Relabelings *apiextensionsv1.JSON `json:"relabelings"`
	// +optional
	MetricRelabelings *apiextensionsv1.JSON `json:"metricRelabelings"`
	// +optional
	AdditionalConfigs *apiextensionsv1.JSON `json:"additionalConfigs"`
	// +optional
	AdditionalEndpointConfigs *apiextensionsv1.JSON `json:"additionalEndpointConfigs"`
}

type PromLabelProxyKubeRBACProxy struct {
	Enabled bool `json:"enabled"`
	// +optional
	Config *apiextensionsv1.JSON `json:"config"`
	Image  PromLabelProxyImage   `json:"image"`
	//+optional
	ExtraArgs []string `json:"extraArgs"`
	Port      int32    `json:"port"`
	// +optional
	ContainerSecurityContext *core.SecurityContext `json:"containerSecurityContext"`
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// +optional
	VolumeMounts []core.VolumeMount `json:"volumeMounts"`
}

type PromLabelProxyClickhouse struct {
	Enabled        bool   `json:"enabled"`
	Version        string `json:"version"`
	DeletionPolicy string `json:"deletionPolicy"`
	DeploymentMode string `json:"deploymentMode"`
	// +optional
	TLS PromLabelProxyClickhouseTLS `json:"tls"`
	// +optional
	Storage PromLabelProxyStorage `json:"storage"`
	// +optional
	ClusterTopology PromLabelProxyClickhouseClusterTopology `json:"clusterTopology"`
	// +optional
	S3 PromLabelProxyClickhouseS3 `json:"s3"`
}

type PromLabelProxyClickhouseClusterTopology struct {
	// +optional
	ClickHouseKeeper PromLabelProxyClickhouseKeeper `json:"clickHouseKeeper"`
	// +optional
	Cluster PromLabelProxyClickhouseCluster `json:"cluster"`
}

type PromLabelProxyClickhouseKeeper struct {
	ExternallyManaged bool `json:"externallyManaged"`
	Replicas          int  `json:"replicas"`
	// +optional
	Storage PromLabelProxyStorage `json:"storage"`
}

type PromLabelProxyClickhouseCluster struct {
	Name     string `json:"name"`
	Shards   int    `json:"shards"`
	Replicas int    `json:"replicas"`
	// +optional
	PodTemplate PromLabelProxyClickhousePodTemplate `json:"podTemplate"`
	// +optional
	Storage PromLabelProxyStorage `json:"storage"`
}

type PromLabelProxyClickhousePodTemplate struct {
	// +optional
	Spec PromLabelProxyClickhousePodSpec `json:"spec"`
}

type PromLabelProxyClickhousePodSpec struct {
	// +optional
	Containers []core.Container `json:"containers"`
	// +optional
	InitContainers []core.Container `json:"initContainers"`
}

type PromLabelProxyClickhouseTLS struct {
	// +optional
	ClientCaCertificateRefs []PromLabelProxyClickhouseCertRef `json:"clientCaCertificateRefs"`
}

type PromLabelProxyClickhouseCertRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type PromLabelProxyClickhouseS3 struct {
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
	//+optional
	SkipVerify int32 `json:"skipVerify"`
}

type PromLabelProxyStorage struct {
	//+optional
	StorageClassName string `json:"storageClassName"`
	//+optional
	AccessModes []string `json:"accessModes"`
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PromLabelProxyList is a list of PromLabelProxies
type PromLabelProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PromLabelProxy CRD objects
	Items []PromLabelProxy `json:"items,omitempty"`
}
