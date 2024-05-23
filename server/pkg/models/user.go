package models

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	UserId    ID     `json:"user_id" gorm:"primarykey;column:id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Balance   int    `json:"balance" gorm:"column:balance;default:0"`
}

type UserResponse struct {
	User
	Balance float64 `json:"balance"`
}

// BeforeSave is a gorm hook that is called before saving a user
func (u *User) BeforeSave(tx *gorm.DB) error {
	u.FirstName = strings.ToLower(u.FirstName)
	u.LastName = strings.ToLower(u.LastName)

	return nil
}

type UserModel struct {
	baseModel
}

// NewUserModel creates a new UserModel
func NewUserModel() *UserModel {
	return &UserModel{
		baseModel{
			model: DB.Model(&User{}),
		},
	}
}

// GetUsers retrieves all users from the users table
func (m *UserModel) GetUsers() (*[]User, error) {
	var users []User

	err := m.model.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// GetUserByID retrieves a user by their ID
func (m *UserModel) GetUserByID(userID ID) (*User, error) {
	var user User

	err := m.model.Where("id = ?", userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// InsertUser inserts a new user into the users table
func (m *UserModel) InsertUser(req *CreateUserRequest) (bool, error) {
	result := m.model.Create(req)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
