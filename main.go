package main

import (
	"net/http"
	"fmt"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":80", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello Example")
	
}
