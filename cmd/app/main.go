package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/tasksService"
	"awesomeProject/internal/web/tasks"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&tasksService.Task{})

	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlersWithBaseURL(e, strictHandler, "/api")

	// Выводим все зарегистрированные маршруты
	for _, route := range e.Routes() {
		fmt.Printf("Method: %v, Path: %v\n", route.Method, route.Path)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
