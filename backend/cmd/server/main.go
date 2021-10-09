package main

import (
	"net/http"
	"time"

	"github.com/krastomer/treeshop-cpe327/backend/pkg/http/rest"
)

func main() {
	handler := rest.NewHandler()

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()

	// db := mariadb.NewRepository()
	// s := db.GetAll()
	// fmt.Println(s)

}
