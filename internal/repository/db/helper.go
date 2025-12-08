package db

import (
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/pkg/ent"
)

// fromEntToEntityUser はentのユーザー情報を共通モデルにコピーする。
func fromEntToEntityUser(ue *ent.User, um *entity.User) {
	um.ID = ue.ID
	um.Email = ue.Email
	um.Password = ue.Password
	um.CreatedAt = ue.CreatedAt
	um.UpdatedAt = ue.UpdatedAt
}
