package db

import (
	"context"
	"fmt"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/repository"
	"rest-api-codegen/pkg/ent"
	"rest-api-codegen/pkg/ent/user"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, err := ur.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	ue := &entity.User{}
	fromEntToEntityUser(u, ue)
	return ue, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, ue *entity.User) error {
	u, err := ur.client.User.
		Create().
		SetEmail(ue.Email).
		SetPassword(ue.Password).
		Save(ctx)
	if err != nil {
		return err
	}
	fromEntToEntityUser(u, ue)
	return nil
}

func (ur *userRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := ur.client.User.
		Query().
		Where(user.Email(email)).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	return exists, nil
}
