package handlers

import (
	"awesomeProject/internal/userService"
	"awesomeProject/internal/web/users"
	"context"
	"fmt"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:       &user.ID,
			Password: &user.Password,
			Username: &user.Username,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil || request.Body.Password == nil {
		return nil, fmt.Errorf("password is required")
	}

	userToCreate := userService.User{
		Password: *request.Body.Password,
		Username: *request.Body.Username,
	}

	createdUser, err := h.Service.PostUser(userToCreate)
	if err != nil {
		return nil, err
	}

	return users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Password: &createdUser.Password,
		Username: &createdUser.Username,
	}, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := uint(request.Id)
	if err := h.Service.DeleteUserByID(userID); err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := uint(request.Id)
	userToUpdate := userService.User{
		Password: *request.Body.Password,
		Username: *request.Body.Username,
	}

	updatedUser, err := h.Service.PatchUserByID(userID, userToUpdate)
	if err != nil {
		return nil, err
	}

	return users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Password: &updatedUser.Password,
		Username: &updatedUser.Username,
	}, nil
}
