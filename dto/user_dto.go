package dto

import "time"

type UserDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserDTO struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserListDTO struct {
	Users []UserDTO `json:"users"`
	Total int       `json:"total"`
}
