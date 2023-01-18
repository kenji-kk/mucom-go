package handler

import (
	"net/http"
	"context"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/labstack/echo"
	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/pkg/utils"
)

type AuthHandler interface {
	Hello(echo.Context) error
	Signup(echo.Context) error
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

func (haAuth *authHandler) Signup(c echo.Context) error {
	ctx := context.Background()
	user := new(models.User)
	if err := utils.ReadRequest(c, user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, err := haAuth.usAuth.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	
	return c.JSON(http.StatusCreated, createdUser)
}
