package middleware

import (
	"net/http"

	"github.com/zufardhiyaulhaq/chart-exporter/pkg/client"
)

func StatusHandler(client client.KubernetesClient) http.Handler {
	ok, err := client.GetStatus()
	if ok {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			return
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	})
}
