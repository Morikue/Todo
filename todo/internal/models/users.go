package models

import "todo/pkg/grpc_stubs/users"

type UserDTO struct {
	ID       int    `json:"id,omitempty" example:"1"`
	Username string `json:"username" example:"username"`
	Password string `json:"password,omitempty" example:"password"`
	Email    string `json:"email,omitempty" example:"user@example.com"`
}

func NewEmptyUserDTO() *UserDTO {
	return &UserDTO{}
}

func (d *UserDTO) FromGRPC(in *users.UserDTO) *UserDTO {
	d.ID = int(in.Id)
	d.Email = in.Email
	d.Username = in.Username
	d.Password = in.Password
	return d
}
