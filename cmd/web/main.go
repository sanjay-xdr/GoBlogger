package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjay-xdr/goblogger/internals/config"
	"github.com/sanjay-xdr/goblogger/internals/driver"
	"github.com/sanjay-xdr/goblogger/internals/render"
)

var configurations *config.AppConfig

func main() {
	fmt.Println("Setting up project structure")

	render.CreateTemplateCache()

	db, err := driver.ConnectDb("host=localhost port=5432 dbname=GoBlogger user=postgres password=sanjay")

	if err != nil {
		log.Fatal("Not able to connect to db ")
	}

	defer db.Close()

	// fmt.Print(logger)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
