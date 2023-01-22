//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/kenji-kk/mucom-go/internal/interface/handler"
	"github.com/kenji-kk/mucom-go/internal/repository"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/kenji-kk/mucom-go/pkg/db/postgres"

	"github.com/google/wire"
)

func InitializeRootHandlers() handler.RootHandlers {
	wire.Build(
		// handler
		handler.NewRootHandlers,
		handler.NewAuthHandler,

		// usecase
		usecase.NewAuthUsecase,

		// repository
		repository.NewAuthRepository,

		// DB
		mysql.NewMysqlDB,
	)
	return nil
}
