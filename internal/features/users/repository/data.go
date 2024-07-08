package repository

import (
	"news-app-be23/internal/features/articles/repository"
	"news-app-be23/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string               `json:"username"`
	Password string               `json:"password"`
	Email    string               `json:"email"`
	Articles []repository.Article `gorm:"foreignKey:UserId"`
}

func (u *User) toUserEntity() users.User {
	return users.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}

func toUserData(input users.User) User {
	return User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}
}
