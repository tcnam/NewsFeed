package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tcnam/golangbe/pkg/render"
)

func WriteJson(w http.ResponseWriter, statuts int, v any) error {
	w.WriteHeader(statuts)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// default http.HandlerFunc don't have error so we need customize version
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listtenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listtenAddr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	log.Printf("JSON API server running on port: %s\n", s.listenAddr)

	router.HandleFunc("/home", makeHTTPHandleFunc(s.handleHome))
	router.HandleFunc("/about", makeHTTPHandleFunc(s.handleAbout))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("method not allowed: %s", r.Method)
	}
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) error {
	render.RenderTemplate(w, "home.html")
	return nil
}

func (s *APIServer) handleAbout(w http.ResponseWriter, r *http.Request) error {
	render.RenderTemplate(w, "about.html")
	return nil
}
