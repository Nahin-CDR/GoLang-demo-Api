package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	//http.HandleFunc("/api/data", CommentList)

	// Define API routes
	router.HandleFunc("/api/data", CommentList).Methods("GET")
	router.HandleFunc("/", HelloWorld).Methods("GET")
	router.HandleFunc("/greet/{name}", Greet).Methods("GET")

	// Start the web server
	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
	///> available base link : http://localhost:8080
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

func CommentList(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"postID": 1,
			"id":     1,
			"name":   "Nahin",
			"email":  "nahin.cdr@gmail.com",
			"body":   "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English",
		},
		{
			"postId": 1,
			"id":     5,
			"name":   "vero eaque aliquid doloribus et culpa",
			"email":  "Hayden@althea.biz",
			"body":   "harum non quasi et ratione\ntempore iure ex voluptates in ratione\nharum architecto fugit inventore cupiditate\nvoluptates magni quo et",
		},
		{
			"postId": 1,
			"id":     2,
			"name":   "quo vero reiciendis velit similique earum",
			"email":  "Jayne_Kuhic@sydney.com",
			"body":   "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et",
		},
		{
			"postId": 1,
			"id":     3,
			"name":   "odio adipisci rerum aut animi",
			"email":  "Nikita@garfield.biz",
			"body":   "quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
