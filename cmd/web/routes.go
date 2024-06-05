package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanjay-xdr/goblogger/internals/handlers"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(loggerfun)
	mux.Get("/", handlers.HomePage)

	return mux
}
