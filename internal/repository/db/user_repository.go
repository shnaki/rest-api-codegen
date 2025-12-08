package db

import (
	"context"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/pkg/ent"
	"rest-api-codegen/pkg/ent/user"
)

type IUserRepository interface {
	GetUserByEmail(um *entity.User, email string) error
	CreateUser(um *entity.User) error
}

type userRepository struct {
	client *ent.Client
}

func (ur *userRepository) GetUserByEmail(um *entity.User, email string) error {
	u, err := ur.client.User.
		Query().
		Where(user.Email(email)).
		Only(context.Background())
	if err != nil {
		return err
	}
	copyUser(u, um)
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
	copyUser(u, um)
	return nil
}

// CopyUser はentのユーザー情報を共通モデルにコピーする。
func copyUser(ue *ent.User, um *entity.User) {
	um.ID = ue.ID
	um.Email = ue.Email
	um.Password = ue.Password
	um.CreatedAt = ue.CreatedAt
	um.UpdatedAt = ue.UpdatedAt
}
