# AppsCode OpenTelemetry Stack

[AppsCode OpenTelemetry Stack](https://github.com/ops-center) - AppsCode OpenTelemetry Stack

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/appscode-otel-stack --version=v2026.4.30
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2026.4.30
```

## Introduction

This chart deploys AppsCode OpenTelemetry Stack on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.25+

## Installing the Chart

To install/upgrade the chart with the release name `appscode-otel-stack`:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2026.4.30
```

The command deploys AppsCode OpenTelemetry Stack on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `appscode-otel-stack`:

```bash
$ helm uninstall appscode-otel-stack -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `appscode-otel-stack` chart and their default values.

|                                     Parameter                                      | Description |                                          Default                                          |
|------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------------|
| clusterName                                                                        |             | <code>ace</code>                                                                          |
| opentelemetry-operator.manager.featureGatesMap.operator.targetallocator.mtls       |             | <code>true</code>                                                                         |
| opentelemetry-operator.manager.createRbacPermissions                               |             | <code>true</code>                                                                         |
| collectors.daemon.resources.limits.cpu                                             |             | <code>500m</code>                                                                         |
| collectors.daemon.resources.limits.memory                                          |             | <code>1Gi</code>                                                                          |
| collectors.daemon.resources.requests.cpu                                           |             | <code>200m</code>                                                                         |
| collectors.daemon.resources.requests.memory                                        |             | <code>512Mi</code>                                                                        |
| collectors.daemon.suffix                                                           |             | <code>agent</code>                                                                        |
| collectors.daemon.enabled                                                          |             | <code>true</code>                                                                         |
| collectors.daemon.scrape_configs_file                                              |             | <code>"config/kubelet_scrape_configs.yaml"</code>                                         |
| collectors.daemon.labels.otel-collector-type                                       |             | <code>daemonset</code>                                                                    |
| collectors.daemon.targetAllocator.enabled                                          |             | <code>true</code>                                                                         |
| collectors.daemon.targetAllocator.serviceAccount                                   |             | <code>targetallocator</code>                                                              |
| collectors.daemon.targetAllocator.image                                            |             | <code>ghcr.io/open-telemetry/opentelemetry-operator/target-allocator:main</code>          |
| collectors.daemon.targetAllocator.allocationStrategy                               |             | <code>per-node</code>                                                                     |
| collectors.daemon.targetAllocator.prometheusCR.enabled                             |             | <code>true</code>                                                                         |
| collectors.daemon.targetAllocator.prometheusCR.podMonitorSelector                  |             | <code>{}</code>                                                                           |
| collectors.daemon.targetAllocator.prometheusCR.scrapeInterval                      |             | <code>"30s"</code>                                                                        |
| collectors.daemon.targetAllocator.prometheusCR.serviceMonitorSelector              |             | <code>{}</code>                                                                           |
| collectors.daemon.config.receivers.k8sobjects.auth_type                            |             | <code>serviceAccount</code>                                                               |
| collectors.daemon.config.exporters.otlp.endpoint                                   |             | <code>opentelemetry-kube-stack-gateway-collector.monitoring.svc.cluster.local:4317</code> |
| collectors.daemon.config.exporters.otlp.tls.insecure                               |             | <code>true</code>                                                                         |
| collectors.daemon.config.exporters.debug.verbosity                                 |             | <code>detailed</code>                                                                     |
| collectors.daemon.config.processors.memory_limiter.check_interval                  |             | <code>1s</code>                                                                           |
| collectors.daemon.config.processors.memory_limiter.limit_percentage                |             | <code>75</code>                                                                           |
| collectors.daemon.config.processors.memory_limiter.spike_limit_percentage          |             | <code>20</code>                                                                           |
| collectors.daemon.config.processors.batch.send_batch_size                          |             | <code>256</code>                                                                          |
| collectors.daemon.config.processors.batch.send_batch_max_size                      |             | <code>512</code>                                                                          |
| collectors.daemon.config.processors.batch.timeout                                  |             | <code>5s</code>                                                                           |
| collectors.daemon.presets.logsCollection.enabled                                   |             | <code>true</code>                                                                         |
| collectors.daemon.presets.hostMetrics.enabled                                      |             | <code>false</code>                                                                        |
| collectors.daemon.presets.kubeletMetrics.enabled                                   |             | <code>false</code>                                                                        |
| collectors.daemon.presets.kubernetesAttributes.enabled                             |             | <code>true</code>                                                                         |
| collectors.cluster.resources.limits.cpu                                            |             | <code>500m</code>                                                                         |
| collectors.cluster.resources.limits.memory                                         |             | <code>1Gi</code>                                                                          |
| collectors.cluster.resources.requests.cpu                                          |             | <code>200m</code>                                                                         |
| collectors.cluster.resources.requests.memory                                       |             | <code>512Mi</code>                                                                        |
| collectors.cluster.suffix                                                          |             | <code>gateway</code>                                                                      |
| collectors.cluster.enabled                                                         |             | <code>true</code>                                                                         |
| collectors.cluster.labels.otel-collector-type                                      |             | <code>deployment</code>                                                                   |
| collectors.cluster.image.repository                                                |             | <code>rokibulhasan114/opentelemetry-collector-ace</code>                                  |
| collectors.cluster.image.tag                                                       |             | <code>v0.1.0</code>                                                                       |
| collectors.cluster.image.pullPolicy                                                |             | <code>Always</code>                                                                       |
| collectors.cluster.presets.clusterMetrics.enabled                                  |             | <code>false</code>                                                                        |
| collectors.cluster.config.receivers.otlp.protocols.grpc.endpoint                   |             | <code>0.0.0.0:4317</code>                                                                 |
| collectors.cluster.config.receivers.otlp.protocols.grpc.include_metadata           |             | <code>true</code>                                                                         |
| collectors.cluster.config.receivers.otlp.protocols.grpc.max_recv_msg_size_mib      |             | <code>32</code>                                                                           |
| collectors.cluster.config.receivers.otlp.protocols.http.endpoint                   |             | <code>0.0.0.0:4318</code>                                                                 |
| collectors.cluster.config.receivers.otlp.protocols.http.include_metadata           |             | <code>true</code>                                                                         |
| collectors.cluster.config.processors.memory_limiter.check_interval                 |             | <code>1s</code>                                                                           |
| collectors.cluster.config.processors.memory_limiter.limit_percentage               |             | <code>75</code>                                                                           |
| collectors.cluster.config.processors.memory_limiter.spike_limit_percentage         |             | <code>20</code>                                                                           |
| collectors.cluster.config.processors.transform/tenant.error_mode                   |             | <code>ignore</code>                                                                       |
| collectors.cluster.config.processors.batch/tenant.send_batch_size                  |             | <code>256</code>                                                                          |
| collectors.cluster.config.processors.batch/tenant.send_batch_max_size              |             | <code>512</code>                                                                          |
| collectors.cluster.config.exporters.debug.verbosity                                |             | <code>detailed</code>                                                                     |
| collectors.cluster.config.exporters.prometheusremotewrite.endpoint                 |             | <code>"https://10.2.0.187/api/v1/receive"</code>                                          |
| collectors.cluster.config.exporters.prometheusremotewrite.auth.authenticator       |             | <code>headers_setter</code>                                                               |
| collectors.cluster.config.exporters.prometheusremotewrite.target_info.enabled      |             | <code>false</code>                                                                        |
| collectors.cluster.config.exporters.prometheusremotewrite.tls.ca_file              |             | <code>/etc/certs/ca.crt</code>                                                            |
| collectors.cluster.config.exporters.prometheusremotewrite.tls.cert_file            |             | <code>/etc/certs/tls.crt</code>                                                           |
| collectors.cluster.config.exporters.prometheusremotewrite.tls.key_file             |             | <code>/etc/certs/tls.key</code>                                                           |
| collectors.cluster.config.exporters.prometheusremotewrite.tls.insecure_skip_verify |             | <code>false</code>                                                                        |
| collectors.cluster.config.exporters.clickhouse.endpoint                            |             | <code>"https://10.2.0.187/?database=default&secure=true"</code>                           |
| collectors.cluster.config.exporters.clickhouse.username                            |             | <code>"ace"</code>                                                                        |
| collectors.cluster.config.exporters.clickhouse.password                            |             | <code>"QEFNeAa7hUGjJjVY"</code>                                                           |
| collectors.cluster.config.exporters.clickhouse.database                            |             | <code>"default"</code>                                                                    |
| collectors.cluster.config.exporters.clickhouse.tls.ca_file                         |             | <code>/etc/certs/ca.crt</code>                                                            |
| collectors.cluster.config.exporters.clickhouse.tls.cert_file                       |             | <code>/etc/certs/tls.crt</code>                                                           |
| collectors.cluster.config.exporters.clickhouse.tls.key_file                        |             | <code>/etc/certs/tls.key</code>                                                           |
| collectors.cluster.config.exporters.clickhouse.tls.insecure_skip_verify            |             | <code>false</code>                                                                        |
| collectors.cluster.config.exporters.clickhouse.create_schema                       |             | <code>false</code>                                                                        |
| collectors.cluster.config.exporters.clickhouse.logs_table_name                     |             | <code>otel_logs</code>                                                                    |
| collectors.cluster.config.exporters.clickhouse.timeout                             |             | <code>10s</code>                                                                          |
| collectors.cluster.config.exporters.clickhouse.sending_queue.queue_size            |             | <code>1000</code>                                                                         |
| collectors.cluster.config.exporters.clickhouse.retry_on_failure.enabled            |             | <code>true</code>                                                                         |
| collectors.cluster.config.exporters.clickhouse.retry_on_failure.initial_interval   |             | <code>5s</code>                                                                           |
| collectors.cluster.config.exporters.clickhouse.retry_on_failure.max_interval       |             | <code>30s</code>                                                                          |
| collectors.cluster.config.exporters.clickhouse.retry_on_failure.max_elapsed_time   |             | <code>300s</code>                                                                         |
| collectors.cluster.config.exporters.otlp.endpoint                                  |             | <code>opentelemetry-kube-stack-gateway-collector.monitoring.svc.cluster.local:4317</code> |
| collectors.cluster.config.exporters.otlp.tls.insecure                              |             | <code>true</code>                                                                         |
| instrumentation.enabled                                                            |             | <code>false</code>                                                                        |
| opAMPBridge.enabled                                                                |             | <code>false</code>                                                                        |
| kubernetesServiceMonitors.enabled                                                  |             | <code>true</code>                                                                         |
| kubeApiServer.enabled                                                              |             | <code>false</code>                                                                        |
| kubelet.enabled                                                                    |             | <code>false</code>                                                                        |
| kubeControllerManager.enabled                                                      |             | <code>false</code>                                                                        |
| kubeDns.enabled                                                                    |             | <code>false</code>                                                                        |
| kubeEtcd.enabled                                                                   |             | <code>false</code>                                                                        |
| kubeScheduler.enabled                                                              |             | <code>false</code>                                                                        |
| kubeProxy.enabled                                                                  |             | <code>false</code>                                                                        |
| kubeStateMetrics.enabled                                                           |             | <code>false</code>                                                                        |
| nodeExporter.enabled                                                               |             | <code>false</code>                                                                        |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2026.4.30 --set clusterName=ace
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2026.4.30 --values values.yaml
```
