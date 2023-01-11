package handler

import (
	"github.com/kenji-kk/mucom-go/server/internal/usecase"
)

type Handler interface {

}

type handler struct {
	ucAuth usecase.AuthUsecase
}

