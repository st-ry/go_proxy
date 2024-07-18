package proxy

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func ServeProxy() {
	http.HandleFunc("/", rootHandler)
	if err := http.ListenAndServe(":18888", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	distPort := 28888
	requestURL := fmt.Sprintf("http://localhost:%d", distPort)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dist response code: %s\n", string(body))
	proxyRes := "hello, " + string(body)
	fmt.Fprint(w, proxyRes)
}