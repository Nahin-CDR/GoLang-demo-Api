package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/", HelloWorld).Methods("GET")
	router.HandleFunc("/greet/{name}", Greet).Methods("GET")

	// Start the web server
	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// HelloWorld handles the root route
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,Nahin the Coder")
}

// Greet handles the /greet/{name} route
func Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}
