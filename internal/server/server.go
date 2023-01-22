package server

import (
	"github.com/labstack/echo"

	"github.com/kenji-kk/mucom-go/internal/injection"
	"github.com/kenji-kk/mucom-go/internal/interface/handler"
	"github.com/kenji-kk/mucom-go/pkg/logger"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{echo: echo.New()}
}

func (s *Server) Run() error {
	// roothandlers生成
	rootHandlers := injection.InitializeRootHandlers()

	s.MapHandler(rootHandlers)

	logger.Logger.Info("server start")
	s.echo.Logger.Fatal(s.echo.Start(":8080"))

	return nil
}

func (s *Server) MapHandler(rootHandlers handler.RootHandlers) error {
	s.echo.POST("/signup", rootHandlers.Signup)

	return nil
}
