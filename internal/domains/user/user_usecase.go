package user

import (
	"context"
	"errors"
	"go-clean-architecture/internal/domains/user/dto"
	"go-clean-architecture/internal/domains/user/entity"
	"go-clean-architecture/pkg"
	"gorm.io/gorm"
	"net/http"
)

type userUseCase struct {
	userRepository IUserRepository
}

func (u *userUseCase) CreateUser(ctx context.Context, createUserRequest *dto.UpsertUserRequest) error {

	hashedPassword, err := pkg.HashPassword(createUserRequest.Password)
	if err != nil {
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	newUser := &entity.User{
		Email:    createUserRequest.Email,
		Name:     createUserRequest.Name,
		Username: createUserRequest.Username,
		Password: hashedPassword,
	}

	err = u.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &pkg.CustomError{
				Code:    http.StatusConflict,
				Message: "User with that email already exists!",
			}
		}

		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (u *userUseCase) GetUser(ctx context.Context, id string) (*dto.UserResponse, error) {
	user, err := u.userRepository.FindUser(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &pkg.CustomError{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return nil, &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if user == nil {
		return nil, &pkg.CustomError{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil

}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	_, err := u.userRepository.FindUser(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pkg.CustomError{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = u.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, id string, updateUserRequest *dto.UpsertUserRequest) error {
	user, err := u.userRepository.FindUser(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pkg.CustomError{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	user.Name = updateUserRequest.Name
	hashedPassword, err := pkg.HashPassword(updateUserRequest.Password)
	if err != nil {
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.Password = hashedPassword
	user.Email = updateUserRequest.Email
	user.Name = updateUserRequest.Name
	user.Username = updateUserRequest.Username

	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := u.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, &pkg.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	var usersList = make([]dto.UserResponse, len(users)-1)

	for _, user := range users {
		usersList = append(usersList, dto.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		})
	}
	return usersList, nil
}

func NewUserUseCase(userRepository IUserRepository) IUserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
