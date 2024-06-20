package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanjay-xdr/goblogger/internals/handlers"
)

func Routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(loggerfunc)
	mux.Use(SessionLoad)
	mux.Use(Auth)
	mux.Get("/", handlers.Repo.HomePage)

	//auth
	mux.Get("/login", handlers.Repo.Login)
	mux.Post("/login", handlers.PostLogin)
	mux.Get("/signup", handlers.SignUp)
	mux.Post("/signup", handlers.Repo.PostSignUp)

	//users
	mux.Get("/user/{id}", handlers.GetUserById)
	mux.Post("/user/{id}", handlers.UpdateUserById)

	//blogs
	mux.Get("/blogs", handlers.Repo.GetAllBlogs)
	mux.Get("/blog/{id}", handlers.Repo.GetBlogById)
	mux.Get("/user/{id}/createblog", handlers.CreateBlog)
	mux.Post("/user/{id}/createblog", handlers.Repo.PostCreateBlog)
	mux.Post("/comment", handlers.Repo.CreateComment)

	return mux
}
