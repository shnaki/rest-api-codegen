package repository

import (
	"context"
	"rest-api-codegen/internal/entity"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, ue *entity.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}
