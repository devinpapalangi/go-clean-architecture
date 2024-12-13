package user

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/common"
	"go-clean-architecture/internal/domains/user/dto"
	"go-clean-architecture/pkg"
	"net/http"
)

type UserHTTP struct {
	userUseCase IUserUseCase
}

// Get User godoc
// Get User swagger
//
//	@Summary		Get a user
//	@Description	Get a user by its ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"User ID"
//	@Success		200				{object}	dto.SingleUserResponse
//	@Failed			404 {object}	common.ErrorResponse
//
//	@Security		ApiKeyAuth
//	@in				header
//	@name			Authorization
//	@description	Description for what is this security definition being used
//
//	@Router			/v1/user/{id} [get]
func (u *UserHTTP) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userUseCase.GetUser(c, id)

	if err != nil {
		pkg.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SingleUserResponse{
		Data:    *user,
		Message: "Success getting user",
	})
}

// Create User godoc
//
//	@Summary		Create a user
//	@Description	Create a new user by providing details
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UpsertUserRequest	true	"Add User"
//	@Success		201		{object}	common.NoDataResponse
//	@Failed			404								{object}	common.ErrorResponse
//	@Failed			400								{object}	common.ErrorResponse
//	@Failed			500								{object}	common.ErrorResponse
//
//	@Security		ApiKeyAuth
//	@in				header
//	@name			Authorization
//	@description	Description for what is this security definition being used
//
//	@Router			/v1/user [post]
func (u *UserHTTP) CreateUser(ctx *gin.Context) {
	var request dto.UpsertUserRequest

	err := pkg.HandleRequestValidation(ctx, &request)
	if err != nil {
		pkg.HandleError(ctx, err)
		return
	}

	err = u.userUseCase.CreateUser(ctx, &request)
	if err != nil {
		pkg.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, common.NoDataResponse{
		Message: "User created successfully",
	})
	return
}

// Delete User godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by its ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"User ID"
//	@Success		200				{object}	common.NoDataResponse
//	@Failed			400 {object}	response.ErrorResponse
//
//	@Security		ApiKeyAuth
//	@in				header
//	@name			Authorization
//	@description	Description for what is this security definition being used
//
//	@Router			/v1/user/{id} [delete]
func (u *UserHTTP) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.userUseCase.DeleteUser(c, id)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, common.NoDataResponse{
		Message: "Success deleting user!",
	})
}

// Update User godoc
//
//	@Summary		Update a user
//	@Description	Update a user data
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string					true	"User ID"
//	@Param			user			body		dto.UpsertUserRequest	true	"Add User"
//	@Success		200				{object}	common.NoDataResponse
//	@Failed			400 {object}	response.ErrorResponse
//
//	@Security		ApiKeyAuth
//	@in				header
//	@name			Authorization
//	@description	Description for what is this security definition being used
//
//	@Router			/v1/user/{id} [put]
func (u *UserHTTP) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var request dto.UpsertUserRequest

	err := pkg.HandleRequestValidation(c, &request)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}

	err = u.userUseCase.UpdateUser(c, id, &request)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, common.NoDataResponse{
		Message: "Success updating user!",
	})
}

// Get All Users godoc
//
//	@Summary		Get all users data
//	@Description	Get all users data except its password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.ListBaseResponse[dto.UserResponse]
//	@Security		ApiKeyAuth
//		@in				header
//		@name			Authorization
//		@description	Description for what is this security definition being used
//
//	@Router			/v1/user [get]
func (u *UserHTTP) GetAllUsers(c *gin.Context) {
	users, err := u.userUseCase.GetAllUsers(c)

	if err != nil {
		pkg.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, common.ListBaseResponse[dto.UserResponse]{
		Data:    users,
		Message: "success retrieving all users",
	})

}

func NewUserHTTP(userUseCase IUserUseCase) IUserHTTP {
	return &UserHTTP{
		userUseCase: userUseCase,
	}
}

func SetupRoutes(router *gin.RouterGroup, handler IUserHTTP) {
	userRoutes := router.Group("/v1/user")

	userRoutes.POST("", handler.CreateUser)
	userRoutes.GET("", handler.GetAllUsers)
	userRoutes.GET("/:id", handler.GetUser)
	userRoutes.DELETE("/:id", handler.DeleteUser)
	userRoutes.PUT("/:id", handler.UpdateUser)
}
