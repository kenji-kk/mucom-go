package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kenji-kk/mucom-go/internal/interface/handler"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{echo: echo.New()}
}

func (s *Server) Run() error {
	// roothandlers生成
	// rootHandlers := injection.InitializeRootHandlers()
	// s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3000"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))
	// s.MapHandler(rootHandlers)
	// logger.Logger.Info("server start")
	// s.echo.Use(middleware.Logger())
	// s.echo.Use(middleware.CORS())
	// s.echo.Logger.Fatal(s.echo.Start(":8080"))

	return nil
}

func (s *Server) MapHandler(rootHandlers handler.RootHandlers) error {
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Logger())

	s.echo.POST("/signup", rootHandlers.Signup)
	s.echo.POST("/signin", rootHandlers.Signin)
	s.echo.GET("/reviews", rootHandlers.GetReviews)

	return nil
}
