package injection

import (
	"github.com/kenji-kk/mucom-go/server/internal/interface/handler"
	"github.com/kenji-kk/mucom-go/server/internal/usecase"
	"github.com/kenji-kk/mucom-go/server/internal/repository"
	

	"github.com/google/wire"
)

func InitializeHandlers() handler.Handlers {
		wire.Build(
			// handler
			handler.NewHandlers,
			handler.NewAuthHandler,

			// usecase
			usecase.NewAuthUsecase,

			// repository
			repository.NewAuthRepository,

		)
    return nil
}
