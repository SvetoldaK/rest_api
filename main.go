package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := Task{
		Task:   requestBody.Task,
		IsDone: requestBody.IsDone,
	}

	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
