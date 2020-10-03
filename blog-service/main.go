package main

import (
	"github.com/go-microservices/blog-service/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
