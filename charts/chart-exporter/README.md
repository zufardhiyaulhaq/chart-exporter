# chart-exporter

![Version: 1.0.1](https://img.shields.io/badge/Version-1.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.0.0](https://img.shields.io/badge/AppVersion-2.0.0-informational?style=flat-square)

chart-exporter, export helm information to prometheus

**Homepage:** <https://github.com/zufardhiyaulhaq/chart-exporter>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| image.name | string | `"zufardhiyaulhaq/chart-exporter"` |  |
| image.tag | string | `"2.0.0"` |  |
| pullPolicy | string | `"Always"` |  |
| serviceMonitor.enabled | bool | `true` |  |
