package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "rnt9hc:rnt9hc@tcp(localhost:3306)/ecomerce")
	if err != nil {
		log.Fatalf("Error when connecting to the database: %s\n", err)
	}

	result, err := db.Query("Select * from user")

	if err != nil {
		log.Printf("Data can not be fetched: %s\n", err)
	}

	for result.Next() {
		var (
			id           int
			name         string
			dob          string
			phone_number string
			email        string
			password     string
			active       int
			created_at   string
		)
		err := result.Scan(&id, &name, &dob, &phone_number, &email, &password, &active, &created_at)

		if err != nil {
			log.Printf("Record error: %s", err)
		}

		fmt.Fprintf(w, "Id: %d, name: %s, dob: %s, phone_number: %s, email: %s, password: %s, active: %d, created_at:%s\n", id, name, dob, phone_number, email, password, active, created_at)
	}
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
