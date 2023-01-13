package usecase

import (
	"github.com/kenji-kk/mucom-go/internal/repository"
)

type AuthUsecase interface {
	Hello() string
}

type authUsecase struct {
	reAuth repository.AuthRepository
}

func NewAuthUsecase (reAuth repository.AuthRepository) AuthUsecase {
	return &authUsecase{reAuth}
}

func (usAuth *authUsecase) Hello() string{
	return usAuth.reAuth.Hello()
}
