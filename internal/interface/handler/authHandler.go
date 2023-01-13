package handler

import (
	"net/http"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/labstack/echo"
)

type AuthHandler interface {
	Hello(c echo.Context) error
}

type authHandler struct {
	usAuth usecase.AuthUsecase
}

func NewAuthHandler (usAuth usecase.AuthUsecase) AuthHandler {
	return &authHandler{usAuth}
}

func (haAuth *authHandler) Hello(c echo.Context) error{
	hello := haAuth.usAuth.Hello()
	return c.String(http.StatusOK, hello)
}
