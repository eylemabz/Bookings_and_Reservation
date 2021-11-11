package main

import (
	"fmt"
	"net/http"

	"github.com/eylemabz/go-course/pkg/handlers"
)

var portNumber = ":1313"

//main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
