package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chrissxMedia/cm3.go"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	var reqs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "erwin_requests",
		Help: "Number of HTTP requests by IP.",
	}, []string{"ip"})
	cm3.HandleMetrics(reqs)
	cm3.ListenAndServeHttp(":8080", func(w http.ResponseWriter, r *http.Request) {
		ip := cm3.RemoteIp(r)
		// TODO: break cm3.RemoteIp
		ip = strings.ReplaceAll(strings.ReplaceAll(ip, "[", ""), "]", "")
		reqs.WithLabelValues(ip).Inc()
		fmt.Fprintln(w, ip)
	})
}
