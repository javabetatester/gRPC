package handlers

import (
	"grpc-app/controller"
	"grpc-app/service"
	"grpc-app/repository"
	"database/sql"
)

type GRPCHandler struct {
	UserController *controller.UserController
}

func NewGRPCHandler(db *sql.DB) *GRPCHandler {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &GRPCHandler{
		UserController: userController,
	}
}
