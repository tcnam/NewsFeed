package handlers

import (
	"net/http"

	"github.com/tcnam/golangbe/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
