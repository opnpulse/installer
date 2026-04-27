{{/*
Contruct a temporary URL for Cassandra. This URL ends with a comma.
*/}}
{{- define "cassandra.url.list" -}}
{{- printf "%s-cassandra.%s.svc:%s," (include "inbox-server-distributed.fullname" .) .Release.Namespace "9042" }}
{{- end -}}


{{/*
Contruct a temporary URLfor Elastic Search. This URL ends with a comma
*/}}
{{- define "opensearch.url.list" -}}
{{- printf "%s-opensearch.%s.svc:%s," (include "inbox-server-distributed.fullname" .) .Release.Namespace "9200" }}
{{- end -}}
