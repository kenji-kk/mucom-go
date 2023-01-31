package handler

import (
	"context"
	"github.com/kenji-kk/mucom-go/internal/const/rest/cookie"
	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/kenji-kk/mucom-go/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler interface {
	Signup(echo.Context) error
	Signin(echo.Context) error
}

type authHandler struct {
	usAuth usecase.AuthUsecase
}

func NewAuthHandler(usAuth usecase.AuthUsecase) AuthHandler {
	return &authHandler{usAuth}
}

func (haAuth *authHandler) Signup(c echo.Context) error {
	ctx := context.Background()
	user := new(models.User)
	if err := utils.ReadRequest(c, user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, jws, err := haAuth.usAuth.Signup(ctx, user)
	if err != nil || jws == "" {
		return c.JSON(http.StatusBadRequest, err)
	}

	utils.WriteCookie(c, cookie.JWSCookieName, jws, 1)


	return c.JSON(http.StatusCreated, createdUser)
}

func (haAuth *authHandler) Signin(c echo.Context) error {
	ctx := context.Background()
	user := new(models.User)
	if err := utils.ReadRequest(c, user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	singinUser, jws, err := haAuth.usAuth.Signin(ctx, user)
	if err != nil || jws == "" {
		return c.JSON(http.StatusBadRequest, err)
	}

	utils.WriteCookie(c, cookie.JWSCookieName, jws, 1)

	return c.JSON(http.StatusOK, singinUser)
}
