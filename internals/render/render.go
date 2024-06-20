package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/sanjay-xdr/goblogger/internals/models"
)

var pathToTemplates = "./templates"

func RenderTemplate(w http.ResponseWriter, html string, data *models.TemplateData) {

	// fmt.Print(data)
	var tc map[string]*template.Template

	//check if template already exists in cache or  not TODO:
	tc, err := CreateTemplateCache()

	if err != nil {
		log.Fatal("Pta nhi kuch toh krb hai ", err)
	}

	t, ok := tc[html]

	if !ok {
		log.Println("mapping se")
		log.Print("Something went wrong")
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		// log.Fatal("I am here")
		log.Fatal("Something Went Wrong1", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("I am here  2")
		log.Fatal("Something Went Wrong", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	var tc = map[string]*template.Template{}

	tmpl, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates)) //this returns template/about.page.html

	if err != nil {
		return tc, err
	}

	for _, page := range tmpl {

		name := filepath.Base(page)
		fmt.Print(name)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return tc, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates)) //find the layout
		if err != nil {
			return tc, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates)) //attach the layout result to existing template
			if err != nil {
				return tc, err
			}
		}

		tc[name] = ts

	}

	return tc, nil

}
