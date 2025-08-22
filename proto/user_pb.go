package proto

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type UpdateUserRequest struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

type ListUsersRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type UserResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ListUsersResponse struct {
	Users []*UserResponse `json:"users"`
	Total int64           `json:"total"`
}

type Empty struct{}
