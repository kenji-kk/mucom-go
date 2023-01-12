package handler

import (
	"github.com/kenji-kk/mucom-go/server/internal/usecase"
)

type AuthHandler interface {

}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler (authUsecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{authUsecase}
}
