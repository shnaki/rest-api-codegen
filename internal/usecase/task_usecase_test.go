package usecase

import (
	"errors"
	"testing"

	"rest-api-codegen/internal/entity"
	repositorymock "rest-api-codegen/mock/repository/mock"

	"go.uber.org/mock/gomock"
)

func TestTaskUsecase_GetAllTasks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	expected := []*entity.Task{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	mr.EXPECT().GetAllTasks(gomock.Any(), uint64(10)).Return(expected, nil)

	got, err := tu.GetAllTasks(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0].ID != 1 || got[1].Title != "b" {
		t.Fatalf("unexpected result: %+v", got)
	}
}

func TestTaskUsecase_GetAllTasks_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	mr.EXPECT().GetAllTasks(gomock.Any(), uint64(10)).Return(nil, errors.New("db"))

	_, err := tu.GetAllTasks(10)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestTaskUsecase_GetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	mr.EXPECT().GetTaskByID(gomock.Any(), uint64(10), uint64(5)).Return(&entity.Task{ID: 5, Title: "x"}, nil)
	got, err := tu.GetTaskByID(10, 5)
	if err != nil || got.ID != 5 {
		t.Fatalf("unexpected: %+v, err=%v", got, err)
	}
}

func TestTaskUsecase_GetTaskByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	mr.EXPECT().GetTaskByID(gomock.Any(), uint64(10), uint64(5)).Return(nil, errors.New("db"))
	_, err := tu.GetTaskByID(10, 5)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestTaskUsecase_CreateUpdateDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	// Create
	task := &entity.Task{Title: "new"}
	mr.EXPECT().CreateTask(gomock.Any(), task).Return(nil)
	if err := tu.CreateTask(task); err != nil {
		t.Fatalf("create unexpected err: %v", err)
	}

	// Update
	mr.EXPECT().UpdateTask(gomock.Any(), gomock.Any(), uint64(10), uint64(1)).Return(nil)
	if err := tu.UpdateTask(&entity.Task{Title: "u"}, 10, 1); err != nil {
		t.Fatalf("update unexpected err: %v", err)
	}

	// Delete
	mr.EXPECT().DeleteTask(gomock.Any(), uint64(10), uint64(1)).Return(nil)
	if err := tu.DeleteTask(10, 1); err != nil {
		t.Fatalf("delete unexpected err: %v", err)
	}
}

func TestTaskUsecase_CreateUpdateDelete_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repositorymock.NewMockTaskRepository(ctrl)
	tu := NewTaskUsecase(mr)

	mr.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(errors.New("db"))
	if err := tu.CreateTask(&entity.Task{Title: "x"}); err == nil {
		t.Fatalf("expected error")
	}

	mr.EXPECT().UpdateTask(gomock.Any(), gomock.Any(), uint64(10), uint64(1)).Return(errors.New("db"))
	if err := tu.UpdateTask(&entity.Task{Title: "u"}, 10, 1); err == nil {
		t.Fatalf("expected error")
	}

	mr.EXPECT().DeleteTask(gomock.Any(), uint64(10), uint64(1)).Return(errors.New("db"))
	if err := tu.DeleteTask(10, 1); err == nil {
		t.Fatalf("expected error")
	}
}
