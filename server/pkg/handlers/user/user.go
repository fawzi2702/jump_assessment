package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/pkg/models"
	"github.com/this_is_iz/jump_server/pkg/response"
	"github.com/this_is_iz/jump_server/pkg/transformers"
)

func GetUsers(c *gin.Context) {
	userModel := models.NewUserModel()

	users, err := userModel.GetUsers()
	if err != nil {
		response.InternalServerError(c)
		return
	}

	response.Ok(c, transformers.TransformUsers(*users))
}

func GetUser(c *gin.Context) {
	userModel := models.NewUserModel()

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.BadRequest(c, "invalid user_id param")
		return
	}

	user, err := userModel.GetUserByID(userID)
	if err != nil {
		response.InternalServerError(c)
		return
	} else if user == nil {
		response.NotFound(c, "user not found")
		return
	}

	response.Ok(c, transformers.TransformUser(*user))
}

func CreateUser(c *gin.Context) {
	userModel := models.NewUserModel()

	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
		return
	}

	ok, err := userModel.InsertUser(&req)
	if err != nil || !ok {
		response.InternalServerError(c)
		return
	}

	response.NoContent(c)
}
