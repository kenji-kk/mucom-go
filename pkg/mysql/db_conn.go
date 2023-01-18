package mysql

import (
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)

func NewMysqlDB() (*sqlx.DB) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "go_grpc:password@tcp(mysql:3306)/go_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("db接続完了")

	cmdU := `CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE,
		hashedpassword LONGBLOB NOT NULL,
		salt LONGBLOB NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)
		`

	_, err = db.Exec(cmdU)
	count := 0
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				panic(err)
			}
			_, err = db.Exec(cmdU)
		}
	}
	fmt.Println("success creating table")


	return db
}
