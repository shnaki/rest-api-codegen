package repository

import (
	"rest-api-codegen/internal/entity"
)

type IUserRepository interface {
	GetUserByEmail(um *entity.User, email string) error
	CreateUser(um *entity.User) error
	UserExistsByEmail(email string) (bool, error)
}
