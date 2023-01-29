package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
	"golang.org/x/crypto/bcrypt"
	"go.uber.org/zap"

	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/internal/repository"
	"github.com/kenji-kk/mucom-go/pkg/logger"
)

type AuthUsecase interface {
	Signup(context.Context, *models.User) (*models.User, string, error)
	Signin(context.Context, *models.User) (*models.User, string, error)
}

type authUsecase struct {
	reAuth repository.AuthRepository
}

func NewAuthUsecase(reAuth repository.AuthRepository) AuthUsecase {
	return &authUsecase{reAuth}
}

func (usAuth *authUsecase) Signup(ctx context.Context, user *models.User) (*models.User, string, error) {
	createdUser, err := usAuth.reAuth.CreateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}

	// passwordの値を空白にする
	createdUser.Password = ""
	createdUser.HashedPassword = []byte{}

	jws := createJWT(createdUser.Id.String())

	return createdUser, jws, err

}

func (usAuth *authUsecase) Signin(ctx context.Context, user *models.User) (*models.User, string, error) {
	extractedUser, err := usAuth.reAuth.GetUserByEmail(ctx, user)
	if err != nil {
		return nil, "", err
	}

	userPassword := append([]byte(user.Password), extractedUser.Salt...)

	// パスワード比較
	err = bcrypt.CompareHashAndPassword(extractedUser.HashedPassword, userPassword)
	if err != nil {
		logger.Logger.Error("An error occurred while comparing passward", zap.Error(err))
		return nil, "", err
	}

	// passwordの値を空白にする
	extractedUser.Password = ""
	extractedUser.HashedPassword = []byte{}

	jws := createJWT(user.Id.String())

	return extractedUser, jws, err
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

