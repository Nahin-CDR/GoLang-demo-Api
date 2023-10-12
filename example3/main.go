package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// / define an Employee Model
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

func EmployeeProvider(w http.ResponseWriter, r *http.Request) {

	employee := []Employee{
		{
			ID:       1,
			Name:     "John Doe",
			Position: "Software Engineer",
		},
		{
			ID:       2,
			Name:     "Alice Smith",
			Position: "Product Manager",
		},
		{
			ID:       3,
			Name:     "Bob Johnson",
			Position: "QA Tester",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	/// Define a handler function to serve Employee details
	http.HandleFunc("/employee", EmployeeProvider)

	fmt.Println("Server is running on : 8080")
	/// start the web Server
	http.ListenAndServe(":8080", nil)

}
