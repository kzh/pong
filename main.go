package main

import (
	"log"
	"net/http"
)

func main() {
	static := http.FileServer(http.Dir("static"))
	http.Handle("/", static)
	log.Fatal(http.ListenAndServe(":80", nil))
}
