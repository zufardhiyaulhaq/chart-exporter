module github.com/zufardhiyaulhaq/chart-exporter

go 1.15

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/norwoodj/helm-docs v1.5.0 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/sirupsen/logrus v1.8.1
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)

replace k8s.io/client-go => k8s.io/client-go v0.20.4
