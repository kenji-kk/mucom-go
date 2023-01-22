package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"github.com/kenji-kk/mucom-go/internal/server"
	_ "github.com/labstack/echo/middleware"
)

func main() {
	// ポート番号 8080 で待ち受けを開始
	server := server.NewServer()
	server.Run()
}
