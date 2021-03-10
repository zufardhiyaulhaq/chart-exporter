# chart-exporter

![Version: 1.0.1](https://img.shields.io/badge/Version-1.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.0.0](https://img.shields.io/badge/AppVersion-2.0.0-informational?style=flat-square)

chart-exporter, export helm information like chart name and chart version to prometheus. All information is retrieve from [Helm best practice](https://helm.sh/docs/chart_best_practices/labels/) for label and annotations.

## Installation

### Helm
Please read README.md in charts folder for more information.
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/chart-exporter --name-template chart-exporter
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| image.name | string | `"zufardhiyaulhaq/chart-exporter"` |  |
| image.tag | string | `"2.0.0"` |  |
| pullPolicy | string | `"Always"` |  |
| serviceMonitor.enabled | bool | `true` |  |
