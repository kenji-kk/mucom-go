package repository

import (
	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Hello() string
	CreateUser(*models.User) error
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository (db *sqlx.DB) AuthRepository {
	return &authRepository{db: db}
}

func (reAuth *authRepository) Hello() string{
	return "Hello World"
}

func (reAuth *authRepository) CreateUser(user *models.User) error {
	return nil
}
