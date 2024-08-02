package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tcnam/golangbe/pkg/config"
	"github.com/tcnam/golangbe/pkg/handlers"
	"github.com/tcnam/golangbe/pkg/render"
)

// func divide(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		x float32 = 100.00
// 		y float32 = 2.00
// 	)
// 	var f, err = divideValues(x, y)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}
// 	fmt.Fprintf(w, "%f divided by %f is %f", x, y, f)
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		var err error = errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	var result float32 = x / y
// 	return result, nil
// }

func main() {
	const (
		server string = "localhost:8080"
	)

	var app config.AppConfig
	tc, err := render.InitTemplateCache()
	if err != nil {
		log.Printf("Couldn't create template cache")
	}
	app.TemplateCache = tc
	render.NewConfig(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	fmt.Printf("Starting the application at %s\n", server)
	http.HandleFunc("/", repo.HomeHandler)
	http.HandleFunc("/about", repo.AboutHandler)

	log.Fatal(http.ListenAndServe(server, nil))
}
