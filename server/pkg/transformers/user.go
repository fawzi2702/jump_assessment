package transformers

import (
	"github.com/this_is_iz/jump_server/internal/utils"
	"github.com/this_is_iz/jump_server/pkg/models"
)

// TransformUser transforms a user model into a user response
func TransformUser(u models.User) models.UserResponse {
	balance, err := utils.AddDecimalPoint(u.Balance)
	if err != nil {
		balance = 0
	}

	return models.UserResponse{
		User:    u,
		Balance: balance,
	}
}

// TransformUsers transforms a slice of user models into a slice of user responses
func TransformUsers(users []models.User) []models.UserResponse {
	return utils.Map(users, TransformUser)
}
