package controller

import (
	"context"
	"grpc-app/dto"
	"grpc-app/proto"
	"grpc-app/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	userDTO, err := c.userService.CreateUser(ctx, &dto.CreateUserDTO{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:    userDTO.ID,
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}, nil
}

func (c *UserController) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	userDTO, err := c.userService.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:    userDTO.ID,
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}, nil
}

func (c *UserController) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UserResponse, error) {
	userDTO, err := c.userService.UpdateUser(ctx, req.Id, &dto.UpdateUserDTO{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:    userDTO.ID,
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}, nil
}

func (c *UserController) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.Empty, error) {
	err := c.userService.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (c *UserController) ListUsers(ctx context.Context, req *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	userListDTO, err := c.userService.ListUsers(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	users := make([]*proto.UserResponse, len(userListDTO.Users))
	for i, user := range userListDTO.Users {
		users[i] = &proto.UserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return &proto.ListUsersResponse{
		Users: users,
		Total: int64(userListDTO.Total),
	}, nil
}
