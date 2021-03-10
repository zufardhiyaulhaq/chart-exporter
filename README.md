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

### Example Metrics
```
chart_metrics{api_version="apps/v1",chart_name="chart-exporter",chart_version="1.0.1",deployment_name="chart-exporter",namespace="infrastructure"} 1
chart_metrics{api_version="apps/v1",chart_name="grafana",chart_version="5.8.16",deployment_name="prometheus-grafana",namespace="infrastructure"} 1
chart_metrics{api_version="apps/v1",chart_name="kiali-operator",chart_version="1.28.0",deployment_name="kiali-operator",namespace="infrastructure"} 1
chart_metrics{api_version="apps/v1",chart_name="kube-state-metrics",chart_version="2.9.4",deployment_name="prometheus-kube-state-metrics",namespace="infrastructure"} 1
chart_metrics{api_version="apps/v1",chart_name="kubernetes-dashboard",chart_version="3.0.2",deployment_name="kubernetes-dashboard",namespace="infrastructure"} 1
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| image.name | string | `"zufardhiyaulhaq/chart-exporter"` |  |
| image.tag | string | `"2.0.0"` |  |
| pullPolicy | string | `"Always"` |  |
| serviceMonitor.enabled | bool | `true` |  |
