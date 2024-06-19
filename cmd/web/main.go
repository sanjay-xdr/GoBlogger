package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanjay-xdr/goblogger/internals/config"
	"github.com/sanjay-xdr/goblogger/internals/driver"
	"github.com/sanjay-xdr/goblogger/internals/handlers"
	"github.com/sanjay-xdr/goblogger/internals/render"
)

var configurations config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false //if true wont work with https site
	configurations.Session = sessionManager
	fmt.Println("Setting up project structure")

	render.CreateTemplateCache()

	db, err := driver.ConnectDb("host=localhost port=5432 dbname=GoBlogger user=postgres password=sanjay")

	if err != nil {
		log.Fatal("Not able to connect to db ")
	}

	repo := handlers.NewRepo(&configurations, db)
	handlers.NewHandlers(repo)

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
