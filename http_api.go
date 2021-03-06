package main

import (
	"encoding/json"
	"strconv"

	//"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User is a struct that represents a User in our applciation
type User struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	//fmt.Println("HTTP API Go Project starts here!!")

	//setup router
	router := mux.NewRouter()

	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	router.HandleFunc("/posts/{id}", patchpost).Methods("PATCH")

	router.HandleFunc("/posts/{id}", deletepost).Methods("DELETE")

	//start webserver
	http.ListenAndServe(":5000", router)

}

func addItem(w http.ResponseWriter, r *http.Request) { // this is a route-hanler

	//get item value from JSON Body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to INT"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to INT"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//get the value from the json body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func patchpost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to INT"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//get the current value
	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func deletepost(w http.ResponseWriter, r *http.Request) {

	// get the ID of the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		//there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to INT"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//delete the post from slice --> this is delete slice syntax
	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)
}
