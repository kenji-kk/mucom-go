package usecase

import (
	"github.com/kenji-kk/mucom-go/server/internal/repository"
)

type AuthUsecase interface {

}

type authUsecase struct {
	authRepository repository.AuthRepository
}

func NewAuthUsecase (authRepository repository.AuthRepository) AuthUsecase {
	return &authUsecase{authRepository}
}

func (a *authUsecase) createUser() {
	
}
