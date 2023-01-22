package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/internal/repository"
)

type AuthUsecase interface {
	CreateUser(context.Context, *models.User) (*models.User, string, error)
}

type authUsecase struct {
	reAuth repository.AuthRepository
}

func NewAuthUsecase(reAuth repository.AuthRepository) AuthUsecase {
	return &authUsecase{reAuth}
}

func (usAuth *authUsecase) CreateUser(ctx context.Context, user *models.User) (*models.User, string, error) {
	createdUser, err := usAuth.reAuth.CreateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}

	// passwordの値を空白にする
	createdUser.Password = ""

	jws := createJWT(createdUser.Id.String())

	return createdUser, jws, err

}

func createJWT(userID string) string {
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
