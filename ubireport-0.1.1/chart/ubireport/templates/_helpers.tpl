{{/*
Expand the name of the chart. (ubireport)
*/}}
{{- define "ubireport.name" -}}
{{- default .Chart.Name .Values.ubireport.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
This is default ubireport (63 chars k8s name limited.)
*/}}
{{- define "ubireport.fullnameOverride" -}}
{{- if .Values.ubireport.fullnameOverride }}
{{- .Values.ubireport.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.ubireport.fullnameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}


