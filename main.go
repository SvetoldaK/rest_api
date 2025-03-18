package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type RequestBody struct {
	Task string `json:"task"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, %s", task)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	task = requestBody.Task
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
