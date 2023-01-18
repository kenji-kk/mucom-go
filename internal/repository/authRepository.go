package repository

import (
	"fmt"
	"context"
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

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db}
}

func (reAuth *authRepository) Hello() string{
	return "Hello World"
}

func (reAuth *authRepository) CreateUser(ctx context.Context, user *models.User) error {

	if err := reAuth.db.QueryRowxContext(ctx, )
}
