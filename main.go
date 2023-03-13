package main

import (
	"github.com/tencypress/mygogin/api"
	"log"
	"net/http"
)

func main() {
	router := api.InitRouter()

	server := &http.Server{
		Addr:    "localhost:80",
		Handler: router,
	}

	log.Printf("[info] start http server listening %s", "localhost:80")
	server.ListenAndServe()
}
