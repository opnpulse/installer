# Thanos Operator

[Thanos Operator by AppsCode](https://github.com/opnpulse/thanos-operator) - Thanos Operator for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/thanos-operator --version=v2026.4.30
$ helm upgrade -i thanos-operator appscode/thanos-operator -n kubeops --create-namespace --version=v2026.4.30
```

## Introduction

This chart deploys a Thanos Operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `thanos-operator`:

```bash
$ helm upgrade -i thanos-operator appscode/thanos-operator -n kubeops --create-namespace --version=v2026.4.30
```

The command deploys a Thanos Operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `thanos-operator`:

```bash
$ helm uninstall thanos-operator -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `thanos-operator` chart and their default values.

|                    Parameter                     |                                                   Description                                                   |                    Default                    |
|--------------------------------------------------|-----------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| manager.replicas                                 |                                                                                                                 | <code>1</code>                                |
| manager.image.repository                         |                                                                                                                 | <code>ghcr.io/opnpulse/thanos-operator</code> |
| manager.image.tag                                |                                                                                                                 | <code>v2026.4.24</code>                       |
| manager.image.pullPolicy                         |                                                                                                                 | <code>IfNotPresent</code>                     |
| manager.env                                      | # Environment variables #                                                                                       | <code>[]</code>                               |
| manager.envOverrides                             | # Env overrides (--set manager.envOverrides.VAR=value) # Same name in env above: this value takes precedence. # | <code>{}</code>                               |
| manager.imagePullSecrets                         | # Image pull secrets #                                                                                          | <code>[]</code>                               |
| manager.podSecurityContext.runAsNonRoot          |                                                                                                                 | <code>true</code>                             |
| manager.securityContext.allowPrivilegeEscalation |                                                                                                                 | <code>false</code>                            |
| manager.securityContext.seccompProfile.type      |                                                                                                                 | <code>RuntimeDefault</code>                   |
| manager.resources.limits.cpu                     |                                                                                                                 | <code>500m</code>                             |
| manager.resources.limits.memory                  |                                                                                                                 | <code>128Mi</code>                            |
| manager.resources.requests.cpu                   |                                                                                                                 | <code>10m</code>                              |
| manager.resources.requests.memory                |                                                                                                                 | <code>64Mi</code>                             |
| manager.affinity                                 | # Manager pod's affinity #                                                                                      | <code>{}</code>                               |
| manager.nodeSelector                             | # Manager pod's node selector #                                                                                 | <code>{}</code>                               |
| manager.tolerations                              | # Manager pod's tolerations #                                                                                   | <code>[]</code>                               |
| rbacHelpers.enable                               | Install convenience admin/editor/viewer roles for CRDs                                                          | <code>false</code>                            |
| crd.enable                                       | Install CRDs with the chart                                                                                     | <code>false</code>                            |
| crd.keep                                         | Keep CRDs when uninstalling                                                                                     | <code>true</code>                             |
| metrics.enable                                   |                                                                                                                 | <code>true</code>                             |
| metrics.port                                     | Metrics server port                                                                                             | <code>8443</code>                             |
| certManager.enable                               |                                                                                                                 | <code>false</code>                            |
| prometheus.enable                                |                                                                                                                 | <code>false</code>                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i thanos-operator appscode/thanos-operator -n kubeops --create-namespace --version=v2026.4.30 --set manager.replicas=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i thanos-operator appscode/thanos-operator -n kubeops --create-namespace --version=v2026.4.30 --values values.yaml
```
