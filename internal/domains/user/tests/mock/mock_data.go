package mocks

import "go-clean-architecture/internal/domains/user/dto"

var (
	MockUpsertRequestDto = dto.UpsertUserRequest{
		Email:    "johndoe@gmail.com",
		Username: "johndoe",
		Name:     "John Doe",
		Password: "123456",
	}

	MockUserResponseDto = &dto.UserResponse{
		ID:       "ctd0c79gm80jksnkss21",
		Name:     "John Doe",
		Username: "johndoe",
		Email:    "johndoe@gmail.com",
	}

	MockBulkUsersResponse = []dto.UserResponse{
		{
			ID:       "ctd0c79gm80jksnkss21",
			Name:     "John Doe",
			Username: "johndoe",
			Email:    "johndoe@gmail.com",
		},
		{
			ID:       "ctd2c79gm80gssak4sdkaw",
			Name:     "John Doe2",
			Username: "johndoe2",
			Email:    "johndoe2@gmail.com",
		},
		{
			ID:       "ctd3c79asdk2k123",
			Name:     "John Doe3",
			Username: "johndoe3",
			Email:    "johndoe3@gmail.com",
		},
	}

	MockUsersRouterPath = "/api/v1/user"
	MockUserDetailID    = "ctd0c79gm80jksnkss21"
)
