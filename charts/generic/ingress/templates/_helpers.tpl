{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 24 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- $name := default .Values.entity.name -}}
{{- if $name }}
{{- printf "%s-%s" .Chart.Name $name | trunc 63 | trimSuffix "-" -}}
{{- else }}
{{- printf "%s" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end }}
{{- end -}}
