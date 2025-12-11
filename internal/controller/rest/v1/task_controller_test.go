package v1

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"rest-api-codegen/internal/controller/rest/middleware/jwt"
	"rest-api-codegen/internal/entity"
	usecasemock "rest-api-codegen/mock/usecase/mock"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

// helper to build ctx with userID via middleware.SuccessHandler
func ctxWithUserID(t *testing.T, userID uint64) context.Context {
	t.Helper()
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	token := &jwtv5.Token{Claims: jwtv5.MapClaims{"user_id": float64(userID)}}
	c.Set("user", token)
	jwt.SuccessHandler(c)
	return c.Request().Context()
}

func TestTaskController_GetAllTasks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)

	ctx := ctxWithUserID(t, 99)
	tasks := []*entity.Task{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}
	tu.EXPECT().GetAllTasks(uint64(99)).Return(tasks, nil)

	resp, err := tc.GetAllTasks(ctx, GetAllTasksRequestObject{})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	r, ok := resp.(GetAllTasks200JSONResponse)
	if !ok || len(r) != 2 || r[0].Id != 1 || r[1].Title != "b" {
		t.Fatalf("unexpected response: %#v (%T)", resp, resp)
	}
}

func TestTaskController_GetAllTasks_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)
	ctx := ctxWithUserID(t, 99)
	tu.EXPECT().GetAllTasks(uint64(99)).Return(nil, errors.New("db"))
	resp, _ := tc.GetAllTasks(ctx, GetAllTasksRequestObject{})
	if _, ok := resp.(GetAllTasks500ApplicationProblemPlusJSONResponse); !ok {
		t.Fatalf("expected 500 response, got %T", resp)
	}
}

func TestTaskController_CreateTask_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)
	ctx := ctxWithUserID(t, 7)
	body := TaskRequest{Title: "new"}
	tu.EXPECT().CreateTask(gomock.Any()).DoAndReturn(func(tk *entity.Task) error { tk.ID = 10; return nil })
	resp, err := tc.CreateTask(ctx, CreateTaskRequestObject{Body: &body})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	r, ok := resp.(CreateTask201JSONResponse)
	if !ok || r.Id != 10 || r.Title != "new" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestTaskController_CreateTask_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)
	ctx := ctxWithUserID(t, 7)
	body := TaskRequest{Title: "new"}
	tu.EXPECT().CreateTask(gomock.Any()).Return(errors.New("db"))
	resp, _ := tc.CreateTask(ctx, CreateTaskRequestObject{Body: &body})
	if _, ok := resp.(CreateTask500ApplicationProblemPlusJSONResponse); !ok {
		t.Fatalf("expected 500, got %T", resp)
	}
}

func TestTaskController_GetTaskByID_Update_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)
	ctx := ctxWithUserID(t, 3)

	// Get by ID
	tu.EXPECT().GetTaskByID(uint64(3), uint64(11)).Return(&entity.Task{ID: 11, Title: "t"}, nil)
	gr, err := tc.GetTaskByID(ctx, GetTaskByIDRequestObject{TaskID: 11})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if r, ok := gr.(GetTaskByID200JSONResponse); !ok || r.Id != 11 {
		t.Fatalf("unexpected response: %#v", gr)
	}

	// Update
	tu.EXPECT().UpdateTask(gomock.Any(), uint64(3), uint64(11)).Return(nil)
	ur, err := tc.UpdateTask(ctx, UpdateTaskRequestObject{TaskID: 11, Body: &TaskRequest{Title: "u"}})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if r, ok := ur.(UpdateTask200JSONResponse); !ok || r.Title != "u" {
		t.Fatalf("unexpected response: %#v", ur)
	}

	// Delete
	tu.EXPECT().DeleteTask(uint64(3), uint64(11)).Return(nil)
	dr, err := tc.DeleteTask(ctx, DeleteTaskRequestObject{TaskID: 11})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if _, ok := dr.(DeleteTask204Response); !ok {
		t.Fatalf("unexpected response: %#v", dr)
	}
}

func TestTaskController_GetTaskByID_Update_Delete_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tu := usecasemock.NewMockITaskUsecase(ctrl)
	tc := NewTaskController(tu)
	ctx := ctxWithUserID(t, 3)

	tu.EXPECT().GetTaskByID(uint64(3), uint64(11)).Return(nil, errors.New("db"))
	if r, _ := tc.GetTaskByID(ctx, GetTaskByIDRequestObject{TaskID: 11}); r == nil {
		t.Fatalf("expected error response")
	}

	tu.EXPECT().UpdateTask(gomock.Any(), uint64(3), uint64(11)).Return(errors.New("db"))
	if r, _ := tc.UpdateTask(ctx, UpdateTaskRequestObject{TaskID: 11, Body: &TaskRequest{Title: "u"}}); r == nil {
		t.Fatalf("expected error response")
	}

	tu.EXPECT().DeleteTask(uint64(3), uint64(11)).Return(errors.New("db"))
	if r, _ := tc.DeleteTask(ctx, DeleteTaskRequestObject{TaskID: 11}); r == nil {
		t.Fatalf("expected error response")
	}
}
