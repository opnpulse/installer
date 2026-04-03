#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

CERT_MANAGER_CERT_MANAGER_TAG=${CERT_MANAGER_CERT_MANAGER_TAG:-v1.19.2}
FLUXCD_HELM_CONTROLLER_TAG=${FLUXCD_HELM_CONTROLLER_TAG:-v1.3.0}
FLUXCD_SOURCE_CONTROLLER_TAG=${FLUXCD_SOURCE_CONTROLLER_TAG:-v1.6.2}
KMODULES_CUSTOM_RESOURCES_TAG=${KMODULES_CUSTOM_RESOURCES_TAG:-v0.34.0}
KMODULES_RESOURCE_METADATA_TAG=${KMODULES_RESOURCE_METADATA_TAG:-master}
KUBEOPS_CSI_DRIVER_CACERTS_TAG=${KUBEOPS_CSI_DRIVER_CACERTS_TAG:-v0.5.0}
KUBEOPS_EXTERNAL_DNS_OPERATOR_TAG=${KUBEOPS_EXTERNAL_DNS_OPERATOR_TAG:-v0.3.0}
KUBEOPS_OPERATOR_SHARD_MANAGER_TAG=${KUBEOPS_OPERATOR_SHARD_MANAGER_TAG:-v0.5.0}
KUBEOPS_PETSET_TAG=${KUBEOPS_PETSET_TAG:-v0.0.16}
KUBEOPS_SIDEKICK_TAG=${KUBEOPS_SIDEKICK_TAG:-v0.0.14}
KUBEOPS_TASKQUEUE_TAG=${KUBEOPS_TASKQUEUE_TAG:-v0.0.3}
OPEN_POLICY_AGENT_GATEKEEPER_TAG=${OPEN_POLICY_AGENT_GATEKEEPER_TAG:-v3.14.0}
OPEN_VIZ_APIMACHINERY_TAG=${OPEN_VIZ_APIMACHINERY_TAG:-v0.0.8}
PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR=${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR:-v0.87.1}
X_HELM_APIMACHINERY_TAG=${X_HELM_APIMACHINERY_TAG:-v0.0.18}

crd-importer \
    --no-description \
    --input=https://github.com/kubeops/sidekick/raw/${KUBEOPS_SIDEKICK_TAG}/crds/apps.k8s.appscode.com_sidekicks.yaml \
    --out=./charts/sidekick/crds

crd-importer \
    --no-description \
    --input=https://github.com/cert-manager/cert-manager/releases/download/${CERT_MANAGER_CERT_MANAGER_TAG}/cert-manager.crds.yaml \
    --gk=Certificate.cert-manager.io --gk=Issuer.cert-manager.io \
    --out=./charts/sidekick/crds
