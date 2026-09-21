{{- define "otel-nats.labels" -}}
app.kubernetes.io/name: otel-nats
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
{{- end -}}

{{- define "otel-nats.forwarder.selector" -}}
app.kubernetes.io/name: otel-nats
app.kubernetes.io/component: forwarder
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Name of the Service the vendored nats subchart creates. Mirrors its own
fullname template, so the forwarder addresses NATS wherever the release lands.
*/}}
{{- define "otel-nats.natsService" -}}
{{- if contains "nats" .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-nats" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end }}

{{/*
Client URL for the in-cluster NATS Service, unless one is configured.
*/}}
{{- define "otel-nats.natsClientURL" -}}
{{- if .Values.forwarder.natsClientURL -}}
{{- .Values.forwarder.natsClientURL -}}
{{- else -}}
{{- printf "tls://%s.%s.svc:4222" (include "otel-nats.natsService" .) .Release.Namespace -}}
{{- end -}}
{{- end }}

{{/*
Host the workload gateways dial. Defaults to the in-cluster FQDN, which is a
SAN on the gateway certificate and which b3 pins to the load balancer IP with
a hostAlias. That works in both ip and domain mode, so no public DNS record is
needed. Override only when a real public name exists.
*/}}
{{- define "otel-nats.publicHostname" -}}
{{- if .Values.publicHostname -}}
{{- .Values.publicHostname -}}
{{- else -}}
{{- printf "%s.%s.svc.cluster.local" (include "otel-nats.natsService" .) .Release.Namespace -}}
{{- end -}}
{{- end }}

{{/*
Base URL of the prom-label-proxy forwarder listener.
*/}}
{{- define "otel-nats.plpBase" -}}
{{- printf "https://prom-label-proxy.%s.svc:%v" .Values.promLabelProxyNamespace .Values.promLabelProxyPort -}}
{{- end }}

{{/*
Stream replica count. A stream cannot have more replicas than there are nodes.
*/}}
{{- define "otel-nats.streamReplicas" -}}
{{- if .Values.nats.cluster.enabled -}}
{{- min 3 .Values.nats.cluster.replicas -}}
{{- else -}}
1
{{- end -}}
{{- end }}
