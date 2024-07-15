package proxy

import (
	"fmt"
	"log"
	"net/http"
)

func ServeProxy() {
	http.HandleFunc("/", rootHandler)
	if err := http.ListenAndServe(":18888", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}