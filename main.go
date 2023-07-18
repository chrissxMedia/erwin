package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/chrissxMedia/cm3.go"
)

func main() {
	// TODO: prometheus
	cm3.ListenAndServeHttp(":8080", func(w http.ResponseWriter, r *http.Request) {
        if realIp, hasRealIp := r.Header["X-Real-Ip"]; hasRealIp && len(realIp) == 1 {
            fmt.Fprintln(w, realIp[0])
        } else {
            fmt.Fprintln(w, regexp.MustCompile(`:\d+$`).ReplaceAllString(r.RemoteAddr, ""))
        }
    })
}
