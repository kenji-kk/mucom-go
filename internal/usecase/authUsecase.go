package usecase

import (
	"context"
	"time"
	"github.com/dgrijalva/jwt-go"

	"github.com/kenji-kk/mucom-go/internal/repository"
	"github.com/kenji-kk/mucom-go/internal/models"
)

type AuthUsecase interface {
	Hello() string
	CreateUser(context.Context, *models.User) (*models.User, string, error)
}

type authUsecase struct {
	reAuth repository.AuthRepository
}

func NewAuthUsecase (reAuth repository.AuthRepository) AuthUsecase {
	return &authUsecase{reAuth}
}

func (usAuth *authUsecase) Hello() string {
	return usAuth.reAuth.Hello()
}

func (usAuth *authUsecase) CreateUser(ctx context.Context, user *models.User) (*models.User, string, error) {
	createdUser, err := usAuth.reAuth.CreateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}

	// passwordの値を空白にする
	createdUser.Password = ""

	// JWTの生成
	jws := createJWT(createdUser.Id.String())

	return createdUser, jws, err
	

}

func createJWT(userID string) string{
	// Claimsオブジェクト作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロード作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	jws, _ := token.SignedString([]byte("SECRET_KEY"))

	return jws
}
