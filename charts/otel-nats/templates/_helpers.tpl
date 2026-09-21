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

{{- define "otel-nats.natsClientURL" -}}
{{- if .Values.forwarder.natsClientURL -}}
{{- .Values.forwarder.natsClientURL -}}
{{- else -}}
{{- printf "tls://%s.%s.svc:4222" (include "otel-nats.natsService" .) .Release.Namespace -}}
{{- end -}}
{{- end }}

{{/*
Defaults to the in-cluster FQDN: a SAN on the gateway cert that b3 pins to the
load balancer IP with a hostAlias, so no public DNS record is needed.
*/}}
{{- define "otel-nats.publicHostname" -}}
{{- if .Values.publicHostname -}}
{{- .Values.publicHostname -}}
{{- else -}}
{{- printf "%s.%s.svc.cluster.local" (include "otel-nats.natsService" .) .Release.Namespace -}}
{{- end -}}
{{- end }}

{{- define "otel-nats.plpBase" -}}
{{- printf "https://prom-label-proxy.%s.svc:%v" .Values.promLabelProxyNamespace .Values.promLabelProxyPort -}}
{{- end }}

{{/*
A stream cannot have more replicas than there are nodes.
*/}}
{{- define "otel-nats.streamReplicas" -}}
{{- if .Values.nats.cluster.enabled -}}
{{- min 3 .Values.nats.cluster.replicas -}}
{{- else -}}
1
{{- end -}}
{{- end }}
