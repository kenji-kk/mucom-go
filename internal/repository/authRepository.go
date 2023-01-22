package repository

import (
	"context"
	"crypto/rand"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/kenji-kk/mucom-go/internal/const/sql"
	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/pkg/logger"
)

type AuthRepository interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db}
}

func (reAuth *authRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	u := new(models.User)

	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}
	user.Salt = salt

	// hashPassword作成
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Error("An error occurred while creating the hashedPassword", zap.Error(err))
		return nil, err
	}
	user.HashedPassword = hashedPassword

	// ユーザ作成
	if err := reAuth.db.QueryRowxContext(ctx, sql.CreateUserQuery, &user.UserName, &user.Email, &user.Salt, &user.HashedPassword).StructScan(u); err != nil {
		logger.Logger.Error("An error occurred while inserting user-data in DB", zap.Error(err))
		return nil, err
	}

	return u, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		logger.Logger.Error("An error occurred while creating the salt", zap.Error(err))
		return nil, err
	}
	return salt, nil
}
