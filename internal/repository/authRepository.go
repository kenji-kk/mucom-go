package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/pkg/mysql"
)

type AuthRepository interface {
	Hello() string
	CreateUser(*models.User) error
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository() AuthRepository {
	db, err := mysql.NewMysqlDB()
	if err !=nil {
		fmt.Println(err)
	}
	return &authRepository{db}
}

func (reAuth *authRepository) Hello() string{
	return "Hello World"
}

func (reAuth *authRepository) CreateUser(user *models.User) error {
	return nil
}
