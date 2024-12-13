package tests

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-clean-architecture/common"
	"go-clean-architecture/internal/domains/user"
	"go-clean-architecture/internal/domains/user/dto"
	mocks "go-clean-architecture/internal/domains/user/tests/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUserRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mocks.MockUserUseCase)
	mockUseCase.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	handler := user.NewUserHTTP(mockUseCase)

	expectedResponse := common.NoDataResponse{
		Message: "User created successfully",
	}

	jsonMockDto, _ := json.Marshal(mocks.MockUpsertRequestDto)
	stringifyJsonMockDto := strings.NewReader(string(jsonMockDto))

	jsonMockExpectedResponse, _ := json.Marshal(expectedResponse)
	jsonMockExpectedResponseStr := string(jsonMockExpectedResponse)

	router.POST(mocks.MockUsersRouterPath, handler.CreateUser)

	req, _ := http.NewRequest("POST", mocks.MockUsersRouterPath, stringifyJsonMockDto)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, jsonMockExpectedResponseStr, w.Body.String())
}

func TestGetAllUserRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mocks.MockUserUseCase)

	mockUseCase.On("GetAllUsers", mock.Anything).Return(mocks.MockBulkUsersResponse, nil)

	router := gin.Default()
	handler := user.NewUserHTTP(mockUseCase)

	expectedResponse := common.ListBaseResponse[dto.UserResponse]{
		Data:    mocks.MockBulkUsersResponse,
		Message: "success retrieving all users",
	}

	jsonMockExpectedResponse, _ := json.Marshal(expectedResponse)
	jsonMockExpectedResponseStr := string(jsonMockExpectedResponse)
	path := mocks.MockUsersRouterPath
	req, _ := http.NewRequest("GET", path, nil)

	router.GET(path, handler.GetAllUsers)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, jsonMockExpectedResponseStr, w.Body.String())
}

func TestGetUserRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mocks.MockUserUseCase)

	mockUseCase.On("GetUser", mock.Anything, mock.Anything).Return(mocks.MockUserResponseDto, nil)

	router := gin.Default()
	handler := user.NewUserHTTP(mockUseCase)

	expectedResponse := dto.SingleUserResponse{
		Data:    *mocks.MockUserResponseDto,
		Message: "Success getting user",
	}

	jsonMockExpectedResponse, _ := json.Marshal(expectedResponse)
	jsonMockExpectedResponseStr := string(jsonMockExpectedResponse)

	path := fmt.Sprintf("%s/%s", mocks.MockUsersRouterPath, mocks.MockUserDetailID)

	req, _ := http.NewRequest("GET", path, nil)

	router.GET(path, handler.GetUser)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, jsonMockExpectedResponseStr, w.Body.String())
}

func TestUpdateUserRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mocks.MockUserUseCase)

	mockUseCase.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	handler := user.NewUserHTTP(mockUseCase)

	expectedResponse := common.NoDataResponse{
		Message: "Success updating user!",
	}

	jsonMockDto, _ := json.Marshal(mocks.MockUpsertRequestDto)
	stringifyJsonMockDto := strings.NewReader(string(jsonMockDto))

	jsonMockExpectedResponse, _ := json.Marshal(expectedResponse)
	jsonMockExpectedResponseStr := string(jsonMockExpectedResponse)

	path := fmt.Sprintf("%s/%s", mocks.MockUsersRouterPath, mocks.MockUserDetailID)

	router.PUT(path, handler.UpdateUser)

	req, _ := http.NewRequest("PUT", path, stringifyJsonMockDto)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, jsonMockExpectedResponseStr, w.Body.String())
}

func TestDeleteUserRoutes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(mocks.MockUserUseCase)
	mockUseCase.On("DeleteUser", mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	handler := user.NewUserHTTP(mockUseCase)

	expectedResponse := common.NoDataResponse{
		Message: "Success deleting user!",
	}

	jsonMockExpectedResponse, _ := json.Marshal(expectedResponse)
	jsonMockExpectedResponseStr := string(jsonMockExpectedResponse)

	path := fmt.Sprintf("%s/%s", mocks.MockUsersRouterPath, mocks.MockUserDetailID)

	router.DELETE(path, handler.DeleteUser)

	req, _ := http.NewRequest("DELETE", path, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, jsonMockExpectedResponseStr, w.Body.String())

}
