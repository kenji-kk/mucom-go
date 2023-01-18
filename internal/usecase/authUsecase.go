package usecase

import (
	"context"

	"github.com/kenji-kk/mucom-go/internal/repository"
	"github.com/kenji-kk/mucom-go/internal/models"
)

type AuthUsecase interface {
	Hello() string
	CreateUser(context.Context, *models.User) (*models.User, error)
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

func (usAuth *authUsecase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return usAuth.reAuth.CreateUser(ctx, user)
}
