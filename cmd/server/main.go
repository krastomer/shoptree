package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	r := http.NewServeMux()

	buildHandler := http.FileServer(http.Dir("web/build"))
	r.Handle("/", buildHandler)

	s := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
