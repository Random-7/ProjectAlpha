package main

import (
	"log"
	"net/http"

	"github.com/Random-7/ProjectAlpha/internal/router"
)

func main() {

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router.SetRoutes(),
	}

	log.Println("Starting listen")
	err := srv.ListenAndServe()
	log.Fatal(err)

}
