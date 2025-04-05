package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	taskService2 "awesomeProject/internal/taskService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&taskService2.Task{})

	repo := taskService2.NewTaskRepository(database.DB)
	service := taskService2.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
