package v1

import (
	"rest-api-codegen/internal/entity"
)

// fromEntityToTaskResponse はタスクエンティティをTaskResponseにコピーする。
func fromEntityToTaskResponse(taskEntity *entity.Task, taskResponse *TaskResponse) {
	taskResponse.Id = taskEntity.ID
	taskResponse.Title = taskEntity.Title
	taskResponse.CreatedAt = taskEntity.CreatedAt.String()
	taskResponse.UpdatedAt = taskEntity.UpdatedAt.String()
}
