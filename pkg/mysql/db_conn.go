package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewMysqlDB() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "go_grpc:password@tcp(mysql:3306)/go_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	db, err := sqlx.Connect("mysql", dataSourceName)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
