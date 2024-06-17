package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/sanjay-xdr/goblogger/internals/Db"
	"github.com/sanjay-xdr/goblogger/internals/config"
	"github.com/sanjay-xdr/goblogger/internals/models"
	"github.com/sanjay-xdr/goblogger/internals/render"
)

type Repositry struct {
	App    *config.AppConfig
	DbConn *Db.PostgresDBRepo
}

var Repo *Repositry

// Creates a new Repo
func NewRepo(a *config.AppConfig, db *sql.DB) *Repositry {
	return &Repositry{
		App:    a,
		DbConn: Db.NewPostgresRepo(db, a),
	}
}

// set the Above Repo Variable
func NewHandlers(r *Repositry) {
	Repo = r
}

func (m *Repositry) Login(w http.ResponseWriter, r *http.Request) {

	//return the login page
	render.RenderTemplate(w, "login.page.html", "")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {

	// login user

	err := r.ParseForm()
	if err != nil {
		fmt.Print("Could not parse form")
	}

	r.Form.Get("")
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	//return the signup page
	render.RenderTemplate(w, "signup.page.html", "")
}

func (m *Repositry) PostSignUp(w http.ResponseWriter, r *http.Request) {

	// signup the  user

	err := r.ParseForm()
	if err != nil {
		fmt.Print("Could not parse form")
	}

	fmt.Print("About to sign up")
	err = m.DbConn.InsertIntoUser(models.User{
		FirstName: r.Form.Get("firstName"),
		LastName:  r.Form.Get("lastName"),
		Email:     r.Form.Get("email"),
		Password:  r.Form.Get("password"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		fmt.Print("NOt able to insert into the DB", err)
	}

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
