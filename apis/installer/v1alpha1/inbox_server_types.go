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
	ResourceKindInboxServer = "InboxServer"
	ResourceInboxServer     = "inboxserver"
	ResourceInboxServers    = "inboxservers"
)

// InboxServer defines the schama for InboxServer operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type InboxServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InboxServerSpec `json:"spec,omitempty"`
}

// InboxServerSpec is the schema for Inbox server values file
type InboxServerSpec struct {
	NameOverride     string                 `json:"nameOverride"`
	FullnameOverride string                 `json:"fullnameOverride"`
	Dns              DnsSpec                `json:"dns"`
	James            JamesSpec              `json:"james"`
	ServiceMonitor   ServiceMonitorSpec     `json:"serviceMonitor"`
	Ingress          IngressSpec            `json:"ingress"`
	AdminIngress     IngressSpec            `json:"adminIngress"`
	Postgresql       PostgresqlSpec         `json:"postgresql"`
	Migrate          InboxServerMigrateSpec `json:"migrate"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type JamesSpec struct {
	ReplicaJmapInstanceCount     int          `json:"replicaJmapInstanceCount"`
	ReplicaImapSmtpInstanceCount int          `json:"replicaImapSmtpInstanceCount"`
	Image                        string       `json:"image"`
	Env                          JamesEnv     `json:"env"`
	TLS                          TLSSpec      `json:"tls"`
	Secret                       *JamesSecret `json:"secret,omitempty"`
}

type JamesEnv struct {
	JamesMessageSize      string                    `json:"jamesMessageSize"`
	JamesDkimSignSmtp     string                    `json:"jamesDkimSignSmtp"`
	JamesDkimSignDomain   string                    `json:"jamesDkimSignDomain"`
	JamesKeyStorePassword string                    `json:"jamesKeyStorePassword"`
	JamesHELOMessage      string                    `json:"jamesHELOMessage"`
	JvmOpts               string                    `json:"jvmOpts"`
	Glowroot              GlowrootSpec              `json:"glowroot"`
	JamesResources        core.ResourceRequirements `json:"jamesResources"`
}

type PostgresqlSpec struct {
	Enabled  bool                `json:"enabled"`
	Auth     *PostgresqlAuth     `json:"auth,omitempty"`
	External *ExternalPostgresql `json:"external,omitempty"`
}

type PostgresqlAuth struct {
	SecretName string `json:"secretName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
}

type ExternalPostgresql struct {
	Addr string `json:"addr"`
}

type InboxServerMigrateSpec struct {
	Image              ImageReference            `json:"image"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	SecurityContext    *core.SecurityContext     `json:"securityContext"`
	Resources          core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	Tolerations  []core.Toleration `json:"tolerations"`
	Affinity     *core.Affinity    `json:"affinity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InboxServerList is a list of InboxServers
type InboxServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InboxServer CRD objects
	Items []InboxServer `json:"items,omitempty"`
}
