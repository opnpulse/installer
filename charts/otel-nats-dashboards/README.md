# OTel NATS Dashboards

[OTel NATS Dashboards by AppsCode](https://github.com/opnpulse/otel-nats) - Grafana dashboards for the NATS JetStream telemetry buffer

## TL;DR;

```bash
$ helm repo add appscode-charts oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode-charts/otel-nats-dashboards --version=v2026.9.22
$ helm upgrade -i  appscode-charts/otel-nats-dashboards -n  --create-namespace --version=v2026.9.22
```

## Introduction

This chart deploys OTel NATS dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.26+

## Installing the Chart

To install/upgrade the chart with the release name ``:

```bash
$ helm upgrade -i  appscode-charts/otel-nats-dashboards -n  --create-namespace --version=v2026.9.22
```

The command deploys OTel NATS dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the ``:

```bash
$ helm uninstall  -n 
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `otel-nats-dashboards` chart and their default values.

|     Parameter      |                                                                 Description                                                                 |            Default             |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------|
| namespace          | Where Grafana's sidecar looks for ConfigMaps. Empty follows the release namespace; set it if kube-prometheus-stack watches a different one. | <code>""</code>                |
| sidecarLabel.key   |                                                                                                                                             | <code>grafana_dashboard</code> |
| sidecarLabel.value |                                                                                                                                             | <code>"1"</code>               |
| folder             | Folder hint stored as an annotation on the ConfigMap, picked up by the sidecar to organize dashboards in Grafana.                           | <code>"otel-nats"</code>       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i  appscode-charts/otel-nats-dashboards -n  --create-namespace --version=v2026.9.22 --set sidecarLabel.key=grafana_dashboard
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i  appscode-charts/otel-nats-dashboards -n  --create-namespace --version=v2026.9.22 --values values.yaml
```
