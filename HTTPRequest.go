package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// Resource struct
type Resource struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello Motto"))
}

//Init books var as a slice Resource struct
var resources []Resource

//Get all resources
func getResources(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resources)
}

func getResource(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	//Loop through resources and find with id
	for _, item := range resources {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Resource{})
}

func createResource(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var resource Resource
	//request
	_ = json.NewDecoder(r.Body).Decode(&resource)
	resource.ID = strconv.Itoa(rand.Intn(1000000))
	resources = append(resources, resource)
	json.NewEncoder(writer).Encode(resource)
}

func updateResources(writer http.ResponseWriter, r *http.Request) {

}
func DeleteResources(writer http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Router
	router := mux.NewRouter()

	resources = append(resources, Resource{ID: "1", Isbn: "448734", Title: "Resource", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	resources = append(resources, Resource{ID: "2", Isbn: "448732", Title: "Resource", Author: &Author{Firstname: "Chris", Lastname: "Rock"}})

	// Routes consist of a path and a handler function.
	//router.HandleFunc("/create/resources", RootEndpoint).Methods("GET")
	router.HandleFunc("/create", getResources).Methods("GET")

	router.HandleFunc("/create/{id}", getResource).Methods("GET")

	router.HandleFunc("/create", createResource).Methods("POST")

	router.HandleFunc("/create/{id}", updateResources).Methods("PUT")

	router.HandleFunc("/create/{id}", DeleteResources).Methods("Delete")

	// Bind to a port and pass our router in, if it fail we'll throw an error
	//Start servers
	log.Fatal(http.ListenAndServe(":12349", router))
}
