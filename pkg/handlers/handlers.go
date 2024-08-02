package handlers

import (
	"net/http"

	"github.com/tcnam/golangbe/pkg/config"
	"github.com/tcnam/golangbe/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
