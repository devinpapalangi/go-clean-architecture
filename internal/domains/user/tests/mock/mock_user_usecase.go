package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-clean-architecture/internal/domains/user/dto"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) CreateUser(ctx context.Context, createUserRequest *dto.UpsertUserRequest) error {
	args := m.Called(ctx, createUserRequest)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUser(ctx context.Context, id string) (*dto.UserResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}

func (m *MockUserUseCase) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserUseCase) UpdateUser(ctx context.Context, id string, updateUserRequest *dto.UpsertUserRequest) error {
	args := m.Called(ctx, id, updateUserRequest)
	return args.Error(0)
}

func (m *MockUserUseCase) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]dto.UserResponse), args.Error(1)
}
