package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
