package usecase

import (
	"context"
	"errors"
	"os"
	"testing"

	"rest-api-codegen/internal/entity"
	repositorymock "rest-api-codegen/mock/repository/mock"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUsecase_Signup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	email := "test@example.com"
	pass := "password123"

	ur.EXPECT().UserExistsByEmail(gomock.Any(), email).Return(false, nil)
	ur.EXPECT().CreateUser(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, u *entity.User) error {
		// Simulate DB setting ID
		u.ID = 42
		return nil
	})

	user, err := uu.Signup(entity.User{Email: email, Password: pass})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID != 42 || user.Email != email || user.Password != "" {
		t.Fatalf("unexpected user returned: %+v", user)
	}
}

func TestUserUsecase_Signup_AlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	email := "dup@example.com"
	ur.EXPECT().UserExistsByEmail(gomock.Any(), email).Return(true, nil)

	_, err := uu.Signup(entity.User{Email: email, Password: "x"})
	if !errors.Is(err, ErrUserAlreadyExists) {
		t.Fatalf("expected ErrUserAlreadyExists, got %v", err)
	}
}

func TestUserUsecase_Signup_RepoErrorOnExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	ur.EXPECT().UserExistsByEmail(gomock.Any(), gomock.Any()).Return(false, errors.New("boom"))

	_, err := uu.Signup(entity.User{Email: "a@b.c", Password: "x"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestUserUsecase_Signup_RepoErrorOnCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	email := "new@example.com"
	ur.EXPECT().UserExistsByEmail(gomock.Any(), email).Return(false, nil)
	ur.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("db create error"))

	_, err := uu.Signup(entity.User{Email: email, Password: "pw"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestUserUsecase_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	// Prepare stored user with bcrypt hash
	plain := "secretPW!"
	hash, _ := bcrypt.GenerateFromPassword([]byte(plain), 10)
	stored := &entity.User{ID: 7, Email: "login@example.com", Password: string(hash)}

	ur.EXPECT().GetUserByEmail(gomock.Any(), stored.Email).Return(stored, nil)

	// SECRET must be set for signing
	t.Setenv("SECRET", "testsecret")

	tok, err := uu.Login(entity.User{Email: stored.Email, Password: plain})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok == "" {
		t.Fatalf("expected non-empty token")
	}
	// Basic sanity: token should parse with same secret
	_, parseErr := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if parseErr != nil {
		t.Fatalf("token should be parseable: %v", parseErr)
	}
}

func TestUserUsecase_Login_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	ur.EXPECT().GetUserByEmail(gomock.Any(), "no@ex.com").Return(nil, errors.New("not found"))

	_, err := uu.Login(entity.User{Email: "no@ex.com", Password: "x"})
	if err == nil || err.Error() != "user not found: no@ex.com" {
		t.Fatalf("unexpected err: %v", err)
	}
}

func TestUserUsecase_Login_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := repositorymock.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	hash, _ := bcrypt.GenerateFromPassword([]byte("correct"), 10)
	ur.EXPECT().GetUserByEmail(gomock.Any(), "u@e.com").Return(&entity.User{ID: 1, Email: "u@e.com", Password: string(hash)}, nil)

	_, err := uu.Login(entity.User{Email: "u@e.com", Password: "wrong"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
