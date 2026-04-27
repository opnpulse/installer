# Inbox Server

[Inbox Server by AppsCode](https://github.com/ops-center/james-project) - Inbox Server by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/inbox-server-distributed --version=v2025.12.25
$ helm upgrade -i inbox-server-distributed appscode/inbox-server-distributed -n monitoring --create-namespace --version=v2025.12.25
```

## Introduction

This chart deploys Inbox Server on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `inbox-server-distributed`:

```bash
$ helm upgrade -i inbox-server-distributed appscode/inbox-server-distributed -n monitoring --create-namespace --version=v2025.12.25
```

The command deploys Inbox Server on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `inbox-server-distributed`:

```bash
$ helm uninstall inbox-server-distributed -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `inbox-server-distributed` chart and their default values.

|                 Parameter                 |                                                                Description                                                                 |                      Default                      |
|-------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| nameOverride                              | Default values for inbox-agent. This is a YAML-formatted file. Declare variables to be passed into your templates. Overrides name template | <code>''</code>                                   |
| fullnameOverride                          | Overrides fullname template                                                                                                                | <code>''</code>                                   |
| dns.domain                                | # Default domain for all emailing components jmapUrl: "jmap.example.com" adminUrl: "admin.example.com" smtpHostname: "smtp.example.com"    | <code>''</code>                                   |
| dns.emailDomain                           | in ip mode, the emailDomain is ace.internal, otherwise same as domain                                                                      | <code>''</code>                                   |
| dns.targetIPs                             |                                                                                                                                            | <code>[]</code>                                   |
| james.replicaJmapInstanceCount            |                                                                                                                                            | <code>1</code>                                    |
| james.replicaImapSmtpInstanceCount        |                                                                                                                                            | <code>0</code>                                    |
| james.image                               |                                                                                                                                            | <code>ghcr.io/appscode/inbox-server:latest</code> |
| james.secret.jwtPublicKey                 |                                                                                                                                            | <code>''</code>                                   |
| james.secret.jwtPrivateKey                |                                                                                                                                            | <code>''</code>                                   |
| james.secret.adminJWTPublicKey            |                                                                                                                                            | <code>''</code>                                   |
| james.secret.dkimPrivateKey               |                                                                                                                                            | <code>''</code>                                   |
| james.env.jamesCassandraKeyspace          |                                                                                                                                            | <code>sandbox_james</code>                        |
| james.env.jamesCassandraCacheKeyspace     |                                                                                                                                            | <code>sandbox_james_cache</code>                  |
| james.env.jamesEsMailboxIndex             |                                                                                                                                            | <code>mailbox_v1</code>                           |
| james.env.jamesEsClusterName              |                                                                                                                                            | <code>change-me</code>                            |
| james.env.jamesEsHostScheme               |                                                                                                                                            | <code>http</code>                                 |
| james.env.jamesEsSslValidationStrategy    |                                                                                                                                            | <code>default</code>                              |
| james.env.jamesEsHostNameVerifier         |                                                                                                                                            | <code>default</code>                              |
| james.env.jamesEsNbShards                 |                                                                                                                                            | <code>5</code>                                    |
| james.env.jamesEsNbReplica                |                                                                                                                                            | <code>1</code>                                    |
| james.env.jamesEsMailboxReadAlias         |                                                                                                                                            | <code>read-mailbox</code>                         |
| james.env.jamesEsMailboxWriteAlias        |                                                                                                                                            | <code>write-mailbox</code>                        |
| james.env.jamesRabbitHost                 |                                                                                                                                            | <code>inbox-server-distributed-rabbitmq</code>    |
| james.env.jamesMessageSize                |                                                                                                                                            | <code>25M</code>                                  |
| james.env.jamesDkimSignSmtp               |                                                                                                                                            | <code>dkimselector</code>                         |
| james.env.jamesDkimSignDomain             |                                                                                                                                            | <code>mail.example.com</code>                     |
| james.env.cassandraReplicationFactor      |                                                                                                                                            | <code>1</code>                                    |
| james.env.jamesKeyStorePassword           |                                                                                                                                            | <code>james72laBalle</code>                       |
| james.env.jamesHELOMessage                |                                                                                                                                            | <code>change-me</code>                            |
| james.env.jvmOpts                         |                                                                                                                                            | <code>-Xms3g -Xmx3g</code>                        |
| james.env.glowroot.enabled                |                                                                                                                                            | <code>false</code>                                |
| james.env.jamesResources.limits.cpu       |                                                                                                                                            | <code>'2'</code>                                  |
| james.env.jamesResources.limits.memory    |                                                                                                                                            | <code>4Gi</code>                                  |
| james.env.jamesResources.requests.cpu     |                                                                                                                                            | <code>'1'</code>                                  |
| james.env.jamesResources.requests.memory  |                                                                                                                                            | <code>4Gi</code>                                  |
| james.tls.secretName                      | secretName: the-name-of-a-secret                                                                                                           | <code>''</code>                                   |
| serviceMonitor.enabled                    |                                                                                                                                            | <code>false</code>                                |
| serviceMonitor.additionalLabels.change-me |                                                                                                                                            | <code>change-me</code>                            |
| serviceMonitor.interval                   | Scrape interval. Use Prometheus default value if not specified                                                                             | <code>30s</code>                                  |
| ingress.enabled                           |                                                                                                                                            | <code>false</code>                                |
| ingress.className                         |                                                                                                                                            | <code>''</code>                                   |
| ingress.annotations                       |                                                                                                                                            | <code>{}</code>                                   |
| ingress.tls.secretName                    |                                                                                                                                            | <code>the-name-of-a-secret</code>                 |
| adminIngress.enabled                      |                                                                                                                                            | <code>false</code>                                |
| adminIngress.className                    |                                                                                                                                            | <code>''</code>                                   |
| adminIngress.annotations                  |                                                                                                                                            | <code>{}</code>                                   |
| adminIngress.tls.secretName               |                                                                                                                                            | <code>the-name-of-a-secret</code>                 |
| cassandra.image.repository                |                                                                                                                                            | <code>cassandra</code>                            |
| cassandra.image.pullPolicy                |                                                                                                                                            | <code>IfNotPresent</code>                         |
| cassandra.image.tag                       | Overrides the image tag whose default is the chart appVersion.                                                                             | <code>4.1.3</code>                                |
| cassandra.livenessProbe.failureThreshold  |                                                                                                                                            | <code>5</code>                                    |
| cassandra.livenessProbe.periodSeconds     |                                                                                                                                            | <code>3</code>                                    |
| cassandra.livenessProbe.timeoutSeconds    |                                                                                                                                            | <code>20</code>                                   |
| cassandra.secret.username                 |                                                                                                                                            | <code>''</code>                                   |
| cassandra.secret.password                 |                                                                                                                                            | <code>''</code>                                   |
| opensearch.enabled                        |                                                                                                                                            | <code>true</code>                                 |
| opensearch.image.repository               |                                                                                                                                            | <code>opensearchproject/opensearch</code>         |
| opensearch.image.pullPolicy               |                                                                                                                                            | <code>IfNotPresent</code>                         |
| opensearch.image.tag                      | Overrides the image tag whose default is the chart appVersion.                                                                             | <code>2.1.0</code>                                |
| opensearch.secret.username                |                                                                                                                                            | <code>''</code>                                   |
| opensearch.secret.password                |                                                                                                                                            | <code>''</code>                                   |
| rabbitmq.image.repository                 |                                                                                                                                            | <code>rabbitmq</code>                             |
| rabbitmq.image.pullPolicy                 |                                                                                                                                            | <code>IfNotPresent</code>                         |
| rabbitmq.image.tag                        | Overrides the image tag whose default is the chart appVersion.                                                                             | <code>3.12.1-management</code>                    |
| rabbitmq.secret.username                  |                                                                                                                                            | <code>''</code>                                   |
| rabbitmq.secret.password                  |                                                                                                                                            | <code>''</code>                                   |
| s3.enabled                                |                                                                                                                                            | <code>false</code>                                |
| s3.image.repository                       |                                                                                                                                            | <code>ghcr.io/appscode/cloudserver</code>         |
| s3.image.pullPolicy                       |                                                                                                                                            | <code>IfNotPresent</code>                         |
| s3.image.tag                              | Overrides the image tag whose default is the chart appVersion.                                                                             | <code>8.8.20</code>                               |
| s3.endpoint                               |                                                                                                                                            | <code>''</code>                                   |
| s3.region                                 |                                                                                                                                            | <code>''</code>                                   |
| s3.accessKeyId                            |                                                                                                                                            | <code>''</code>                                   |
| s3.secretKey                              |                                                                                                                                            | <code>''</code>                                   |
| tika.enabled                              |                                                                                                                                            | <code>false</code>                                |
| tika.image.repository                     |                                                                                                                                            | <code>apache/tika</code>                          |
| tika.image.pullPolicy                     |                                                                                                                                            | <code>IfNotPresent</code>                         |
| tika.image.tag                            | Overrides the image tag whose default is the chart appVersion.                                                                             | <code>2.8.0.0</code>                              |
| service.type                              |                                                                                                                                            | <code>ClusterIP</code>                            |
| distro.openshift                          | Set true, if installed in OpenShift                                                                                                        | <code>false</code>                                |
| distro.ubi                                | Set operator or all to use ubi images                                                                                                      | <code>""</code>                                   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i inbox-server-distributed appscode/inbox-server-distributed -n monitoring --create-namespace --version=v2025.12.25 --set nameOverride=''
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i inbox-server-distributed appscode/inbox-server-distributed -n monitoring --create-namespace --version=v2025.12.25 --values values.yaml
```
