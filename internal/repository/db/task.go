package db

import (
	"context"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/repository"
	"rest-api-codegen/pkg/ent"
	"rest-api-codegen/pkg/ent/task"
	"rest-api-codegen/pkg/ent/user"
)

type taskRepository struct {
	client *ent.Client
}

func (tr *taskRepository) GetAllTasks(ctx context.Context, userID uint64) ([]*entity.Task, error) {
	tasks, err := tr.client.Task.Query().
		Where(task.HasOwnerWith(user.ID(userID))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	resultTasks := make([]*entity.Task, 0, len(tasks))
	for _, entTask := range tasks {
		t := &entity.Task{}
		fromEntToEntityTask(entTask, t)
		resultTasks = append(resultTasks, t)
	}
	return resultTasks, nil
}

func (tr *taskRepository) GetTaskByID(ctx context.Context, userID uint64, taskID uint64) (*entity.Task, error) {
	t, err := tr.client.Task.Query().
		Where(
			task.HasOwnerWith(
				user.ID(userID),
			),
			task.ID(taskID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	te := &entity.Task{}
	fromEntToEntityTask(t, te)
	return te, nil
}

func (tr *taskRepository) CreateTask(ctx context.Context, te *entity.Task) error {
	t, err := tr.client.Task.Create().
		SetTitle(te.Title).
		SetOwnerID(te.UserID).
		Save(ctx)
	if err != nil {
		return err
	}
	fromEntToEntityTask(t, te)
	return nil
}

func (tr *taskRepository) UpdateTask(ctx context.Context, te *entity.Task, userID uint64, taskID uint64) error {
	t, err := tr.client.Task.
		UpdateOneID(taskID).
		Where(
			task.UserID(userID),
		).
		SetTitle(te.Title).
		Save(ctx)
	if err != nil {
		return err
	}
	fromEntToEntityTask(t, te)
	return nil
}

func (tr *taskRepository) DeleteTask(ctx context.Context, userID uint64, taskID uint64) error {
	return tr.client.Task.
		DeleteOneID(taskID).
		Where(
			task.UserID(userID),
		).
		Exec(ctx)
}

func NewTaskRepository(client *ent.Client) repository.TaskRepository {
	return &taskRepository{
		client: client,
	}
}
