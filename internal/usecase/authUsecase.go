package usecase

import (
	"github.com/kenji-kk/mucom-go/internal/repository"
	"github.com/kenji-kk/mucom-go/internal/models"
)

type AuthUsecase interface {
	Hello() string
	CreateUser(*models.User) error
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

func (usAuth *authUsecase) CreateUser(user *models.User) error {
	return nil
	// usAuth.reAuth.CreateUser(user)
}
