package mysql

import (
	_ "github.com/jackc/pgx/v4/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"

	"github.com/kenji-kk/mucom-go/pkg/logger"
)

func NewMysqlDB() (*sqlx.DB) {
	dataSourceName := "host=postgres port=5432 user=postgres dbname=app_db sslmode=disable password=password"

	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	logger.Logger.Info("success connecting DB")

	return db
}
