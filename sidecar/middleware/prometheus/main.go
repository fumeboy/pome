package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func init() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		addr := fmt.Sprintf("0.0.0.0:%d", 8888)
		http.ListenAndServe(addr, nil)
	}()
}
