package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/tasksService"
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/tasks"
	"awesomeProject/internal/web/users"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&tasksService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	tasksRepo := tasksService.NewTaskRepository(database.DB)
	TasksService := tasksService.NewService(tasksRepo)
	TasksHandler := handlers.NewTaskHandler(TasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	UsersService := userService.NewUserService(usersRepo)
	UsersHandler := handlers.NewUserHandler(UsersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(TasksHandler, nil)
	tasks.RegisterHandlersWithBaseURL(e, strictTasksHandler, "/api")

	strictUsersHandler := users.NewStrictHandler(UsersHandler, nil)
	users.RegisterHandlersWithBaseURL(e, strictUsersHandler, "/api")

	// Выводим все зарегистрированные маршруты
	for _, route := range e.Routes() {
		fmt.Printf("Method: %v, Path: %v\n", route.Method, route.Path)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
