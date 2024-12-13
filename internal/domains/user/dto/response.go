package dto

import "go-clean-architecture/common"

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SingleUserResponse = common.BaseResponse[UserResponse]

type ListUserResponse = common.BaseResponse[[]UserResponse]
