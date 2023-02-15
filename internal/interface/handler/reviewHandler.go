package handler

import (
	"log"
	"net/http"

	"github.com/kenji-kk/mucom-go/internal/models"
	"github.com/labstack/echo/v4"
)

type ReviewHandler interface {
	// CheckToken(echo.Context) error
	GetReviews(echo.Context) error
}
type reviewHnadler struct {
	// usReview usecase.ReveiwUsecase
}

func NewReviewHandler() ReviewHandler {
	return &reviewHnadler{}
}

func (haReview *reviewHnadler) GetReviews(c echo.Context) error {
	log.Print("-- --reviews-- --")
	reviewResult := &models.ReviewResultEntity{Reviews: []string{"item1", "item2", "item3"}, Error: nil}
	return c.JSON(http.StatusOK, reviewResult)

}
