package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//Setup Router
func Router() *mux.Router {
	router := mux.NewRouter()
	//Mock endpoint that matches RootEndpoint Func
	router.HandleFunc("/create", RootEndpoint).Methods("GET")
	return router
}

func TestCreateEndpoint(t *testing.T) {
	//Create a request
	request, _ := http.NewRequest("GET", "/create", nil)
	//Create a response
	response := httptest.NewRecorder()
	//Passing request and recording response !
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Expected a response !")
	assert.Equal(t, "Hello Motto", response.Body.String(), "Incorrect Body Found !")
}
