package rest

import (
	"github.com/kenji-kk/mucom-go/internal/interface/handler"
	"github.com/labstack/echo"
)

type HandlerInfo struct {
	Method     string
	Path       string
  Handler    echo.HandlerFunc
}

type SomeHandlerInfo []HandlerInfo

func NewSomeHandlerInfo(rootHandlers handler.RootHandlers) SomeHandlerInfo {
	someHandlerInfo := SomeHandlerInfo{
		{
			"GET",
			"/",
			rootHandlers.Hello,
		},
		{
			"POST",
			"/signup",
			rootHandlers.Signup,
		},
	}

	return someHandlerInfo
}
