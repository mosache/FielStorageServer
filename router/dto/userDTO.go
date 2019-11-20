package dto

import "FileStorageServer/model"

//UserDTO UserDTO
type UserDTO struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

//NewUserDTO NewUserDTO
func NewUserDTO(user *model.User) *UserDTO {
	var intance = new(UserDTO)
	intance.UserID = user.ID
	intance.UserName = user.UserName
	return intance
}
