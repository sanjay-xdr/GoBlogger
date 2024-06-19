package handlers

import (
	"database/sql"
	"fmt"
	"log"
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

	render.RenderTemplate(w, "login.page.html", &models.TemplateData{})
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
	render.RenderTemplate(w, "signup.page.html", &models.TemplateData{})
}

func (m *Repositry) PostSignUp(w http.ResponseWriter, r *http.Request) {

	// signup the  user

	err := r.ParseForm()
	if err != nil {
		fmt.Print("Could not parse form")
	}

	fmt.Print("About to sign up")
	userid, err := m.DbConn.InsertIntoUser(models.User{
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

	m.App.Session.Put(r.Context(), "userid", userid)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// return the user details page
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	//update the userby id
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

func (m *Repositry) GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	//render all blogs

	blogs, err := m.DbConn.GetAllBlogs()
	if err != nil {
		log.Fatal("SOmethign wrong ")
	}
	// finalData := make(map[string]interface{})
	// finalData["blogs"] = data
	// fmt.Print(data)

	data := &models.TemplateData{
		Data: map[string]interface{}{
			"blogs": blogs,
		},
	}
	render.RenderTemplate(w, "blogs.page.html", data)

}

func (m *Repositry) GetBlogById(w http.ResponseWriter, r *http.Request) {
	//return the blog by id with comments

	//TODO: get the id from the paramenter and
	//find that blog and comments with it and display here
	userid, ok := m.App.Session.Get(r.Context(), "userid").(int)

	var mapdata = make(map[string]int)
	mapdata["userid"] = userid

	if !ok {
		log.Println("Can't get userid from the session")
	}
	render.RenderTemplate(w, "blog.page.html", &models.TemplateData{
		IntMap: mapdata,
	})
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "createblog.page.html", &models.TemplateData{})
}

func (m *Repositry) PostCreateBlog(w http.ResponseWriter, r *http.Request) {

	//Here Add the Blog to database
	//TODO: get userid in session

	err := r.ParseForm()

	if err != nil {
		log.Fatal("COuld not parse form something went wrong")
	}

	heading := r.Form.Get("heading")
	SubHeading := r.Form.Get("subheading")
	content := r.Form.Get("content")

	m.DbConn.InsertIntoBlogs(models.Blog{

		Heading:    heading,
		SubHeading: SubHeading,
		Content:    content,
		UserId:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	http.Redirect(w, r, "/blogs", http.StatusSeeOther)

}
