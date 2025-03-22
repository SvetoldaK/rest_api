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

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task Task
	if result := DB.First(&task, id); result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updateData RequestBody
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateData.Task != "" {
		task.Task = updateData.Task
	}
	task.IsDone = updateData.IsDone

	if result := DB.Save(&task); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	result := DB.Delete(&Task{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Task не найдена", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
