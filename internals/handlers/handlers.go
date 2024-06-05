package handlers

import (
	"net/http"

	"github.com/sanjay-xdr/goblogger/internals/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html")

}
