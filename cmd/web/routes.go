package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanjay-xdr/goblogger/internals/handlers"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(loggerfunc)
	mux.Get("/", handlers.HomePage)

	//auth
	mux.Get("/login", handlers.Login)
	mux.Post("/login", handlers.PostLogin)
	mux.Get("/signup", handlers.SignUp)
	mux.Post("/signup", handlers.PostSignUp)

	//users
	mux.Get("/user/{id}", handlers.GetUserById)
	mux.Post("/user/{id}", handlers.UpdateUserById)

	//blogs
	mux.Get("/blog", handlers.GetAllBlogs)
	mux.Get("/blog/{id}", handlers.GetBlogById)

	return mux
}
