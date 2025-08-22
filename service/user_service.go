package service

import (
	"context"
	"grpc-app/dto"
	"grpc-app/model"
	"grpc-app/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserDTO) (*dto.UserDTO, error)
	GetUser(ctx context.Context, id int64) (*dto.UserDTO, error)
	UpdateUser(ctx context.Context, id int64, req *dto.UpdateUserDTO) (*dto.UserDTO, error)
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, limit, offset int) (*dto.UserListDTO, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, req *dto.CreateUserDTO) (*dto.UserDTO, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*dto.UserDTO, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int64, req *dto.UpdateUserDTO) (*dto.UserDTO, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	updatedUser, err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context, limit, offset int) (*dto.UserListDTO, error) {
	users, err := s.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	return &dto.UserListDTO{
		Users: userDTOs,
		Total: len(userDTOs),
	}, nil
}
