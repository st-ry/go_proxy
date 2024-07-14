package proxy

import (
	"fmt"
	"net/http"
)

func serveProxy() {
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}