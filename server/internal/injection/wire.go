package injection

import (
	"github.com/kenji-kk/mucom-go/server/internal/usecase"

	"github.com/google/wire"
)

func InitializeEvent() usecase.AuthUsecase {
		wire.Build()
    return usecase.WauthUsecase{}
}
