package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		fmt.Fprintln(w, regexp.MustCompile(`:\d+$`).ReplaceAllString(r.RemoteAddr, ""))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
