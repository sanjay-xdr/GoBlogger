package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjay-xdr/goblogger/internals/config"
	"github.com/sanjay-xdr/goblogger/internals/render"
)

var configurations *config.AppConfig

func main() {
	fmt.Print("Setting up project structure")

	render.CreateTemplateCache()

	// fmt.Print(logger)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
