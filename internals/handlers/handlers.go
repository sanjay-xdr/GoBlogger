package handlers

import (
	"net/http"

	"github.com/sanjay-xdr/goblogger/internals/render"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//return the login page
}

func PostLogin(w http.ResponseWriter, r *http.Request) {

	// login user
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	//return the signup page
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {

	// signup the  user

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// return the user details page
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	//update the userby id
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html")

}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	//render all blogs

}

func GetBlogById(w http.ResponseWriter, r *http.Request) {
	//return the blog by id with comments
}
