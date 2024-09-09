package main

import (
	"net/http"
)

func main() {

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "oi todo mundo")
		http.ServeFile(w, r, "index.html")
	}

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8083", nil)
}
