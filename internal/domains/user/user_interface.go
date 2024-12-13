package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/domains/user/dto"
	"go-clean-architecture/internal/domains/user/entity"
)

type IUserHTTP interface {
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindUser(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type IUserUseCase interface {
	CreateUser(ctx context.Context, createUserRequest *dto.UpsertUserRequest) error
	GetUser(ctx context.Context, id string) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, updateUserRequest *dto.UpsertUserRequest) error
	GetAllUsers(ctx context.Context) ([]dto.UserResponse, error)
}
