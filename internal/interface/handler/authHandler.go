package handler

import (
	"net/http"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/labstack/echo"
	// "github.com/kenji-kk/mucom-go/internal/models"
)

type AuthHandler interface {
	Hello(echo.Context) error
	// Signup(echo.Context) error
}

type authHandler struct {
	usAuth usecase.AuthUsecase
}

func NewAuthHandler(usAuth usecase.AuthUsecase) AuthHandler {
	return &authHandler{usAuth}
}

func (haAuth *authHandler) Hello(c echo.Context) error {
	hello := haAuth.usAuth.Hello()
	return c.String(http.StatusOK, hello)
}

// func (haAuth *authHandler) Signup(c echo.Context) error {
// 	user := new(models.User)

// }
