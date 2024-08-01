package main

import (
	"log"
	"net/http"

	"github.com/tcnam/golangbe/pkg/handlers"
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

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	log.Fatal(http.ListenAndServe(server, nil))
}
