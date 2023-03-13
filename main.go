package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s %s %s from %s (%s) to %s",
			r.Method, r.URL.Path, r.Proto, r.RemoteAddr, r.UserAgent(), r.Host))
		fmt.Fprintln(w, regexp.MustCompile(`:\d+$`).ReplaceAllString(r.RemoteAddr, ""))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
