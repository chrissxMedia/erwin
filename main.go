package main

import (
	"fmt"
	"net/http"

	"github.com/chrissxMedia/cm3.go"
)

func main() {
	// TODO: prometheus
	cm3.ListenAndServeHttp(":8080", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, cm3.RemoteIp(r))
	})
}
