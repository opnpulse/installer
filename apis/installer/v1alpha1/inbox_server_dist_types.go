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
	ResourceKindInboxServerDistributed = "InboxServerDistributed"
	ResourceInboxServerDistributed     = "inboxserver"
	ResourceInboxServerDistributeds    = "inboxservers"
)

// InboxServerDistributed defines the schama for InboxServerDistributed operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type InboxServerDistributed struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InboxServerDistributedSpec `json:"spec,omitempty"`
}

// InboxServerDistributedSpec is the schema for Inbox server values file
type InboxServerDistributedSpec struct {
	NameOverride     string               `json:"nameOverride"`
	FullnameOverride string               `json:"fullnameOverride"`
	Dns              DnsSpec              `json:"dns"`
	James            JamesDistributedSpec `json:"james"`
	ServiceMonitor   ServiceMonitorSpec   `json:"serviceMonitor"`
	Ingress          IngressSpec          `json:"ingress"`
	AdminIngress     IngressSpec          `json:"adminIngress"`
	Cassandra        CassandraSpec        `json:"cassandra"`
	Opensearch       OpensearchSpec       `json:"opensearch"`
	Rabbitmq         RabbitmqSpec         `json:"rabbitmq"`
	S3               S3Spec               `json:"s3"`
	Tika             TikaSpec             `json:"tika"`
	Service          ServiceSpec          `json:"service"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type DnsSpec struct {
	Domain      string   `json:"domain"`
	EmailDomain string   `json:"emailDomain"`
	TargetIPs   []string `json:"targetIPs"`
}

type JamesDistributedSpec struct {
	ReplicaJmapInstanceCount     int                 `json:"replicaJmapInstanceCount"`
	ReplicaImapSmtpInstanceCount int                 `json:"replicaImapSmtpInstanceCount"`
	Image                        string              `json:"image"`
	Env                          JamesDistributedEnv `json:"env"`
	TLS                          TLSSpec             `json:"tls"`
	Secret                       *JamesSecret        `json:"secret,omitempty"`
}

type JamesDistributedEnv struct {
	JamesCassandraKeyspace       string                    `json:"jamesCassandraKeyspace"`
	JamesCassandraCacheKeyspace  string                    `json:"jamesCassandraCacheKeyspace"`
	JamesEsMailboxIndex          string                    `json:"jamesEsMailboxIndex"`
	JamesEsClusterName           string                    `json:"jamesEsClusterName"`
	JamesEsHostScheme            string                    `json:"jamesEsHostScheme"`
	JamesEsSslValidationStrategy string                    `json:"jamesEsSslValidationStrategy"`
	JamesEsHostNameVerifier      string                    `json:"jamesEsHostNameVerifier"`
	JamesEsNbShards              int                       `json:"jamesEsNbShards"`
	JamesEsNbReplica             int                       `json:"jamesEsNbReplica"`
	JamesEsMailboxReadAlias      string                    `json:"jamesEsMailboxReadAlias"`
	JamesEsMailboxWriteAlias     string                    `json:"jamesEsMailboxWriteAlias"`
	JamesRabbitHost              string                    `json:"jamesRabbitHost"`
	JamesMessageSize             string                    `json:"jamesMessageSize"`
	JamesDkimSignSmtp            string                    `json:"jamesDkimSignSmtp"`
	JamesDkimSignDomain          string                    `json:"jamesDkimSignDomain"`
	CassandraReplicationFactor   int                       `json:"cassandraReplicationFactor"`
	JamesKeyStorePassword        string                    `json:"jamesKeyStorePassword"`
	JamesHELOMessage             string                    `json:"jamesHELOMessage"`
	JvmOpts                      string                    `json:"jvmOpts"`
	Glowroot                     GlowrootSpec              `json:"glowroot"`
	JamesResources               core.ResourceRequirements `json:"jamesResources"`
}

type JamesSecret struct {
	AdminJWTPublicKey string `json:"adminJWTPublicKey"`
	JwtPublicKey      string `json:"jwtPublicKey"`
	JwtPrivateKey     string `json:"jwtPrivateKey"`
	DkimPrivateKey    string `json:"dkimPrivateKey"`
}

type GlowrootSpec struct {
	Enabled bool `json:"enabled"`
}

type ServiceMonitorSpec struct {
	Enabled          bool              `json:"enabled"`
	AdditionalLabels map[string]string `json:"additionalLabels"`
	Interval         string            `json:"interval"`
}

type IngressSpec struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	TLS         TLSSpec           `json:"tls"`
}

type TLSSpec struct {
	SecretName string `json:"secretName"`
}

type CassandraSpec struct {
	Image         ImageReference  `json:"image"`
	LivenessProbe core.Probe      `json:"livenessProbe"`
	Secret        CassandraSecret `json:"secret"`
}

type CassandraSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OpensearchSpec struct {
	Enabled bool             `json:"enabled"`
	Image   ImageReference   `json:"image"`
	Secret  OpensearchSecret `json:"secret"`
}

type OpensearchSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RabbitmqSpec struct {
	Image  ImageReference `json:"image"`
	Secret RabbitmqSecret `json:"secret"`
}

type RabbitmqSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type S3Spec struct {
	Enabled     bool           `json:"enabled"`
	Image       ImageReference `json:"image"`
	Endpoint    string         `json:"endpoint"`
	Region      string         `json:"region"`
	AccessKeyId string         `json:"accessKeyId"`
	SecretKey   string         `json:"secretKey"`
}

type TikaSpec struct {
	Enabled bool           `json:"enabled"`
	Image   ImageReference `json:"image"`
}

type ServiceSpec struct {
	Type string `json:"type"`
}

type ImageReference struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	Tag        string `json:"tag"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InboxServerDistributedList is a list of InboxServerDistributeds
type InboxServerDistributedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InboxServerDistributed CRD objects
	Items []InboxServerDistributed `json:"items,omitempty"`
}
