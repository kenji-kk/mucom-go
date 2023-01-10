package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

func main() {
	// ポート番号 8080 で待ち受けを開始
	fmt.Println("server start")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
