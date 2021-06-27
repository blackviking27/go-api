package api

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Defining the user struct
type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Payload  *Payload `json:"payload"`
}

// Defining payload
type Payload struct {
	Lang string `json:"lang"`
	Str  string `json:"str"`
}

// Get a list of users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get a single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Getting all the params in the url
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// if we dont find the book
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	// decoding the json data as the User struct
	_ = json.NewDecoder(r.Body).Decode(&user)

	// convert the username to hex string and convert it to string type
	user.ID = string(hex.EncodeToString([]byte(user.Username)))
	users = append(users, user) // append to users
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			// remove the current item
			users = append(users[:index], users[index+1:]...)

			//making a new item with updated details
			var user User
			// decoding the json data as the User struct
			_ = json.NewDecoder(r.Body).Decode(&user)
			// convert the username to hex string and convert it to string type
			user.ID = string(hex.EncodeToString([]byte(user.Username)))
			users = append(users, user) // append to users
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			// remove the current item
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

//creating the users slice
var users []User

// functions which are suppose to be exported need to start with a capital letter
func StartApi() {
	// Initialise router
	router := mux.NewRouter()

	// adding the mock data
	// You can skip this part if you want
	users = append(users, User{ID: "1", Name: "John Wick", Username: "john", Payload: &Payload{Lang: "python", Str: "python"}})
	users = append(users, User{ID: "2", Name: "Juhn Weak", Username: "juhn", Payload: &Payload{Lang: "c", Str: "c"}})

	// Defining routes
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// starting the server
	// log.Fatal is just gonna show us the error if there is one
	log.Fatal(http.ListenAndServe(":8000", router))
}
