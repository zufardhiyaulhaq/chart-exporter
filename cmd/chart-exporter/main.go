package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/client"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/collector"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/middleware"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/settings"
)

func main() {
	log.Println("Starting chart-exporter")
	settings := settings.NewSettings()

	client := client.KubernetesClient{}
	client.Start(settings)

	chartCollector := collector.NewChartCollector(client)
	prometheus.MustRegister(chartCollector)

	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/healthz", middleware.StatusHandler(client))
	router.Handle("/readyz", middleware.StatusHandler(client))

	http.ListenAndServe(":9125", router)
}
