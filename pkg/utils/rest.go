package utils

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

func WriteCookie(c echo.Context, cookieName, cookievalue string, expiresTime time.Duration) error {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookievalue,
		Expires:  time.Now().Add(expiresTime * time.Hour),
		HttpOnly: true,
	}
	
  c.SetCookie(cookie)
  return nil
}
