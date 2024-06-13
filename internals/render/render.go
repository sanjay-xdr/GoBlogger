package render

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var pathToTemplates = "./templates"

func RenderTemplate(w http.ResponseWriter, html string, data string) {

	var tc map[string]*template.Template

	//check if template already exists in cache or  not TODO:
	tc, _ = CreateTemplateCache()

	t, ok := tc[html]

	if !ok {
		log.Print("Something went wrong")
	}

	// buf := new(bytes.Buffer)
	err := t.Execute(w, data)

	// _, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Something Went Wrong")
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

		tc[name] = ts

	}

	return tc, nil

}
