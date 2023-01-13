package server

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/kenji-kk/mucom-go/internal/server/rest"
	"github.com/kenji-kk/mucom-go/internal/injection"
)

type Server struct {
	echo     *echo.Echo

}

func NewServer() *Server {
	return &Server{echo: echo.New()}
}

func (s *Server) Run() error {

	rootHandlers := injection.InitializeRootHandlers()
	someHandlerInfo := rest.NewSomeHandlerInfo(rootHandlers)
	if err := s.MapHandler(s.echo, someHandlerInfo); err != nil {
		return err
	}

	fmt.Println("server start")
	s.echo.Logger.Fatal(s.echo.Start(":8080"))

	return nil

}

func (s *Server) MapHandler(e *echo.Echo, someHandlerInfo rest.SomeHandlerInfo) error {
	for _, handlerInfo := range someHandlerInfo {
		if handlerInfo.Method == "GET" {
			e.GET(handlerInfo.Path, handlerInfo.Handler)
			continue
		}
		
		if handlerInfo.Method == "POST" {
			e.POST(handlerInfo.Path, handlerInfo.Handler)
			continue
		}
	}

	return nil
}
