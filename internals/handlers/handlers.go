package handlers

import (
	"fmt"
	"net/http"

	"github.com/sanjay-xdr/goblogger/internals/render"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//return the login page
	render.RenderTemplate(w, "login.page.html", "")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {

	// login user

	err := r.ParseForm()
	if err != nil {
		fmt.Print("Could not parse form");
	}

	r.Form.Get("")
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	//return the signup page
	render.RenderTemplate(w, "signup.page.html", "")
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {

	// signup the  user

	
	err := r.ParseForm()
	if err != nil {
		fmt.Print("Could not parse form");
	}

	r.Form.Get("")

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// return the user details page
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	//update the userby id
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html", "")

}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	//render all blogs
	render.RenderTemplate(w, "blogs.page.html", "")

}

func GetBlogById(w http.ResponseWriter, r *http.Request) {
	//return the blog by id with comments
}
