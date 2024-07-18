package dist

import (
	"fmt"
	"log"
	"net/http"
)

func ServeDist(){
	http.HandleFunc("/", rootHandler)
	if err := http.ListenAndServe(":28888", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "dist")
}