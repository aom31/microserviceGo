package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("hello server")
	})
	http.ListenAndServe(":9090", nil)
}
