package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"rest-api-codegen/internal/entity"
	"rest-api-codegen/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mock/usecase/mock/mock_$GOFILE -package=usecasemock

var ErrUserAlreadyExists = errors.New("user already exists")

type IUserUsecase interface {
	Signup(u entity.User) (entity.User, error)
	Login(u entity.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur: ur}
}

func (uu *userUsecase) Signup(ue entity.User) (entity.User, error) {
	// 登録済みかどうかをチェックする。
	exists, err := uu.ur.UserExistsByEmail(context.Background(), ue.Email)
	if err != nil {
		return entity.User{}, err
	}
	if exists {
		return entity.User{}, ErrUserAlreadyExists
	}

	// GenerateFromPassword の第2引数は暗号化の複雑化を表す。
	hash, err := bcrypt.GenerateFromPassword([]byte(ue.Password), 10)
	if err != nil {
		return entity.User{}, err
	}
	newUser := entity.User{Email: ue.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(context.Background(), &newUser); err != nil {
		return entity.User{}, fmt.Errorf("error creating user: %s", ue.Email)
	}
	resUser := entity.User{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(ue entity.User) (string, error) {
	storedUser, err := uu.ur.GetUserByEmail(context.Background(), ue.Email)
	if err != nil {
		return "", fmt.Errorf("user not found: %s", ue.Email)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password),
		[]byte(ue.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		// トークンの有効期限は12時間としておく。
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
