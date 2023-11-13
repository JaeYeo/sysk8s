{{/*
Expand the name of the chart. (veleroApi, veleroApp)
*/}}
{{- define "SysMasterK8sVelero.veleroApi.name" -}}
{{- default .Chart.Name .Values.veleroApi.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- define "SysMasterK8sVelero.veleroApp.name" -}}
{{- default .Chart.Name .Values.veleroApp.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
This is default velero Api name: veleroApi (63 chars k8s name limited.)
*/}}
{{- define "veleroApi.fullname" -}}
{{- if .Values.veleroApi.fullnameOverride }}
{{- .Values.veleroApi.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.veleroApi.fullnameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
This is default velero App name: veleroApp (63 chars k8s name limited.)
*/}}
{{- define "veleroApp.fullname" -}}
{{- if .Values.veleroApp.fullnameOverride }}
{{- .Values.veleroApp.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.veleroApp.fullnameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "SysMasterK8sVelero.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Velero Api labels / Selector labels
*/}}
{{- define "veleroApi.labels" -}}
helm.sh/chart: {{ include "SysMasterK8sVelero.chart" . }}
{{ include "veleroApi.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "veleroApi.selectorLabels" -}}
app.kubernetes.io/name: {{ include "SysMasterK8sVelero.veleroApi.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{/*
Velero App labels / Selector labels
*/}}
{{- define "veleroApp.labels" -}}
helm.sh/chart: {{ include "SysMasterK8sVelero.chart" . }}
{{ include "veleroApp.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "veleroApp.selectorLabels" -}}
app.kubernetes.io/name: {{ include "SysMasterK8sVelero.veleroApp.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

