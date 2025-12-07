package repository

import (
	"context"
	"rest-api-codegen/ent"
	"rest-api-codegen/ent/user"
	"rest-api-codegen/model"
)

type IUserRepository interface {
	GetUserByEmail(um *model.User, email string) error
	CreateUser(um *model.User) error
}

type userRepository struct {
	client *ent.Client
}

func (ur *userRepository) GetUserByEmail(um *model.User, email string) error {
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

func (ur *userRepository) CreateUser(um *model.User) error {
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
func copyUser(ue *ent.User, um *model.User) {
	um.ID = ue.ID
	um.Email = ue.Email
	um.Password = ue.Password
	um.CreatedAt = ue.CreatedAt
	um.UpdatedAt = ue.UpdatedAt
}
