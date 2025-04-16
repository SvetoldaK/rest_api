package handlers

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/tasksService"
	"awesomeProject/internal/web/tasks"
	"context"
	"fmt"
)

type TaskHandler struct {
	Service *tasksService.TaskService
}

func NewTaskHandler(service *tasksService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	if request.Params.UserId != nil {
		userTasks, err := h.Service.GetTasksByUserID(uint(*request.Params.UserId))
		if err != nil {
			return nil, err
		}
		var response tasks.GetTasks200JSONResponse
		for _, tsk := range userTasks {
			id := tsk.ID
			response = append(response, tasks.Task{
				Id:     &id,
				Task:   tsk.Task,
				IsDone: tsk.IsDone,
				UserId: tsk.UserID,
			})
		}
		return response, nil
	}

	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasks200JSONResponse
	for _, tsk := range allTasks {
		id := tsk.ID
		response = append(response, tasks.Task{
			Id:     &id,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
			UserId: tsk.UserID,
		})
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	taskToCreate := models.Task{
		Task:   request.Body.Task,
		IsDone: request.Body.IsDone,
		UserID: request.Body.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	id := createdTask.ID
	response := tasks.PostTasks201JSONResponse{
		Id:     &id,
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
		UserId: createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := uint(request.Id)
	if err := h.Service.DeleteTaskByID(taskID); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)
	if request.Body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	taskToUpdate := models.Task{
		Task:   request.Body.Task,
		IsDone: request.Body.IsDone,
		UserID: request.Body.UserId,
	}
	updatedTask, err := h.Service.UpdateTaskByID(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	id := updatedTask.ID
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &id,
		Task:   updatedTask.Task,
		IsDone: updatedTask.IsDone,
		UserId: updatedTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userTasks, err := h.Service.GetTasksByUserID(request.Id)
	if err != nil {
		return nil, err
	}
	var response tasks.GetUsersIdTasks200JSONResponse
	for _, tsk := range userTasks {
		id := tsk.ID
		response = append(response, tasks.Task{
			Id:     &id,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
			UserId: tsk.UserID,
		})
	}
	return response, nil
}
