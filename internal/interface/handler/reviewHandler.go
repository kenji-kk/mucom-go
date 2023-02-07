package handler

import (
	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/kenji-kk/mucom-go/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ReviewHandler interface {
	// CheckToken(echo.Context) error
	GetReviews(echo.Context) *models.ReviewResultEntity
}
type reviewHnadler struct {
	usReview usecase.ReveiwUsecase
}

func NewReviewHandler(usAuth usecase.ReveiwUsecase) ReviewHandler {
	// usReview :=
	// return &ReviewHandler{usReview}
}

func (haReview *reviewHnadler) GetReviews(c echo.Context) *models.ReviewResultEntity {
	return &models.ReviewResultEntity{Reviews: []string{"item1", "item2", "item3"}, Error: nil}
}
