package repository

import (
	"context"
	"rest-api-codegen/internal/entity"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mock/repository/mock/mock_user_repository.go -package=repositorymock

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, ue *entity.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}
