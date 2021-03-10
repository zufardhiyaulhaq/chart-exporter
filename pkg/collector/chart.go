package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/client"
)

const DefaultMetricsValue float64 = 1

type ChartCollector struct {
	chartMetric *prometheus.Desc
	client      client.KubernetesClient
}

func (collector *ChartCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.chartMetric
}

func (collector *ChartCollector) Collect(channel chan<- prometheus.Metric) {
	deployments := collector.client.GetDeployments()

	for _, deployment := range deployments {
		channel <- prometheus.MustNewConstMetric(collector.chartMetric, prometheus.CounterValue, DefaultMetricsValue, deployment.Namespace, deployment.Name, deployment.ChartName, deployment.ChartVersion, deployment.APIVersion)
	}
}

func NewChartCollector(client client.KubernetesClient) *ChartCollector {
	return &ChartCollector{
		chartMetric: prometheus.NewDesc("chart_metrics",
			"metrics about charts",
			[]string{"namespace", "deployment_name", "chart_name", "chart_version", "api_version"}, nil,
		),
		client: client,
	}
}
