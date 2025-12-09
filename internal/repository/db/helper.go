package db

import (
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/pkg/ent"
)

// fromEntToEntityUser はentのユーザー情報をエンティティにコピーする。
func fromEntToEntityUser(entUser *ent.User, userEntity *entity.User) {
	userEntity.ID = entUser.ID
	userEntity.Email = entUser.Email
	userEntity.Password = entUser.Password
	userEntity.CreatedAt = entUser.CreatedAt
	userEntity.UpdatedAt = entUser.UpdatedAt
}

// fromEntToEntityTask はentのタスク情報をエンティティにコピーする。
func fromEntToEntityTask(entTask *ent.Task, taskEntity *entity.Task) {
	taskEntity.ID = entTask.ID
	taskEntity.Title = entTask.Title
	taskEntity.CreatedAt = entTask.CreatedAt
	taskEntity.UpdatedAt = entTask.UpdatedAt
	taskEntity.UserID = entTask.UserID
}
