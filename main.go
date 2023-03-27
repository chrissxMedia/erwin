package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/chrissxMedia/cm3.go"
)

func main() {
	cm3.ListenAndServeHttp(":8080", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, regexp.MustCompile(`:\d+$`).ReplaceAllString(r.RemoteAddr, ""))
	})
}
