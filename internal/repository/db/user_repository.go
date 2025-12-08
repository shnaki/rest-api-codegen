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

func NewUserRepository(client *ent.Client) repository.IUserRepository {
	return &userRepository{
		client: client,
	}
}

func (ur *userRepository) GetUserByEmail(um *entity.User, email string) error {
	u, err := ur.client.User.
		Query().
		Where(user.Email(email)).
		Only(context.Background())
	if err != nil {
		return err
	}
	fromEntToEntityUser(u, um)
	return nil
}

func (ur *userRepository) CreateUser(um *entity.User) error {
	u, err := ur.client.User.
		Create().
		SetEmail(um.Email).
		SetPassword(um.Password).
		Save(context.Background())
	if err != nil {
		return err
	}
	fromEntToEntityUser(u, um)
	return nil
}

func (ur *userRepository) UserExistsByEmail(email string) (bool, error) {
	exists, err := ur.client.User.
		Query().
		Where(user.Email(email)).
		Exist(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}
	return exists, nil
}
